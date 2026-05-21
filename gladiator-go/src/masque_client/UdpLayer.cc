#include "UdpLayer.h"
#include "Config.h"
#include "FdWatcher.h"
#include "FlatBuffer.h"
#include "MasqueClient.h"
#include "PM.h"
#include "TransportLayer.h"
#include "debug.h"
#include "network.h"
#include "util.h"
#include <atomic>
#include <cassert>
#include <cerrno>
#include <cstdio>
#include <iostream>
#include <sys/types.h>
#include <unistd.h>
std::atomic_bool UdpLayer::mtu_acquired_ = false;

std::atomic_uint16_t UdpLayer::mtu_ = 1500;

UdpLayer::UdpLayer(int family, MasqueClient *client)
    : client_(client), ip_family_(family) {

  socket_fd_ = network::create_nonblock_socket(family, SOCK_DGRAM, IPPROTO_UDP);
  if (socket_fd_ == -1) {
    throw "fail to create socket!";
  }

  // create watcher for socket_fd
  socket_read_watcher_ =
      new FdWatcher(client_->get_ev_loop(), socket_fd_, this);
}

void UdpLayer::set_local_address(const network::Address &addr) {
  local_addr_ = addr;
}
void UdpLayer::set_remote_address(const network::Address &addr) {
  remote_addr_ = addr;
}

UdpLayer::~UdpLayer() {

  if (socket_fd_ > 0) {
    close(socket_fd_);
    socket_fd_ = -1;
  }
  delete tx_buffer_;
  delete rx_buffer_;
  delete socket_read_watcher_;
}

void UdpLayer::start() {
  // bind local address

  if (!client_->get_config()->quiet) {

    ngtcp2::debug::log_printf(
        nullptr, "client[%s] UDP start: %s == %s", client_->get_id().c_str(),
        local_addr_.to_string().c_str(), remote_addr_.to_string().c_str());
  }

  if (-1 == bind(socket_fd_, &local_addr_.su.sa, local_addr_.len)) {
    perror("Fail to bind socket!");
    exit(-1);
  }
  if (connect(socket_fd_, &remote_addr_.su.sa, remote_addr_.len) < 0) {
    perror("Fail to connect socket!");
    exit(-1);
  }

  if (!client_->get_config()->quiet) {
    // call getsockname to get the port value.
    if (getsockname(socket_fd_, (struct sockaddr *)&local_addr_.su.sa,
                    &local_addr_.len) < 0) {
      perror("getsockname");
      exit(EXIT_FAILURE);
    }
    ngtcp2::debug::log_printf(
        nullptr, "client[%s] UDP start: %s -- %s", client_->get_id().c_str(),
        local_addr_.to_string().c_str(), remote_addr_.to_string().c_str());
  }

  network::fd_set_recv_ecn(socket_fd_, ip_family_);
  network::fd_set_ip_mtu_discover(socket_fd_, ip_family_);
  network::fd_set_ip_dontfrag(socket_fd_, ip_family_);
  network::fd_set_udp_gro(socket_fd_);

  if (client_->get_config()->max_trans_unit != 0) {
    mtu_ = client_->get_config()->max_trans_unit;
    mtu_acquired_ = true;
  }

  if (!mtu_acquired_) {
    mtu_acquired_ = true; // ensure it only get once for all clients.
    // let's get the MTU of the interface which udp socket associated
    char buffer[128];
    auto ifname = network::get_interface_name(socket_fd_, ip_family_, buffer,
                                              sizeof(buffer));
    if (ifname[0] != '\0') {
      mtu_ = network::get_mtu(ifname);
      if (!client_->get_config()->quiet) {
        ngtcp2::debug::log_printf(nullptr, "The MTU of %s: %d", ifname,
                                  mtu_.load());
      }
      if (mtu_ <= 0 or mtu_ > 1500) {
        mtu_ = 1500;
      }
    } else {
      mtu_ = 1420;
    }
  }

  if (ip_family_ == AF_INET6) {
    max_udp_payload_size_ = mtu_.load() - 80;
  } else if (ip_family_ == AF_INET) {
    max_udp_payload_size_ = mtu_.load() - 36;
  } else {
    max_udp_payload_size_ = 1300;
  }

  tx_buffer_ = new FlatBuffer(mtu_.load());
  rx_buffer_ = new FlatBuffer(mtu_.load());

  socket_read_watcher_->start();
}

