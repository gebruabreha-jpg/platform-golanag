#ifndef INNERLAYER_H
#define INNERLAYER_H

#include "TransportLayer.h"

#pragma once

class MasqueClient;
class OuterLayer;
class UdpLayer;

class InnerLayer : public TransportLayer {
public:
  // Lazy allocation in proxy mode
  InnerLayer(MasqueClient *client, const char *remote_host_name,
             const network::Address &local, const network::Address &remote,
             const std::vector<HttpRequest> *requests,
             int max_udp_payload_size);
  ~InnerLayer();

  const LayerType get_type() const final {
    return TransportLayer::LayerType::INNER;
  }

  void on_udp_data_ready_to_read(const uint8_t *buffer, size_t data_len) final;
  void on_quic_streams_can_be_opened(uint64_t max_streams) final;
  void on_quic_connection_closed() final;
  void on_tunnel_data_ready_to_read(const uint8_t *buffer, size_t data_len);

  void set_tunnel_mode(bool is_tunnel_mode);

  OuterLayer *get_tunnel_layer() const;
  UdpLayer *get_udp_layer() const;

  void build_and_submit_http3_request(HttpRequest *req, int64_t stream_id);

  void set_masque_stream_id(int64_t stream_id) {
    masque_stream_id_ = stream_id;
  }

  int64_t get_masque_stream_id() const { return masque_stream_id_; }

private:
  bool is_in_tunnel_mode_ = false;
  const std::vector<HttpRequest> *requests_ = nullptr;
  int64_t masque_stream_id_ = -1;
};

#endif