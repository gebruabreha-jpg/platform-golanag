# Client Wrapper API

Flask-based REST API for managing and monitoring HTTP/3 client test executions with automatic retry capabilities, comprehensive logging, and certificate retrieval for MASQUE proxy testing.

## Overview

The Client Wrapper provides a REST API to:
- Execute HTTP/3 client tests with configurable parameters
- Automatically retry failed tests until duration requirements are met
- Retrieve and parse TLS certificates from target servers
- Track task execution across multiple attempts with detailed logging
- Download, manage, and clean up execution logs
- Aggregate statistics and metrics from test runs

## Architecture

```
Client Network (23.0.0.0/8)          Proxy Network              Server Network (192.0.0.1/24)
┌─────────────────────────┐              ┌──────┐              ┌─────────────────┐
│  Client Wrapper API     │              │      │              │                 │
│  ┌──────────────────┐   │   QUIC/UDP   │ PCG  │   QUIC/UDP   │  Target Server  │
│  │  http3client     │───┼──────────────┤MASQUE├──────────────┤   HTTP/3        │
│  │  cert-retriever  │   │   Tunnel     │Proxy │   Proxied    │   192.0.0.x     │
│  └──────────────────┘   │              │      │              │                 │
│  23.x.x.x               │              └──────┘              └─────────────────┘
└─────────────────────────┘
```

## Features

### Automatic Retry Logic
- Detects `--duration` parameter in test arguments
- If a test fails before completing its duration, automatically restarts it
- Adjusts remaining duration for each retry attempt
- Continues until total execution time meets requested duration
- Tracks all attempts with individual statistics

### Task Management
- Create, monitor, and terminate test tasks
- Track task state across multiple execution attempts
- Filter tasks by date range
- Aggregate statistics across all tasks

### Log Management
- Structured log storage in `/var/log/http3_client/`
- Compressed logs (gzip) for each attempt
- Download logs as tarball archives (.tgz)
- Automatic log cleanup after 7 days (configurable)
- Per-task and global log management
- Log metadata and summaries in JSON format

### Output Processing
- Parses gladiator/http3client logs to extract metrics
- Strips ANSI escape codes for clean output
- Extracts: throughput, packets, connections, errors, payload bytes

### Certificate Retrieval
- Fetch TLS certificates from target servers
- Parse certificate chains with SNI, issuer, validity dates
- Separate endpoint for certificate operations

## API Endpoints

### Tasks

#### Create Task
```http
POST /tasks
Content-Type: application/json

{
  "arguments": ["--duration=300", "--target=192.0.0.5:443", "--pvd-server=10.0.0.2:443"],
  "auto_retry": true
}
```

**Response (202 Accepted):**
```json
{
  "task_id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "initializing",
  "message": "Task created and started",
  "duration_detected": true,
  "auto_retry": true,
  "has_duration": true,
  "duration_seconds": 300,
  "created_at": "2026-02-03T10:00:00.000000"
}
```

#### Get Task Status
```http
GET /tasks/{task_id}
```

**Response (200 OK):**
```json
{
  "task_id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "completed_with_retries",
  "arguments": ["--duration=300", "--target=192.0.0.5:443"],
  "created_at": "2026-02-03T10:00:00.000000",
  "total_attempts": 3,
  "successful_attempts": 2,
  "failed_attempts": 1,
  "total_execution_time": 305.42,
  "requested_duration": 300,
  "achieved_duration_ratio": 1.018,
  "has_duration": true,
  "auto_retry_enabled": true,
  "last_attempt": {
    "attempt_id": "550e8400-e29b-41d4-a716-446655440000_attempt_3",
    "status": "completed",
    "execution_time": 150.5,
    "returncode": 0,
    "start_time": "2026-02-03T10:02:35.000000",
    "end_time": "2026-02-03T10:05:05.500000"
  },
  "attempts_summary": [
    {
      "attempt_id": "550e8400-e29b-41d4-a716-446655440000_attempt_1",
      "status": "failed",
      "execution_time": 45.2,
      "returncode": 1,
      "start_time": "2026-02-03T10:00:00.000000",
      "end_time": "2026-02-03T10:00:45.200000"
    },
    {
      "attempt_id": "550e8400-e29b-41d4-a716-446655440000_attempt_2",
      "status": "completed",
      "execution_time": 109.72,
      "returncode": 0,
      "start_time": "2026-02-03T10:00:47.200000",
      "end_time": "2026-02-03T10:02:36.920000"
    },
    {
      "attempt_id": "550e8400-e29b-41d4-a716-446655440000_attempt_3",
      "status": "completed",
      "execution_time": 150.5,
      "returncode": 0,
      "start_time": "2026-02-03T10:02:35.000000",
      "end_time": "2026-02-03T10:05:05.500000"
    }
  ]
}
```

