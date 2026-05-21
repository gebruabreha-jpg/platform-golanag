#include "OuterLayer.h"

#include "Capsule.h"
#include "Http3Connection.h"
#include "HttpRequest.h"
#include "InnerLayer.h"
#include "MasqueClient.h"
#include "MasqueStreamInfo.h"
#include "PM.h"
#include "QuicConnection.h"
#include "TransportLayer.h"
#include "debug.h"
#include "util.h"
#include <algorithm>
#include <cassert>
#include <cstdint>
#include <cstring>
#include <ngtcp2/ngtcp2.h>
#include <sstream>

namespace {

std::string strip_spaces(const std::string &str) {
  size_t first = str.find_first_not_of(" \t\n\r\f\v");
  if (first == std::string::npos)
    return ""; // String contains only whitespace

  size_t last = str.find_last_not_of(" \t\n\r\f\v");
  return str.substr(first, (last - first + 1));
}

std::pair<std::string, std::map<std::string, std::string>>
parse_http_header(const std::string &input) {
  std::pair<std::string, std::map<std::string, std::string>> result;
  std::istringstream iss(input);
  std::string token;

  // read the value of the header
  if (std::getline(iss, token, ';')) {
    result.first = strip_spaces(token);
  }

  // read the parameters of the header
  while (std::getline(iss, token, ';')) {
    std::istringstream token_stream(token);
    std::string key, value;
    if (std::getline(token_stream, key, '=') &&
        std::getline(token_stream >> std::ws, value)) {
      result.second[strip_spaces(key)] = strip_spaces(value);
    } else {
      // Handle parsing error if needed
    }
  }

  return result;
}
} // namespace

OuterLayer::OuterLayer(MasqueClient *client, const char *remote_host_name,
                       const network::Address &local,
                       const network::Address &remote, int max_udp_payload_size)
    : TransportLayer(client, remote_host_name, local, remote) {
  quic_ = new QuicConnection(this, remote_host_name, local, remote,
                             max_udp_payload_size);

  tx_buffer_ = new FlatBuffer(max_udp_payload_size);
}

OuterLayer::~OuterLayer() noexcept { delete tx_buffer_; }

void OuterLayer::on_udp_data_ready_to_read(const uint8_t *buffer,
                                           size_t data_len) {
  bool is_short_header = (buffer[0] & 0x80) == 0;

  InnerLayer *inner = nullptr;

  if (is_short_header) {
    ngtcp2_cid vcid{};
    get_quic()->get_scid(
        vcid); // get the data len, the inner QUIC has the same cid length

    for (size_t i = 0; i < std::min(vcid.datalen, data_len - 1); i++) {
      vcid.data[i] = buffer[i + 1];
    }

    auto it = vcid_to_inner_map_.find(vcid);
    if (it != vcid_to_inner_map_.end()) {
      inner = it->second;
    }
  }

  if (is_short_header && inner) {
    quic_aware_mode_ingress_forward(buffer, data_len, inner);
  } else {
    quic_->on_udp_data_ready_to_read(buffer, data_len);
  }
};

void OuterLayer::on_quic_connection_closed() {
  if (stopped_) {
    return;
  }
  stopped_ = true;

  if (udp_layer_) {
    udp_layer_->stop();
  }
  // inform client
  client_->on_quic_connection_closed(this);
}

void OuterLayer::on_quic_streams_can_be_opened(uint64_t max_streams) {

  if (http3_ && !http3_->has_started()) {
    http3_->start();
    open_streams_and_submit_requests();
  }
}

void OuterLayer::open_streams_and_submit_requests() {

  for (const auto &ele : get_config()->target_servers_to_requests) {
    auto &[host, port] = ele.first;
    HttpRequest req;
    req.scheme = "https";
    req.method = "CONNECT";
    req.authority = client_->get_config()->proxy_server;
    req.protocol = "connect-udp";
    req.authority += ":";
    req.authority += client_->get_config()->proxy_server_port;
    req.path = "/.well-known/masque/udp/";
    req.path += host;
    req.path += "/";
    req.path += port;
    req.path += "/";

    auto new_req = http3_->new_request(req);
    int64_t stream_id = -1;
    quic_->open_bidi_stream(stream_id, new_req);

    MasqueStreamInfo info = {
        .stream_id = stream_id, .host = host, .port = port};

    quarter_stream_id_to_inner_info_map_[stream_id >> 2] = info;

    build_and_submit_http3_request(new_req, stream_id);
    http3_->send_data();
  }
}

