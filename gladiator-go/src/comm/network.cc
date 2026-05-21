
#include "network.h"
#include "debug.h"
#include <arpa/inet.h>
#include <cassert>
#include <cstdint>
#include <cstdio>
#include <cstring>
#include <fcntl.h>
#include <ifaddrs.h>
#include <iostream>
#include <net/if.h>
#include <netdb.h>
#include <netinet/ip.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <sys/sysinfo.h>
#include <sys/types.h>
#include <unistd.h>
// local functions
namespace {
void cidr_to_subnet_mask(uint32_t &ipv4_mask, uint8_t cidr) {
  ipv4_mask = 0;
  for (uint8_t i = 0; i < cidr; ++i) {
    ipv4_mask |= (1 << (31 - i));
  }
}

void cidr_to_subnet_mask(in6_addr &ipv6_mask, uint8_t cidr) {
  ipv6_mask = {};
  int8_t index = cidr / 8;
  if (index < 0) {
    return;
  }
  for (int i = 0; i < index; ++i) {
    ipv6_mask.s6_addr[i] = 0xFF;
  }

  for (int i = 0; i < (cidr % 8); ++i) {
    ipv6_mask.s6_addr[index] |= (1 << (7 - i));
  }
}

} // namespace

namespace network {

Address::Address(int family, const char *ip_str, uint16_t port) {
  su.sa.sa_family = family;
  if (family == AF_INET6) {
    inet_pton(family, ip_str, &su.in6.sin6_addr);
    su.in6.sin6_port = port;
    len = sizeof(su.in6);
  } else if (family == AF_INET) {
    inet_pton(family, ip_str, &su.in.sin_addr);
    su.in.sin_port = port;
    len = sizeof(su.in);
  }
}

int Address::get_family() const { return su.sa.sa_family; }

std::string Address::to_string() const {
  char ip_str[200];
  ip_str[0] = '[';
  auto family = get_family();
  if (family == AF_INET) {
    inet_ntop(family, &su.in.sin_addr, &ip_str[1], len);
  } else {
    inet_ntop(family, &su.in6.sin6_addr, &ip_str[1], len);
  }
  auto ip_str_len = strlen(ip_str);
  snprintf(&ip_str[ip_str_len], 8, "]:%u",
           htons((family == AF_INET) ? su.in.sin_port : su.in6.sin6_port));

  return ip_str;
}

Address Address::operator=(const Address &other) {
  memcpy(this, &other, sizeof(Address));
  return *this;
}

std::vector<std::string> get_ipv4_in_cidr_network(const std::string &cidr,
                                                  int n) {
  size_t pos = cidr.find('/');
  if (pos == std::string::npos) {
    std::cerr << "Invalid CIDR format" << std::endl;
    return {};
  }

  std::string network_addr = cidr.substr(0, pos);
  int cidr_prefix_len = std::stoi(cidr.substr(pos + 1));

  unsigned int ipv4_mask;
  cidr_to_subnet_mask(ipv4_mask, cidr_prefix_len);
  // std::cout << "mask " << std::hex << ipv4_mask << std::endl;

  unsigned int network_ipv4 = inet_network(network_addr.c_str());
  // std::cout << "ipv4 " << std::hex << network_ipv4 << std::endl;

  unsigned int masked_ipv4_address = network_ipv4 & ipv4_mask;
  unsigned int start_ipv4 = masked_ipv4_address;
  unsigned int end_ipv4 = masked_ipv4_address + (~ipv4_mask);

  //assert(start_ipv4 != end_ipv4);

  // std::cout << std::hex << "start:" << std::min(start_ipv4,end_ipv4) <<
  // "\tend:" << std::max(start_ipv4,end_ipv4) << std::endl;

  std::vector<std::string> ip_list;
  // Note: need to skip the network address and the network broadcast address
  for (unsigned int i = std::min(start_ipv4, end_ipv4) + 1;
       i <= std::max(start_ipv4, end_ipv4) - 1; ++i) {
    struct in_addr addr;
    addr.s_addr = htonl(i);
    ip_list.insert(ip_list.begin(), std::string(inet_ntoa(addr)));
    if (--n <= 0) {
      break;
    }
  }
  if (32 == cidr_prefix_len){
    ip_list.push_back(network_addr);
  }

  return ip_list;
}

int compare_ipv6(const in6_addr &a, const in6_addr &b) {
  for (int i = 0; i < 16; i++) {
    if (a.s6_addr[i] > b.s6_addr[i]) {
      return 1;
    } else if (a.s6_addr[i] < b.s6_addr[i]) {
      return -1;
    }
  }
  return 0;
}

static void add_and_carry(uint8_t *v, int idx) {
  if (idx < 0) {
    return;
  }

  v[idx] += 1;

  if (v[idx] == 0) {
    add_and_carry(v, idx - 1);
  }
}

void ipv6_increase(in6_addr &ipv6) { add_and_carry(ipv6.s6_addr, 15); }

std::vector<std::string> get_ipv6_in_cidr_network(const std::string &cidr,
                                                  int n) {
  // parse the cidr network address
  size_t pos = cidr.find('/');
  if (pos == std::string::npos) {
    std::cerr << "Invalid CIDR format" << std::endl;
    return {};
  }

  std::string network_addr = cidr.substr(0, pos);
  int cidr_prefix_len = std::stoi(cidr.substr(pos + 1));

  in6_addr cidr_network;
  inet_pton(AF_INET6, network_addr.c_str(), &cidr_network);

  in6_addr ipv6_mask;
  cidr_to_subnet_mask(ipv6_mask, cidr_prefix_len);

  in6_addr ipv6_start = cidr_network;
  in6_addr ipv6_end = cidr_network;

  for (int i = 0; i < 16; i++) {
    ipv6_start.s6_addr[i] &= ipv6_mask.s6_addr[i];
    ipv6_end.s6_addr[i] = ipv6_start.s6_addr[i] | (~(ipv6_mask.s6_addr[i]));
  }

  ipv6_start.s6_addr[15] = 1;

  std::vector<std::string> ip_list;

  for (in6_addr ipv6 = ipv6_start; compare_ipv6(ipv6, ipv6_end) <= 0 && n--;
       ipv6_increase(ipv6)) {
    char ipv6_str[128];
    inet_ntop(AF_INET6, &ipv6, ipv6_str, INET6_ADDRSTRLEN);
    ip_list.push_back(ipv6_str);
  }

  return ip_list;
}

std::tuple<bool, bool> get_udp_addresses_by_name(const char *domain_name,
                                                 const char *service_name,
                                                 struct sockaddr &ipv4,
                                                 struct sockaddr &ipv6) {
  struct addrinfo *result;
  ipv4.sa_family = 0;
  ipv6.sa_family = 0;