#### List All Tasks
```http
GET /tasks
GET /tasks?start_date=1706950800&end_date=1706954400
```

#### Get Tasks Summary
```http
GET /tasks/summary
GET /tasks/summary?start_date=1706950800&end_date=1706954400
```

**Response (200 OK):**
```json
{
  "total": 150,
  "running": 5,
  "initializing": 2,
  "completed": 80,
  "failed": 10,
  "task failed": 3,
  "completed_with_retries": 35,
  "partial_completion": 8,
  "failed_all_attempts": 5,
  "failed_before_duration": 2,
  "terminated": 0,
  "timeout": 0,
  "total_attempts": 250,
  "successful_attempts": 200,
  "failed_attempts": 50,
  "total_execution_time": 45000.5,
  "average_attempts_per_task": 1.67,
  "total_http3_errors": 15,
  "total_masque_errors": 8,
  "total_tunnel_tx_pkts": 5000000,
  "total_tunnel_rx_pkts": 5000000,
  "total_forward_tx_pkts": 2500000,
  "total_forward_rx_pkts": 2500000,
  "total_payload_bytes": 104857600000,
  "total_total_time": 45000500
}
```

#### Delete Task
```http
DELETE /tasks/{task_id}
```

**Response (200 OK):**
```json
{
  "message": "Task 550e8400-e29b-41d4-a716-446655440000 successfully terminated",
  "terminated_processes": [
    "550e8400-e29b-41d4-a716-446655440000_attempt_2"
  ]
}
```

### Logs

#### Download All Logs
```http
GET /logs
```

**Response:** Binary tarball file (.tgz) containing all logs

#### Download Task Logs
```http
GET /logs/{task_id}
```

**Response:** Binary tarball file (.tgz) containing logs for specific task

#### Get Logs Information
```http
GET /logs/info
```

**Response (200 OK):**
```json
{
  "total_size": 524288000,
  "total_size_human": "500.00 MB",
  "total_tasks": 150,
  "total_log_files": 450,
  "log_directory": "/var/log/http3_client",
  "retention_days": 7,
  "tasks": [
    {
      "task_id": "550e8400-e29b-41d4-a716-446655440000",
      "size_bytes": 3145728,
      "size_human": "3.00 MB",
      "log_files": 3,
      "status": "completed_with_retries",
      "created_at": "2026-02-03T10:00:00.000000",
      "attempts": 3
    }
  ]
}
```

#### Delete All Logs
```http
DELETE /logs
```

**Response (200 OK):**
```json
{
  "message": "All logs deleted successfully"
}
```

#### Delete Task Logs
```http
DELETE /logs/{task_id}
```

**Response (200 OK):**
```json
{
  "message": "Logs for task 550e8400-e29b-41d4-a716-446655440000 deleted successfully"
}
```

### Certificates

#### Retrieve Certificates
```http
POST /certificates
Content-Type: application/json

{
  "arguments": {
    "proxy": "10.0.0.1:443",
    "ipv4-pool": "23.0.0.1/24"
  }
}
```

