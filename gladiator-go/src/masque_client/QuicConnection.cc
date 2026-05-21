#include "QuicConnection.h"
#include "Http3Connection.h"
#include "InnerLayer.h"
#include "MasqueClient.h"
#include "OuterLayer.h"
#include "PM.h"
#include "TransportLayer.h"
#include "debug.h"
#include "util.h"
#include <algorithm>
#include <cassert>
#include <cstdint>
#include <ngtcp2/ngtcp2.h>

// Define the callback functions
namespace {
ngtcp2_conn *get_conn(ngtcp2_crypto_conn_ref *conn_ref) {
  auto c = static_cast<QuicConnection *>(conn_ref->user_data);
  return c->get_quic_conn();
}

int handshake_completed_cb(ngtcp2_conn *conn, void *user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);

  return quic->on_handshake_completed();
}

int recv_stream_data_cb(ngtcp2_conn *conn, uint32_t flags, int64_t stream_id,
                        uint64_t offset, const uint8_t *data, size_t datalen,
                        void *user_data, void *stream_user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  return quic->on_recv_stream_data(flags, stream_id, offset, data, datalen,
                                   stream_user_data);
}

int acked_stream_data_offset_cb(ngtcp2_conn *conn, int64_t stream_id,
                                uint64_t offset, uint64_t datalen,
                                void *user_data, void *stream_user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  return quic->on_acked_stream_data_offset(stream_id, offset, datalen,
                                           stream_user_data);
}

void rand_cb(uint8_t *dest, size_t destlen, const ngtcp2_rand_ctx *rand_ctx) {
  auto dis = std::uniform_int_distribution<uint8_t>();
  std::random_device seed;
  auto engine = std::mt19937(seed());
  std::generate(dest, dest + destlen,
                [&engine, &dis]() { return dis(engine); });
}

int get_new_connection_id_cb(ngtcp2_conn *conn, ngtcp2_cid *cid, uint8_t *token,
                             size_t cidlen, void *user_data) {
  if (ngtcp2::util::generate_secure_random(cid->data, cidlen) != 0) {
    return NGTCP2_ERR_CALLBACK_FAILURE;
  }

  auto quic = static_cast<QuicConnection *>(user_data);
  auto config = quic->get_config();

  cid->datalen = cidlen;
  if (ngtcp2_crypto_generate_stateless_reset_token(
          token, config->static_secret.data(), config->static_secret.size(),
          cid) != 0) {
    return NGTCP2_ERR_CALLBACK_FAILURE;
  }

  return 0;
}

int extend_max_bidi_streams_cb(ngtcp2_conn *conn, uint64_t max_streams,
                               void *user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  return quic->on_extend_max_bidi_streams(max_streams);
}

int recv_datagram_cb(ngtcp2_conn *conn, uint32_t flags, const uint8_t *data,
                     size_t datalen, void *user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  return quic->on_recv_datagram(flags, data, datalen);
}

int ack_datagram_cb(ngtcp2_conn *conn, uint64_t dgram_id, void *user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  return quic->on_ack_datagram(dgram_id);
}

int key_update_cb(ngtcp2_conn *conn, uint8_t *rx_secret, uint8_t *tx_secret,
                  ngtcp2_crypto_aead_ctx *rx_aead_ctx, uint8_t *rx_iv,
                  ngtcp2_crypto_aead_ctx *tx_aead_ctx, uint8_t *tx_iv,
                  const uint8_t *current_rx_secret,
                  const uint8_t *current_tx_secret, size_t secretlen,
                  void *user_data) {
  auto quic = static_cast<QuicConnection *>(user_data);
  if (quic->on_key_update(rx_secret, tx_secret, rx_aead_ctx, rx_iv, tx_aead_ctx,
                          tx_iv, current_rx_secret, current_tx_secret,
                          secretlen)) {
    return NGTCP2_ERR_CALLBACK_FAILURE;
  }
  return 0;
}

int recv_crypto_data_cb(ngtcp2_conn *conn,
                        ngtcp2_encryption_level encryption_level,
                        uint64_t offset, const uint8_t *data, size_t datalen,
                        void *user_data) {

  // ngtcp2::debug::print_crypto_data(encryption_level, data, datalen);

  return ngtcp2_crypto_recv_crypto_data_cb(conn, encryption_level, offset, data,
                                           datalen, user_data);
}

int stream_close_cb(ngtcp2_conn *conn, uint32_t flags, int64_t stream_id,
                    uint64_t app_error_code, void *user_data,
                    void *stream_user_data) {

  auto quic = static_cast<QuicConnection *>(user_data);
  quic->on_stream_close(flags, stream_id, app_error_code, stream_user_data);
  return 0;
}

