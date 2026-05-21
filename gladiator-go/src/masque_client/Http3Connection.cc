#include "Http3Connection.h"

#include "Capsule.h"
#include "HttpRequest.h"
#include "InnerLayer.h"
#include "MasqueClient.h"
#include "OuterLayer.h"
#include "PM.h"
#include "PvdClient.h"
#include "QuicConnection.h"
#include "TransportLayer.h"
#include "debug.h"
#include "util.h"
#include <cassert>
#include <cstdint>
#include <cstdio>

#include <fstream>
#include <iostream>
#include <nghttp3/nghttp3.h>
#include <ngtcp2/ngtcp2.h>

namespace {
int acked_stream_data_cb(nghttp3_conn *conn, int64_t stream_id,
                         uint64_t datalen, void *conn_user_data,
                         void *stream_user_data) {
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_acked_stream_data(stream_id, datalen, stream_user_data);
}
int stream_close_cb(nghttp3_conn *conn, int64_t stream_id,
                    uint64_t app_error_code, void *conn_user_data,
                    void *stream_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_stream_close(stream_id, app_error_code, stream_user_data);
}
int recv_data_cb(nghttp3_conn *conn, int64_t stream_id, const uint8_t *data,
                 size_t datalen, void *conn_user_data, void *stream_user_data) {
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_recv_data(stream_id, data, datalen, stream_user_data);
}
int deferred_consume_cb(nghttp3_conn *conn, int64_t stream_id, size_t consumed,
                        void *conn_user_data, void *stream_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);

  return http->on_deferred_consume(stream_id, consumed, stream_user_data);
}
int begin_headers_cb(nghttp3_conn *conn, int64_t stream_id,
                     void *conn_user_data, void *stream_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_begin_headers(stream_id, stream_user_data);
}
int recv_header_cb(nghttp3_conn *conn, int64_t stream_id, int32_t token,
                   nghttp3_rcbuf *name, nghttp3_rcbuf *value, uint8_t flags,
                   void *conn_user_data, void *stream_user_data) {

  // ngtcp2::debug::print_http_header(stream_id, name, value, flags);
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_recv_header(stream_id, token, name, value, flags,
                              stream_user_data);
  return 0;
}
int end_headers_cb(nghttp3_conn *conn, int64_t stream_id, int fin,
                   void *conn_user_data, void *stream_user_data) {

  // ngtcp2::debug::print_http_end_headers(stream_id);
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_end_headers(stream_id, fin, stream_user_data);
  return 0;
}
int end_stream_cb(nghttp3_conn *conn, int64_t stream_id, void *conn_user_data,
                  void *stream_user_data) {
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);

  return http->on_end_stream(stream_id, stream_user_data);
}
int stop_sending_cb(nghttp3_conn *conn, int64_t stream_id,
                    uint64_t app_error_code, void *conn_user_data,
                    void *stream_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);

  return http->on_stop_sending(stream_id, app_error_code, stream_user_data);
}
int reset_stream_cb(nghttp3_conn *conn, int64_t stream_id,
                    uint64_t app_error_code, void *conn_user_data,
                    void *stream_user_data) {
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);

  return http->on_reset_stream(stream_id, app_error_code, stream_user_data);
}
int shutdown_cb(nghttp3_conn *conn, int64_t id, void *conn_user_data) {
  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_shutdown(id);
}
int recv_settings_cb(nghttp3_conn *conn, const nghttp3_settings *settings,
                     void *conn_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  return http->on_recv_settings(settings);
}

nghttp3_ssize build_client_cid_registration_cb(
    nghttp3_conn *conn, int64_t stream_id, nghttp3_vec *vec, size_t veccnt,
    uint32_t *pflags, void *conn_user_data, void *stream_user_data) {

  Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  HttpRequest *req = static_cast<HttpRequest *>(stream_user_data);

  if (!http->is_data_ready()) {
    return NGHTTP3_ERR_WOULDBLOCK;
  }

  vec[0].base = req->data_body;
  vec[0].len = req->data_len;

  if (req->data_end) {
    *pflags |= NGHTTP3_DATA_FLAG_EOF;
  }

  http->pause_data_read();
  return 1;
}

