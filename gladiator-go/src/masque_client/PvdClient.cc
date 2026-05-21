#include "PvdClient.h"
#include "Http3Connection.h"
#include "OuterLayer.h"
#include "PM.h"
#include "QuicConnection.h"
#include "debug.h"
#include "util.h"
#include <jansson.h>

PvdClient::PvdClient(MasqueClient *client, const char *remote_host_name,
                     const network::Address &local,
                     const network::Address &remote, int max_udp_payload_size)
    : TransportLayer(client, remote_host_name, local, remote) {
  quic_ = new QuicConnection(this, remote_host_name, local, remote,
                             max_udp_payload_size);
}

PvdClient::~PvdClient() {}

void PvdClient::on_udp_data_ready_to_read(const uint8_t *buffer,
                                          size_t data_len) {
  quic_->on_udp_data_ready_to_read(buffer, data_len);
}
void PvdClient::on_quic_streams_can_be_opened(uint64_t max_streams) {
  // http3 submit requests with this stream id
  if (http3_ && !http3_->has_started()) {
    http3_->start();
    HttpRequest req;

    req.method = "GET";
    req.scheme = "https";
    req.authority = client_->get_config()->pvd_server;
    req.path = "/.well-known/pvd";
    req.accept = "application/pvd+json";

    auto new_req = http3_->new_request(req);
    int64_t stream_id = -1;
    quic_->open_bidi_stream(stream_id, new_req);
    build_and_submit_http3_request(new_req, stream_id);
    http3_->send_data();
  }
}
void PvdClient::on_quic_connection_closed() {
  stopped_ = true;
  if (udp_layer_) {
    udp_layer_->stop();
  }
  // inform client
  client_->on_quic_connection_closed(this);
}

void PvdClient::build_and_submit_http3_request(HttpRequest *req,
                                               int64_t stream_id) {
  std::array<nghttp3_nv, 5> nva{
      ngtcp2::util::make_nv_nn(":method", req->method),
      ngtcp2::util::make_nv_nn(":scheme", req->scheme),
      ngtcp2::util::make_nv_nn(":authority", req->authority),
      ngtcp2::util::make_nv_nn(":path", req->path),
      ngtcp2::util::make_nv_nn("accept", req->accept),
  };

  http3_->submit_http_request(req, stream_id, nva.data(), nva.size());
}

void PvdClient::on_response_header_recv(nghttp3_rcbuf *name,
                                        nghttp3_rcbuf *value,
                                        int64_t stream_id) {

  auto namebuf = nghttp3_rcbuf_get_buf(name);
  auto valuebuf = nghttp3_rcbuf_get_buf(value);

  // check the header
  if ((strcmp(reinterpret_cast<char *>(namebuf.base), ":status") == 0) &&
      (valuebuf.base)[0] != '2') {
    // negotication failure. proxy do not support
    ngtcp2::debug::log_printf(nullptr,
                              "PVD Server does not support query service!");
    quic_->disconnect(false);
  }

  if ((strcmp(reinterpret_cast<char *>(namebuf.base), "content-type") == 0) &&
      (strcmp(reinterpret_cast<char *>(valuebuf.base),
              "application/pvd+json") != 0)) {
    ngtcp2::debug::log_printf(nullptr,
                              "The content-type is not application/pvd+json");
    quic_->disconnect(false);
  }
}

void PvdClient::on_recv_data(const uint8_t *data, size_t len) {

  for (size_t i = 0; i < len; i++) {
    pvd_data_.push_back(data[i]);
  }
}

void PvdClient::on_recv_data_completed() {
  // all pvd data have been received
  if (!parse_pvd_data() || client_->get_config()->pvd_only) {
    return;
  }

  if (masque_proxy_info_.empty()) {
    ngtcp2::debug::log_printf(nullptr, "No PVD data found in the response");
    return;
  }

  // By default, the IPv4 proxy is used
  for (auto &masque_proxy_info : masque_proxy_info_) {

    if (!masque_proxy_info.host.empty() && !masque_proxy_info.is_ipv6) {
      auto proxy_address = network::get_remote_address(
          masque_proxy_info.host.c_str(), masque_proxy_info.port.c_str(),
          masque_proxy_info.is_ipv6 ? AF_INET6 : AF_INET);
      get_client()->set_proxy_server_address(proxy_address,
                                             masque_proxy_info.host.c_str());
      break;
    }
  }
}

bool PvdClient::parse_pvd_data() {

  auto json = json_loadb(reinterpret_cast<const char *>(pvd_data_.data()),
                         pvd_data_.size(), 0, NULL);
  if (!json) {
    ngtcp2::debug::log_printf(nullptr, "Failed to parse PVD JSON data");
    json_decref(json);
    return false;
  }

  json_t *proxies = json_object_get(json, "proxies");
  if (!proxies) {
    ngtcp2::debug::log_printf(nullptr, "Failed to get PVD object");
    json_decref(json);
    return false;
  }

  size_t array_length = json_array_size(proxies);

  for (size_t i = 0; i < array_length; i++) {
    json_t *element = json_array_get(proxies, i);
    const char *protocol =
        json_string_value(json_object_get(element, "protocol"));

    const char *proxy = json_string_value(json_object_get(element, "proxy"));

    HttpRequest req;
    if (!req.parse_uri(proxy)) {
      ngtcp2::debug::log_printf(nullptr, "Could not parse URI: %s\n", proxy);
      json_decref(json);
      return false;
    }

    MasqueProxyInfo masque_proxy_info;
    masque_proxy_info.host = req.host_name;
    masque_proxy_info.port = req.port;
    masque_proxy_info.path = req.path;
    if (req.authority.c_str()[0] == '[') {
      masque_proxy_info.is_ipv6 = true;
    } else {
      masque_proxy_info.is_ipv6 = false;
    }
    masque_proxy_info_.push_back(masque_proxy_info);

    if (!get_config()->quiet) {
      ngtcp2::debug::log_printf(
          nullptr, "Proxy[%d] protocol: %s host: %s port: %s path: %s", i,
          protocol, masque_proxy_info.host.c_str(),
          masque_proxy_info.port.c_str(), masque_proxy_info.path.c_str());
    }
  }
  json_decref(json);

  return true;
}