int stream_reset_cb(ngtcp2_conn *conn, int64_t stream_id, uint64_t final_size,
                    uint64_t app_error_code, void *user_data,
                    void *stream_user_data) {
  auto c = static_cast<QuicConnection *>(user_data);

  if (c->on_stream_reset(stream_id) != 0) {
    return NGTCP2_ERR_CALLBACK_FAILURE;
  }

  return 0;
}

void set_callback_functions(ngtcp2_callbacks &callback) {

  callback.client_initial = ngtcp2_crypto_client_initial_cb;

  callback.recv_client_initial = nullptr; // for server

  callback.recv_crypto_data = recv_crypto_data_cb;

  callback.handshake_completed =
      handshake_completed_cb; // QUIC cryptographic handshake completed

  callback.recv_version_negotiation = nullptr;

  callback.encrypt = ngtcp2_crypto_encrypt_cb;

  callback.decrypt = ngtcp2_crypto_decrypt_cb;

  callback.hp_mask = ngtcp2_crypto_hp_mask_cb;

  callback.recv_stream_data =
      recv_stream_data_cb; // rx stream data for application

  callback.acked_stream_data_offset =
      acked_stream_data_offset_cb; // inform application acked offset

  callback.stream_open = nullptr; // remote endpoint open stream

  callback.stream_close =
      stream_close_cb; // stream is closed; inform application

  callback.recv_stateless_reset = nullptr;

  callback.recv_retry = ngtcp2_crypto_recv_retry_cb;

  callback.extend_max_local_streams_bidi =
      extend_max_bidi_streams_cb; // inform application can open bidi streams

  callback.extend_max_local_streams_uni =
      nullptr; // inform app can open uidi stream

  callback.rand = rand_cb;

  callback.get_new_connection_id = get_new_connection_id_cb;

  callback.remove_connection_id =
      nullptr; // notify app the connection is not used by the remote

  callback.update_key = key_update_cb; // ngtcp2_crypto_update_key_cb;

  callback.path_validation = nullptr;

  callback.select_preferred_addr = nullptr;

  callback.stream_reset = nullptr; // stream reset by remote

  callback.extend_max_remote_streams_bidi = nullptr;

  callback.extend_max_remote_streams_uni = nullptr;

  callback.extend_max_stream_data =
      nullptr; // max offset of stream data can send is increased

  callback.dcid_status = nullptr; // destination connection id status changed

  callback.handshake_confirmed =
      nullptr; // both sides agree that handshake finished

  callback.recv_new_token = nullptr; // new token from server

  callback.delete_crypto_aead_ctx = ngtcp2_crypto_delete_crypto_aead_ctx_cb;

  callback.delete_crypto_cipher_ctx = ngtcp2_crypto_delete_crypto_cipher_ctx_cb;

  callback.recv_datagram = recv_datagram_cb; // Datagram frame is received

  callback.ack_datagram =
      ack_datagram_cb; // Datagram frame is acked by remote endpoint

  callback.lost_datagram = nullptr; // Datagram frame is declared lost

  callback.get_path_challenge_data = ngtcp2_crypto_get_path_challenge_data_cb;

  callback.stream_stop_sending =
      nullptr; // no longer reads from a stream before it receive all data

  callback.stream_reset = stream_reset_cb;

  callback.version_negotiation = ngtcp2_crypto_version_negotiation_cb;

  callback.recv_rx_key = nullptr;

  callback.recv_tx_key = nullptr;

  callback.tls_early_data_rejected = nullptr;
}

} // namespace

