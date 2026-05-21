#ifndef PM_H
#define PM_H

#pragma once

#include <atomic>
#include <string>

class PM {
public:
  PM();
  ~PM();

  void inc_created_connects();
  void inc_closed_connects();

  void inc_tunnel_created_connects();
  void inc_tunnel_closed_connects();

  void inc_opened_streams();
  void inc_closed_streams();

  void inc_udp_tx_bytes(int bytes);
  void inc_udp_rx_bytes(int bytes);

  void inc_user_rx_bytes(int bytes);

  void inc_inner_quic_timeout(bool handshake_completed);
  void inc_tunnel_quic_timeout(bool handshake_completed);

  void inc_quic_tunnel_rx_packets();
  void inc_quic_tunnel_tx_packets();

  void inc_quic_forwarding_rx_packets();
  void inc_quic_forwarding_tx_packets();

  int64_t get_udp_rx_bytes() const;
  int64_t get_udp_tx_bytes() const;

  int64_t get_created_conns() const;
  int64_t get_closed_conns() const;

  int64_t get_opened_streams() const;
  int64_t get_closed_streams() const;

  int64_t get_tunnel_created_conns() const;
  int64_t get_tunnel_closed_conns() const;

  int64_t get_user_rx_bytes() const;

  int64_t get_inner_quic_timeout(bool handshake_completed) const;
  int64_t get_tunnel_quic_timeout(bool handshake_completed) const;


  int64_t get_quic_tunnel_rx_packets() const;
  int64_t get_quic_tunnel_tx_packets() const;

  int64_t get_quic_forwarding_rx_packets() const;
  int64_t get_quic_forwarding_tx_packets() const;

private:
  std::atomic_int64_t created_connects_ = 0;
  std::atomic_int64_t closed_connects_ = 0;
  std::atomic_int64_t opened_streams_ = 0;
  std::atomic_int64_t closed_streams_ = 0;

  std::atomic_int64_t tunnel_created_connects_ = 0;
  std::atomic_int64_t tunnel_closed_connects_ = 0;

  std::atomic_int64_t udp_tx_bytes_ = 0;
  std::atomic_int64_t udp_rx_bytes_ = 0;

  std::atomic_int64_t inner_quic_connect_failure_times_ = 0;
  std::atomic_int64_t inner_quic_timeout_times_ = 0;
  std::atomic_int64_t tunnel_quic_connect_failure_times_ = 0;
  std::atomic_int64_t tunnel_quic_timeout_times_ = 0;

  // user payload bytes
  std::atomic_int64_t user_payload_rx_bytes_ = 0;

  // rx packets by quic tunnel
  std::atomic_int64_t quic_tunnel_rx_packets_ = 0;
  // tx packets by quic tunnel
  std::atomic_int64_t quic_tunnel_tx_packets_ = 0;

  // rx packets by quic forwarding
  std::atomic_int64_t quic_forwarding_rx_packets_ = 0;
  // tx packets by quic forwarding
  std::atomic_int64_t quic_forwarding_tx_packets_ = 0;

  // pvd requests
  std::atomic_int64_t pvd_requests_ = 0;
  // pvd responses
  std::atomic_int64_t pvd_responses_ = 0;
};

#endif