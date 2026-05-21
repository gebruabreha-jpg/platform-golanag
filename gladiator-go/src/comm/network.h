/*
 * ngtcp2
 *
 * Copyright (c) 2017 ngtcp2 contributors
 * Copyright (c) 2016 nghttp2 contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 * LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
#ifndef NETWORK_H
#define NETWORK_H

#include <cstdint>
#include <netinet/in.h>
#include <optional>
#include <string>
#include <sys/types.h>
#include <sys/un.h>
#include <tuple>
#include <vector>

namespace network {

enum network_error {
  NETWORK_ERR_OK = 0,
  NETWORK_ERR_FATAL = -10,
  NETWORK_ERR_SEND_BLOCKED = -11,
  NETWORK_ERR_CLOSE_WAIT = -12,
  NETWORK_ERR_RETRY = -13,
  NETWORK_ERR_DROP_CONN = -14,
};

union in_addr_union {
  in_addr in;
  in6_addr in6;
};

union sockaddr_union {
  sockaddr_storage storage;
  sockaddr sa;
  sockaddr_in6 in6;
  sockaddr_in in;
};

struct Address {
  Address(int family, const char *ip_str, uint16_t port = 0);
  Address() = default;
  int get_family() const;
  std::string to_string() const;
  Address operator=(const Address &other);


  socklen_t len;
  union sockaddr_union su;
  uint32_t ifindex;  
};

enum class AppProtocol {
  H3,
  HQ,
};

constexpr uint8_t HQ_ALPN[] = "\xahq-interop";
constexpr uint8_t HQ_ALPN_V1[] = "\xahq-interop";

constexpr uint8_t H3_ALPN[] = "\x2h3";
constexpr uint8_t H3_ALPN_V1[] = "\x2h3";

// msghdr_get_ecn gets ECN bits from |msg|.  |family| is the address
// family from which packet is received.
unsigned int msghdr_get_ecn(msghdr *msg, int family);

// fd_set_recv_ecn sets socket option to |fd| so that it can receive
// ECN bits.
void fd_set_recv_ecn(int fd, int family);

// fd_set_ip_mtu_discover sets IP(V6)_MTU_DISCOVER socket option to
// |fd|.
void fd_set_ip_mtu_discover(int fd, int family);

// fd_set_ip_dontfrag sets IP(V6)_DONTFRAG socket option to |fd|.
void fd_set_ip_dontfrag(int fd, int family);

// fd_set_udp_gro sets UDP_GRO socket option to |fd|.
void fd_set_udp_gro(int fd);

std::optional<Address> msghdr_get_local_addr(msghdr *msg, int family);

// msghdr_get_udp_gro returns UDP_GRO value from |msg|.  If UDP_GRO is
// not found, or UDP_GRO is not supported, this function returns 0.
size_t msghdr_get_udp_gro(msghdr *msg);

void set_port(Address &dst, network::Address &src);

// get_local_addr stores preferred local address (interface address)
// in |iau| for a given destination address |remote_addr|.
int get_local_addr(in_addr_union &iau, const Address &remote_addr);

// addreq returns true if |sa| and |iau| contain the same address.
bool addreq(const sockaddr *sa, const in_addr_union &iau);

// get n IPv4 addresses from cidr network, e.g. cidr = "192.168.0.1/24"
std::vector<std::string> get_ipv4_in_cidr_network(const std::string &cidr,
                                                  int n);

// get n IPv6 addresses from cidr network, e.g. cidr =
// "2001:0db8:85a3::8a2e:0370:7334/64"
std::vector<std::string> get_ipv6_in_cidr_network(const std::string &cidr,
                                                  int n);

// get UDP IPv4 and IPv6 addresses of the service
std::tuple<bool, bool> get_udp_addresses_by_name(const char *domain_name,
                                                 const char *service_name,
                                                 struct sockaddr &ipv4,
                                                 struct sockaddr &ipv6);

int make_socket_nonblocking(int fd);
int create_nonblock_socket(int domain, int type, int protocol);
const char *get_interface_name(int socket_fd, int family, char* buffer, size_t bufsize);
int get_mtu(const char *ifname);

// get remote Address
Address get_remote_address(const char *domain_name, const char *service_name,
                           int family);

} // namespace network

#endif // NETWORK_H