QuicConnection::QuicConnection(TransportLayer *owner, const char *remote_host,
                               const network::Address &local,
                               const network::Address &remote,
                               int max_udp_payload_size)
    : owner_(owner), local_addr_(local),
      remote_addr_(remote), conn_ref_{get_conn, this},
      config_(owner->get_config()) {

  idle_timer_watcher_ = new TimerWatcher(owner_->get_ev_loop(), this);
  ngtcp2_ccerr_default(&last_error_);
  ngtcp2_cid scid, dcid;

  scid.datalen = 17;
  ngtcp2::util::generate_secure_random(scid.data, scid.datalen);

  dcid.datalen = 18;
  ngtcp2::util::generate_secure_random(dcid.data, dcid.datalen);

  auto path = ngtcp2_path{
      {&local_addr_.su.sa, local_addr_.len},
      {&remote_addr_.su.sa, remote_addr_.len},
      nullptr,
  };

  ngtcp2_settings settings;
  ngtcp2_settings_default(&settings);

  settings.log_printf = config_->quiet ? nullptr : ngtcp2::debug::log_printf;

  settings.cc_algo = config_->cc_algo;
  settings.initial_ts = ngtcp2::util::timestamp();
  settings.initial_rtt = config_->initial_rtt;
  settings.max_window = config_->max_window;
  settings.max_stream_window = config_->max_stream_window;

  if (max_udp_payload_size) {
    settings.max_tx_udp_payload_size = max_udp_payload_size;
    settings.no_tx_udp_payload_size_shaping = 1;
  }

  settings.handshake_timeout = config_->handshake_timeout;
  settings.no_pmtud = config_->no_pmtud;
  settings.ack_thresh = config_->ack_thresh;
  if (config_->initial_pkt_num == UINT32_MAX) {
    auto dis = std::uniform_int_distribution<uint32_t>(0, INT32_MAX);
    auto engine = ngtcp2::util::make_mt19937();
    settings.initial_pkt_num = dis(engine);
  } else {
    settings.initial_pkt_num = config_->initial_pkt_num;
  }

  settings.original_version = config_->version;

  ngtcp2_transport_params params;
  ngtcp2_transport_params_default(&params);
  params.initial_max_stream_data_bidi_local =
      config_->max_stream_data_bidi_local;
  params.initial_max_stream_data_bidi_remote =
      config_->max_stream_data_bidi_remote;
  params.initial_max_stream_data_uni = config_->max_stream_data_uni;
  params.initial_max_data = config_->max_data;
  params.initial_max_streams_bidi = config_->max_streams_bidi;
  params.initial_max_streams_uni = config_->max_streams_uni;
  params.max_idle_timeout = config_->timeout;
  params.active_connection_id_limit = 8;
  params.grease_quic_bit = 1;

  // if (owner_->is_outer_layer()) {
  params.max_datagram_frame_size = 65535;
  //}
  if (owner_->is_inner_layer() && owner_->get_client()->get_outer_layer()) {
    params.max_udp_payload_size = max_udp_payload_size - 128;
  } else {
    params.max_udp_payload_size = max_udp_payload_size;
  }

  // define the callback functions
  auto callbacks = ngtcp2_callbacks{};

  set_callback_functions(callbacks);

  // create ngtcp2_conn object
  auto rv =
      ngtcp2_conn_client_new(&quic_conn_, &dcid, &scid, &path, config_->version,
                             &callbacks, &settings, &params, nullptr, this);

  if (rv != 0) {
    ngtcp2::debug::log_printf(nullptr, "ngtcp2_conn_client_new: %s\n",
                              ngtcp2_strerror(rv));
    return;
  }
  if (settings.log_printf) {
    settings.log_printf(
        this, "Client[%s] %s QuicConnection object is created! remote: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        remote_addr_.to_string().c_str());
  }

  bool early_data = false;

  if (tls_session_.init(early_data, *config_->tls_context, remote_host,
                        conn_ref(), config_->version, network::AppProtocol::H3,
                        *config_) != 0) {
    return;
  }

  ngtcp2_conn_set_tls_native_handle(quic_conn_,
                                    tls_session_.get_native_handle());
}

QuicConnection::~QuicConnection() {
  if (quic_conn_) {
    ngtcp2_conn_del(quic_conn_);
  }
  idle_timer_watcher_->stop();
  delete idle_timer_watcher_;
}

void QuicConnection::on_timeout(TimerWatcher *watcher) {

  auto now = ngtcp2::util::timestamp();
  if (auto rv = ngtcp2_conn_handle_expiry(quic_conn_, now); rv != 0) {
    if (!config_->quiet) {
      ngtcp2::debug::log_printf(nullptr,
                                "client[%s] %s ngtcp2_conn_handle_expiry: %s\n",
                                owner_->get_client_id().c_str(),
                                owner_->get_layer_str(), ngtcp2_strerror(rv));
    }

    ngtcp2_ccerr_set_liberr(&last_error_, rv, nullptr, 0);
    owner_->get_client()->on_quic_timeout(handshake_completed_, owner_);
    disconnect(true);
    return;
  }

  if (!handshake_completed_) {
    handshake();
  } else {
    auto http3 = owner_->get_http3();
    if (http3) {
      http3->send_data();
    }
  }

  update_timer();
};

void QuicConnection::on_udp_data_ready_to_read(const uint8_t *buffer,
                                               size_t data_len) {

  recv_data(buffer, data_len);
  auto http3 = owner_->get_http3();
  if (http3) {
    http3->send_data();
  }
}

const Config *QuicConnection::get_config() const { return config_; }

int QuicConnection::on_handshake_completed() {
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "client[%s] %s QUIC handshake completed!",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }

  owner_->on_quic_handshake_completed();
  handshake_completed_ = true;
  if (owner_->is_inner_layer()) {
    auto pm = owner_->get_client()->get_pm();
    pm->inc_created_connects();

    // inform tunnel layer if it exists
    auto tunnel = owner_->get_client()->get_outer_layer();
    if (tunnel) {
      auto masque_stream_id =
          static_cast<InnerLayer *>(owner_)->get_masque_stream_id();
      tunnel->on_inner_quic_handshake_completed(this, masque_stream_id);
    }

  } else if (owner_->is_outer_layer()) {
    auto pm = owner_->get_client()->get_pm();
    pm->inc_tunnel_created_connects();
  } else if (owner_->is_pvd_client()) {
    auto pm = owner_->get_client()->get_pm();
    pm->inc_created_connects();
  }

  return 0;
}

