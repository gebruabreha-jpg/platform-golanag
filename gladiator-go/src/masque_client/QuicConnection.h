#ifndef QUICCONNECTION_H
#define QUICCONNECTION_H

#include "EventHandler.h"
#include "FdWatcher.h"
#include "FlatBuffer.h"
#include "TimerWatcher.h"
#include "TransportLayer.h"
#include "network.h"
#include "tls_client_session_quictls.h"
#include <ngtcp2/ngtcp2_crypto.h>
#pragma once

class TimerWatcher;
class TransportLayer;
class FlatBuffer;

class QuicConnection : public EventHandler {
public:
  struct Stats {
    uint64_t total_tx_stream_bytes = 0;
  };
  QuicConnection(TransportLayer *owner, const char *remote_host,
                 const network::Address &local, const network::Address &remote,
                 int max_udp_payload_size);
  virtual ~QuicConnection();

  ngtcp2_conn *get_quic_conn() const { return quic_conn_; }

  void on_timeout(TimerWatcher *watcher) final;
  void on_read_ready(FdWatcher *watcher) final{}; // do not care fd ready event

  void on_udp_data_ready_to_read(const uint8_t *buffer, size_t data_len);

  const Config *get_config() const;

  void handshake();

  // callback functions
  int on_handshake_completed();

  int on_recv_stream_data(uint32_t flags, int64_t stream_id, uint64_t offset,
                          const uint8_t *data, size_t datalen,
                          void *stream_user_data);

  int on_acked_stream_data_offset(int64_t stream_id, uint64_t offset,
                                  uint64_t datalen, void *stream_user_data);

  int on_extend_max_bidi_streams(uint64_t max_streams);

  int on_recv_datagram(uint32_t flags, const uint8_t *data, size_t datalen);

  int on_ack_datagram(uint64_t dgram_id);

  int open_uni_stream(int64_t &stream_id);
  int open_bidi_stream(int64_t &stream_id, void *per_stream_user_data);

  uint64_t get_left_uidi_streams();

  uint64_t get_left_uni_streams();
  int open_uni_stream(int64_t &stream_id, void *per_stream_user_data);

  int on_key_update(uint8_t *rx_secret, uint8_t *tx_secret,
                    ngtcp2_crypto_aead_ctx *rx_aead_ctx, uint8_t *rx_iv,
                    ngtcp2_crypto_aead_ctx *tx_aead_ctx, uint8_t *tx_iv,
                    const uint8_t *current_rx_secret,
                    const uint8_t *current_tx_secret, size_t secretlen);

  int on_stream_close(uint32_t flags, int64_t stream_id,
                      uint64_t app_error_code, void *stream_user_data);

  int on_stream_reset(int64_t stream_id);

  // size_t send_data();
  size_t recv_data(const uint8_t *buffer, size_t data_len);

  size_t get_max_data_left();
  size_t get_send_quantum();

  using TX_VEC_ARRAY = std::array<ngtcp2_vec, TransportLayer::TX_BUFF_VEC_NUM>;

  struct Stream_Data {
    TX_VEC_ARRAY vec_array;
    size_t vec_num = 0;
    int64_t stream_id = -1;
    int fin = 0;
  };

  Stream_Data &get_stream_data_ref();

  void disconnect(bool silently);

  void on_upper_layer_stream_closed(int64_t stream_id);

  int stop_sending(int64_t stream_id, uint64_t app_error_code);
  int reset_stream(int64_t stream_id, uint64_t app_error_code);

  bool is_handshake_completed() const { return handshake_completed_; }

  void write_datagram_data(FlatBuffer *buffer);

  bool get_scid(ngtcp2_cid &cid);

  bool get_dcid(ngtcp2_cid_token &cid);

  void update_timer();

private:
  // return the bytes of stream data have been written to tx_buffer sucessfuly
  // size_t send_stream_data();

  TransportLayer *const owner_ = nullptr;
  TimerWatcher *idle_timer_watcher_ = nullptr;

  ngtcp2_crypto_conn_ref *conn_ref() { return &conn_ref_; }

  network::Address local_addr_;
  network::Address remote_addr_;

  ngtcp2_conn *quic_conn_ = nullptr;
  ngtcp2_ccerr last_error_;

  TLSClientSession tls_session_;
  ngtcp2_crypto_conn_ref conn_ref_;

  const struct Config *config_ = nullptr;

  Stream_Data stream_data_{};

  int opened_streams_ = 0;
  int closed_streams_ = 0;

  bool all_streams_closed() const {
    return closed_streams_ >= opened_streams_ && opened_streams_ > 0;
  }

  // for inner layer's tunnel mode it return the tx buffer of tunnel layer, the
  // offset , max payload size and the tunnel layer for non-tunnel mode it
  // return the tx buffer of udp layer, the offset (0) and max payload size, the
  // tunnel layer is nullptr
  FlatBuffer *get_tx_buffer(int &offset, int &max_payload_size,
                            OuterLayer *&tunnel);

  bool handshake_completed_ = false;

  ngtcp2_cid scid_{};
  ngtcp2_cid_token dcid_{};
};

#endif
