package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"http-service/internal/iputils"

	"github.com/quic-go/quic-go"
)

// ===== Command Line Flags =====
var (
	flagProxy      = flag.String("proxy", "", "MASQUE proxy host:port (required)")
	flagSNI        = flag.String("sni", "", "Override TLS SNI/ServerName for proxy certificate verification")
	flagAllowInsec = flag.Bool("allow-insecure", true, "Proceed even if certificate verification fails (WARNING)")
	flagProtocol   = flag.String("protocol", "udp", "Protocol to use: udp or tcp")
	flagIPv4Pool   = flag.String("ipv4-pool", "", "CIDR notation for IPv4 pool (e.g., 192.168.1.0/24)")
	flagMask       = flag.String("mask", "", "Subnet mask (alternative to CIDR)")
	flagDebug      = flag.Bool("debug", false, "Enable debug output")
	flagMerge      = flag.Bool("merge", true, "Merge duplicate certificates and group source IPs")
	flagRampup     = flag.Int("rampup", 0, "Number of tasks per second (0 for unlimited)")
	flagTLSVersion = flag.String("tls-version", "auto", "TLS version to use: auto, 1.0, 1.1, 1.2, or 1.3")
)

// ===== JSON Certificate Structures =====
type CertificateInfo struct {
	Subject struct {
		CommonName string   `json:"commonName,omitempty"`
		DNSSANs    []string `json:"dnsSANs,omitempty"`
		IPSANs     []string `json:"ipSANs,omitempty"`
	} `json:"subject"`
	Issuer struct {
		CommonName string `json:"commonName,omitempty"`
	} `json:"issuer"`
	SerialNumber   string `json:"serialNumber"`
	NotBefore      string `json:"notBefore"`
	NotAfter       string `json:"notAfter"`
	AuthorityKeyID string `json:"authorityKeyId,omitempty"`
	SubjectKeyID   string `json:"subjectKeyId,omitempty"`
	Raw            string `json:"raw,omitempty"`
	IsRoot         bool   `json:"isRoot"`
}

type CertChainInfo struct {
	Who           string           `json:"who"`
	Type          string           `json:"type"`
	Certificates  []CertificateInfo `json:"certificates"`
	IsComplete    bool             `json:"isComplete"`
	MissingRoot   bool             `json:"missingRoot"`
	ServerSNI     string           `json:"serverSNI,omitempty"`
	ChainLength   int              `json:"chainLength"`
}

type CertificateResult struct {
	SourceIP    string           `json:"sourceIP,omitempty"`
	SourceIPs   []string         `json:"sourceIPs,omitempty"`
	Protocol    string           `json:"protocol"`
	Presented   *CertChainInfo   `json:"presented,omitempty"`
	Verified    []CertChainInfo  `json:"verified,omitempty"`
	Error       string           `json:"error,omitempty"`
}

type MergedCertificateResult struct {
	SourceIPs []string         `json:"sourceIPs"`
	Protocol  string           `json:"protocol"`
	Presented *CertChainInfo   `json:"presented,omitempty"`
	Verified  []CertChainInfo  `json:"verified,omitempty"`
	Error     string           `json:"error,omitempty"`
}

// ==========================================================
// ===== JSON Certificate Helpers =====
// ==========================================================

// certToJSON converts an x509 certificate to the JSON-friendly CertificateInfo structure
func certToJSON(cert *x509.Certificate) CertificateInfo {
	var info CertificateInfo
	info.Subject.CommonName = cert.Subject.CommonName
	info.Subject.DNSSANs = cert.DNSNames
	for _, ip := range cert.IPAddresses {
		info.Subject.IPSANs = append(info.Subject.IPSANs, ip.String())
	}
	info.Issuer.CommonName = cert.Issuer.CommonName
	info.SerialNumber = cert.SerialNumber.Text(16)
	info.NotBefore = cert.NotBefore.Format(time.RFC3339)
	info.NotAfter = cert.NotAfter.Format(time.RFC3339)
	if len(cert.AuthorityKeyId) > 0 {
		info.AuthorityKeyID = strings.ToUpper(hex.EncodeToString(cert.AuthorityKeyId))
	}
	if len(cert.SubjectKeyId) > 0 {
		info.SubjectKeyID = strings.ToUpper(hex.EncodeToString(cert.SubjectKeyId))
	}
	info.Raw = strings.TrimSpace(string(pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})))
	
	// Determine if this is a root certificate (self-signed and CA)
	info.IsRoot = cert.IsCA && bytes.Equal(cert.RawSubject, cert.RawIssuer)
	
	return info
}