int QuicConnection::on_recv_stream_data(uint32_t flags, int64_t stream_id,
                                        uint64_t offset, const uint8_t *data,
                                        size_t datalen,
                                        void *stream_user_data) {
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s on_recv_stream_data for "
                              "stream:%ld, flags:%u stream_data:%p",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), stream_id, flags,
                              stream_user_data);
  }

  if (!config_->quiet) {
    ngtcp2::debug::print_stream_data(stream_id, data, datalen);
  }

  auto http3 = owner_->get_http3();

  if (http3) {

    auto consumed_bytes =
        http3->read_stream_data(stream_id, data, datalen, flags);

    ngtcp2_conn_extend_max_stream_offset(quic_conn_, stream_id, consumed_bytes);
    ngtcp2_conn_extend_max_offset(quic_conn_, consumed_bytes);
  }

  return 0;
}

int QuicConnection::on_acked_stream_data_offset(int64_t stream_id,
                                                uint64_t offset,
                                                uint64_t datalen,
                                                void *stream_user_data) {
  // ngtcp2::debug::log_printf(
  //     nullptr,
  //     "======================= on_acked_stream_data_offset! stream_id:%ld "
  //     "offset:%lu datalen:%lu",
  //     stream_id, offset, datalen);

  auto http3 = owner_->get_http3();

  if (auto rv =
          nghttp3_conn_add_ack_offset(http3->get_conn(), stream_id, datalen);
      rv != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s ERROR nghttp3_conn_add_ack_offset: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        nghttp3_strerror(rv));

    return -1;
  }

  return 0;
}

int QuicConnection::on_extend_max_bidi_streams(uint64_t max_streams) {

  owner_->on_quic_streams_can_be_opened(max_streams);

  return 0;
}

int QuicConnection::on_recv_datagram(uint32_t flags, const uint8_t *data,
                                     size_t datalen) {
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%s] %s on_recv_datagram!",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str());
  }

  // for tunnel layer, the datagram should be forward to inner layer.
  if (owner_->is_outer_layer()) {
    OuterLayer *tunnel = static_cast<OuterLayer *>(owner_);
    tunnel->tunnel_mode_ingress_forward(flags, data, datalen);
  }

  return 0;
}

int QuicConnection::on_ack_datagram(uint64_t dgram_id) {
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%s] %s on_ack_datagram: %lu!",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), dgram_id);
  }
  return 0;
}

int QuicConnection::on_key_update(
    uint8_t *rx_secret, uint8_t *tx_secret, ngtcp2_crypto_aead_ctx *rx_aead_ctx,
    uint8_t *rx_iv, ngtcp2_crypto_aead_ctx *tx_aead_ctx, uint8_t *tx_iv,
    const uint8_t *current_rx_secret, const uint8_t *current_tx_secret,
    size_t secretlen) {
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Updating traffic key\n");
  }

  auto crypto_ctx = ngtcp2_conn_get_crypto_ctx(quic_conn_);
  auto aead = &crypto_ctx->aead;
  auto keylen = ngtcp2_crypto_aead_keylen(aead);
  auto ivlen = ngtcp2_crypto_packet_protection_ivlen(aead);

  //++nkey_update_;

  std::array<uint8_t, 64> rx_key, tx_key;

  if (ngtcp2_crypto_update_key(quic_conn_, rx_secret, tx_secret, rx_aead_ctx,
                               rx_key.data(), rx_iv, tx_aead_ctx, tx_key.data(),
                               tx_iv, current_rx_secret, current_tx_secret,
                               secretlen) != 0) {
    return -1;
  }

  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "application_traffic rx secret ");
    ngtcp2::debug::print_secrets(rx_secret, secretlen, rx_key.data(), keylen,
                                 rx_iv, ivlen);
    ngtcp2::debug::log_printf(nullptr, "application_traffic tx secret ");
    ngtcp2::debug::print_secrets(tx_secret, secretlen, tx_key.data(), keylen,
                                 tx_iv, ivlen);
  }

  return 0;
}

