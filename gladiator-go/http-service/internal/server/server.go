// Package that implements a simple HTTP server for testing purposes
// key features: rate limiting, brocken TCP support, like ignoring packets with FIN flag.
package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
    "golang.org/x/sys/unix"
	"crypto/rand"
	"encoding/hex"
	xrate "golang.org/x/time/rate"
)

// Task represents a single HTTP request
// Fields:
//   ID: ID of the task
//   URL: URL of the task
//   SourceIP: Source IP of the task
//   DestIP: Destination IP of the task
//   SourcePort: Source port of the task
//   DestPort: Destination port of the task
//   StartTime: Timestamp of the start of the task
//   BytesSent: Number of bytes sent
//   BytesReceived: Number of bytes received 
//   PacketsSent: Number of packets sent
//   PacketsReceived: Number of packets received
//   RSTTime: Time to trigger TCP RST
//   conn: TCP connection
//   mu: Mutex for synchronizing access to the task
type Task struct {
	ID              string
	URL             string
	SourceIP        string
	DestIP          string
	SourcePort      string
	DestPort        string
	StartTime       time.Time
	BytesSent       int64
	BytesReceived   int64
	PacketsSent     int64
	PacketsReceived int64
  IgnoreFIN       bool
	Bandwidth       *xrate.Limiter
	RSTTime         time.Duration
	conn            *net.TCPConn
	mu              sync.Mutex
}

// Server represents the HTTP server
// Fields:
//   Port: Port number of the server (80 by default)
//   Debug: Enable debug mode (false by default)
//   tasks: Map of tasks
//   taskLock: Mutex for synchronizing access to the tasks
type Server struct {
	Port     int
	Debug    bool
	tasks    map[string]*Task
	taskLock sync.RWMutex
}

// NewServer creates a new Server
func NewServer(port int, debug bool) *Server {
	return &Server{
		Port:  port,
		Debug: debug,
		tasks: make(map[string]*Task),
	}
}

// logf logs a message if Debug is true
func (s *Server) logf(format string, args ...interface{}) {
	if s.Debug {
		log.Printf("[DEBUG] "+format, args...)
	}
}

// Start starts the server and implements the following endpoints:
//   /seconds/:duration - keeps the connection open and sends progress information to the client for the specified duration
//   /size/:size?unit - sends random data to the client until the desired size is reached
//   /tasks - returns a list of tasks
//   /tasks/:id - returns details of the specified task
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/seconds/", s.secondsHandler)
	mux.HandleFunc("/size/", s.sizeHandler)
	mux.HandleFunc("/tasks", s.listTasksHandler)
	mux.HandleFunc("/tasks/", s.taskDetailHandler)

	s.logf("Starting server on :%d", s.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), mux)
}

