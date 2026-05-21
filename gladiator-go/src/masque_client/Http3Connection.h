#ifndef HTTP3CONNECTION_H
#define HTTP3CONNECTION_H

#include "EventHandler.h"
#include "HttpRequest.h"
#include "TimerWatcher.h"
#include "TransportLayer.h"
#include <cstdint>
#include <vector>
#pragma once
class TransportLayer;
struct nghttp3_conn;
struct nghttp3_rcbuf;
struct nghttp3_settings;
struct nghttp3_vec;
struct ngtcp2_cid;
struct nghttp3_nv;

class Http3Connection : public EventHandler {
public:
  // Lazy allocation! Please create instance after QUIC handshake completed!
  // The lifecycle of a HTTP/3 object commences upon establishment of a
  // bidirectional QUIC stream and concludes when the stream is terminated.
  Http3Connection(TransportLayer *owner, int64_t ctrl_stream_id,
                  int64_t qpack_enc_stream_id, int64_t qpack_dec_stream_id);
  virtual ~Http3Connection();

  void start();
  void stop();
  void send_data();

  void on_timeout(TimerWatcher *watcher) final;
  void on_read_ready(FdWatcher *watcher) final{}; // do not care fd ready event

  // submit requests, they are stored in internal ring queue
  // need to call write stream to write them to http3_vec buffer
  int submit_http_request(HttpRequest *req, int64_t stream_id, nghttp3_nv *nv, size_t nvlen);

  int write_stream_data(int64_t &stream_id, int &fin, nghttp3_vec *vec,
                        size_t vec_array_size);

  // will trigger the nghttp3 callback functions
  ssize_t read_stream_data(int64_t stream_id, const uint8_t *data,
                           size_t datalen, uint32_t flags);

  // int add_write_offset(int64_t stream_id, size_t bytes);

  int add_acked_stream_offset(int64_t stream_id, size_t data_len);

  // callback functions
  int on_acked_stream_data(int64_t stream_id, uint64_t datalen,
                           void *stream_user_data);

  int on_stream_close(int64_t stream_id, uint64_t app_error_code,
                      void *stream_user_data);

  int on_recv_data(int64_t stream_id, const uint8_t *data, size_t datalen,
                   void *stream_user_data);
  int on_deferred_consume(int64_t stream_id, size_t consumed,
                          void *stream_user_data);
  int on_begin_headers(int64_t stream_id, void *stream_user_data);
  int on_recv_header(int64_t stream_id, int32_t token, nghttp3_rcbuf *name,
                     nghttp3_rcbuf *value, uint8_t flags,
                     void *stream_user_data);
  int on_end_headers(int64_t stream_id, int fin, void *stream_user_data);
  int on_end_stream(int64_t stream_id, void *stream_user_data);
  int on_stop_sending(int64_t stream_id, uint64_t app_error_code,
                      void *stream_user_data);
  int on_reset_stream(int64_t stream_id, uint64_t app_error_code,
                      void *stream_user_data);
  int on_shutdown(int64_t id);
  int on_recv_settings(const nghttp3_settings *settings);

  nghttp3_conn *get_conn() const { return http3_conn_; }

  // void on_quic_write_stream(int64_t stream_id, ssize_t error,
  //                           ssize_t stream_data_len);

  HttpRequest *new_request(const HttpRequest &reqest);

  bool has_started() const { return started_; }

  void disconnect();

  //void submit_register_client_cid_message(int64_t stream_id);

  void pause_data_read();
  void resume_data_read(int64_t stream_id);

  bool is_data_ready() const;

  HttpRequest *get_request(int64_t stream_id);

private:
  TransportLayer *const owner_ = nullptr;
  nghttp3_conn *http3_conn_ = nullptr;
  bool started_ = false;
  std::vector<HttpRequest *> requests_;

  TimerWatcher *close_wait_timer_ = nullptr;

  bool data_ready_ = false;
};

#endif