**Response (202 Accepted):**
```json
{
  "task_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

#### Get Certificate Task
```http
GET /certificates/{task_id}
```

**Response (200 OK):**
```json
{
  "task_id": "660e8400-e29b-41d4-a716-446655440001",
  "status": "completed",
  "arguments": {
    "proxy": "10.0.0.1:443"
  },
  "command": "sleep 10; /usr/app/bin/cert-retriever -proxy 10.0.0.1:443 -ipv4-pool 23.0.0.1/24",
  "start_time": "2026-02-03T10:00:00.000000",
  "end_time": "2026-02-03T10:00:15.500000",
  "output": "...",
  "error": "",
  "certificates_data": {
    "different_certs": 1,
    "server_sni": "example.com",
    "chain_length": "3",
    "is_complete": true,
    "missing_root": false,
    "certificates": [
      [
        {
          "sni": "example.com",
          "issuer": "Let's Encrypt Authority X3",
          "valid_from": "2026-01-01T00:00:00Z",
          "valid_until": "2026-04-01T00:00:00Z"
        }
      ]
    ],
    "errors": []
  }
}
```

#### List All Certificate Tasks
```http
GET /certificates
```

#### Get Last Certificate Task
```http
GET /certificates/last
```

## Task Status Values

| Status | Description |
|--------|-------------|
| `initializing` | Task created, not yet started |
| `running` | Task currently executing |
| `completed` | Single attempt completed successfully |
| `completed_with_retries` | Task completed after one or more retries, duration requirement met |
| `partial_completion` | Task stopped but didn't reach full duration (some attempts succeeded) |
| `failed` | Single attempt failed (no duration specified) |
| `failed_all_attempts` | All retry attempts failed |
| `failed_before_duration` | Single attempt failed before reaching duration |
| `task failed` | Exception occurred during task execution |
| `timeout` | Task exceeded timeout limit |
| `terminated` | Task manually terminated via DELETE |

## HTTP/3 Client Parameters

### MASQUE Proxy Configuration

| Parameter | Description | Example |
|-----------|-------------|---------|
| `--connect-udp=<proxy:port>` | Enable MASQUE client mode. Sends HTTP3 CONNECT request to proxy server to establish tunnel to target | `--connect-udp=10.0.0.1:443` |
| `--pvd-server=<pvd:port>` | Connect to PVD server to get proxy information. Mutually exclusive with `--connect-udp` | `--pvd-server=10.0.0.2:443` |
| `--pvd-only` | Only connect to PVD server, get proxy info, then close without connecting to MASQUE proxy | `--pvd-only` |
| `--quic-forwarding` | Request QUIC-Aware forwarding service from MASQUE proxy. Ignored if `--connect-udp` not provided | `--quic-forwarding` |

### Client Configuration

| Parameter | Description | Example |
|-----------|-------------|---------|
| `--clients-num=<N>` | Number of clients. Each client establishes a connection to HTTP server. Default: 1 | `--clients-num=100` |
| `--ipv4-pool=<network/mask>` | Local IPv4 address pool for client connections. Auto-detected if not specified | `--ipv4-pool=23.0.0.1/24` |
| `--ipv6-pool=<network/mask>` | Local IPv6 address pool for client connections | `--ipv6-pool=2001:db8::/64` |
| `--timeout=<N>` | QUIC connection idle timeout in seconds | `--timeout=30` |
| `--connect-speed=<N>` | Rate of new connections per second in multi-client scenarios. Default: 10. Ignored for single client | `--connect-speed=50` |
| `--mtu=<N>` | Maximum Transmission Unit. Auto-detected if not set | `--mtu=1500` |

### Traffic Generation

| Parameter | Description | Example |
|-----------|-------------|---------|
| `--duration=<N>` | Continue traffic until elapsed time exceeds duration. Supports units: H/h (hours), M/m (minutes), S/s (seconds), or no suffix (seconds) | `--duration=300`, `--duration=1.5h`, `--duration=30m` |
| `--request-dup-factor=<N>` | Duplication times of requests per connection. Increases QUIC streams (one per request). Default: 1 | `--request-dup-factor=10` |

## Certificate Retriever Parameters

### Connection Configuration

| Parameter | Description | Example |
|-----------|-------------|---------|
| `-proxy=<host:port>` | MASQUE proxy host:port (required) | `-proxy=10.0.0.1:443` |
| `-protocol=<udp\|tcp>` | Protocol to use: udp or tcp. Default: udp | `-protocol=udp` |
| `-ipv4-pool=<CIDR>` | CIDR notation for IPv4 pool | `-ipv4-pool=23.0.0.1/24` |
| `-mask=<mask>` | Subnet mask (alternative to CIDR) | `-mask=255.255.255.0` |

### TLS Configuration

| Parameter | Description | Example |
|-----------|-------------|---------|
| `-sni=<hostname>` | Override TLS SNI/ServerName for proxy certificate verification | `-sni=proxy.example.com` |
| `-tls-version=<version>` | TLS version to use: auto, 1.0, 1.1, 1.2, or 1.3. Default: auto | `-tls-version=1.3` |
| `-allow-insecure` | Proceed even if certificate verification fails (WARNING). Default: true | `-allow-insecure` |

### Performance & Output

| Parameter | Description | Example |
|-----------|-------------|---------|
| `-rampup=<N>` | Number of tasks per second (0 for unlimited) | `-rampup=10` |
| `-merge` | Merge duplicate certificates and group source IPs. Default: true | `-merge` |
| `-debug` | Enable debug output | `-debug` |

## Usage Examples

### Task Management

#### Basic HTTP/3 Test (No Retry)
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": ["--target=192.0.0.5:443", "--requests=1000"],
    "auto_retry": false
  }'
```

