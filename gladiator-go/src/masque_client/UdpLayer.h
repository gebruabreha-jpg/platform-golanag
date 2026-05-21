#ifndef UDPLAYER_H
#define UDPLAYER_H

#include "EventHandler.h"
#include "FdWatcher.h"
#include "FlatBuffer.h"
#include "MasqueClient.h"
#include "network.h"
#include "template.h"
#include <atomic>
#include <cstdint>
#include <cstring>
#include <ev.h>
#pragma once

class UdpLayer : public EventHandler {
public:
  struct Stats {
    uint64_t total_tx_bytes = 0;
    uint64_t total_rx_bytes = 0;
    uint64_t total_tx_udp_pkts = 0;
    uint64_t total_rx_udp_pkts = 0;
  };

  enum class DIRECTION : uint8_t { RX, TX, BI };

  UdpLayer(int family, MasqueClient *client);
  virtual ~UdpLayer();

  void on_timeout(TimerWatcher *watcher) final{}; // do not care
  void on_read_ready(FdWatcher *watcher) final;

  void set_local_address(const network::Address &addr);
  void set_remote_address(const network::Address &addr);
  const network::Address &get_remote_address() const { return remote_addr_; }

  const network::Address &get_local_address() const { return local_addr_; }

  int get_socket_fd() const { return socket_fd_; }

  void start();
  void stop(DIRECTION direction = DIRECTION::BI);

  // return bytes sent out
  size_t send_data();

  void recv_data();

  FlatBuffer *get_tx_buffer() { return tx_buffer_; }

  FlatBuffer *get_rx_buffer() { return rx_buffer_; }

  const struct Stats &get_stats() const { return stats_; }

  int get_max_udp_payload_size() const { return max_udp_payload_size_; }

  void set_upper_layer(TransportLayer *upper_layer) {
    upper_layer_ = upper_layer;
  }

private:
  MasqueClient *const client_ = nullptr;
  network::Address local_addr_;
  network::Address remote_addr_;

  int socket_fd_ = -1;
  int ip_family_ = AF_INET;
  FdWatcher *socket_read_watcher_ = nullptr;

  FlatBuffer *tx_buffer_ = nullptr;

  FlatBuffer *rx_buffer_ = nullptr;

  TransportLayer *upper_layer_ = nullptr;

  // this default value will be updated when socket is created and MTU is
  // gotten!
  int32_t max_udp_payload_size_ = 1500 - 80;

  static std::atomic_bool mtu_acquired_; // TODO: remove it.

  static std::atomic_uint16_t mtu_;

  Stats stats_;
};

#endif