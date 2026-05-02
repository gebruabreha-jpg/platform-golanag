FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
# Build both components
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-server ./cmd/http-server
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-client ./cmd/http-client

# Runtime stage
FROM alpine:3.18

# Install required tools
RUN apk add --no-cache \
    iptables \
    tcpdump \
    tshark \
    wget \
    && rm -rf /var/cache/apk/*

# Copy built binary
COPY --from=builder /http-server /http-server
COPY --from=builder /http-client /http-client
# Expose ports
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/http-server"]