// getCertChainInfo converts a chain of x509 certificates to CertChainInfo with completeness detection
func getCertChainInfo(certs []*x509.Certificate, who, chainType string) CertChainInfo {
	chainInfo := CertChainInfo{
		Who:          who,
		Type:         chainType,
		ChainLength:  len(certs),
		IsComplete:   false,
		MissingRoot:  true,
	}
	
	if len(certs) == 0 {
		return chainInfo
	}
	
	// Convert all certificates in the chain to JSON format
	for _, cert := range certs {
		chainInfo.Certificates = append(chainInfo.Certificates, certToJSON(cert))
	}
	
	// Store server SNI (from leaf certificate's CommonName or first DNS SAN)
	if len(certs) > 0 {
		leafCert := certs[0]
		chainInfo.ServerSNI = leafCert.Subject.CommonName
		
		// If CommonName is empty but there are DNS SANs, use the first one
		if chainInfo.ServerSNI == "" && len(leafCert.DNSNames) > 0 {
			chainInfo.ServerSNI = leafCert.DNSNames[0]
		}
	}
	
	// Determine if chain ends with a root certificate
	if len(certs) > 0 {
		lastCert := certs[len(certs)-1]
		chainInfo.IsComplete = lastCert.IsCA && 
			bytes.Equal(lastCert.RawSubject, lastCert.RawIssuer)
		chainInfo.MissingRoot = !chainInfo.IsComplete
	}
	
	return chainInfo
}

// debugPrint logs debug messages if debug flag is enabled
func debugPrint(format string, args ...interface{}) {
	if *flagDebug {
		log.Printf("[DEBUG] "+format, args...)
	}
}

// ==========================================================
// ===== Rate Limiter =====
// ==========================================================

type RateLimiter struct {
	tokens    chan struct{}
	closeChan chan struct{}
}

// NewRateLimiter creates a rate limiter that controls task execution speed
func NewRateLimiter(tasksPerSecond int) *RateLimiter {
	rl := &RateLimiter{
		tokens:    make(chan struct{}, tasksPerSecond),
		closeChan: make(chan struct{}),
	}

	if tasksPerSecond > 0 {
		// Pre-fill the token bucket
		for i := 0; i < tasksPerSecond; i++ {
			rl.tokens <- struct{}{}
		}

		// Start token replenishment goroutine
		go func() {
			ticker := time.NewTicker(time.Second / time.Duration(tasksPerSecond))
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					select {
					case rl.tokens <- struct{}{}:
						// Token added successfully
					default:
						// Token bucket is full, skip
					}
				case <-rl.closeChan:
					return
				}
			}
		}()
	}

	return rl
}

// Acquire waits for a token if rate limiting is enabled
func (rl *RateLimiter) Acquire() {
	if *flagRampup > 0 {
		<-rl.tokens
	}
}

// Close stops the rate limiter's token replenishment goroutine
func (rl *RateLimiter) Close() {
	close(rl.closeChan)
}

// ==========================================================
// ===== Certificate Hashing and Comparison =====
// ==========================================================

// hashCertChain creates a SHA256 hash of certificate chain contents for comparison
func hashCertChain(presented *CertChainInfo, verified []CertChainInfo) string {
	h := sha256.New()
	
	if presented != nil {
		for _, cert := range presented.Certificates {
			h.Write([]byte(cert.Raw))
		}
	}
	
	for _, chain := range verified {
		for _, cert := range chain.Certificates {
			h.Write([]byte(cert.Raw))
		}
	}
	
	return hex.EncodeToString(h.Sum(nil))
}

