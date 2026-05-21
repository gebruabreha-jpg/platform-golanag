package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
	"golang.org/x/time/rate"
	"http-service/internal/iputils"
)

type Stats struct {
	Requests        int64
	BytesSent       int64
	BytesReceived   int64
	PacketsSent     int64
	PacketsReceived int64
	Errors          int64
	Timeouts        int64
}

type BandwidthConfig struct {
	UploadLimit   *rate.Limiter
	DownloadLimit *rate.Limiter
}

type IPPool struct {
	IPs     []string
	current int
	mu      sync.Mutex
}

func (p *IPPool) GetNextIP() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.IPs) == 0 {
		return ""
	}

	ip := p.IPs[p.current]
	p.current = (p.current + 1) % len(p.IPs)
	return ip
}

type Client struct {
	IPPool      *IPPool
	HTTPClient  *http.Client
	Debug       bool
	Quiet       bool
	RST         bool
	Bandwidth   BandwidthConfig
	stats       Stats
}

func New(cidr string, debug, quiet, rst bool, uploadBW, downloadBW string) (*Client, error) {
	ips, err := iputils.GenerateIPsFromCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("failed to generate IP pool: %w", err)
	}

	uploadLimit, err := parseBandwidth(uploadBW)
	if err != nil {
		return nil, fmt.Errorf("invalid upload bandwidth: %w", err)
	}

	downloadLimit, err := parseBandwidth(downloadBW)
	if err != nil {
		return nil, fmt.Errorf("invalid download bandwidth: %w", err)
	}

	transport := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  true,
	}

	return &Client{
		IPPool: &IPPool{
			IPs:     ips,
			current: 0,
		},
		HTTPClient: &http.Client{
			Transport: transport,
		},
		Debug: debug,
		Quiet: quiet,
		RST:   rst,
		Bandwidth: BandwidthConfig{
			UploadLimit:   uploadLimit,
			DownloadLimit: downloadLimit,
		},
	}, nil
}

func parseBandwidth(bw string) (*rate.Limiter, error) {
    if bw == "" {
        return nil, nil
    }

    bw = strings.ToLower(bw)
    var bitsPerSec float64
    var multiplier float64 = 1

    switch {
    case strings.HasSuffix(bw, "kbps"):
        multiplier = 1000 // 1 kbps = 1000 bps
        bw = strings.TrimSuffix(bw, "kbps")
    case strings.HasSuffix(bw, "mbps"):
        multiplier = 1000 * 1000 // 1 mbps = 1,000,000 bps
        bw = strings.TrimSuffix(bw, "mbps")
    case strings.HasSuffix(bw, "gbps"):
        multiplier = 1000 * 1000 * 1000 // 1 gbps = 1,000,000,000 bps
        bw = strings.TrimSuffix(bw, "gbps")
    case strings.HasSuffix(bw, "bps"):
        multiplier = 1
        bw = strings.TrimSuffix(bw, "bps")
    default:
        return nil, fmt.Errorf("invalid bandwidth unit")
    }

    arate, err := strconv.ParseFloat(bw, 64)
    if err != nil {
        return nil, err
    }

    bitsPerSec = arate * multiplier
    if bitsPerSec <= 0 {
        return nil, fmt.Errorf("bandwidth must be positive")
    }

    // Convert bits to bytes (8 bits = 1 byte)
    bytesPerSec := int(bitsPerSec / 8)
    return rate.NewLimiter(rate.Limit(bytesPerSec), bytesPerSec), nil
}

type throttledReader struct {
	reader   io.Reader
	limiter  *rate.Limiter
	ctx      context.Context
	client   *Client
	isUpload bool
}

func (t *throttledReader) Read(p []byte) (int, error) {
	if t.limiter != nil {
		err := t.limiter.WaitN(t.ctx, len(p))
		if err != nil {
			return 0, err
		}
	}

	n, err := t.reader.Read(p)
	if err != nil {
		return n, err
	}

	if t.isUpload {
		atomic.AddInt64(&t.client.stats.BytesSent, int64(n))
		atomic.AddInt64(&t.client.stats.PacketsSent, 1)
	} else {
		atomic.AddInt64(&t.client.stats.BytesReceived, int64(n))
		atomic.AddInt64(&t.client.stats.PacketsReceived, 1)
	}

	return n, nil
}