void OuterLayer::build_and_submit_http3_request(HttpRequest *req,
                                                int64_t stream_id) {
  std::array<nghttp3_nv, 7> nva{
      ngtcp2::util::make_nv_nn(":method", req->method),
      ngtcp2::util::make_nv_nn(":protocol", req->protocol),
      ngtcp2::util::make_nv_nn(":scheme", req->scheme),
      ngtcp2::util::make_nv_nn(":authority", req->authority),
      ngtcp2::util::make_nv_nn(":path", req->path.c_str()),
      ngtcp2::util::make_nv_nn("capsule-protocol", "?1")};

  size_t nvlen = 6;

  if (get_config()->request_quic_forwarding) {
    nva[6] = ngtcp2::util::make_nv_nn("proxy-quic-forwarding",
                                      "?1; accept-transform=identity");
    nvlen = 7;
  }

  http3_->submit_http_request(req, stream_id, nva.data(), nvlen);
}

void OuterLayer::on_response_header_recv(nghttp3_rcbuf *name,
                                         nghttp3_rcbuf *value,
                                         int64_t stream_id) {

  auto namebuf = nghttp3_rcbuf_get_buf(name);
  auto valuebuf = nghttp3_rcbuf_get_buf(value);

  // check the header
  if ((strcmp(reinterpret_cast<char *>(namebuf.base), ":status") == 0)) {
    if ((valuebuf.base)[0] == '2') {
      auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
      client_->on_tunnel_setup_ready_async(stream_id, info);
    }
    else {
      // negotiation failure. proxy do not support
      ngtcp2::debug::log_printf(nullptr, "Proxy does not support MASQUE!");
      quic_->disconnect(false);
    }
  }

  if ((strcmp(reinterpret_cast<char *>(namebuf.base), "capsule-protocol") ==
       0) &&
      (strcmp(reinterpret_cast<char *>(valuebuf.base), "?1") == 0)) {
    auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
    info.support_capsule_protocol = true;
  }

  if (strcmp(reinterpret_cast<char *>(namebuf.base), "proxy-quic-forwarding") ==
      0) {

    auto header = parse_http_header(reinterpret_cast<char *>(valuebuf.base));

    // check its value
    if (header.first == "?1") {

      // check the transform parameter
      // refer to
      // https://ietf-wg-masque.github.io/draft-ietf-masque-quic-proxy/draft-ietf-masque-quic-proxy.html#transforms
      if (auto p = header.second.find("transform"); p != header.second.end()) {

        auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
        info.support_quic_forwarding = true;

        if (!get_config()->quiet) {
          ngtcp2::debug::log_printf(nullptr,
                                    "Proxy supports QUIC forwarding! %s = %s",
                                    p->first.c_str(), p->second.c_str());
        }
      }

      // get scramble-key
      if (auto p = header.second.find("scramble-key");
          p != header.second.end()) {

        if (!get_config()->quiet) {
          ngtcp2::debug::log_printf(nullptr, "%s = %s", p->first.c_str(),
                                    p->second.c_str());
        }
      }
    }
  }
}

void OuterLayer::tunnel_mode_egress_forward() {

  if (tx_buffer_ && tx_buffer_->len()) {
    quic_->write_datagram_data(tx_buffer_);
    get_client()->get_pm()->inc_quic_tunnel_tx_packets();
  }
}

void OuterLayer::tunnel_mode_ingress_forward(uint32_t flags,
                                             const uint8_t *data,
                                             size_t datalen) {
  uint64_t stream_quarter;
  uint8_t *buff = const_cast<uint8_t *>(data);

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr,
        "Client[%s] Tunner Layer ingress forward datagram:%p, datalen:%lu",
        get_client_id().c_str(), data, datalen);
  }

  // decode and strip off the stream id & context id
  auto bytes =
      ngtcp2::util::decode_var_len_integer(stream_quarter, buff, datalen);

  buff += bytes;
  datalen -= bytes;

  uint64_t context_id;
  bytes = ngtcp2::util::decode_var_len_integer(context_id, buff, datalen);
  buff += bytes;
  datalen -= bytes;
  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] Tunner Layer stream-id/4:%lu context-id:%lu\n",
        get_client_id().c_str(), stream_quarter, context_id);
  }

  // forward left bytes to inner layer.
  auto inner = get_inner_layer(stream_quarter);
  if (inner) {
    auto quic = inner->get_quic();
    quic->recv_data(buff, datalen);
    get_client()->get_pm()->inc_quic_tunnel_rx_packets();
  }
}