nghttp3_ssize tunnel_data_cb(nghttp3_conn *conn, int64_t stream_id,
                             nghttp3_vec *vec, size_t veccnt, uint32_t *pflags,
                             void *conn_user_data, void *stream_user_data) {

  // Http3Connection *http = static_cast<Http3Connection *>(conn_user_data);
  // HttpRequest *req = static_cast<HttpRequest *>(stream_user_data);

  return NGHTTP3_ERR_WOULDBLOCK;
}

} // namespace

Http3Connection::Http3Connection(TransportLayer *owner, int64_t ctrl_stream_id,
                                 int64_t qpack_enc_stream_id,
                                 int64_t qpack_dec_stream_id)
    : owner_(owner) {

  nghttp3_settings settings;
  nghttp3_settings_default(&settings);
  settings.qpack_max_dtable_capacity = 4_k;
  settings.qpack_blocked_streams = 100;

  // if (owner_->is_outer_layer()) {
  settings.h3_datagram = 1;
  // }

  auto mem = nghttp3_mem_default();

  nghttp3_callbacks callbacks{.acked_stream_data = acked_stream_data_cb,
                              .stream_close = stream_close_cb,
                              .recv_data = recv_data_cb,
                              .deferred_consume = deferred_consume_cb,
                              .begin_headers = begin_headers_cb,
                              .recv_header = recv_header_cb,
                              .end_headers = end_headers_cb,
                              .stop_sending = stop_sending_cb,
                              .end_stream = end_stream_cb,
                              .reset_stream = reset_stream_cb,
                              .shutdown = shutdown_cb,
                              .recv_settings = recv_settings_cb};

  if (auto rv = nghttp3_conn_client_new(&http3_conn_, &callbacks, &settings,
                                        mem, this);
      rv != 0) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s nghttp3_conn_client_new: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), nghttp3_strerror(rv));

    owner_->get_quic()->disconnect(false);
    return;
  }

  if (auto rv = nghttp3_conn_bind_control_stream(http3_conn_, ctrl_stream_id);
      rv != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s nghttp3_conn_bind_control_stream: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        nghttp3_strerror(rv));

    owner_->get_quic()->disconnect(false);
    return;
  }

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%s] %s http: control stream=%lu",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), ctrl_stream_id);
  }

  if (auto rv = nghttp3_conn_bind_qpack_streams(
          http3_conn_, qpack_enc_stream_id, qpack_dec_stream_id);
      rv != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s nghttp3_conn_bind_qpack_streams: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        nghttp3_strerror(rv));

    owner_->get_quic()->disconnect(false);
    return;
  }

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s http: QPACK streams encoder=%ld decoder=%ld",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        qpack_enc_stream_id, qpack_dec_stream_id);
  }
}

Http3Connection::~Http3Connection() {
  nghttp3_conn_del(http3_conn_);
  for (auto req : requests_) {
    delete req;
  }
  requests_.clear();
  delete close_wait_timer_;
}