func (c *Client) DoRequest(ctx context.Context, url string) error {
	atomic.AddInt64(&c.stats.Requests, 1)
	start := time.Now()

	ip := c.IPPool.GetNextIP()
	if !c.Quiet && c.Debug {
		log.Printf("Using IP: %s for %s", ip, url)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		atomic.AddInt64(&c.stats.Errors, 1)
		return fmt.Errorf("failed to create request: %w", err)
	}

	transport := c.HTTPClient.Transport.(*http.Transport).Clone()
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		atomic.AddInt64(&c.stats.PacketsSent, 1)
		atomic.AddInt64(&c.stats.PacketsReceived, 1)

		d := &net.Dialer{
			LocalAddr: &net.TCPAddr{IP: net.ParseIP(ip)},
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			Control: func(network, address string, conn syscall.RawConn) error {
				if !c.RST {
					return nil
				}
				return conn.Control(func(fd uintptr) {
					unix.SetsockoptLinger(int(fd), unix.SOL_SOCKET, unix.SO_LINGER, &unix.Linger{
						Onoff:  1,
						Linger: 0,
					})
				})
			},
		}
		return d.DialContext(ctx, network, addr)
	}

	client := &http.Client{
		Transport: transport,
	}

	reqBytes, _ := httputil.DumpRequestOut(req, false)
	headerSize := len(reqBytes)
	
	if c.Bandwidth.UploadLimit != nil {
		err := c.Bandwidth.UploadLimit.WaitN(ctx, headerSize)
		if err != nil {
			atomic.AddInt64(&c.stats.Timeouts, 1)
			return fmt.Errorf("upload bandwidth limit timeout: %w", err)
		}
	}
	
	atomic.AddInt64(&c.stats.BytesSent, int64(headerSize))
	atomic.AddInt64(&c.stats.PacketsSent, 1)

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			atomic.AddInt64(&c.stats.Timeouts, 1)
		} else {
			atomic.AddInt64(&c.stats.Errors, 1)
		}
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respHeaderBytes, _ := httputil.DumpResponse(resp, false)
	headerSize = len(respHeaderBytes)
	
	if c.Bandwidth.DownloadLimit != nil {
		err := c.Bandwidth.DownloadLimit.WaitN(ctx, headerSize)
		if err != nil {
			atomic.AddInt64(&c.stats.Timeouts, 1)
			return fmt.Errorf("download bandwidth limit timeout: %w", err)
		}
	}
	
	atomic.AddInt64(&c.stats.BytesReceived, int64(headerSize))
	atomic.AddInt64(&c.stats.PacketsReceived, 1)

	bodyReader := &throttledReader{
		reader:   resp.Body,
		limiter:  c.Bandwidth.DownloadLimit,
		ctx:      ctx,
		client:   c,
		isUpload: false,
	}

	_, err = io.Copy(io.Discard, bodyReader)
	if err != nil {
		atomic.AddInt64(&c.stats.Errors, 1)
		return fmt.Errorf("failed to read response: %w", err)
	}

	if !c.Quiet && c.Debug {
		log.Printf("Completed %s in %.2fs (Sent: %d bytes, Recv: %d bytes)",
			url, time.Since(start).Seconds(),
			atomic.LoadInt64(&c.stats.BytesSent),
			atomic.LoadInt64(&c.stats.BytesReceived))
	}

	return nil
}

func (c *Client) GetStats() Stats {
	return Stats{
		Requests:        atomic.LoadInt64(&c.stats.Requests),
		BytesSent:       atomic.LoadInt64(&c.stats.BytesSent),
		BytesReceived:   atomic.LoadInt64(&c.stats.BytesReceived),
		PacketsSent:     atomic.LoadInt64(&c.stats.PacketsSent),
		PacketsReceived: atomic.LoadInt64(&c.stats.PacketsReceived),
		Errors:          atomic.LoadInt64(&c.stats.Errors),
		Timeouts:        atomic.LoadInt64(&c.stats.Timeouts),
	}
}