void UdpLayer::stop(DIRECTION direction) {

  if (direction == DIRECTION::RX or direction == DIRECTION::BI) {
    // no longer read incoming packet
    socket_read_watcher_->stop();
  }

  if (direction == DIRECTION::TX or direction == DIRECTION::BI) {
    if (socket_fd_ > 0) {
      close(socket_fd_);
      socket_fd_ = -1;
    }
  }
}

size_t UdpLayer::send_data() {

  if (tx_buffer_->len() == 0 or socket_fd_ == -1) {
    return 0;
  }

  assert(tx_buffer_->len() >= 0 &&
         tx_buffer_->len() < static_cast<ssize_t>(64_k));

  ssize_t sent_bytes = 0;

  do {

    int sending_bytes = tx_buffer_->len() < max_udp_payload_size_
                            ? tx_buffer_->len()
                            : max_udp_payload_size_;

    sent_bytes = send(socket_fd_, tx_buffer_->data(), sending_bytes, 0);
  } while (sent_bytes == -1 && errno == EINTR);

  if (sent_bytes < 0) {

    perror("UdpLayer::send_data() fail!");
    return 0;

  } else {
    stats_.total_tx_bytes += sent_bytes;
    stats_.total_tx_udp_pkts++;
    // reset client's timer
    client_->reset_idle_timer();
  }
  if (!client_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "client[%s] %s UDP sent %lu bytes",
                              client_->get_id().c_str(),
                              local_addr_.to_string().c_str(), sent_bytes);
  }

  client_->get_pm()->inc_udp_tx_bytes(sent_bytes);

  return sent_bytes;
}

void UdpLayer::recv_data() {

  ssize_t recv_bytes = 0;

  int max_retry_times = 0;

  do {

    if (socket_fd_ < 0) {
      return;
    }

    recv_bytes = recv(socket_fd_, rx_buffer_->data(), rx_buffer_->size(), 0);

    if (recv_bytes < 0) {

      if (errno == EAGAIN) {
        if (!client_->get_config()->quiet) {
          // it's expected behavior.
          ngtcp2::debug::log_printf(nullptr, "UDP no more data coming");
        }

        break; // quit while and wait for the next time the fd ready event.

      } else if (errno != EINTR) {
        perror("UdpLayer::recv_data() Error:");
        stop();
      }
      if (!client_->get_config()->quiet) {
        ngtcp2::debug::log_printf(nullptr,
                                  "client[%s] %s UDP recv() return error: %ld",
                                  client_->get_id().c_str(),
                                  local_addr_.to_string().c_str(), recv_bytes);
      }

    } else {

      rx_buffer_->set_data_len(recv_bytes);

      stats_.total_rx_bytes += recv_bytes;
      stats_.total_rx_udp_pkts++;
      if (!client_->get_config()->quiet) {
        ngtcp2::debug::log_printf(nullptr, "client[%s] %s UDP recv %lu bytes",
                                  client_->get_id().c_str(),
                                  local_addr_.to_string().c_str(), recv_bytes);
      }

      client_->get_pm()->inc_udp_rx_bytes(recv_bytes);

      // feed the received packet to upper layer!
      upper_layer_->on_udp_data_ready_to_read(rx_buffer_->data(), recv_bytes);
    }
  } while (max_retry_times-- > 0);
}

void UdpLayer::on_read_ready(FdWatcher *watcher) {
  if (client_->get_status() != MasqueClient::Status::STOPPED) {
    recv_data();
  }
}
