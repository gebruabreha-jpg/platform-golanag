#include "InnerLayer.h"
#include "Http3Connection.h"
#include "HttpRequest.h"
#include "MasqueClient.h"
#include "OuterLayer.h"
#include "QuicConnection.h"
#include "debug.h"
#include "util.h"
#include <filesystem>
#include <string>

InnerLayer::InnerLayer(MasqueClient *client, const char *remote_host_name,
                       const network::Address &local,
                       const network::Address &remote,
                       const std::vector<HttpRequest> *requests,
                       int max_udp_payload_size)
    : TransportLayer(client, remote_host_name, local, remote) {
  quic_ = new QuicConnection(this, remote_host_name, local, remote,
                             max_udp_payload_size);
  requests_ = requests;
}

InnerLayer::~InnerLayer() {}

void InnerLayer::on_udp_data_ready_to_read(const uint8_t *buffer,
                                           size_t data_len) {

  quic_->on_udp_data_ready_to_read(buffer, data_len);
};

void InnerLayer::on_quic_connection_closed() {
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

void InnerLayer::on_tunnel_data_ready_to_read(const uint8_t *buffer,
                                              size_t data_len) {

  // quic read the data
  quic_->recv_data(buffer, data_len);
};

void InnerLayer::on_quic_streams_can_be_opened(uint64_t max_streams) {
  // http3 submit requests with this stream id
  if (http3_ && !http3_->has_started()) {
    http3_->start();
    for (auto &req : *requests_) {

      auto new_req = http3_->new_request(req);
      int64_t stream_id = -1;
      quic_->open_bidi_stream(stream_id, new_req);

      auto &output_dir = client_->get_config()->output_dir;
      if (!output_dir.empty()) {

        std::filesystem::path dirname(output_dir);

        std::filesystem::path req_path(new_req->path);
        auto req_file = req_path.filename();
        auto prefix = new_req->authority + "_" + client_->get_id() + "_" +
                      std::to_string(stream_id) + "_";

        auto filename =
            req_file.empty() ? "index.html" : req_file.generic_string();

        auto full_path = dirname / std::filesystem::path(prefix + filename);

        if (!get_config()->quiet) {
          ngtcp2::debug::log_printf(nullptr, "will save data to %s\n",
                                    full_path.c_str());
        }

        // open file for the path
        if (new_req->output) {
          new_req->output->open(full_path, std::ios::binary);
        }
      }

      build_and_submit_http3_request(new_req, stream_id);
    }
    http3_->send_data();
  }
}

void InnerLayer::build_and_submit_http3_request(HttpRequest *req,
                                                int64_t stream_id) {

  std::array<nghttp3_nv, 5> nva{
      ngtcp2::util::make_nv_nn(":method", req->method),
      ngtcp2::util::make_nv_nn(":scheme", req->scheme),
      ngtcp2::util::make_nv_nn(":authority", req->authority),
      ngtcp2::util::make_nv_nn(":path", req->path),
      ngtcp2::util::make_nv_nn("user-agent", "nghttp3/ngtcp2 client"),
  };

  http3_->submit_http_request(req, stream_id, nva.data(), nva.size());
}

void InnerLayer::set_tunnel_mode(bool is_tunnel_mode) {
  is_in_tunnel_mode_ = is_tunnel_mode;
}

OuterLayer *InnerLayer::get_tunnel_layer() const {
  return client_->get_outer_layer();
}

UdpLayer *InnerLayer::get_udp_layer() const {
  if (udp_layer_) {
    return udp_layer_;
  }

  return client_->get_outer_layer()->get_udp_layer();
}