// mergeCertificateResults groups results with identical certificate chains and errors
func mergeCertificateResults(results []CertificateResult) []MergedCertificateResult {
	mergedMap := make(map[string]*MergedCertificateResult)
	
	for _, result := range results {
		// Create a unique key based on certificate content, protocol, and error
		hash := hashCertChain(result.Presented, result.Verified)
		key := fmt.Sprintf("%s|%s|%s", hash, result.Protocol, result.Error)
		
		if existing, exists := mergedMap[key]; exists {
			// Merge source IPs from duplicate results
			if result.SourceIP != "" {
				existing.SourceIPs = append(existing.SourceIPs, result.SourceIP)
			}
		} else {
			// Create new merged result
			merged := &MergedCertificateResult{
				SourceIPs: []string{},
				Protocol:  result.Protocol,
				Presented: result.Presented,
				Verified:  result.Verified,
				Error:     result.Error,
			}
			if result.SourceIP != "" {
				merged.SourceIPs = append(merged.SourceIPs, result.SourceIP)
			}
			mergedMap[key] = merged
		}
	}
	
	// Convert map to slice
	var mergedResults []MergedCertificateResult
	for _, merged := range mergedMap {
		mergedResults = append(mergedResults, *merged)
	}
	
	return mergedResults
}

// ==========================================================
// ===== TLS Config with Certificate Collection =====
// ==========================================================

// certCollector collects certificates during TLS handshake
type certCollector struct {
	SourceIP       string
	Protocol       string
	PresentedCerts []*x509.Certificate
	VerifiedChains [][]*x509.Certificate
}

// buildProxyTLS creates a TLS configuration with certificate collection hooks
func (c *certCollector) buildProxyTLS(proxySNI string, allowInsecure bool) *tls.Config {
	roots := systemRoots()
	
	// Parse TLS version from flag
	var minVersion uint16
	switch strings.ToLower(*flagTLSVersion) {
	case "auto":
		// Use zero value (0) to let Go auto-negotiate
		minVersion = 0
	case "1.0":
		minVersion = tls.VersionTLS10
	case "1.1":
		minVersion = tls.VersionTLS11
	case "1.2":
		minVersion = tls.VersionTLS12
	case "1.3":
		minVersion = tls.VersionTLS13
	default:
		log.Printf("WARNING: Invalid TLS version %s, using auto-negotiation", *flagTLSVersion)
		minVersion = 0
	}
	
	tlsConf := &tls.Config{
		ServerName:         proxySNI,
		NextProtos:         []string{"h3"},
		InsecureSkipVerify: allowInsecure,
	}
	
	// Only set MinVersion if explicitly requested (not "auto")
	if minVersion > 0 {
		tlsConf.MinVersion = minVersion
	}

	// Hook into certificate verification to capture certificates
	tlsConf.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		// Parse raw certificates to x509.Certificate slice
		var presented []*x509.Certificate
		for _, rb := range rawCerts {
			cert, err := x509.ParseCertificate(rb)
			if err == nil {
				presented = append(presented, cert)
			}
		}

		c.PresentedCerts = presented
		c.VerifiedChains = verifiedChains

		// Optional manual verification for logging purposes
		if len(presented) > 0 {
			leaf := presented[0]
			rest := presented[1:]
			opts := x509.VerifyOptions{DNSName: proxySNI, Roots: roots, Intermediates: x509.NewCertPool()}
			for _, ic := range rest {
				opts.Intermediates.AddCert(ic)
			}
			if _, err := leaf.Verify(opts); err != nil {
				debugPrint("WARNING: Proxy certificate verification failed for SNI '%s': %v", proxySNI, err)
			}
		}

		return nil
	}
	return tlsConf
}

// toCertificateResult converts collected certificates to the result structure
func (c *certCollector) toCertificateResult() CertificateResult {
	result := CertificateResult{
		SourceIP: c.SourceIP,
		Protocol: c.Protocol,
	}

	if len(c.PresentedCerts) > 0 {
		presentedInfo := getCertChainInfo(c.PresentedCerts, "Proxy", "presented")
		result.Presented = &presentedInfo
	}

	if len(c.VerifiedChains) > 0 {
		for i, chain := range c.VerifiedChains {
			verifiedInfo := getCertChainInfo(chain, fmt.Sprintf("Proxy Chain %d", i), "verified")
			result.Verified = append(result.Verified, verifiedInfo)
		}
	}

	return result
}

// systemRoots retrieves the system certificate pool
func systemRoots() *x509.CertPool {
	roots, err := x509.SystemCertPool()
	if err != nil || roots == nil {
		roots = x509.NewCertPool()
	}
	return roots
}

// ==========================================================
// ===== Certificate Retrieval Logic =====
// ==========================================================