// function that handles the /seconds/:duration endpoint
//
func (s *Server) secondsHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/seconds/"), "?")
	secondsStr := pathParts[0]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil || seconds < 1 {
		http.Error(w, "Invalid duration. Use /seconds/INTEGER", http.StatusBadRequest)
		s.logf("secondsHandler error: invalid duration %s", secondsStr)
		return
	}

	task := &Task{
		ID:        generateTaskID(),
		URL:       r.URL.String(),
		StartTime: time.Now(),
		DestIP:    getLocalIP(),
		DestPort:  strconv.Itoa(s.Port),
	}

	if host, port, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		task.SourceIP = host
		task.SourcePort = port
	}
	ignoreFIN := r.URL.Query().Get("ignore_fin") == "true"
	task.IgnoreFIN = ignoreFIN

	s.taskLock.Lock()
	s.tasks[task.ID] = task
	s.taskLock.Unlock()

	defer func() {
		s.taskLock.Lock()
		delete(s.tasks, task.ID)
		s.taskLock.Unlock()
	}()

	if bw := r.URL.Query().Get("bw"); bw != "" {
		rate, err := parseBandwidth(bw)
		if err != nil {
			http.Error(w, "Invalid bandwidth format", http.StatusBadRequest)
			s.logf("secondsHandler error: invalid bandwidth %s", bw)
			return
		}
		task.Bandwidth = xrate.NewLimiter(xrate.Limit(rate), rate)
		s.logf("Bandwidth limiting enabled: %d bytes/sec", rate)
	} else {
		task.Bandwidth = xrate.NewLimiter(xrate.Inf, 0)
	}

	if rst := r.URL.Query().Get("RST"); rst != "" {
		rstSec, err := strconv.Atoi(rst)
		if err != nil || rstSec < 0 {
			http.Error(w, "Invalid RST time", http.StatusBadRequest)
			s.logf("secondsHandler error: invalid RST time %s", rst)
			return
		}
		task.RSTTime = time.Duration(rstSec) * time.Second
		s.logf("TCP RST scheduled after %d seconds", rstSec)
	}

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		s.logf("secondsHandler error: connection hijacking not supported")
		return
	}

	conn, bufRW, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "Hijacking failed", http.StatusInternalServerError)
		s.logf("secondsHandler error: hijack failed: %v", err)
		return
	}
	defer conn.Close()

	tcpConn := conn.(*net.TCPConn)
	task.conn = tcpConn

	if task.IgnoreFIN {
	// Set up connection in order to ignore FIN flag
	rawConn, err := tcpConn.SyscallConn()
	if err != nil {
		s.logf("Error obtaining raw connection: %v", err)
		return
	}
	rawConn.Control(func(fd uintptr) {
	// Disable the standard behaviour.
		unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_LINGER, 0)
	})
	s.logf("Seting up ignore FIN on task: %s", task.ID)
	}

	if task.RSTTime > 0 {
		time.AfterFunc(task.RSTTime, func() {
			task.mu.Lock()
			defer task.mu.Unlock()
			s.logf("Triggering TCP RST for task %s", task.ID)
			tcpConn.SetLinger(0)
			tcpConn.Close()
		})
	}

	bufRW.WriteString("HTTP/1.1 200 OK\r\n")
	bufRW.WriteString("Content-Type: text/plain\r\n")
	bufRW.WriteString("Transfer-Encoding: chunked\r\n")
	bufRW.WriteString("\r\n")
	bufRW.Flush()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < seconds; i++ {
		select {
		case <-r.Context().Done():
			s.logf("Client disconnected after %d seconds", i)
			time.Sleep(60 * time.Second)
			return
		case <-ticker.C:
			chunk := fmt.Sprintf("Second %d/%d\n%s", i+1, seconds)

			if task.Bandwidth != nil {
				s.logf("Waiting %d bytes for task %s", len(chunk), task.ID)
				task.Bandwidth.WaitN(r.Context(), len(chunk))
			}

			task.mu.Lock()
			task.BytesSent += int64(len(chunk))
			task.PacketsSent++
			task.mu.Unlock()

			bufRW.WriteString(fmt.Sprintf("%x\r\n%s\r\n", len(chunk), chunk))
			bufRW.Flush()
			s.logf("Sent second %d/%d to task %s", i+1, seconds, task.ID)
		}
	}

	bufRW.WriteString("0\r\n\r\n")
	bufRW.Flush()
}