void Http3Connection::send_data() {

  if (not started_ or owner_->is_stopped()) {
    return;
  }

  auto quic = owner_->get_quic();

  if (not quic->is_handshake_completed() or
      owner_->get_udp_layer() == nullptr) {
    return;
  }

  auto max_udp_payload_size =
      owner_->get_udp_layer()->get_max_udp_payload_size();

  auto &stream_data_ref = quic->get_stream_data_ref();

  auto ts = ngtcp2::util::timestamp();

  int64_t stream_id = -1;
  int fin = 0;

  auto quic_conn = quic->get_quic_conn();
  auto udp = owner_->get_udp_layer();
  auto tx_buffer = udp->get_tx_buffer();

  uint8_t *dest_buf = tx_buffer->data();

  // for inner layer's tunnel mode, the tx_buffer is outer layer's buffer
  OuterLayer *outer_layer = nullptr;
  int extra_bytes = 0; // for stream-id, context id
  if (owner_->is_inner_layer()) {
    outer_layer = owner_->get_client()->get_outer_layer();
    if (outer_layer) {

      int64_t masque_stream_id =
          static_cast<InnerLayer *>(owner_)->get_masque_stream_id();

      if (outer_layer->is_quic_forwarding_tx_ready(masque_stream_id) == false) {
        tx_buffer = outer_layer->get_tx_buffer();
        dest_buf = tx_buffer->data();
        // store the stream_id / 4;
        auto left_space = tx_buffer->size();
        auto write_bytes = ngtcp2::util::encode_var_len_integer(
            masque_stream_id >> 2, dest_buf, left_space);
        left_space -= write_bytes;
        dest_buf += write_bytes;
        extra_bytes += write_bytes;

        // store the context id
        const uint64_t context_id = 0;
        write_bytes = ngtcp2::util::encode_var_len_integer(context_id, dest_buf,
                                                           left_space);
        left_space -= write_bytes;
        dest_buf += write_bytes;
        extra_bytes += write_bytes;

        // make room for outer layer
        max_udp_payload_size = udp->get_max_udp_payload_size() - 128;
      }
    }
  }

  int64_t pkts_allowed_to_send =
      quic->get_send_quantum() / max_udp_payload_size;

  do {

    // ngtcp2::debug::log_printf(nullptr, "send pkt!");
    ssize_t vec_num = 0;
    if (quic->get_max_data_left() > 0) {
      vec_num = nghttp3_conn_writev_stream(
          http3_conn_, &stream_id, &fin,
          reinterpret_cast<nghttp3_vec *>(stream_data_ref.vec_array.data()),
          stream_data_ref.vec_array.size());
      if (vec_num < 0) {
        ngtcp2::debug::log_printf(
            nullptr, "Client[%s] %s nghttp3_conn_writev_stream: %s",
            owner_->get_client_id().c_str(), owner_->get_layer_str(),
            nghttp3_strerror(vec_num));

        quic->disconnect(false);
        return;
      }
      if (!owner_->get_config()->quiet) {
        ngtcp2::debug::log_printf(
            nullptr, "Client[%s] %s nghttp3_conn_writev_stream: %ld",
            owner_->get_client_id().c_str(), owner_->get_layer_str(), vec_num);
        for (auto i = 0; i < vec_num; i++) {
          auto &vec = stream_data_ref.vec_array[i];
          ngtcp2::debug::print_stream_data(stream_id, vec.base, vec.len);
        }
      }
    }

    uint32_t flags = NGTCP2_WRITE_STREAM_FLAG_MORE;
    if (fin) {
      flags |= NGTCP2_WRITE_STREAM_FLAG_FIN;
    }

    ngtcp2_pkt_info pi;
    ngtcp2_path_storage ps;
    ngtcp2_path_storage_zero(&ps);

    ssize_t ndatalen = 0;

    auto nwrite = ngtcp2_conn_writev_stream(
        quic_conn, &ps.path, &pi, dest_buf, max_udp_payload_size, &ndatalen,
        flags, stream_id, stream_data_ref.vec_array.data(), vec_num, ts);

    if (nwrite < 0) {

      switch (nwrite) {
      case NGTCP2_ERR_CLOSING:
        // ngtcp2::debug::log_printf(
        //     nullptr, "ngtcp2_conn_write_stream: %ld, %ld, %ld %s %ld",
        //     stream_id, vec_num, nwrite, ngtcp2_strerror(nwrite), ndatalen);
        // quic->disconnect(false);
        return;

      case NGTCP2_ERR_STREAM_DATA_BLOCKED:
        assert(ndatalen == -1);
        nghttp3_conn_block_stream(http3_conn_, stream_id);
        continue;
      case NGTCP2_ERR_STREAM_SHUT_WR:
        assert(ndatalen == -1);
        nghttp3_conn_shutdown_stream_write(http3_conn_, stream_id);
        continue;
      case NGTCP2_ERR_WRITE_MORE:
        assert(ndatalen >= 0);
        if (auto rv =
                nghttp3_conn_add_write_offset(http3_conn_, stream_id, ndatalen);
            rv != 0) {
          ngtcp2::debug::log_printf(
              nullptr, "Client[%s] %s nghttp3_conn_add_write_offset: %s",
              owner_->get_client_id().c_str(), owner_->get_layer_str(),
              nghttp3_strerror(rv));

          quic->disconnect(false);
        }
        continue;

      default:
        pkts_allowed_to_send = 0;
        continue;
      }

      assert(ndatalen == -1);

      ngtcp2::debug::log_printf(
          nullptr, "Client[%s] %s ngtcp2_conn_write_stream: other error %s",
          owner_->get_client_id().c_str(), owner_->get_layer_str(),
          ngtcp2_strerror(nwrite));

      quic->disconnect(false);

    } else if (ndatalen >= 0) {
      if (auto rv =
              nghttp3_conn_add_write_offset(http3_conn_, stream_id, ndatalen);
          rv != 0) {
        ngtcp2::debug::log_printf(
            nullptr, "Client[%s] %s nghttp3_conn_add_write_offset: %s",
            owner_->get_client_id().c_str(), owner_->get_layer_str(),
            nghttp3_strerror(rv));
        quic->disconnect(false);
      }
    }

    if (nwrite == 0) {
      // We are congestion limited.
      ngtcp2_conn_update_pkt_tx_time(quic_conn, ts);
      break;

    } else {

      if (outer_layer) {

        auto data = tx_buffer->data();
        bool is_short_header = (data[0] & 0x80) == 0;

        auto masque_stream_id =
            static_cast<InnerLayer *>(owner_)->get_masque_stream_id();

        if (outer_layer->is_quic_forwarding_tx_ready(masque_stream_id) &&
            is_short_header) {
          tx_buffer->set_data_len(nwrite);
          outer_layer->quic_aware_mode_egress_forward(masque_stream_id);
        } else {
          tx_buffer->set_data_len(nwrite + extra_bytes);
          outer_layer->tunnel_mode_egress_forward();
        }
      } else {
        if (!owner_->get_config()->quiet) {
          ngtcp2::debug::log_printf(
              nullptr, "Client[%s] %s Http3Connection::send_data(%lu)",
              owner_->get_client_id().c_str(), owner_->get_layer_str(), nwrite);
        }

        tx_buffer->set_data_len(nwrite);
        udp->send_data();
      }

      ngtcp2_conn_update_pkt_tx_time(quic_conn, ts);
    }

  } while (--pkts_allowed_to_send > 0);
}

