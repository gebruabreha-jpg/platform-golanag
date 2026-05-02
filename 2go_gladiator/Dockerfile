# Base stage with common dependencies
FROM ubuntu:22.04 AS common_stage

# Set UTC for timezone
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/UTC
RUN ln -fs /usr/share/zoneinfo/$TZ /etc/localtime

# Install all dependencies in a single RUN layer to reduce image size
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    iproute2 \
    iputils-ping \
    tcpdump \
    libev-dev \
    cmake \
    ninja-build \
    build-essential \
    git \
    autoconf \
    libtool \
    vim \
    libgtest-dev \
    ethtool \
    libjansson-dev \
    ca-certificates \
    dnsmasq \
    nginx \
    dnsutils \
    curl \
    netcat-openbsd \
    iptables \
    openssl \
    python3 \
    python3-flask \
    python3-scapy \
    supervisor \
    && rm -rf /var/lib/apt/lists/*

# copy nginx and dnsmasq configuration
RUN mkdir -p /etc/nginx/ssl
COPY extra/config/nginx.conf /etc/nginx/
COPY extra/config/dnsmasq.conf /etc/
COPY extra/config/nginx_default /etc/nginx/sites-available/default
COPY extra/config/certificate.crt /etc/nginx/ssl/
COPY extra/config/private.key /etc/nginx/ssl/
COPY extra/config/ca_bundle.crt /etc/nginx/ssl/
# Copy all application files
COPY . /usr/app/

# Build quictls
RUN git clone --depth 1 -b openssl-3.1.4+quic https://github.com/quictls/openssl && \
    cd openssl && \
    ./config enable-tls1_3 --prefix=/usr/local && \
    make -j$(nproc) && \
    make install && \
    cp /usr/local/lib64/libssl.so* /usr/lib/x86_64-linux-gnu/ && \
    cp /usr/local/lib64/libcrypto.so* /usr/lib/x86_64-linux-gnu/ && \
    cd .. && \
    rm -rf openssl

# Build nghttp3
RUN git clone -b v1.1.0 --depth 1 https://github.com/ngtcp2/nghttp3 && \
    mkdir /nghttp3/build && \
    cd /nghttp3/build && \
    cmake -GNinja .. && \
    ninja install && \
    cd / && \
    rm -rf nghttp3

# Build ngtcp2
RUN git clone -b v1.1.0 --depth 1 https://github.com/ngtcp2/ngtcp2 && \
    mkdir /ngtcp2/build && \
    cd /ngtcp2/build && \
    cmake -GNinja .. && \
    ninja install && \
    cd / && \
    rm -rf ngtcp2

WORKDIR /usr/app

# Build the application
RUN mkdir ./build && \
    cd ./build && \
    cmake -GNinja .. && \
    ninja && \
    ldconfig && \
    ./utest/utest

# Build Go binaries in a separate stage
FROM golang:1.25 AS go_build_stage

WORKDIR /http-service
COPY http-service/ .

# Build Go binaries
RUN if [ -f "build.sh" ]; then \
    chmod +x build.sh && ./build.sh; \
    else \
    echo "Building Go binaries manually..." && \
    go build -o ./bin/linux/http-server ./cmd/http-server && \
    go build -o ./bin/linux/http-client ./cmd/http-client; \
    fi

# Runtime base image - minimal with only needed files
FROM ubuntu:22.04 AS runtime_base

# Set UTC for timezone
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/UTC
RUN ln -fs /usr/share/zoneinfo/$TZ /etc/localtime

# Install only runtime dependencies and utilities
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    # System utilities
    iproute2 \
    iputils-ping \
    tcpdump \
    dnsutils \
    netcat-openbsd \
    curl \
    openssl \
    # Application dependencies
    libev4 \
    libjansson4 \
    python3 \
    python3-flask \
    python3-scapy \
    supervisor \
    nginx \
    dnsmasq \
    iptables \
    ethtool \
    vim \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

WORKDIR /usr/app

# Create bin directory
RUN mkdir -p /usr/app/bin

# Copy C++ binaries from common_stage
COPY --from=common_stage /usr/app/build/src/masque_client/http3client /usr/app/bin/
COPY --from=common_stage /usr/app/build/src/http3_server/http3server /usr/app/bin/



# Copy Go binaries from go_build_stage
COPY --from=go_build_stage /http-service/bin/linux/* /usr/app/bin/


# Copy built libraries from common_stage
COPY --from=common_stage /usr/local/lib/ /usr/local/lib/
COPY --from=common_stage /usr/local/lib64/ /usr/local/lib64/
RUN echo /usr/local/lib > /etc/ld.so.conf.d/masque.conf
RUN echo /usr/local/lib64 >> /etc/ld.so.conf.d/masque.conf

RUN ldconfig

# Copy configuration files
RUN mkdir -p /etc/nginx/ssl
COPY --from=common_stage /etc/nginx/ /etc/nginx/
COPY --from=common_stage /etc/dnsmasq.conf /etc/

# Make all binaries executable and ensure they're in PATH
RUN chmod +x /usr/app/bin/* && \
    chmod +x /usr/app/build/* 2>/dev/null || true

# Add bin to PATH
ENV PATH="/usr/app/bin:/usr/app/build:${PATH}"

# Client stage
FROM runtime_base AS client_stage

# Copy client specific files
COPY client_wrapper.py /usr/app/
COPY openapi-client.json /usr/app/openapi.json
COPY index.html /usr/app
COPY http3_client.conf /etc/supervisor/conf.d/
COPY client_autoconfig_st.sh /usr/bin/

RUN mkdir -p /var/log/supervisor && \
    touch /var/log/supervisor/supervisord.log


CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]

# Server stage
FROM runtime_base AS server_stage

# Copy server specific files
COPY http3_server.conf /etc/supervisor/conf.d/
COPY server_autoconfig_st.sh /usr/bin/
COPY README.md /usr/app/
COPY cert.* /usr/app/ 


# Create test files
RUN dd if=/dev/zero of=/usr/app/16K bs=1k count=16 && \
    dd if=/dev/zero of=/usr/app/128K bs=1k count=128 && \
    dd if=/dev/zero of=/usr/app/2M bs=1M count=2 && \
    dd if=/dev/zero of=/usr/app/16M bs=1M count=16  

CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]