#### MASQUE Proxy Test with Direct Connection
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--connect-udp=10.0.0.1:443",
      "--target=192.0.0.5:443",
      "--duration=300",
      "--clients-num=50"
    ],
    "auto_retry": true
  }'
```

#### PVD Server Discovery Test
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--pvd-server=10.0.0.2:443",
      "--target=192.0.0.5:443",
      "--duration=600",
      "--clients-num=100",
      "--connect-speed=20"
    ],
    "auto_retry": true
  }'
```

#### PVD Discovery Only (No Traffic)
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--pvd-server=10.0.0.2:443",
      "--pvd-only"
    ],
    "auto_retry": false
  }'
```

#### QUIC-Aware Forwarding Test
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--connect-udp=10.0.0.1:443",
      "--quic-forwarding",
      "--target=192.0.0.5:443",
      "--duration=1h",
      "--clients-num=200"
    ],
    "auto_retry": true
  }'
```

#### High-Load Multi-Client Test
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--pvd-server=10.0.0.2:443",
      "--target=192.0.0.5:443",
      "--duration=30m",
      "--clients-num=500",
      "--connect-speed=100",
      "--request-dup-factor=5",
      "--timeout=60",
      "--ipv4-pool=23.0.0.1/24"
    ],
    "auto_retry": true
  }'
```

#### IPv6 Test with Custom MTU
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--connect-udp=10.0.0.1:443",
      "--target=192.0.0.5:443",
      "--duration=15m",
      "--ipv6-pool=2001:db8::/64",
      "--mtu=1280",
      "--clients-num=50"
    ],
    "auto_retry": true
  }'
```

#### Long-Running Stress Test
```bash
curl -X POST http://localhost/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": [
      "--pvd-server=10.0.0.2:443",
      "--target=192.0.0.5:443",
      "--duration=24h",
      "--clients-num=1000",
      "--connect-speed=50",
      "--request-dup-factor=10",
      "--timeout=120"
    ],
    "auto_retry": true
  }'
```

#### Check Task Status
```bash
TASK_ID="550e8400-e29b-41d4-a716-446655440000"
curl http://localhost/tasks/$TASK_ID | jq .
```

#### Monitor Running Tasks
```bash
# Get summary
curl http://localhost/tasks/summary | jq .

# List all running tasks
curl http://localhost/tasks | jq '.[] | select(.status == "running")'
```

#### Terminate Running Task
```bash
TASK_ID="550e8400-e29b-41d4-a716-446655440000"
curl -X DELETE http://localhost/tasks/$TASK_ID
```

#### Filter Tasks by Date
```bash
# Get tasks from last hour
START=$(date -d '1 hour ago' +%s)
END=$(date +%s)
curl "http://localhost/tasks?start_date=$START&end_date=$END" | jq .

# Get summary for specific time range
curl "http://localhost/tasks/summary?start_date=$START&end_date=$END" | jq .
```