int QuicConnection::on_stream_close(uint32_t flags, int64_t stream_id,
                                    uint64_t app_error_code,
                                    void *stream_user_data) {

  auto http3_conn = owner_->get_http3()->get_conn();

  if (app_error_code == 0) {
    app_error_code = NGHTTP3_H3_NO_ERROR;
  }
  if (!config_->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s nghttp3_conn_close_stream",
        owner_->get_client_id().c_str(), owner_->get_layer_str());
  }

  // will trigger http3's call back
  auto rv = nghttp3_conn_close_stream(http3_conn, stream_id, app_error_code);
  switch (rv) {
  case 0:
    break;
  case NGHTTP3_ERR_STREAM_NOT_FOUND:
    // We have to handle the case when stream opened but no data is
    // transferred.  In this case, nghttp3_conn_close_stream might
    // return error.
    if (!ngtcp2_is_bidi_stream(stream_id)) {
      assert(!ngtcp2_conn_is_local_stream(quic_conn_, stream_id));
      ngtcp2_conn_extend_max_streams_uni(quic_conn_, 1);
    }
    break;
  default:

    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s nghttp3_conn_close_stream: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), nghttp3_strerror(rv));

    ngtcp2_ccerr_set_application_error(
        &last_error_, nghttp3_err_infer_quic_app_error_code(rv), nullptr, 0);
    return -1;
  }
  return 0;
}

int QuicConnection::on_stream_reset(int64_t stream_id) {
  auto http3 = owner_->get_http3();
  if (http3) {
    if (auto rv =
            nghttp3_conn_shutdown_stream_read(http3->get_conn(), stream_id);
        rv != 0) {
      ngtcp2::debug::log_printf(
          nullptr, "Client[%s] %s nghttp3_conn_shutdown_stream_read: %s ",
          owner_->get_client_id().c_str(), owner_->get_layer_str(),
          nghttp3_strerror(rv));
      ngtcp2_ccerr_set_application_error(
          &last_error_, nghttp3_err_infer_quic_app_error_code(rv), nullptr, 0);
      return -1;
    }
  }
  return 0;
}

void QuicConnection::update_timer() {
  auto expiry = ngtcp2_conn_get_expiry(quic_conn_);
  auto now = ngtcp2::util::timestamp();

  if (expiry <= now) {
    // if (!config_->quiet) {
    //   auto t = static_cast<ev_tstamp>(now - expiry) / NGTCP2_SECONDS;
    //   ngtcp2::debug::log_printf(nullptr, "%s QUIC Timer has already expired:
    //   %f", t, owner_->get_layer_str());
    // }

    idle_timer_watcher_->awake_timer();

    return;
  }

  auto t = static_cast<ev_tstamp>(expiry - now) / NGTCP2_SECONDS;
  // if (!config_->quiet) {
  //  ngtcp2::debug::log_printf(nullptr, "reset timer as %f\n", t);
  //}
  idle_timer_watcher_->reset_timer(t);
}