// retrieveCertificatesQUIC establishes a QUIC connection to collect certificates
func retrieveCertificatesQUIC(proxyAddr, proxySNI, sourceIP string) (*certCollector, error) {
	collector := &certCollector{
		SourceIP: sourceIP,
		Protocol: "quic",
	}

	// Configure QUIC with datagram support
	qconf := &quic.Config{EnableDatagrams: true}
	tlsConf := collector.buildProxyTLS(proxySNI, *flagAllowInsec)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	debugPrint("Dialing QUIC to %s from %s", proxyAddr, sourceIP)
	proxyConn, err := quic.DialAddrEarly(ctx, proxyAddr, tlsConf, qconf)
	if err != nil {
		return nil, fmt.Errorf("QUIC dial failed: %v", err)
	}
	defer proxyConn.CloseWithError(0, "")

	return collector, nil
}

// retrieveCertificatesTCP establishes a TCP/TLS connection to collect certificates
func retrieveCertificatesTCP(proxyAddr, proxySNI, sourceIP string) (*certCollector, error) {
	collector := &certCollector{
		SourceIP: sourceIP,
		Protocol: "tcp",
	}

	tlsConf := collector.buildProxyTLS(proxySNI, *flagAllowInsec)
	tlsConf.NextProtos = []string{"http/1.1"} // Appropriate ALPN for TCP/TLS

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dialer net.Dialer
	if sourceIP != "" {
		debugPrint("Setting local address to %s for TCP connection", sourceIP)
		localAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(sourceIP, "0"))
		if err != nil {
			return nil, fmt.Errorf("failed to resolve local address: %v", err)
		}
		dialer.LocalAddr = localAddr
	}

	debugPrint("Dialing TCP to %s from %s", proxyAddr, sourceIP)
	conn, err := dialer.DialContext(ctx, "tcp", proxyAddr)
	if err != nil {
		return nil, fmt.Errorf("TCP dial failed: %v", err)
	}
	defer conn.Close()

	tlsConn := tls.Client(conn, tlsConf)
	
	// Perform TLS handshake to trigger certificate collection
	err = tlsConn.HandshakeContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("TLS handshake failed: %v", err)
	}

	return collector, nil
}

// ==========================================================
// ===== Worker Pool and Task Management =====
// ==========================================================

type retrievalTask struct {
	SourceIP string
	Protocol string
}

type retrievalResult struct {
	Result CertificateResult
	Error  error
}

// worker processes certificate retrieval tasks
func worker(id int, tasks <-chan retrievalTask, results chan<- retrievalResult, wg *sync.WaitGroup, 
	           proxyAddr, proxySNI string, rateLimiter *RateLimiter) {
	defer wg.Done()

	for task := range tasks {
		rateLimiter.Acquire()
		
		debugPrint("Worker %d processing %s from %s", id, task.Protocol, task.SourceIP)
		
		var collector *certCollector
		var err error

		if task.Protocol == "quic" {
			collector, err = retrieveCertificatesQUIC(proxyAddr, proxySNI, task.SourceIP)
		} else {
			collector, err = retrieveCertificatesTCP(proxyAddr, proxySNI, task.SourceIP)
		}

		if err != nil {
			results <- retrievalResult{
				Result: CertificateResult{
					SourceIP: task.SourceIP,
					Protocol: task.Protocol,
					Error:    err.Error(),
				},
				Error: err,
			}
		} else {
			results <- retrievalResult{
				Result: collector.toCertificateResult(),
				Error:  nil,
			}
		}
	}
}