void Http3Connection::start() { started_ = true; }
void Http3Connection::stop() {}

// callback functions
int Http3Connection::on_acked_stream_data(int64_t stream_id, uint64_t datalen,
                                          void *stream_user_data) {
  // ngtcp2::debug::log_printf(nullptr,
  //                           "=======================Http3Connection::on_acked_"
  //                           "stream_data(%ld,%lu,%p)",
  //                           stream_id, datalen, stream_user_data);

  HttpRequest *req = reinterpret_cast<HttpRequest *>(stream_user_data);
  req->acked_len += datalen;

  // std::cout << *req << std::endl;
  //  release the memory only if all of the data have been acked
  if (req->acked_len >= req->data_len) {
    req->data_len = 0;
    req->acked_len = 0;
  }

  return 0;
}
int Http3Connection::on_stream_close(int64_t stream_id, uint64_t app_error_code,
                                     void *stream_user_data) {

  auto quic_conn = owner_->get_quic()->get_quic_conn();

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s "
                              "Http3Connection::on_stream_close (%ld, %p)",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), stream_id,
                              stream_user_data);
  }

  if (ngtcp2_is_bidi_stream(stream_id)) {
    assert(ngtcp2_conn_is_local_stream(quic_conn, stream_id));

    owner_->get_quic()->on_upper_layer_stream_closed(stream_id);
  } else {
    assert(!ngtcp2_conn_is_local_stream(quic_conn, stream_id));
    ngtcp2_conn_extend_max_streams_uni(quic_conn, 1);
  }

  return 0;
}
int Http3Connection::on_recv_data(int64_t stream_id, const uint8_t *data,
                                  size_t datalen, void *stream_user_data) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s "
                              "Http3Connection::on_recv_data(datalen=%lu)",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), datalen);
  }

  if (owner_->is_pvd_client()) {
    HttpRequest *req = reinterpret_cast<HttpRequest *>(stream_user_data);
    owner_->get_client()->get_pm()->inc_user_rx_bytes(datalen);
    if (req) {
      if (!owner_->get_config()->quiet) {
        ngtcp2::debug::log_printf(nullptr, "Client[%s] %s data length:%ld",
                                  owner_->get_client_id().c_str(),
                                  "pvd client on_recv_data", datalen);
      }
      auto *pvd_client = static_cast<PvdClient *>(owner_);
      pvd_client->on_recv_data(data, datalen);
    }
  } else if (owner_->is_inner_layer()) {

    HttpRequest *req = reinterpret_cast<HttpRequest *>(stream_user_data);
    if (req && req->output) {
      req->output->write(reinterpret_cast<const char *>(data), datalen);
    }
    owner_->get_client()->get_pm()->inc_user_rx_bytes(datalen);
  } else if (owner_->is_outer_layer()) {

    OuterLayer *outer_layer = static_cast<OuterLayer *>(owner_);
    if (outer_layer->is_quic_forwarding_supported(stream_id) &&
        !outer_layer->is_quic_forwarding_ready(stream_id)) {
      // outer layer is during QUIC forwarding negotiation stage.
      // it's the negotiation message.
      outer_layer->on_recv_capsule_data(data, datalen, stream_id);

    } else {

      // inform inner layer, the datagram is ready to read!
      outer_layer->get_inner_layer(stream_id >> 2)
          ->on_tunnel_data_ready_to_read(data, datalen);
    }
  }

  auto quic_conn = owner_->get_quic()->get_quic_conn();
  ngtcp2_conn_extend_max_stream_offset(quic_conn, stream_id, datalen);
  ngtcp2_conn_extend_max_offset(quic_conn, datalen);

  return 0;
}
int Http3Connection::on_deferred_consume(int64_t stream_id, size_t consumed,
                                         void *stream_user_data) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connection::on_deferred_consume",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }

  auto quic_conn = owner_->get_quic()->get_quic_conn();
  ngtcp2_conn_extend_max_stream_offset(quic_conn, stream_id, consumed);
  ngtcp2_conn_extend_max_offset(quic_conn, consumed);
  return 0;
}
int Http3Connection::on_begin_headers(int64_t stream_id,
                                      void *stream_user_data) {

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::print_http_begin_response_headers(stream_id);
  }

  return 0;
}
int Http3Connection::on_recv_header(int64_t stream_id, int32_t token,
                                    nghttp3_rcbuf *name, nghttp3_rcbuf *value,
                                    uint8_t flags, void *stream_user_data) {

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::print_http_header(stream_id, name, value, flags);
  }

  if (owner_->is_outer_layer()) {
    OuterLayer *outer_layer = static_cast<OuterLayer *>(owner_);
    outer_layer->on_response_header_recv(name, value, stream_id);
  } else if (owner_->is_pvd_client()) {
    PvdClient *pvd_client = static_cast<PvdClient *>(owner_);
    pvd_client->on_response_header_recv(name, value, stream_id);
  }

  return 0;
}
int Http3Connection::on_end_headers(int64_t stream_id, int fin,
                                    void *stream_user_data) {

  return 0;
}
int Http3Connection::on_end_stream(int64_t stream_id, void *stream_user_data) {

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connection::on_end_stream (%ld, %p)",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), stream_id,
        stream_user_data);
  }

  auto req = reinterpret_cast<HttpRequest *>(stream_user_data);
  if (req->output && req->output->is_open()) {
    req->output->close();
  }
  if (owner_->is_pvd_client()) {
    PvdClient *pvd_client = static_cast<PvdClient *>(owner_);
    pvd_client->on_recv_data_completed();
  }
  return 0;
}
int Http3Connection::on_stop_sending(int64_t stream_id, uint64_t app_error_code,
                                     void *stream_user_data) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr,
                              "client[%s]=======================%s "
                              "Http3Connection::on_stop_sending",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str());
  }

  auto quic = owner_->get_quic();

  if (quic->stop_sending(stream_id, app_error_code) != 0) {
    owner_->on_stop_rx();
    return NGHTTP3_ERR_CALLBACK_FAILURE;
  }
  owner_->on_stop_rx();
  return 0;
}
int Http3Connection::on_reset_stream(int64_t stream_id, uint64_t app_error_code,
                                     void *stream_user_data) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connection::on_reset_stream",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }

  auto quic = owner_->get_quic();

  if (quic->reset_stream(stream_id, app_error_code) != 0) {
    owner_->on_stop_tx();
    return NGHTTP3_ERR_CALLBACK_FAILURE;
  }

  send_data(); // drain out frames before stop tx.
  owner_->on_stop_tx();
  return 0;
}
int Http3Connection::on_shutdown(int64_t id) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connection::on_shutdown(%lu)",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), id);
  }
  return 0;
}
int Http3Connection::on_recv_settings(const nghttp3_settings *settings) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connection::on_recv_settings",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }
  return 0;
}