void QuicConnection::write_datagram_data(FlatBuffer *buffer) {

  if (owner_->is_inner_layer()) {
    // inner layer only use STREAM!
    assert(false);
    return;
  }

  ngtcp2_pkt_info pi;
  ngtcp2_path_storage ps;
  int accepted;

  uint32_t flags = NGTCP2_WRITE_DATAGRAM_FLAG_NONE;
  auto ts = ngtcp2::util::timestamp();

  auto udp = owner_->get_udp_layer();

  auto tx_buffer = udp->get_tx_buffer();

  ngtcp2_path_storage_zero(&ps);

  do {

    auto ssize = ngtcp2_conn_write_datagram(
        quic_conn_, &ps.path, &pi, tx_buffer->data(), tx_buffer->size(),
        &accepted, flags, 0, buffer->data(), buffer->len(), ts);

    if (ssize > 0) {
      tx_buffer->set_data_len(ssize);
      if (!config_->quiet) {
        ngtcp2::debug::log_printf(
            nullptr,
            "Client[%s] %s QuicConnection::write_datagram_data(%lu) "
            "accepted:%d",
            owner_->get_client_id().c_str(), owner_->get_layer_str(), ssize,
            accepted);
      }

      udp->send_data();

      ngtcp2_conn_update_pkt_tx_time(quic_conn_, ts);

      if (accepted) {
        // the data in buffer is accepted, so quit from the loop
        break;
      } else {
        // Certain frame in QUIC layer is sent out.
        // In this case, need to call ngtcp2_conn_write_datagram again.
        if (!config_->quiet) {
          ngtcp2::debug::log_printf(
              nullptr,
              "Client[%s] %s QuicConnection::write_datagram_data(%lu) the "
              "non-datagram "
              "frame (e.g ACK) is sent.",
              owner_->get_client_id().c_str(), owner_->get_layer_str(), ssize);
        }
      }
    } else if (ssize == 0) {
      // No more data can be sent because of congestion control limit
      // It implies the Congestion Window (CWND) is full.
      // When ACKs are received, the CWND size will be reduced, and
      // consequently, congestion might be diminished. Instead of waiting here,
      // we must keep the thread running to have a chance to receive ACKs.
      // Therefore, we reschedule to write the datagram.
      owner_->get_client()->reschedule_write_datagram(this, buffer);
      if (!config_->quiet) {
        ngtcp2::debug::log_printf(
            nullptr,
            "Client[%s] %s QuicConnection::write_datagram_data() rescheduled.",
            owner_->get_client_id().c_str(), owner_->get_layer_str());
      }
      return;

    } else {
      if (!config_->quiet) {
        ngtcp2::debug::log_printf(
            nullptr,
            "Client[%s] %s QuicConnection::write_datagram_data() return code: "
            "%lu",
            owner_->get_client_id().c_str(), owner_->get_layer_str(), ssize);
      }
      break;
    }
  } while (true);
}

QuicConnection::Stream_Data &QuicConnection::get_stream_data_ref() {
  return stream_data_;
}

size_t QuicConnection::get_max_data_left() {
  // number of bytes that this local endpoint can send in this connection
  // without violating connection-level flow control.
  return ngtcp2_conn_get_max_data_left(quic_conn_);
}

size_t QuicConnection::get_send_quantum() {
  return ngtcp2_conn_get_send_quantum(quic_conn_);
}

size_t QuicConnection::recv_data(const uint8_t *buffer, size_t data_len) {

  if (!config_->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s QuicConnection::recv_data %lu bytes",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), data_len);
  }

  if (!all_streams_closed()) {
    update_timer();
  }

  // for path validation
  auto path = ngtcp2_path{
      {&local_addr_.su.sa, local_addr_.len},
      {&remote_addr_.su.sa, remote_addr_.len},
      nullptr,
  };

  ngtcp2_pkt_info pi;

  // If the ngtcp2_callbacks.recv_stream_data has been defined,
  // it will be called after ngtcp2_conn_read_pkt(...)
  // the upper layer protocol can read data in
  // ngtcp2_callbacks.recv_stream_data(...)
  if (auto rv = ngtcp2_conn_read_pkt(quic_conn_, &path, &pi, buffer, data_len,
                                     ngtcp2::util::timestamp());
      rv != 0) {

    if (!config_->quiet) {
      ngtcp2::debug::log_printf(nullptr,
                                "Client[%s] %s ngtcp2_conn_read_pkt: %s\n",
                                owner_->get_client_id().c_str(),
                                owner_->get_layer_str(), ngtcp2_strerror(rv));
    }
    if (!last_error_.error_code) {
      if (rv == NGTCP2_ERR_CRYPTO) {
        ngtcp2_ccerr_set_tls_alert(
            &last_error_, ngtcp2_conn_get_tls_alert(quic_conn_), nullptr, 0);
      } else {
        ngtcp2_ccerr_set_liberr(&last_error_, rv, nullptr, 0);
      }
    }

    return -1;
  }

  if (!config_->quiet) {
    const char info[] = "Client[%s] %s Received packet: local= %s remote= %s "
                        "ecn=0x%0x %lu bytes";
    ngtcp2::debug::log_printf(
        nullptr, info, owner_->get_client_id().c_str(), owner_->get_layer_str(),
        ngtcp2::util::straddr(path.local.addr, path.local.addrlen).c_str(),
        ngtcp2::util::straddr(path.remote.addr, path.remote.addrlen).c_str(),
        static_cast<uint32_t>(pi.ecn), data_len);
  }

  return 0;
}

int QuicConnection::open_bidi_stream(int64_t &stream_id,
                                     void *per_stream_user_data) {

  auto ret = ngtcp2_conn_open_bidi_stream(quic_conn_, &stream_id,
                                          per_stream_user_data);

  if (ret != 0) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s fail to open bidi stream: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), ngtcp2_strerror(ret));
  }

  ++opened_streams_;
  if (owner_->is_inner_layer() || owner_->is_pvd_client()) {
    owner_->get_client()->get_pm()->inc_opened_streams();
  }
  return ret;
}

