#include "Capsule.h"
#include "debug.h"
#include "util.h"

#include <cstring>

namespace Capsule {

/*
Acknowledge Client CID Capsule {
  Type (i) = ACK_CLIENT_CID
  Length (i)
  Connection ID Length (i)
  Connection ID (0..2040),
  Virtual Connection ID Length (i)
  Virtual Connection ID (0..2040),
}
*/
bool parse_ack_client_cid(const uint8_t *pbuf, size_t &totallen,
                          ngtcp2_cid &scid, ngtcp2_cid &vscid, bool quiet) {

  // decode scid len
  auto bytes =
      ngtcp2::util::decode_var_len_integer(scid.datalen, pbuf, totallen);

  if (bytes == 0) {
    return false;
  }
  pbuf += bytes;
  if (static_cast<size_t>(bytes) > totallen) {
    return false;
  }

  totallen -= bytes;

  memcpy(scid.data, pbuf, scid.datalen);
  pbuf += scid.datalen;
  if (totallen < scid.datalen) {
    return false;
  }
  totallen -= scid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::parse_ack_client_cid scid = %s",
        ngtcp2::util::format_hex(scid.data, scid.datalen).c_str());
  }

  // decode vscid len
  bytes = ngtcp2::util::decode_var_len_integer(vscid.datalen, pbuf, totallen);
  pbuf += bytes;
  if (static_cast<size_t>(bytes) > totallen) {
    return false;
  }
  totallen -= bytes;

  memcpy(vscid.data, pbuf, vscid.datalen);
  pbuf += vscid.datalen;
  if (totallen < vscid.datalen) {
    return false;
  }
  totallen -= vscid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::parse_ack_client_cid vscid = %s",
        ngtcp2::util::format_hex(vscid.data, vscid.datalen).c_str());
  }
  return true;
}

/*
Acknowledge Target CID Capsule {
  Type (i) = ACK_TARGET_CID
  Length (i)
  Connection ID Length (i)
  Connection ID (0..2040),
  Virtual Connection ID Length (i)
  Virtual Connection ID (0..2040),
  Stateless Reset Token Length (i),
  Stateless Reset Token (..),
}
*/
bool parse_ack_target_cid(const uint8_t *pbuf, size_t &totallen,
                          ngtcp2_cid &tcid, ngtcp2_cid &vtcid, bool quiet) {

  // decode the tcid
  auto bytes =
      ngtcp2::util::decode_var_len_integer(tcid.datalen, pbuf, totallen);

  if (bytes == 0) {
    return false;
  }
  pbuf += bytes;

  if (static_cast<size_t>(bytes) > totallen) {
    return false;
  }
  totallen -= bytes;

  memcpy(tcid.data, pbuf, tcid.datalen);
  pbuf += tcid.datalen;
  if (totallen < tcid.datalen) {
    return false;
  }
  totallen -= tcid.datalen;

  if (!quiet or bytes == 0) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::parse_ack_target_cid tcid = %s",
        ngtcp2::util::format_hex(tcid.data, tcid.datalen).c_str());
  }

  // decode the vtcid
  bytes = ngtcp2::util::decode_var_len_integer(vtcid.datalen, pbuf, totallen);
  if (bytes == 0) {
    return false;
  }
  pbuf += bytes;
  if (static_cast<size_t>(bytes) > totallen) {
    return false;
  }
  totallen -= bytes;

  memcpy(vtcid.data, pbuf, vtcid.datalen);
  pbuf += vtcid.datalen;
  if (totallen < tcid.datalen) {
    return false;
  }
  totallen -= vtcid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::parse_ack_target_cid vtcid = %s",
        ngtcp2::util::format_hex(vtcid.data, vtcid.datalen).c_str());
  }

  return true;

  // TODO: decode token?
}

/*
Close CID Capsule {
  Type (i) = CLOSE_CLIENT_CID, CLOSE_TARGET_CID
  Length (i),
  Connection ID (0..2040),
}
*/
void parse_close_x_cid(const uint8_t *pbuf, size_t &totallen, ngtcp2_cid &cid,
                       bool quiet) {

  // decode scid len
  auto bytes =
      ngtcp2::util::decode_var_len_integer(cid.datalen, pbuf, totallen);
  pbuf += bytes;
  totallen -= bytes;

  memcpy(cid.data, pbuf, cid.datalen);
  pbuf += cid.datalen;
  totallen -= cid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::parse_close_x_cid cid = %s",
        ngtcp2::util::format_hex(cid.data, cid.datalen).c_str());
  }
}

/*
Register Client CID Capsule {
  Type (i) = REGISTER_CLIENT_CID
  Length (i),
  Connection ID (0..2040),
}
}
*/
size_t create_register_client_cid(const ngtcp2_cid &cid, uint8_t *pbuf,
                                  bool quiet) {

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::create_register_client_cid cid = %s len=%d",
        ngtcp2::util::format_hex(cid.data, cid.datalen).c_str(), cid.datalen);
  }

  // encode capsule type
  uint8_t *pos = pbuf;
  int left_space = REGISTER_CLIENT_CID_MAX_BUF_SIZE;
  auto bytes = ngtcp2::util::encode_var_len_integer(REGISTER_CLIENT_CID, pos,
                                                    left_space);
  pos += bytes;
  left_space -= bytes;

  // encode the total length
  bytes = ngtcp2::util::encode_var_len_integer(cid.datalen, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode cid content
  memcpy(pos, cid.data, cid.datalen);
  pos += cid.datalen;

  // return total length
  return pos - pbuf;
}

