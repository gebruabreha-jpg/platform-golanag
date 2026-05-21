# HTTP Service Test Suite

A sophisticated client-server system for network performance testing, connection behavior analysis, and HTTP protocol experimentation.

## Features

### Client
- Concurrent HTTP clients with customizable connection pools
- Bandwidth throttling (upload/download limits)
- TCP RST (reset) injection for abnormal termination
- IPv4 address rotation from CIDR ranges
- Real-time statistics and transfer metrics

### Server
- Duration-based (`/seconds/N`) and size-based (`/size/XMB`) test endpoints
- Connection hijacking for low-level TCP control
- Scheduled TCP RST injection
- `ignore_fin` mode for half-open connection simulation
- Dynamic bandwidth throttling
- REST API for active connection monitoring

## Architecture
.
├── cmd/
│ ├── http_client/ # CLI client application
│ └── http_server/ # CLI server application
├── internal/
│ ├── client/ # Core client logic
│ ├── iputils/ # IP address utilities
│ └── server/ # Core server logic


## Installation

```bash
# Build both components
go build -o bin/client ./cmd/http_client
go build -o bin/server ./cmd/http_server

# Start server (default port 80)
./bin/server -port 8080 -debug

# Test endpoints:
curl http://localhost:8080/seconds/10       # 10-second stream
curl http://localhost:8080/size/100MB       # 100MB data dump
curl http://localhost:8080/tasks            # List active connections

# Advanced options:
curl "http://localhost:8080/seconds/30?bw=1mbps&RST=5"  # Bandwidth limit + RST after 5s
curl "http://localhost:8080/size/1GB?ignore_fin=true"   # Ignore FIN packets

# Basic load test (10 concurrent clients)
./bin/client -url http://target:8080/seconds/30 -clients 10

# Bandwidth-limited test
./bin/client -url http://target:8080/size/1GB \
  -upload 500kbps -download 2mbps

# IP rotation test
./bin/client -url http://target:8080/seconds/10 \
  -ipv4_pool 192.168.1.0/24 -clients 50

# RST-enabled destructive test
./bin/client -url http://target:8080/seconds/5 -rst

# Server (throttle responses to 100kbps)
./bin/server -port 8080

# Client (measure actual throughput)
./bin/client -url http://localhost:8080/size/10MB -debug


