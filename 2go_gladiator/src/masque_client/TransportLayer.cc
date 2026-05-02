#include "TransportLayer.h"
#include "Http3Connection.h"
#include "MasqueClient.h"
#include "QuicConnection.h"
#include "debug.h"

TransportLayer::TransportLayer(MasqueClient *client,
                               const char *remote_host_name,
                               const network::Address &local,
                               const network::Address &remote)
    : client_(client) {}

TransportLayer::~TransportLayer() {
  delete http3_;
  delete quic_;
}

QuicConnection *TransportLayer::get_quic() const { return quic_; }

Http3Connection *TransportLayer::get_http3() const { return http3_; }

struct ev_loop *TransportLayer::get_ev_loop() const {
  return client_->get_ev_loop();
}

const Config *TransportLayer::get_config() const {
  return client_->get_config();
}

UdpLayer *TransportLayer::get_udp_layer() const { return udp_layer_; }

MasqueClient *TransportLayer::get_client() const { return client_; }

void TransportLayer::start() { stopped_ = false; }
void TransportLayer::stop() {

  if (stopped_) {
    return;
  }

  stopped_ = true;
  if (http3_) {
    http3_->disconnect();
  }
  if (quic_) {
    quic_->disconnect(false);
  }
}

bool TransportLayer::is_outer_layer() const {
  return get_type() == LayerType::OUTER;
}

bool TransportLayer::is_inner_layer() const {
  return get_type() == LayerType::INNER;
}

bool TransportLayer::is_pvd_client() const {
  return get_type() == LayerType::PVD_CLIENT;
}

void TransportLayer::on_quic_handshake_completed() {

  if (nullptr == http3_) {
    constexpr uint64_t REQUIRED_STREAMS = 3;

    if (quic_->get_left_uidi_streams() < REQUIRED_STREAMS) {
      ngtcp2::debug::log_printf(
          nullptr, "%s peer does not allow 3 unidirectional streams.",
          get_layer_str());
      quic_->disconnect(false);
    }

    int64_t stream_ids[REQUIRED_STREAMS];
    for (uint64_t i = 0; i < REQUIRED_STREAMS; i++) {
      if (auto rv = quic_->open_uni_stream(stream_ids[i]); rv != 0) {
        ngtcp2::debug::log_printf(nullptr,
                                  "%s Fail to open unidirectional stream!\n",
                                  get_layer_str());
        quic_->disconnect(false);
      }
    }

    http3_ =
        new Http3Connection(this, stream_ids[0], stream_ids[1], stream_ids[2]);
  }
}

const char *TransportLayer::get_layer_str() const {
  if (is_inner_layer()) {
    return "Inner Layer";
  } else if (is_outer_layer()) {
    return "Tunnel Layer";
  } else if (is_pvd_client()) {
    return "PVD Client";
  }
  return "Unknown";
}

void TransportLayer::send_data() {

  if (stopped_) {
    return;
  }

  if (quic_ && !quic_->is_handshake_completed()) {
    quic_->handshake();
    quic_->update_timer();
    return;
  }

  if (http3_) {
    http3_->send_data();
    quic_->update_timer();
  }
}

void TransportLayer::on_stop_rx() {

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "%s Stop receiving data!\n",
                              get_layer_str());
  }

  if (udp_layer_) {
    udp_layer_->stop(UdpLayer::DIRECTION::RX);
  }
}

void TransportLayer::on_stop_tx() {
  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "%s Stop sending data!\n",
                              get_layer_str());
  }
  if (udp_layer_) {
    udp_layer_->stop(UdpLayer::DIRECTION::TX);
  }
  stopped_ = true;
}
