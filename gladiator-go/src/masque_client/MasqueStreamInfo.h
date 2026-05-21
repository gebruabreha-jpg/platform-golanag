#ifndef MASQUESTREAMINFO_H
#define MASQUESTREAMINFO_H
#include "MasqueClient.h"
#include <cstdint>
#include <ngtcp2/ngtcp2.h>
#pragma once

class InnerLayer;

struct MasqueStreamInfo {

  int64_t stream_id = -1;

  // connection's information for inner layer
  std::string host;
  std::string port;
  InnerLayer *inner_layer = nullptr;

  // set true if proxy returns a 'capsule-protocol': '?1'
  bool support_capsule_protocol = false;
  // set true if proxy return 'proxy-quic-forwarding': '?1' and support_capsule_protocol==true
  bool support_quic_forwarding = false;

  // set true only if client cid registration process completed
  bool quic_forwording_rx_ready = false;

  // set true only if target cid registration process completed
  bool quic_forwording_tx_ready = false;

  ngtcp2_cid scid{};
  ngtcp2_cid_token dcid{};

  ngtcp2_cid vscid{};
  ngtcp2_cid_token vdcid{};

};

#endif