func randomString(size int64) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// function that handles the /size/:size?unit endpoint
func (s *Server) sizeHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/size/"), "?")
	sizeStr := pathParts[0]

	if len(sizeStr) < 3 {
		http.Error(w, "Size must include value and unit (e.g. 100KB)", http.StatusBadRequest)
		s.logf("sizeHandler error: invalid size format")
		return
	}

	unit := strings.ToUpper(sizeStr[len(sizeStr)-2:])
	numStr := sizeStr[:len(sizeStr)-2]

	size, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil || size < 1 {
		http.Error(w, "Invalid size number", http.StatusBadRequest)
		s.logf("sizeHandler error: invalid size value %s", numStr)
		return
	}

	var bytes int64
	switch unit {
	case "KB":
		bytes = size * 1024
	case "MB":
		bytes = size * 1024 * 1024
	case "GB":
		bytes = size * 1024 * 1024 * 1024
	default:
		http.Error(w, "Invalid unit. Use KB, MB or GB", http.StatusBadRequest)
		s.logf("sizeHandler error: invalid unit %s", unit)
		return
	}

	task := &Task{
		ID:        generateTaskID(),
		URL:       r.URL.String(),
		StartTime: time.Now(),
		DestIP:    getLocalIP(),
		DestPort:  strconv.Itoa(s.Port),
	}

	if host, port, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		task.SourceIP = host
		task.SourcePort = port
	}
	ignoreFIN := r.URL.Query().Get("ignore_fin") == "true"
	task.IgnoreFIN = ignoreFIN

	s.taskLock.Lock()
	s.tasks[task.ID] = task
	s.taskLock.Unlock()

	defer func() {
		s.taskLock.Lock()
		delete(s.tasks, task.ID)
		s.taskLock.Unlock()
	}()

	if bw := r.URL.Query().Get("bw"); bw != "" {
		rate, err := parseBandwidth(bw)
		if err != nil {
			http.Error(w, "Invalid bandwidth format", http.StatusBadRequest)
			s.logf("sizeHandler error: invalid bandwidth %s", bw)
			return
		}
		task.Bandwidth = xrate.NewLimiter(xrate.Limit(rate), rate)
		s.logf("Bandwidth limiting enabled: %d bytes/sec", rate)
	} else {
		task.Bandwidth = xrate.NewLimiter(xrate.Inf, 0)
	}


	if rst := r.URL.Query().Get("RST"); rst != "" {
		rstSec, err := strconv.Atoi(rst)
		if err != nil || rstSec < 0 {
			http.Error(w, "Invalid RST time", http.StatusBadRequest)
			s.logf("sizeHandler error: invalid RST time %s", rst)
			return
		}
		task.RSTTime = time.Duration(rstSec) * time.Second
		s.logf("TCP RST scheduled after %d seconds", rstSec)
	}

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		s.logf("sizeHandler error: hijacking not supported")
		return
	}

	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "Hijacking failed", http.StatusInternalServerError)
		s.logf("sizeHandler error: hijack failed: %v", err)
		return
	}
	defer conn.Close()

	tcpConn := conn.(*net.TCPConn)
	task.conn = tcpConn

  if task.IgnoreFIN {
    // Configurar la conexión para ignorar el FIN (ejemplo en Linux)
    rawConn, err := tcpConn.SyscallConn()
    if err != nil {
        s.logf("Error obtaining raw connection: %v", err)
        return
    }
    rawConn.Control(func(fd uintptr) {
        // Desactivar el comportamiento estándar de cierre
        unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_LINGER, 0)
    })
	s.logf("Setting up ignore FIN on task: %s", task.ID)
  }

	if task.RSTTime > 0 {
		time.AfterFunc(task.RSTTime, func() {
			task.mu.Lock()
			defer task.mu.Unlock()
      s.logf("Triggering TCP RST for task %s (ignore_fin=%v)", task.ID, task.IgnoreFIN)
			tcpConn.SetLinger(0)
			tcpConn.Close()
		})
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: application/octet-stream\r\n"))
	conn.Write([]byte(fmt.Sprintf("Content-Length: %d\r\n", bytes)))
	conn.Write([]byte("\r\n"))

	buf := make([]byte, 32*1024)
	remaining := bytes
	for remaining > 0 {
		chunkSize := min(remaining, int64(len(buf)))

		if task.Bandwidth != nil {
			task.Bandwidth.WaitN(r.Context(), int(chunkSize))
		}

		n, err := io.CopyN(conn, &zeroReader{}, chunkSize)
		if err != nil {
			s.logf("Transfer interrupted at %d/%d bytes: %v", bytes-remaining, bytes, err)
			return
		}

		task.mu.Lock()
		task.BytesSent += n
		task.PacketsSent++
		task.mu.Unlock()

		remaining -= n
		s.logf("Sent %d/%d bytes (%.1f%%)", bytes-remaining, bytes, float64(bytes-remaining)/float64(bytes)*100)
	}
}

