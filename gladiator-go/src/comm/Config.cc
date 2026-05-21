#include "Config.h"
#include "template.h"
#include "util.h"
#include <cstring>
#include <string_view>

Config::Config() {
  tx_loss_prob = 0.;
  rx_loss_prob = 0.;
  fd = -1;
  ciphers = ngtcp2::util::crypto_default_ciphers();
  groups = ngtcp2::util::crypto_default_groups();
  nstreams = 0;
  data = nullptr;
  datalen = 0;
  version = NGTCP2_PROTO_VER_V1;
  timeout = 60 * NGTCP2_SECONDS;
  http_method = std::string_view("GET");
  max_data = 15_m;
  max_stream_data_bidi_local = 6_m;
  max_stream_data_bidi_remote = 6_m;
  max_stream_data_uni = 6_m;
  max_window = 24_m;
  max_stream_window = 16_m;
  max_streams_uni = 100;
  max_streams_bidi = 100;
  cc_algo = NGTCP2_CC_ALGO_CUBIC;
  initial_rtt = NGTCP2_DEFAULT_INITIAL_RTT;
  handshake_timeout = std::min(30 * NGTCP2_SECONDS, timeout);
  ack_thresh = 2;
  initial_pkt_num = UINT32_MAX;
  workers_num = 1;
  clients_num = 1;
  new_conn_per_second = 10;
  exit_on_all_streams_close = false;
  tls_context = nullptr;
  quiet = false;
  max_trans_unit = 0;
  request_quic_forwarding = false;
  pvd_only = false;
  request_duplication_factor = 1;
  duration = 0;
}
