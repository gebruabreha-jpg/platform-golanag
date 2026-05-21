#ifndef PVDCLIENT_H
#define PVDCLIENT_H

#pragma once

#include "TransportLayer.h"

struct nghttp3_rcbuf;

class PvdClient : public TransportLayer {

public:
  PvdClient(MasqueClient *client, const char *remote_host_name,
            const network::Address &local, const network::Address &remote,
            int max_udp_payload_size);
  ~PvdClient();

  const LayerType get_type() const final {
    return TransportLayer::LayerType::PVD_CLIENT;
  }

  void on_udp_data_ready_to_read(const uint8_t *buffer, size_t data_len) final;
  void on_quic_streams_can_be_opened(uint64_t max_streams) final;
  void on_quic_connection_closed() final;

  void on_response_header_recv(nghttp3_rcbuf *name, nghttp3_rcbuf *value,
                               int64_t stream_id);

  void build_and_submit_http3_request(HttpRequest *req, int64_t stream_id);

  void on_recv_data(const uint8_t *data, size_t len);
  void on_recv_data_completed();

  bool parse_pvd_data();

private:
  std::vector<uint8_t> pvd_data_;

  // masque proxy information
  struct MasqueProxyInfo {
    std::string host;
    std::string port;
    std::string path;
    bool is_ipv6 = false;
  };

  std::vector<MasqueProxyInfo> masque_proxy_info_;
};

#endif