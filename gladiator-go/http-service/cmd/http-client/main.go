package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"http-service/internal/client"
	"http-service/internal/common"
)

func main() {
	url := flag.String("url", "", "Target URL to request (required)")
	ipv4Pool := flag.String("ipv4_pool", "", "IPv4 address pool in CIDR notation")
	clients := flag.Int("clients", 1, "Number of concurrent clients")
	timeout := flag.Int("timeout", 0, "Timeout in seconds (0=no timeout)")
	debug := flag.Bool("debug", false, "Enable debug logging")
	quiet := flag.Bool("quiet", false, "Disable all non-error output")
	rst := flag.Bool("rst", false, "Send TCP RST instead of FIN when closing connections")
	uploadBW := flag.String("upload", "", "Upload bandwidth limit (e.g., 1mbps, 100kbps)")
	downloadBW := flag.String("download", "", "Download bandwidth limit (e.g., 10mbps, 1gbps)")
	flag.Parse()

	if *url == "" {
		if !*quiet {
			log.Fatal("Error: --url parameter is required")
		}
		os.Exit(1)
	}

	c, err := client.New(*ipv4Pool, *debug, *quiet, *rst, *uploadBW, *downloadBW)
	if err != nil {
		if !*quiet {
			log.Fatalf("Client initialization failed: %v", err)
		}
		os.Exit(1)
	}

	ctx := context.Background()
	if *timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(*timeout)*time.Second)
		defer cancel()
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(*clients)

	startTime := time.Now()

	for i := 0; i < *clients; i++ {
		go func(id int) {
			defer wg.Done()
			if !*quiet && *debug {
				log.Printf("Client %d started", id)
			}
			if err := c.DoRequest(ctx, *url); err != nil {
				if !*quiet {
					log.Printf("Client %d error: %v", id, err)
				}
			}
		}(i + 1)
	}

	// Goroutine para imprimir las estadísticas en tiempo real
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if *debug {
					stats := c.GetStats()
					log.Printf("Sent: %d bytes in %d packets. Received: %d bytes in %d packets. Errors: %d. Duration: %v",
					    stats.BytesSent, stats.PacketsSent, stats.BytesReceived, stats.PacketsReceived, stats.Errors, time.Since(startTime))
			
				}
			}
		}
	}()

	wg.Wait()
	duration := time.Since(startTime)

	if !*quiet {
		stats := c.GetStats()
		log.Printf("\n=== Transfer Summary ===")
		log.Printf("Requests:      %d", stats.Requests)
		log.Printf("Duration:      %.2f seconds", duration.Seconds())
		log.Printf("Bytes Sent:    %d", stats.BytesSent)
		log.Printf("Bytes Recv:    %d", stats.BytesReceived)
		log.Printf("Packets Sent:  %d", stats.PacketsSent)
		log.Printf("Packets Recv:  %d", stats.PacketsReceived)
		
		if duration.Seconds() > 0 {
			uploadRate, dim := common.FormatRate(float64(stats.BytesSent)*8 / duration.Seconds())
			downloadRate, dim := common.FormatRate(float64(stats.BytesReceived)*8 / duration.Seconds())
			log.Printf("Upload Rate:   %.2f %s", uploadRate, dim)
			log.Printf("Download Rate: %.2f %s", downloadRate, dim)
		}
		
		log.Printf("Errors:        %d", stats.Errors)
		log.Printf("Timeouts:      %d", stats.Timeouts)
		log.Printf("RST Enabled:   %v", *rst)
		
		if *uploadBW != "" {
			log.Printf("Upload Limit:  %s", *uploadBW)
		}
		if *downloadBW != "" {
			log.Printf("Download Limit: %s", *downloadBW)
		}
	}
}