/*
 Register Target CID Capsule {
  Type (i) = REGISTER_TARGET_CID
  Length (i),
  Connection ID Length (i)
  Connection ID (0..2040),
  Stateless Reset Token Length (i),
  Stateless Reset Token (..),
}
*/
size_t create_register_target_cid(const ngtcp2_cid &tcid, uint8_t *pbuf,
                                  bool quiet) {

  // a buffer on stack
  uint8_t buffer[REGISTER_TARGET_CID_MAX_BUF_SIZE];
  uint8_t *pos = buffer;
  int left_space = REGISTER_TARGET_CID_MAX_BUF_SIZE;

  // encode tcid length;
  auto bytes =
      ngtcp2::util::encode_var_len_integer(tcid.datalen, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode tcid content
  memcpy(pos, tcid.data, tcid.datalen);
  pos += tcid.datalen;
  left_space -= tcid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::create_register_target_cid cid = %s",
        ngtcp2::util::format_hex(tcid.data, tcid.datalen).c_str());
  }

  // Stateless Reset Token Length (i),
  // Stateless Reset Token (..),
  bytes = ngtcp2::util::encode_var_len_integer(0, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  auto total_length_of_content = pos - buffer;

  // encode capsule type
  pos = pbuf;
  left_space = REGISTER_TARGET_CID_MAX_BUF_SIZE;
  bytes = ngtcp2::util::encode_var_len_integer(REGISTER_TARGET_CID, pos,
                                               left_space);
  pos += bytes;
  left_space -= bytes;

  // encode the total length
  bytes = ngtcp2::util::encode_var_len_integer(total_length_of_content, pos,
                                               left_space);
  pos += bytes;
  left_space -= bytes;

  auto capsule_header_length = pos - pbuf;

  // copy the content
  memcpy(pos, buffer, total_length_of_content);

  // return total length
  return capsule_header_length + total_length_of_content;
}

/*
 Acknowledge Client VCID Capsule {
  Type (i) = ACK_CLIENT_VCID
  Length (i)
  Connection ID Length (i)
  Connection ID (0..2040),
  Virtual Connection ID Length (i)
  Virtual Connection ID (0..2040),
  Stateless Reset Token Length (i),
  Stateless Reset Token (..),
}
*/
size_t create_ack_client_vcid(const ngtcp2_cid &cid, const ngtcp2_cid &vcid,
                              uint8_t *pbuf, bool quiet) {
  // a buffer on stack
  uint8_t buffer[ACK_CLIENT_VCID_MAX_BUF_SIZE];
  uint8_t *pos = buffer;
  int left_space = ACK_CLIENT_VCID_MAX_BUF_SIZE;

  // encode scid length;
  auto bytes =
      ngtcp2::util::encode_var_len_integer(cid.datalen, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode scid content
  memcpy(pos, cid.data, cid.datalen);
  pos += cid.datalen;
  left_space -= cid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr,
        "Capsule:: cid = "
        "%s",
        ngtcp2::util::format_hex(cid.data, cid.datalen).c_str());
  }

  // encode virtual scid length
  bytes = ngtcp2::util::encode_var_len_integer(vcid.datalen, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode virtual client cid
  memcpy(pos, vcid.data, vcid.datalen);
  pos += vcid.datalen;
  left_space -= vcid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "Capsule::create_ack_client_vcid vcid = %s",
        ngtcp2::util::format_hex(vcid.data, vcid.datalen).c_str());
  }

  // Stateless Reset Token Length (i),
  // Stateless Reset Token (..),
  bytes = ngtcp2::util::encode_var_len_integer(0, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  auto total_length_of_content = pos - buffer;

  // encode capsule type
  pos = pbuf;
  left_space = ACK_CLIENT_VCID_MAX_BUF_SIZE;
  bytes =
      ngtcp2::util::encode_var_len_integer(ACK_CLIENT_VCID, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode the total length
  bytes = ngtcp2::util::encode_var_len_integer(total_length_of_content, pos,
                                               left_space);
  pos += bytes;
  left_space -= bytes;

  auto capsule_header_length = pos - pbuf;

  // copy the content
  memcpy(pos, buffer, total_length_of_content);

  // return total length
  return capsule_header_length + total_length_of_content;
}

/*
Close CID Capsule {
  Type (i) = 0xffe505, 0xffe506
  Length (i),
  Connection ID (0..2040),
}
*/
size_t create_close_x_cid(const ngtcp2_cid &cid, uint64_t type, uint8_t *pbuf,
                          bool quiet) {

  // a buffer on stack
  uint8_t buffer[CLOSE_X_CID_MAX_BUF_SIZE];
  uint8_t *pos = buffer;
  int left_space = CLOSE_X_CID_MAX_BUF_SIZE;

  // encode cid content
  memcpy(pos, cid.data, cid.datalen);
  pos += cid.datalen;
  left_space -= cid.datalen;

  if (!quiet) {
    ngtcp2::debug::log_printf(
        nullptr, "create_close_x_cid type = %s cid = %s len=%d",
        ngtcp2::util::format_hex(type).c_str(),
        ngtcp2::util::format_hex(cid.data, cid.datalen).c_str(), cid.datalen);
  }

  auto total_length_of_content = pos - buffer;

  // encode capsule type
  pos = pbuf;
  left_space = CLOSE_X_CID_MAX_BUF_SIZE;
  auto bytes = ngtcp2::util::encode_var_len_integer(type, pos, left_space);
  pos += bytes;
  left_space -= bytes;

  // encode the total length
  bytes = ngtcp2::util::encode_var_len_integer(total_length_of_content, pos,
                                               left_space);
  pos += bytes;
  left_space -= bytes;

  auto capsule_header_length = pos - pbuf;

  // copy the content
  memcpy(pos, buffer, total_length_of_content);

  // return total length
  return capsule_header_length + total_length_of_content;
}

} // namespace Capsule