  if (getaddrinfo(domain_name, service_name, nullptr, &result) == 0) {

    for (auto *rp = result; rp != nullptr; rp = rp->ai_next) {

      if (rp->ai_socktype != SOCK_DGRAM) {
        continue;
      }

      if (rp->ai_family == AF_INET) {
        memcpy(&ipv4, rp->ai_addr, rp->ai_addrlen);

      } else if (rp->ai_family == AF_INET6) {
        memcpy(&ipv6, rp->ai_addr, rp->ai_addrlen);
      }
    }

    freeaddrinfo(result);
  }

  return {ipv4.sa_family > 0, ipv6.sa_family > 0};
}

Address get_remote_address(const char *domain_name, const char *service_name,
                           int family) {
  Address ipv4;
  Address ipv6;
  ipv4.len = sizeof(ipv4.su.in);
  ipv6.len = sizeof(ipv6.su.in6);
  auto [ipv4_ava, ipv6_ava] = get_udp_addresses_by_name(
      domain_name, service_name, ipv4.su.sa, ipv6.su.sa);

  if (family == AF_INET && ipv4_ava) {
    return ipv4;
  }

  if (family == AF_INET6 && ipv6_ava) {
    return ipv6;
  }

  if (family == AF_INET) {

    ngtcp2::debug::log_printf(nullptr, "Can not get IPv4 address of %s:%s!\n",
                              domain_name, service_name);
  } else {
    ngtcp2::debug::log_printf(nullptr, "Can not get IPv6 address of %s:%s!\n",
                              domain_name, service_name);
  }
  // fatal error
  exit(-1);
};

int make_socket_nonblocking(int fd) {
  int rv;
  int flags;

  while ((flags = fcntl(fd, F_GETFL, 0)) == -1 && errno == EINTR)
    ;
  if (flags == -1) {
    return -1;
  }

  while ((rv = fcntl(fd, F_SETFL, flags | O_NONBLOCK)) == -1 && errno == EINTR)
    ;

  return rv;
}

int create_nonblock_socket(int domain, int type, int protocol) {

  auto fd = socket(domain, type | SOCK_NONBLOCK, protocol);
  if (fd == -1) {
    return -1;
  }
  return fd;
}

const char *get_interface_name(int socket_fd, int family, char* buffer, size_t bufsize) {
  struct sockaddr_in local_addr;
  struct sockaddr_in6 local_addr6;
  socklen_t addrlen = 0;
  buffer[0] = '\0';

  void *addr = &local_addr.sin_addr;

  if (family == AF_INET) {
    addrlen = sizeof(local_addr);
    getsockname(socket_fd, (struct sockaddr *)&local_addr, &addrlen);
    addr = &local_addr.sin_addr;

  } else {
    addrlen = sizeof(local_addr6);
    getsockname(socket_fd, (struct sockaddr *)&local_addr6, &addrlen);
    addr = &local_addr6.sin6_addr;
  }

  // fprintf(stderr,"family=%d,addrlen=%u\n",family,addrlen);

  struct ifaddrs *ifaddr, *ifa;
  if (getifaddrs(&ifaddr) == -1) {
    std::cerr << "Error getting interface addresses" << std::endl;
    return nullptr;
  }

  for (ifa = ifaddr; ifa != NULL; ifa = ifa->ifa_next) {

    if (ifa->ifa_addr == NULL or ifa->ifa_addr->sa_family != family) {
      continue;
    }

    void *p_sin_addr = &((struct sockaddr_in *)ifa->ifa_addr)->sin_addr;
    void *p_sin6_addr = &((struct sockaddr_in6 *)ifa->ifa_addr)->sin6_addr;
    void *p_addr = (family == AF_INET) ? p_sin_addr : p_sin6_addr;
    auto addr_size =
        (family == AF_INET) ? sizeof(struct in_addr) : sizeof(struct in6_addr);

    if (memcmp(addr, p_addr, addr_size) == 0) {
      std::strncpy(buffer,ifa->ifa_name,bufsize);
      buffer[bufsize-1] = 0; // for safe
      break;
    }
  }

  freeifaddrs(ifaddr);
  // ngtcp2::debug::log_printf(nullptr,"Unable to get interface name.");
  return buffer;
}

int get_mtu(const char *ifname) {

  int socket_fd = socket(AF_INET, SOCK_DGRAM, 0);
  if (socket_fd == -1) {
    std::cerr << "Error creating socket" << std::endl;
    return -1;
  }

  struct ifreq ifr;
  strncpy(ifr.ifr_name, ifname, IFNAMSIZ - 1);
  ifr.ifr_name[IFNAMSIZ - 1] = '\0'; // Ensure null-terminated string
  if (ioctl(socket_fd, SIOCGIFMTU, &ifr) == -1) {
    std::cerr << "Error retrieving MTU" << std::endl;
    close(socket_fd);
    return -1;
  }

  close(socket_fd);
  return ifr.ifr_mtu;
}

unsigned int msghdr_get_ecn(msghdr *msg, int family) {
  switch (family) {
  case AF_INET:
    for (auto cmsg = CMSG_FIRSTHDR(msg); cmsg; cmsg = CMSG_NXTHDR(msg, cmsg)) {
      if (cmsg->cmsg_level == IPPROTO_IP &&
#ifdef __APPLE__
          cmsg->cmsg_type == IP_RECVTOS
#else  // !__APPLE__
          cmsg->cmsg_type == IP_TOS
#endif // !__APPLE__
          && cmsg->cmsg_len) {
        return *reinterpret_cast<uint8_t *>(CMSG_DATA(cmsg)) & IPTOS_ECN_MASK;
      }
    }
    break;
  case AF_INET6:
    for (auto cmsg = CMSG_FIRSTHDR(msg); cmsg; cmsg = CMSG_NXTHDR(msg, cmsg)) {
      if (cmsg->cmsg_level == IPPROTO_IPV6 && cmsg->cmsg_type == IPV6_TCLASS &&
          cmsg->cmsg_len) {
        unsigned int tos;

        memcpy(&tos, CMSG_DATA(cmsg), sizeof(int));

        return tos & IPTOS_ECN_MASK;
      }
    }
    break;
  }

  return 0;
}

void fd_set_recv_ecn(int fd, int family) {
  unsigned int tos = 1;
  switch (family) {
  case AF_INET:
    if (setsockopt(fd, IPPROTO_IP, IP_RECVTOS, &tos,
                   static_cast<socklen_t>(sizeof(tos))) == -1) {
      std::cerr << "setsockopt: " << strerror(errno) << std::endl;
    }
    break;
  case AF_INET6:
    if (setsockopt(fd, IPPROTO_IPV6, IPV6_RECVTCLASS, &tos,
                   static_cast<socklen_t>(sizeof(tos))) == -1) {
      std::cerr << "setsockopt: " << strerror(errno) << std::endl;
    }
    break;
  }
}

void fd_set_ip_mtu_discover(int fd, int family) {
#if defined(IP_MTU_DISCOVER) && defined(IPV6_MTU_DISCOVER)
  int val;

  switch (family) {
  case AF_INET:
    val = IP_PMTUDISC_DO;
    if (setsockopt(fd, IPPROTO_IP, IP_MTU_DISCOVER, &val,
                   static_cast<socklen_t>(sizeof(val))) == -1) {
      std::cerr << "setsockopt: IP_MTU_DISCOVER: " << strerror(errno)
                << std::endl;
    }
    break;
  case AF_INET6:
    val = IPV6_PMTUDISC_DO;
    if (setsockopt(fd, IPPROTO_IPV6, IPV6_MTU_DISCOVER, &val,
                   static_cast<socklen_t>(sizeof(val))) == -1) {
      std::cerr << "setsockopt: IPV6_MTU_DISCOVER: " << strerror(errno)
                << std::endl;
    }
    break;
  }
#endif // defined(IP_MTU_DISCOVER) && defined(IPV6_MTU_DISCOVER)
}

void fd_set_ip_dontfrag(int fd, int family) {
#if defined(IP_DONTFRAG) && defined(IPV6_DONTFRAG)
  int val = 1;

  switch (family) {
  case AF_INET:
    if (setsockopt(fd, IPPROTO_IP, IP_DONTFRAG, &val,
                   static_cast<socklen_t>(sizeof(val))) == -1) {
      std::cerr << "setsockopt: IP_DONTFRAG: " << strerror(errno) << std::endl;
    }
    break;
  case AF_INET6:
    if (setsockopt(fd, IPPROTO_IPV6, IPV6_DONTFRAG, &val,
                   static_cast<socklen_t>(sizeof(val))) == -1) {
      std::cerr << "setsockopt: IPV6_DONTFRAG: " << strerror(errno)
                << std::endl;
    }
    break;
  }
#endif // defined(IP_DONTFRAG) && defined(IPV6_DONTFRAG)
}

void fd_set_udp_gro(int fd) {
#ifdef UDP_GRO
  int val = 1;

  if (setsockopt(fd, IPPROTO_UDP, UDP_GRO, &val,
                 static_cast<socklen_t>(sizeof(val))) == -1) {
    static std::atomic_bool has_been_printed = false;
    if (!has_been_printed.load()) {
      has_been_printed.store(true);
      std::cerr << "setsockopt: UDP_GRO: " << strerror(errno) << std::endl;
    }
  }
#endif // UDP_GRO
}

std::optional<network::Address> msghdr_get_local_addr(msghdr *msg, int family) {
  switch (family) {
  case AF_INET:
    for (auto cmsg = CMSG_FIRSTHDR(msg); cmsg; cmsg = CMSG_NXTHDR(msg, cmsg)) {
      if (cmsg->cmsg_level == IPPROTO_IP && cmsg->cmsg_type == IP_PKTINFO) {
        in_pktinfo pktinfo;
        memcpy(&pktinfo, CMSG_DATA(cmsg), sizeof(pktinfo));
        network::Address res{};
        res.ifindex = pktinfo.ipi_ifindex;
        res.len = sizeof(res.su.in);
        auto &sa = res.su.in;
        sa.sin_family = AF_INET;
        sa.sin_addr = pktinfo.ipi_addr;
        return res;
      }
    }
    return {};
  case AF_INET6:
    for (auto cmsg = CMSG_FIRSTHDR(msg); cmsg; cmsg = CMSG_NXTHDR(msg, cmsg)) {
      if (cmsg->cmsg_level == IPPROTO_IPV6 && cmsg->cmsg_type == IPV6_PKTINFO) {
        in6_pktinfo pktinfo;
        memcpy(&pktinfo, CMSG_DATA(cmsg), sizeof(pktinfo));
        network::Address res{};
        res.ifindex = pktinfo.ipi6_ifindex;
        res.len = sizeof(res.su.in6);
        auto &sa = res.su.in6;
        sa.sin6_family = AF_INET6;
        sa.sin6_addr = pktinfo.ipi6_addr;
        return res;
      }
    }
    return {};
  }
  return {};
}

size_t msghdr_get_udp_gro(msghdr *msg) {
  uint16_t gso_size = 0;

#ifdef UDP_GRO
  for (auto cmsg = CMSG_FIRSTHDR(msg); cmsg; cmsg = CMSG_NXTHDR(msg, cmsg)) {
    if (cmsg->cmsg_level == SOL_UDP && cmsg->cmsg_type == UDP_GRO) {
      memcpy(&gso_size, CMSG_DATA(cmsg), sizeof(gso_size));

      break;
    }
  }
#endif // UDP_GRO

  return gso_size;
}

void set_port(network::Address &dst, network::Address &src) {
  switch (dst.su.storage.ss_family) {
  case AF_INET:
    assert(AF_INET == src.su.storage.ss_family);
    dst.su.in.sin_port = src.su.in.sin_port;
    return;
  case AF_INET6:
    assert(AF_INET6 == src.su.storage.ss_family);
    dst.su.in6.sin6_port = src.su.in6.sin6_port;
    return;
  default:
    assert(0);
  }
}

bool addreq(const sockaddr *sa, const network::in_addr_union &iau) {
  switch (sa->sa_family) {
  case AF_INET:
    return memcmp(&reinterpret_cast<const sockaddr_in *>(sa)->sin_addr, &iau.in,
                  sizeof(iau.in)) == 0;
  case AF_INET6:
    return memcmp(&reinterpret_cast<const sockaddr_in6 *>(sa)->sin6_addr,
                  &iau.in6, sizeof(iau.in6)) == 0;
  default:
    assert(0);
    abort();
  }
}

} // namespace network