#### Poll Task Until Completion
```bash
TASK_ID="550e8400-e29b-41d4-a716-446655440000"

while true; do
  STATUS=$(curl -s http://localhost/tasks/$TASK_ID | jq -r .status)
  echo "Status: $STATUS"
  
  if [[ "$STATUS" =~ ^(completed|completed_with_retries|failed|failed_all_attempts|task\ failed|timeout|terminated)$ ]]; then
    echo "Task finished with status: $STATUS"
    curl -s http://localhost/tasks/$TASK_ID | jq .
    break
  fi
  
  sleep 5
done
```

### Log Management

#### Download All Logs
```bash
curl -O -J http://localhost/logs
# Downloads: all_logs_20260203_100000.tgz
```

#### Download Task Logs
```bash
TASK_ID="550e8400-e29b-41d4-a716-446655440000"
curl -O -J http://localhost/logs/$TASK_ID
# Downloads: task_550e8400-e29b-41d4-a716-446655440000_logs_20260203_100000.tgz
```

#### Get Logs Information
```bash
curl http://localhost/logs/info | jq .
```

#### Delete All Logs
```bash
curl -X DELETE http://localhost/logs
```

#### Delete Task Logs
```bash
TASK_ID="550e8400-e29b-41d4-a716-446655440000"
curl -X DELETE http://localhost/logs/$TASK_ID
```

### Certificate Retrieval

#### Basic Certificate Retrieval
```bash
# Start certificate retrieval
curl -X POST http://localhost/certificates \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": {
      "proxy": "10.0.0.1:443",
      "ipv4-pool": "23.0.0.1/24"
    }
  }'

# Get result
CERT_TASK_ID="660e8400-e29b-41d4-a716-446655440001"
curl http://localhost/certificates/$CERT_TASK_ID | jq .certificates_data
```

#### Certificate Retrieval with Custom TLS Settings
```bash
curl -X POST http://localhost/certificates \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": {
      "proxy": "10.0.0.1:443",
      "protocol": "udp",
      "ipv4-pool": "23.0.0.1/24",
      "sni": "proxy.example.com",
      "tls-version": "1.3",
      "allow-insecure": "true"
    }
  }'
```

#### High-Performance Certificate Scan
```bash
curl -X POST http://localhost/certificates \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": {
      "proxy": "10.0.0.1:443",
      "ipv4-pool": "23.0.0.1/24",
      "rampup": "100",
      "merge": "true",
      "debug": "false"
    }
  }'
```

#### Certificate Retrieval with TCP Protocol
```bash
curl -X POST http://localhost/certificates \
  -H "Content-Type: application/json" \
  -d '{
    "arguments": {
      "proxy": "10.0.0.1:443",
      "protocol": "tcp",
      "mask": "255.255.255.0",
      "tls-version": "1.2"
    }
  }'
```

#### List All Certificate Tasks
```bash
curl http://localhost/certificates | jq .
```

#### Get Last Certificate Task
```bash
curl http://localhost/certificates/last | jq .
```

## Log Directory Structure

```
/var/log/http3_client/
├── tasks/
│   ├── {task_id}/
│   │   ├── attempts/
│   │   │   ├── {task_id}_attempt_1.log.gz
│   │   │   ├── {task_id}_attempt_1.meta.json
│   │   │   ├── {task_id}_attempt_2.log.gz
│   │   │   └── {task_id}_attempt_2.meta.json
│   │   ├── metadata.json
│   │   └── summary.json
└── certificates/
    └── {task_id}/
        └── output.json
```

### Log Files

- **{attempt_id}.log.gz**: Compressed execution log with command, stdout, stderr
- **{attempt_id}.meta.json**: Attempt metadata (execution time, return code, timestamps)
- **metadata.json**: Complete task metadata with all attempts
- **summary.json**: Quick summary of all attempts for the task

## Configuration

### Environment Variables
- `HTTP3_CLIENT`: Path to http3client binary (default: `/usr/app/bin/http3client`)
- `CERTIFICATE_RETRIEVER`: Path to cert-retriever binary (default: `/usr/app/bin/cert-retriever`)
- `COMMAND_TIMEOUT`: Maximum timeout per attempt in seconds (default: `3600`)
- `LOG_BASE_DIR`: Base directory for logs (default: `/var/log/http3_client`)
- `LOG_RETENTION_DAYS`: Days to keep logs before cleanup (default: `7`)