// runCertRetrieval is the main orchestration function for certificate retrieval
func runCertRetrieval() {
	if *flagProxy == "" {
		fmt.Fprintln(os.Stderr, "-proxy is required (host:port)")
		os.Exit(2)
	}

	// Validate protocol parameter
	if *flagProtocol != "udp" && *flagProtocol != "tcp" && *flagProtocol != "both" {
		fmt.Fprintf(os.Stderr, "invalid protocol: %s (must be udp, tcp, or both)\n", *flagProtocol)
		os.Exit(2)
	}

	// Validate rampup parameter
	if *flagRampup < 0 {
		fmt.Fprintf(os.Stderr, "invalid rampup: %d (must be >= 0)\n", *flagRampup)
		os.Exit(2)
	}

	// Validate TLS version parameter
	validTLSVersions := map[string]bool{
		"auto": true, "1.0": true, "1.1": true, "1.2": true, "1.3": true,
	}
	if !validTLSVersions[strings.ToLower(*flagTLSVersion)] {
		fmt.Fprintf(os.Stderr, "invalid TLS version: %s (must be auto, 1.0, 1.1, 1.2, or 1.3)\n", *flagTLSVersion)
		os.Exit(2)
	}

	// Parse proxy address and SNI
	proxyHost, _, err := net.SplitHostPort(*flagProxy)
	if err != nil {
		log.Fatalf("invalid -proxy: %v", err)
	}
	proxyAddr := *flagProxy
	proxySNI := proxyHost
	if *flagSNI != "" {
		proxySNI = *flagSNI
	}

	// Generate source IP addresses to use
	var sourceIPs []string
	if *flagIPv4Pool != "" {
		cidr := *flagIPv4Pool
		if *flagMask != "" {
			// Convert IP + mask to CIDR notation
			ip := net.ParseIP(*flagIPv4Pool)
			if ip == nil {
				log.Fatalf("invalid IPv4 pool address: %s", *flagIPv4Pool)
			}
			mask := net.IPMask(net.ParseIP(*flagMask).To4())
			if mask == nil {
				log.Fatalf("invalid mask: %s", *flagMask)
			}
			cidr = fmt.Sprintf("%s/%d", ip.String(), onesCount(mask))
		}
		sourceIPs, err = iputils.GenerateIPsFromCIDR(cidr)
		if err != nil {
			log.Fatalf("failed to generate IPs from CIDR: %v", err)
		}
	} else {
		// Use default IP (empty string for system default)
		sourceIPs = []string{""}
	}

	debugPrint("=== MASQUE Certificate Retrieval Starting ===")
	debugPrint("Proxy: %s, SNI: %s, Protocol: %s", proxyAddr, proxySNI, *flagProtocol)
	debugPrint("TLS Version: %s", *flagTLSVersion)
	debugPrint("Source IPs: %v", sourceIPs)
	debugPrint("Merge duplicates: %v", *flagMerge)
	debugPrint("Rampup: %d tasks/second", *flagRampup)

	// Create retrieval tasks for each source IP and protocol
	var tasks []retrievalTask
	for _, sourceIP := range sourceIPs {
		if *flagProtocol == "udp" || *flagProtocol == "both" {
			tasks = append(tasks, retrievalTask{SourceIP: sourceIP, Protocol: "quic"})
		}
		if *flagProtocol == "tcp" || *flagProtocol == "both" {
			tasks = append(tasks, retrievalTask{SourceIP: sourceIP, Protocol: "tcp"})
		}
	}

	debugPrint("Total tasks to execute: %d", len(tasks))

	// Setup rate limiter if rampup is enabled
	rateLimiter := NewRateLimiter(*flagRampup)
	defer rateLimiter.Close()

	// Create worker pool with appropriate size
	numWorkers := len(tasks)
	if numWorkers > 100 {
		numWorkers = 100 // Limit maximum workers to prevent resource exhaustion
	}
	if *flagRampup > 0 && numWorkers > *flagRampup {
		numWorkers = *flagRampup
	}

	taskChan := make(chan retrievalTask, len(tasks))
	resultChan := make(chan retrievalResult, len(tasks))
	var wg sync.WaitGroup

	// Start worker goroutines
	debugPrint("Starting %d workers", numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, taskChan, resultChan, &wg, proxyAddr, proxySNI, rateLimiter)
	}

	// Send all tasks to workers
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Wait for all workers to complete and close result channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect all results
	var allResults []CertificateResult
	for result := range resultChan {
		allResults = append(allResults, result.Result)
	}

	debugPrint("Completed %d out of %d tasks", len(allResults), len(tasks))

	// Prepare output based on merge flag
	var output interface{}
	if *flagMerge {
		mergedResults := mergeCertificateResults(allResults)
		output = mergedResults
		debugPrint("Merged %d results into %d unique certificate sets", len(allResults), len(mergedResults))
	} else {
		output = allResults
	}

	// Marshal results to JSON
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal results to JSON: %v", err)
	}

	// Output results
	if *flagDebug {
		fmt.Printf("\n=== Certificate Retrieval Results ===\n%s\n", string(jsonData))
	} else {
		fmt.Println(string(jsonData))
	}
}

// onesCount calculates the number of 1 bits in an IP mask
func onesCount(mask net.IPMask) int {
	ones, _ := mask.Size()
	return ones
}

// ===== Main Function =====
func main() {
	flag.Parse()
	runCertRetrieval()
}