InnerLayer *OuterLayer::get_inner_layer(int64_t stream_quarter) const {
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_quarter);
  return info.inner_layer;
}

bool OuterLayer::is_quic_forwarding_supported(int64_t stream_id) const {
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
  return (info.support_quic_forwarding && info.support_capsule_protocol);
}

bool OuterLayer::is_quic_forwarding_ready(int64_t stream_id) const {
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
  return info.quic_forwording_rx_ready && info.quic_forwording_tx_ready;
}

void OuterLayer::on_inner_quic_handshake_completed(QuicConnection *quic,
                                                   int64_t masque_stream_id) {

  auto &info = quarter_stream_id_to_inner_info_map_.at(masque_stream_id >> 2);
  if (!(info.support_quic_forwarding && info.support_capsule_protocol)) {
    return;
  }

  build_register_client_cid_message(info.scid, masque_stream_id);
  http3_->resume_data_read(masque_stream_id);
  http3_->send_data();
}

void OuterLayer::build_register_client_cid_message(const ngtcp2_cid &cid,
                                                   int64_t masque_stream_id) {
  HttpRequest *req = http3_->get_request(masque_stream_id);

  req->data_len = Capsule::create_register_client_cid(cid, req->data_body,
                                                      get_config()->quiet);
}

void OuterLayer::build_register_target_cid_message(const ngtcp2_cid_token &tcid,
                                                   int64_t masque_stream_id) {

  HttpRequest *req = http3_->get_request(masque_stream_id);

  req->data_len = Capsule::create_register_target_cid(tcid.cid, req->data_body,
                                                      get_config()->quiet);
  req->data_end = true; // no more data to send
}

void OuterLayer::build_ack_client_vcid_message(const ngtcp2_cid &cid,
                                               const ngtcp2_cid &vcid,
                                               int64_t masque_stream_id) {
  HttpRequest *req = http3_->get_request(masque_stream_id);

  req->data_len = Capsule::create_ack_client_vcid(cid, vcid, req->data_body,
                                                  get_config()->quiet);
}

void OuterLayer::on_register_client_cid_acked(const ngtcp2_cid &scid,
                                              const ngtcp2_cid &vcid,
                                              int64_t masque_stream_id) {

  auto &info = quarter_stream_id_to_inner_info_map_.at(masque_stream_id >> 2);

  auto quic = info.inner_layer->get_quic();

  quic->get_dcid(info.dcid);

  info.quic_forwording_rx_ready = true;

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "quic_forwording_rx_ready");
  }

  // store vcid from ACK_CLIENT_CID capsule
  auto inner = get_inner_layer(masque_stream_id >> 2);
  vcid_to_inner_map_[vcid] = inner;

  quic->get_scid(info.scid);

  if (0 != memcmp(&info.scid, &scid, scid.datalen)) {
    ngtcp2::debug::log_printf(
        nullptr, "ERROR! - OuterLayer::on_register_client_cid_acked: "
                 "CID sent by client does not equal to CID received from "
                 "ACK_CLIENT_CID capsule");
    assert(0);
    abort();
  }

  // TODO: Verify if there is any conflict with other stored vcids
  info.vscid = vcid;

  build_ack_client_vcid_message(info.scid, info.vscid, masque_stream_id);

  http3_->resume_data_read(masque_stream_id);

  http3_->send_data();

  build_register_target_cid_message(info.dcid, masque_stream_id);

  http3_->resume_data_read(masque_stream_id);

  http3_->send_data();
}

void OuterLayer::on_register_target_cid_acked(const ngtcp2_cid &tcid,
                                              const ngtcp2_cid &vtcid,
                                              int64_t masque_stream_id) {

  auto &info = quarter_stream_id_to_inner_info_map_.at(masque_stream_id >> 2);

  if (0 != memcmp(&info.dcid.cid, &tcid, tcid.datalen)) {
    ngtcp2::debug::log_printf(
        nullptr, "ERROR! - OuterLayer::on_register_target_cid_acked: "
                 "TCID sent by client does not equal to TCID received from "
                 "ACK_TARGET_CID capsule");
    assert(0);
    abort();
  }

  info.vdcid.cid = vtcid;
  info.quic_forwording_tx_ready = true;
  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] Tunner Layer quic_forwording_tx_ready",
        get_client_id().c_str());
  }
}