### Network Configuration
- **Client Network**: `23.0.0.0/8` - Automatically detected from container interface
- **Server Network**: `192.0.0.1/24` - Target servers reside here
- **MASQUE Proxy**: Dual-homed between client and server networks

### IPv4 Pool Auto-Detection
If `--ipv4-pool` is not specified in arguments, the wrapper automatically detects the container's IP address in the `23.0.0.0/8` range and adds `--ipv4-pool=23.x.x.x/24`.

## Metrics Extracted

The wrapper parses client output to extract:

| Metric | Description |
|--------|-------------|
| `payload_bytes` | Total user payload data received |
| `total_time` | Total execution time in milliseconds |
| `http3_errors` | QUIC connections failed to target server |
| `masque_errors` | QUIC connections failed to MASQUE proxy |
| `tunnel_rx_pkts` | QUIC-Tunnel packets received |
| `tunnel_tx_pkts` | QUIC-Tunnel packets transmitted |
| `forward_rx_pkts` | QUIC-Forward packets received |
| `forward_tx_pkts` | QUIC-Forward packets transmitted |
| `udp_throughput_rx` | UDP receive throughput |
| `udp_throughput_tx` | UDP transmit throughput |
| `udp_bytes_rx` | Total UDP bytes received |
| `udp_bytes_tx` | Total UDP bytes transmitted |
| `inner_connections_created` | Inner QUIC connections created |
| `inner_connections_closed` | Inner QUIC connections closed |
| `outer_connections_created` | Outer QUIC connections created |
| `outer_connections_closed` | Outer QUIC connections closed |
| `quic_streams_created` | QUIC streams created |
| `quic_streams_closed` | QUIC streams closed |

## Docker Deployment

```dockerfile
FROM python:3.9-slim

WORKDIR /usr/app

# Copy binaries
COPY bin/http3client /usr/app/bin/
COPY bin/cert-retriever /usr/app/bin/

# Install dependencies
RUN pip install flask

# Copy application
COPY client_wrapper.py /usr/app/
COPY openapi-client.json /usr/app/
COPY index.html /usr/app/

# Create log directory
RUN mkdir -p /var/log/http3_client

EXPOSE 80

CMD ["python", "client_wrapper.py"]
```

Run:
```bash
docker build -t client-wrapper .
docker run -d -p 80:80 --network masque-test client-wrapper
```

## API Documentation

Interactive API documentation is available at:
- Swagger UI: `http://localhost/docs`
- OpenAPI Spec: `http://localhost/openapi.json`

## Troubleshooting

### Task Stuck in "running" State
Check if the process is still alive:
```bash
docker exec -it <container> ps aux | grep http3client
```

Terminate if needed:
```bash
curl -X DELETE http://localhost/tasks/$TASK_ID
```

### No Metrics in Response
- Ensure the client binary outputs logs in the expected format
- Check compressed logs: `curl -O -J http://localhost/logs/$TASK_ID`
- Extract and inspect: `tar -xzf task_*_logs_*.tgz`
- Verify regex patterns in `process_output()` match your client's output

### IPv4 Pool Not Detected
Manually specify in arguments:
```json
{
  "arguments": ["--ipv4-pool=23.0.0.1/24", "--target=192.0.0.5:443"]
}
```

### Retry Not Working
- Ensure `--duration` parameter is present in arguments
- Set `auto_retry: true` in POST request
- Check that duration format is recognized: `--duration=300` or `--duration=300s`

### Logs Growing Too Large
- Adjust `LOG_RETENTION_DAYS` environment variable
- Manually clean up old logs: `curl -X DELETE http://localhost/logs`
- Check log info: `curl http://localhost/logs/info`

### Certificate Retrieval Fails
- Verify proxy is reachable from client network
- Check TLS settings if using custom SNI or TLS version
- Use `-allow-insecure=true` for testing (not recommended for production)
- Enable debug output: `"debug": "true"`

## License

[Your License Here]

## Contributing

[Your Contributing Guidelines Here]