int Http3Connection::add_acked_stream_offset(int64_t stream_id,
                                             size_t data_len) {
  if (auto rv = nghttp3_conn_add_ack_offset(http3_conn_, stream_id, data_len);
      rv != 0) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s nghttp3_conn_add_ack_offset: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), nghttp3_strerror(rv));

    return -1;
  }

  return 0;
}

ssize_t Http3Connection::read_stream_data(int64_t stream_id,
                                          const uint8_t *data, size_t datalen,
                                          uint32_t flags) {
  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr,
        "Client[%s] %s Http3Connection::read_stream_data(%ld,%p,%lu,%u)",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), stream_id,
        data, datalen, flags);
  }

  auto nconsumed =
      nghttp3_conn_read_stream(http3_conn_, stream_id, data, datalen,
                               //  flags);
                               flags & NGTCP2_STREAM_DATA_FLAG_FIN);
  if (nconsumed < 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s nghttp3_conn_read_stream: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        nghttp3_strerror(nconsumed));

    return 0;
  }

  return nconsumed;
}

int Http3Connection::submit_http_request(HttpRequest *req, int64_t stream_id,
                                         nghttp3_nv *nv, size_t nvlen) {

  // submit a request with this stream_id
  auto config = owner_->get_client()->get_config();

  nghttp3_data_reader data_reader{};
  nghttp3_data_reader *pdata_reader = nullptr;
  req->stream_id = stream_id;

  if (!config->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3 submit request with stream-id: %ld",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), stream_id);
  }

  if (owner_->is_outer_layer()) {
    if (config->request_quic_forwarding) {
      data_reader.read_data = build_client_cid_registration_cb;
    } else {
      data_reader.read_data = tunnel_data_cb;
    }
    pdata_reader = &data_reader;
  }

  auto rv =
      nghttp3_conn_submit_request(http3_conn_, stream_id, nv, nvlen,
                                  pdata_reader, const_cast<HttpRequest *>(req));

  if (!config->quiet) {
    ngtcp2::debug::print_http_request_headers(stream_id, nv, nvlen);
  }

  if (rv != 0) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s nghttp3_conn_submit_request: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), nghttp3_strerror(rv));
  }

  return rv;
}