void OuterLayer::quic_aware_mode_egress_forward(int64_t stream_id) {

  auto udp = get_udp_layer();
  auto tx_buffer = udp->get_tx_buffer();
  uint8_t *data = tx_buffer->data();

  data += 1; // skip the header form
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);

  // replace the dcid with vdcid
  memcpy(data, info.vdcid.cid.data, info.vdcid.cid.datalen);
  udp->send_data();

  get_client()->get_pm()->inc_quic_forwarding_tx_packets();

  // the data do not go to quic layer, but need to update its time
  auto quic = get_quic();
  auto ts = ngtcp2::util::timestamp();
  ngtcp2_conn_update_pkt_tx_time(quic->get_quic_conn(), ts);
}

void OuterLayer::quic_aware_mode_ingress_forward(const uint8_t *buffer,
                                                 size_t data_len,
                                                 InnerLayer *inner) {

  uint8_t *data = const_cast<uint8_t *>(buffer);
  data += 1; // skip the header form

  ngtcp2_cid inner_quic_client_cid{};
  inner->get_quic()->get_scid(inner_quic_client_cid);

  memcpy(data, inner_quic_client_cid.data, inner_quic_client_cid.datalen);

  inner->on_udp_data_ready_to_read(buffer, data_len);

  get_client()->get_pm()->inc_quic_forwarding_rx_packets();

  auto quic = get_quic();
  quic->update_timer();
}

bool OuterLayer::is_quic_forwarding_rx_ready(int64_t stream_id) const {
  if (stream_id < 0) {
    return false;
  }
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
  return info.quic_forwording_rx_ready;
}

bool OuterLayer::is_quic_forwarding_tx_ready(int64_t stream_id) const {
  if (stream_id < 0) {
    return false;
  }
  auto &info = quarter_stream_id_to_inner_info_map_.at(stream_id >> 2);
  return info.quic_forwording_tx_ready;
}

void OuterLayer::on_recv_capsule_data(const uint8_t *data, size_t len,
                                      int64_t stream_id) {
  if (!get_config()->quiet) {
    ngtcp2::debug::print_http_data(stream_id, data, len);
  }

  // decode the type
  uint64_t type = 0;
  auto pbuf = data;
  auto left_size = len;
  auto bytes = ngtcp2::util::decode_var_len_integer(type, pbuf, left_size);
  if (bytes == 0) {
    return;
  }
  pbuf += bytes;
  if (left_size < static_cast<size_t>(bytes)) {
    return;
  }
  left_size -= bytes;

  uint64_t total_len = 0;
  bytes = ngtcp2::util::decode_var_len_integer(total_len, pbuf, left_size);
  if (bytes == 0) {
    return;
  }
  pbuf += bytes;
  if (left_size < static_cast<size_t>(bytes)) {
    return;
  }
  left_size -= bytes;

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr,
        "Client[%s] Tunner Layer on_recv_capsule_data type: %lX total len=%lu",
        get_client_id().c_str(), type, total_len);
  }

  switch (type) {
  case Capsule::ACK_CLIENT_CID: {
    ngtcp2_cid scid;
    ngtcp2_cid vscid;
    if (Capsule::parse_ack_client_cid(pbuf, left_size, scid, vscid,
                                      get_config()->quiet)) {
      on_register_client_cid_acked(scid, vscid, stream_id);
    }
  } break;
  case Capsule::ACK_TARGET_CID: {
    ngtcp2_cid tcid;
    ngtcp2_cid vtcid;
    if (Capsule::parse_ack_target_cid(pbuf, left_size, tcid, vtcid,
                                      get_config()->quiet)) {
      on_register_target_cid_acked(tcid, vtcid, stream_id);
    }
  } break;
  case Capsule::CLOSE_CLIENT_CID: {
    ngtcp2_cid scid;
    Capsule::parse_close_x_cid(pbuf, left_size, scid, get_config()->quiet);
    // TODO: close capsule protocol
  } break;
  case Capsule::CLOSE_TARGET_CID: {
    ngtcp2_cid tcid;
    Capsule::parse_close_x_cid(pbuf, left_size, tcid, get_config()->quiet);
    // TODO: close capsule protocol
  } break;
  default:
    break;
  }
}

void OuterLayer::on_stop_tx() {

  // set all inner layer's stopped.
  client_->set_inner_layers_stopped();

  TransportLayer::on_stop_tx();

  // set client stopped
  client_->on_stop();
}
