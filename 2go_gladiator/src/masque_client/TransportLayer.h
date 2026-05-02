#ifndef TRANSPORTLAYER_H
#define TRANSPORTLAYER_H

#include "UdpLayer.h"
#include "network.h"
#include <cstdio>
#pragma once
#include <cstdint>
#include <cstdlib>

class MasqueClient;
class Http3Connection;
class QuicConnection;
class Config;

class TransportLayer {
public:
  enum class LayerType { INNER, OUTER, PVD_CLIENT };

  static constexpr uint8_t TX_BUFF_VEC_NUM = 16;

  TransportLayer(MasqueClient *client, const char *remote_host_name,
                 const network::Address &local, const network::Address &remote);
  virtual ~TransportLayer();

  struct ev_loop *get_ev_loop() const;

  virtual UdpLayer *get_udp_layer() const;

  virtual void on_udp_data_ready_to_read(const uint8_t *buffer,
                                         size_t data_len) = 0;

  virtual void on_quic_handshake_completed();
  // virtual void on_quic_write_stream(int64_t stream_id,ssize_t error, ssize_t
  // stream_data_len); virtual ssize_t on_quic_recv_stream_data(int64_t
  // stream_id,const uint8_t *data, size_t datalen,uint32_t flags);

  virtual void on_quic_connection_closed() = 0;

  virtual void on_quic_streams_can_be_opened(uint64_t max_streams) = 0;

  // when http3 on_stopping_send (ask peer stop sending)
  // It no longer receive byte
  virtual void on_stop_rx();

  // when http3 on_reset
  // It no longer send byte
  virtual void on_stop_tx();

  QuicConnection *get_quic() const;
  Http3Connection *get_http3() const;

  virtual void start();
  virtual void stop();

  bool is_stopped() const { return stopped_; }
  void set_stopped() { stopped_ = true; }
  virtual const LayerType get_type() const = 0;

  virtual void send_data();

  const Config *get_config() const;

  MasqueClient *get_client() const;

  bool is_outer_layer() const;
  bool is_inner_layer() const;
  bool is_pvd_client() const;

  const char *get_layer_str() const;

  const std::string get_client_id() const { return client_->get_id(); }

  void set_udp_layer(UdpLayer *udp_layer) { udp_layer_ = udp_layer; }

protected:
  MasqueClient *const client_ = nullptr;
  QuicConnection *quic_ = nullptr;
  Http3Connection *http3_ = nullptr;

  UdpLayer *udp_layer_ = nullptr;
  std::atomic_bool stopped_ = true;
};

#endif