#ifndef TUNNELLAYER_H
#define TUNNELLAYER_H
#pragma once
#include "MasqueStreamInfo.h"
#include "TransportLayer.h"
#include <cstdint>
#include <cstdlib>
#include <map>
#include <ngtcp2/ngtcp2.h>
#include <unordered_map>

class MasqueClient;
struct nghttp3_rcbuf;
class InnerLayer;
struct ngtcp2_cid_token;

namespace std {
template <> struct hash<ngtcp2_cid> {
  std::size_t operator()(const ngtcp2_cid &cid) const {

    std::size_t hash = 5381;
    for (size_t i = 0; i < cid.datalen; ++i) {
      hash = (hash * 33) + cid.data[i];
    }
    return hash;
  }
};

template <> struct equal_to<ngtcp2_cid> {
  bool operator()(const ngtcp2_cid &lhs, const ngtcp2_cid &rhs) const {
    if (lhs.datalen != rhs.datalen) {
      return false;
    }
    for (size_t i = 0; i < lhs.datalen; ++i) {
      if (lhs.data[i] != rhs.data[i]) {
        return false;
      }
    }
    return true;
  }
};

} // namespace std

class OuterLayer : public TransportLayer {
public:
  OuterLayer(MasqueClient *client, const char *remote_host_name,
             const network::Address &local, const network::Address &remote,
             int max_udp_payload_size);
  ~OuterLayer() noexcept;

  const LayerType get_type() const final {
    return TransportLayer::LayerType::OUTER;
  }

  void on_udp_data_ready_to_read(const uint8_t *buffer, size_t data_len) final;
  void on_quic_streams_can_be_opened(uint64_t max_streams) final;
  void on_quic_connection_closed() final;
  void on_response_header_recv(nghttp3_rcbuf *name, nghttp3_rcbuf *value,
                               int64_t stream_id);

  FlatBuffer *get_tx_buffer() { return tx_buffer_; }

  // outer quic <-> inner quic
  void tunnel_mode_egress_forward();
  void tunnel_mode_ingress_forward(uint32_t flags, const uint8_t *data,
                                   size_t datalen);

  // udp <-> inner quic
  void quic_aware_mode_egress_forward(int64_t stream_id);
  void quic_aware_mode_ingress_forward(const uint8_t *buffer, size_t data_len,
                                       InnerLayer *inner);

  InnerLayer *get_inner_layer(int64_t stream_quarter) const;

  bool is_quic_forwarding_supported(int64_t stream_id) const;

  bool is_quic_forwarding_ready(int64_t stream_id) const;

  bool is_quic_forwarding_rx_ready(int64_t stream_id) const;
  bool is_quic_forwarding_tx_ready(int64_t stream_id) const;

  void on_inner_quic_handshake_completed(QuicConnection *quic,
                                         int64_t masque_stream_id);
  void on_register_client_cid_acked(const ngtcp2_cid &scid,
                                    const ngtcp2_cid &vcid,
                                    int64_t masque_stream_id);
  void on_register_target_cid_acked(const ngtcp2_cid &tcid,
                                    const ngtcp2_cid &vtcid,
                                    int64_t masque_stream_id);

  void open_streams_and_submit_requests();

  void build_and_submit_http3_request(HttpRequest *req, int64_t stream_id);

  void on_recv_capsule_data(const uint8_t *data, size_t len, int64_t stream_id);

  void on_stop_tx() override;

private:
  FlatBuffer *tx_buffer_ = nullptr;

  void build_register_client_cid_message(const ngtcp2_cid &cid,
                                         int64_t masque_stream_id);

  void build_register_target_cid_message(const ngtcp2_cid_token &tcid,
                                         int64_t masque_stream_id);

  void build_ack_client_vcid_message(const ngtcp2_cid &cid,
                                     const ngtcp2_cid &vcid,
                                     int64_t masque_stream_id);

  std::unordered_map<int64_t, MasqueStreamInfo>
      quarter_stream_id_to_inner_info_map_;

  // hash map from vscid to inner layer(for forwarding mode routing)
  std::unordered_map<ngtcp2_cid, InnerLayer *> vcid_to_inner_map_;
};

#endif
