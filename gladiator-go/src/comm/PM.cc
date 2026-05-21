#include "PM.h"
#include "debug.h"
#include <cassert>
#include <cstdio>

PM::PM() {}

PM::~PM() {}

void PM::inc_created_connects() { created_connects_.fetch_add(1); }

void PM::inc_closed_connects() { closed_connects_.fetch_add(1); }

void PM::inc_tunnel_created_connects() {
  tunnel_created_connects_.fetch_add(1);
}

void PM::inc_tunnel_closed_connects() { tunnel_closed_connects_.fetch_add(1); }

void PM::inc_opened_streams() { opened_streams_.fetch_add(1); }
void PM::inc_closed_streams() { closed_streams_.fetch_add(1); }

void PM::inc_udp_tx_bytes(int bytes) { udp_tx_bytes_.fetch_add(bytes); }

void PM::inc_udp_rx_bytes(int bytes) { udp_rx_bytes_.fetch_add(bytes); }

void PM::inc_user_rx_bytes(int bytes) {
  user_payload_rx_bytes_.fetch_add(bytes);
}

void PM::inc_inner_quic_timeout(bool handshake_completed) {
  if (handshake_completed) {
    inner_quic_timeout_times_.fetch_add(1);
  } else {
    inner_quic_connect_failure_times_.fetch_add(1);
  }
}
void PM::inc_tunnel_quic_timeout(bool handshake_completed) {
  if (handshake_completed) {
    tunnel_quic_timeout_times_.fetch_add(1);
  } else {
    tunnel_quic_connect_failure_times_.fetch_add(1);
  }
}

void PM::inc_quic_tunnel_rx_packets() { quic_tunnel_rx_packets_.fetch_add(1); }
void PM::inc_quic_tunnel_tx_packets() { quic_tunnel_tx_packets_.fetch_add(1); }

void PM::inc_quic_forwarding_rx_packets() {
  quic_forwarding_rx_packets_.fetch_add(1);
}
void PM::inc_quic_forwarding_tx_packets() {
  quic_forwarding_tx_packets_.fetch_add(1);
}

int64_t PM::get_udp_rx_bytes() const { return udp_rx_bytes_.load(); }

int64_t PM::get_udp_tx_bytes() const { return udp_tx_bytes_.load(); }

int64_t PM::get_created_conns() const { return created_connects_.load(); }
int64_t PM::get_closed_conns() const { return closed_connects_.load(); }

int64_t PM::get_tunnel_created_conns() const {
  return tunnel_created_connects_.load();
}
int64_t PM::get_tunnel_closed_conns() const {
  return tunnel_closed_connects_.load();
}

int64_t PM::get_opened_streams() const { return opened_streams_.load(); }
int64_t PM::get_closed_streams() const { return closed_streams_.load(); }

int64_t PM::get_user_rx_bytes() const { return user_payload_rx_bytes_.load(); }

int64_t PM::get_inner_quic_timeout(bool handshake_completed) const {
  return handshake_completed ? inner_quic_timeout_times_.load()
                             : inner_quic_connect_failure_times_.load();
}
int64_t PM::get_tunnel_quic_timeout(bool handshake_completed) const {
  return handshake_completed ? tunnel_quic_timeout_times_.load()
                             : tunnel_quic_connect_failure_times_.load();
}
int64_t PM::get_quic_tunnel_rx_packets() const {
  return quic_tunnel_rx_packets_.load();
}
int64_t PM::get_quic_tunnel_tx_packets() const {
  return quic_tunnel_tx_packets_.load();
}

int64_t PM::get_quic_forwarding_rx_packets() const {
  return quic_forwarding_rx_packets_.load();
}
int64_t PM::get_quic_forwarding_tx_packets() const {
  return quic_forwarding_tx_packets_.load();
}