uint64_t QuicConnection::get_left_uni_streams() {
  return ngtcp2_conn_get_streams_uni_left(quic_conn_);
}

int QuicConnection::open_uni_stream(int64_t &stream_id,
                                    void *per_stream_user_data) {
  auto ret =
      ngtcp2_conn_open_uni_stream(quic_conn_, &stream_id, per_stream_user_data);

  if (ret != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s fail to open uni stream, error: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        ngtcp2_strerror(ret));
  }

  return ret;
}

int QuicConnection::open_uni_stream(int64_t &stream_id) {
  auto rv = ngtcp2_conn_open_uni_stream(quic_conn_, &stream_id, this);
  if (rv < 0) {
    ngtcp2::debug::log_printf(nullptr,
                              "Client[%s] %s ngtcp2_conn_open_uni_stream: %s",
                              owner_->get_client_id().c_str(),
                              owner_->get_layer_str(), ngtcp2_strerror(rv));
  }
  return rv;
}

uint64_t QuicConnection::get_left_uidi_streams() {
  return ngtcp2_conn_get_streams_uni_left(quic_conn_);
}

void QuicConnection::handshake() {

  if (handshake_completed_) {
    return;
  }

  int offset = 0;
  int max_payload_size = 0;
  OuterLayer *tunnel = nullptr;

  FlatBuffer *tx_buffer = get_tx_buffer(offset, max_payload_size, tunnel);

  ssize_t datalen = 0;
  uint32_t flags = NGTCP2_WRITE_STREAM_FLAG_MORE;
  ngtcp2_pkt_info pi;
  ngtcp2_path_storage ps;
  auto ts = ngtcp2::util::timestamp();
  ngtcp2_path_storage_zero(&ps);

  auto bytes = ngtcp2_conn_writev_stream(
      quic_conn_, &ps.path, &pi, tx_buffer->data() + offset, max_payload_size,
      &datalen, flags, -1, nullptr, 0, ts);
  ngtcp2_conn_update_pkt_tx_time(quic_conn_, ts);

  if (bytes > 0) {

    if (!config_->quiet) {
      ngtcp2::debug::log_printf(
          nullptr, "client[%s] %s QUIC handshaking..., bytes:%lu",
          owner_->get_client_id().c_str(), owner_->get_layer_str(), bytes);
    }

    if (tunnel) {
      tx_buffer->set_data_len(bytes + offset);
      tunnel->tunnel_mode_egress_forward();
    } else {
      tx_buffer->set_data_len(bytes);
      owner_->get_udp_layer()->send_data();
    }
  }
}

void QuicConnection::disconnect(bool silently) {

  idle_timer_watcher_->stop();

  if (silently) {

    if (!owner_->is_stopped()) {
      owner_->on_quic_connection_closed();
    }

    return;
  }

  if (!quic_conn_ || ngtcp2_conn_in_closing_period(quic_conn_) ||
      ngtcp2_conn_in_draining_period(quic_conn_)) {
    return;
  }

  int offset = 0;
  int max_payload_size = 0;
  OuterLayer *tunnel = nullptr;

  FlatBuffer *tx_buffer = get_tx_buffer(offset, max_payload_size, tunnel);

  if (!handshake_completed_) {
    if (!config_->quiet) {
      ngtcp2::debug::log_printf(
          nullptr, "client[%s] %s has not completed QUIC handshake!",
          owner_->get_client_id().c_str(), owner_->get_layer_str());
    }
    if (!owner_->is_stopped()) {
      owner_->on_quic_connection_closed();
    }
    return;
  }

  // inform peer to close connection
  ngtcp2_path_storage ps;
  ngtcp2_path_storage_zero(&ps);
  ngtcp2_pkt_info pi;
  auto ts = ngtcp2::util::timestamp();

  auto nwrite = ngtcp2_conn_write_connection_close(
      quic_conn_, &ps.path, &pi, tx_buffer->data() + offset, max_payload_size,
      &last_error_, ts);
  ngtcp2_conn_update_pkt_tx_time(quic_conn_, ts);

  if (nwrite < 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s ngtcp2_conn_write_connection_close: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        ngtcp2_strerror(nwrite));
    return;
  }

  if (nwrite == 0) {
    return;
  }

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s ngtcp2_conn_write_connection_close, %lu bytes",
        owner_->get_client_id().c_str(), owner_->get_layer_str(), nwrite);
  }

  if (tunnel) {
    tx_buffer->set_data_len(nwrite + offset);
    tunnel->tunnel_mode_egress_forward();
  } else {
    tx_buffer->set_data_len(nwrite);
    owner_->get_udp_layer()->send_data();
  }

  if (!owner_->is_stopped()) {
    owner_->on_quic_connection_closed();
  }
}

