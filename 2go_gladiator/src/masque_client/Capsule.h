#ifndef CAPSULE_H_
#define CAPSULE_H_

#include <ngtcp2/ngtcp2.h>

namespace Capsule {
constexpr uint64_t REGISTER_CLIENT_CID = 0xFFE500;
constexpr uint64_t REGISTER_TARGET_CID = 0xFFE501;
constexpr uint64_t ACK_CLIENT_CID = 0xFFE502;
constexpr uint64_t ACK_CLIENT_VCID = 0xFFE503;
constexpr uint64_t ACK_TARGET_CID = 0xFFE504;
constexpr uint64_t CLOSE_CLIENT_CID = 0xFFE505;
constexpr uint64_t CLOSE_TARGET_CID = 0xFFE506;

constexpr size_t REGISTER_CLIENT_CID_MAX_BUF_SIZE = 8 + 8 + 8 + 20;
constexpr size_t REGISTER_TARGET_CID_MAX_BUF_SIZE = 8 + 8 + 8 + 20 + 8 + 16;
constexpr size_t ACK_CLIENT_VCID_MAX_BUF_SIZE = 8 + 8 + 8 + 20 + 8 + 20;
constexpr size_t CLOSE_X_CID_MAX_BUF_SIZE = 8 + 8 + 8 + 20;

bool parse_ack_client_cid(const uint8_t *pbuf, size_t &totallen,
                          ngtcp2_cid &scid, ngtcp2_cid &vscid, bool quiet);

bool parse_ack_target_cid(const uint8_t *pbuf, size_t &totallen,
                          ngtcp2_cid &tcid, ngtcp2_cid &vtcid, bool quiet);

void parse_close_x_cid(const uint8_t *pbuf, size_t &totallen, ngtcp2_cid &cid,
                       bool quiet);

size_t create_register_client_cid(const ngtcp2_cid &cid, uint8_t *pbuf,
                                  bool quiet);

size_t create_register_target_cid(const ngtcp2_cid &tcid, uint8_t *pbuf,
                                  bool quiet);

size_t create_ack_client_vcid(const ngtcp2_cid &cid, const ngtcp2_cid &vcid,
                              uint8_t *pbuf, bool quiet);

size_t create_close_x_cid(const ngtcp2_cid &cid, uint64_t type, uint8_t *pbuf,
                          bool quiet);

} // namespace Capsule

#endif /* CAPSULE_H_ */