HttpRequest *Http3Connection::new_request(const HttpRequest &request) {
  HttpRequest *req = new HttpRequest();
  req->method = request.method;
  req->scheme = request.scheme;
  req->authority = request.authority;
  req->path = request.path;
  req->protocol = request.protocol;
  req->host_name = request.host_name;
  req->port = request.port;
  req->accept = request.accept;
  req->output = new std::ofstream();
  // deep copy
  if (request.data_len > 0) {
    memcpy(req->data_body, request.data_body,
           std::min(MAX_HTTP3_BODY_DATA_LEN, request.data_len));
    req->data_len = request.data_len;
  }
  requests_.push_back(req);
  return req;
}

void Http3Connection::disconnect() {

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s Http3Connect::disconnect()",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }

  nghttp3_conn_submit_shutdown_notice(http3_conn_);
  if (close_wait_timer_ == nullptr) {
    close_wait_timer_ = new TimerWatcher(owner_->get_ev_loop(), this);
    close_wait_timer_->reset_timer(0.01); // wait 10ms
  }
}

void Http3Connection::on_timeout(TimerWatcher *watcher) {
  watcher->stop();

  if (!owner_->get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%s] %s Http3Connect Shut Down",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str());
  }

  nghttp3_conn_shutdown(http3_conn_);

  delete close_wait_timer_;

  close_wait_timer_ = nullptr;

  owner_->get_client()->on_http3_closed(owner_);
}

void Http3Connection::pause_data_read() { data_ready_ = false; }

void Http3Connection::resume_data_read(int64_t stream_id) {
  data_ready_ = true;
  auto rv = nghttp3_conn_resume_stream(http3_conn_, stream_id);
  if (rv != 0) {
    ngtcp2::debug::log_printf(nullptr, "Client[%s] %s resume_data_read: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), nghttp3_strerror(rv));
  }
}

bool Http3Connection::is_data_ready() const { return data_ready_; }

HttpRequest *Http3Connection::get_request(int64_t stream_id) {
  for (auto req : requests_) {
    if (req->stream_id == stream_id) {
      return req;
    }
  }

  return nullptr;
}