void QuicConnection::on_upper_layer_stream_closed(int64_t stream_id) {
  closed_streams_++;

  if (owner_->is_inner_layer() || owner_->is_pvd_client()) {
    owner_->get_client()->get_pm()->inc_closed_streams();
  }

  if (closed_streams_ >= opened_streams_) {
    if (!config_->quiet) {
      ngtcp2::debug::log_printf(
          nullptr,
          "Client[%s] %s QUIC close stream:%ld (%ld/%ld), close connection.",
          owner_->get_client_id().c_str(), owner_->get_layer_str(), stream_id,
          owner_->get_client()->get_pm()->get_opened_streams(),
          owner_->get_client()->get_pm()->get_closed_streams());
    }

    // if all streams are closed, then close the connection
    disconnect(false);
  }
}

int QuicConnection::stop_sending(int64_t stream_id, uint64_t app_error_code) {
  if (auto rv = ngtcp2_conn_shutdown_stream_read(quic_conn_, 0, stream_id,
                                                 app_error_code);
      rv != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s ngtcp2_conn_shutdown_stream_read: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        ngtcp2_strerror(rv));
    return -1;
  }
  return 0;
}

int QuicConnection::reset_stream(int64_t stream_id, uint64_t app_error_code) {
  if (auto rv = ngtcp2_conn_shutdown_stream_write(quic_conn_, 0, stream_id,
                                                  app_error_code);
      rv != 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Client[%s] %s ngtcp2_conn_shutdown_stream_write: %s",
        owner_->get_client_id().c_str(), owner_->get_layer_str(),
        ngtcp2_strerror(rv));
    return -1;
  }

  return 0;
}

bool QuicConnection::get_scid(ngtcp2_cid &cid) {

  if (scid_.datalen > 0) {
    cid = scid_;
    return true;
  }

  // TODO: figure out in what condition, there are multiple SCIDs.
  //       and which one is being used.
  auto n = ngtcp2_conn_get_scid(quic_conn_, nullptr);

  ngtcp2_cid *id = new ngtcp2_cid[n];

  assert(ngtcp2_conn_get_scid(quic_conn_, id) > 0);

  cid = id[0];
  scid_ = cid;

  delete[] id;

  return cid.datalen > 0;
}

bool QuicConnection::get_dcid(ngtcp2_cid_token &dcid) {

  if (dcid_.cid.datalen > 0) {
    dcid = dcid_;
    return true;
  }

  auto n = ngtcp2_conn_get_active_dcid(quic_conn_, nullptr);

  ngtcp2_cid_token *id = new ngtcp2_cid_token[n];
  assert(ngtcp2_conn_get_active_dcid(quic_conn_, id) > 0);

  dcid = id[0];
  dcid_ = dcid;

  delete[] id;

  return dcid.cid.datalen > 0;
}

FlatBuffer *QuicConnection::get_tx_buffer(int &offset, int &max_payload_size,
                                          OuterLayer *&tunnel) {
  auto udp = owner_->get_udp_layer();
  auto tx_buffer = udp->get_tx_buffer();
  uint8_t *dest_buf = tx_buffer->data();
  max_payload_size = udp->get_max_udp_payload_size();

  // for inner layer's tunnel mode, the tx_buffer is tunnel layer's buffer
  tunnel = nullptr;
  offset = 0; // for stream-id, context id
  if (owner_->is_inner_layer()) {
    tunnel = owner_->get_client()->get_outer_layer();
    if (tunnel) {
      tx_buffer = tunnel->get_tx_buffer();
      dest_buf = tx_buffer->data();
      // store the stream_id / 4;
      int64_t stream_id =
          static_cast<InnerLayer *>(owner_)->get_masque_stream_id();

      auto left_space = tx_buffer->size();
      auto write_bytes = ngtcp2::util::encode_var_len_integer(
          stream_id >> 2, dest_buf, left_space);
      left_space -= write_bytes;
      dest_buf += write_bytes;
      offset += write_bytes;

      // store the context id
      const uint64_t context_id = 0;
      write_bytes = ngtcp2::util::encode_var_len_integer(context_id, dest_buf,
                                                         left_space);
      left_space -= write_bytes;
      dest_buf += write_bytes;
      offset += write_bytes;

      // make room for outer layer
      max_payload_size = udp->get_max_udp_payload_size() - 128;
    }
  }

  return tx_buffer;
}