func (s *Server) listTasksHandler(w http.ResponseWriter, r *http.Request) {
	s.taskLock.RLock()
	defer s.taskLock.RUnlock()

	tasks := make([]map[string]interface{}, 0, len(s.tasks))
	for id, task := range s.tasks {
		task.mu.Lock()
		tasks = append(tasks, map[string]interface{}{
			"id":          id,
			"url":         task.URL,
			"source_ip":   task.SourceIP,
			"source_port": task.SourcePort,
			"dest_ip":     task.DestIP,
			"dest_port":   task.DestPort,
			"start_time":  task.StartTime.Format(time.RFC3339),
			"duration":    time.Since(task.StartTime).String(),
		})
		task.mu.Unlock()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (s *Server) taskDetailHandler(w http.ResponseWriter, r *http.Request) {
	taskID := strings.TrimPrefix(r.URL.Path, "/tasks/")

	s.taskLock.RLock()
	task, exists := s.tasks[taskID]
	s.taskLock.RUnlock()
	fmt.Println(task)
	fmt.Println("eooo")
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		task.mu.Lock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":                task.ID,
			"url":               task.URL,
			"source_ip":         task.SourceIP,
			"source_port":       task.SourcePort,
			"dest_ip":           task.DestIP,
			"dest_port":         task.DestPort,
			"start_time":        task.StartTime.Format(time.RFC3339),
			"duration":          time.Since(task.StartTime).String(),
			"bytes_sent":        task.BytesSent,
			"bytes_received":    task.BytesReceived,
			"packets_sent":      task.PacketsSent,
			"packets_received":  task.PacketsReceived,
			"bandwidth_limit":   task.Bandwidth.Limit(),
			"rst_scheduled":     task.RSTTime > 0,
			"rst_remaining":     (task.RSTTime - time.Since(task.StartTime)).String(),
		})
		task.mu.Unlock()

	case http.MethodPatch:
		fmt.Fprintf(w, "TODO: PATCH /tasks/%s", taskID)
		var update struct {
			Bandwidth *string `json:"bandwidth"`
			RST       *int    `json:"rst"`
		}

		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		task.mu.Lock()
		defer task.mu.Unlock()

		if update.Bandwidth != nil {
			rate, err := parseBandwidth(*update.Bandwidth)
			if err != nil {
				http.Error(w, "Invalid bandwidth format", http.StatusBadRequest)
				return
			}
			task.Bandwidth = xrate.NewLimiter(xrate.Limit(rate), rate)
			s.logf("Updated bandwidth for task %s to %v", task.ID, *update.Bandwidth)
		}

		if update.RST != nil {
			if *update.RST <= 0 {
				task.RSTTime = 0
				s.logf("Cancelled RST for task %s", task.ID)
			} else {
				newRST := time.Duration(*update.RST) * time.Second
				remaining := newRST - time.Since(task.StartTime)
				if remaining > 0 {
					task.RSTTime = newRST
					go func() {
						time.Sleep(remaining)
						task.mu.Lock()
						defer task.mu.Unlock()
						if task.conn != nil {
							s.logf("Triggering updated TCP RST for task %s", task.ID)
							task.conn.SetLinger(0)
							task.conn.Close()
						}
					}()
					s.logf("Updated RST for task %s to %v", task.ID, newRST)
				}
			}
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func generateTaskID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func parseBandwidth(bw string) (int, error) {
	bw = strings.ToLower(bw)
	var multiplier int

	switch {
	case strings.HasSuffix(bw, "kbps"):
		multiplier = 1024
		bw = strings.TrimSuffix(bw, "kbps")
	case strings.HasSuffix(bw, "mbps"):
		multiplier = 1024 * 1024
		bw = strings.TrimSuffix(bw, "mbps")
	case strings.HasSuffix(bw, "bps"):
		multiplier = 1
		bw = strings.TrimSuffix(bw, "bps")
	case strings.HasSuffix(bw, "gps"):
		multiplier = 1
		bw = strings.TrimSuffix(bw, "gbps")
	default:
		return 0, fmt.Errorf("invalid bandwidth unit")
	}

	rate, err := strconv.Atoi(bw)
	if err != nil {
		return 0, err
	}

	return rate * multiplier, nil
}

type zeroReader struct{}

func (z *zeroReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
