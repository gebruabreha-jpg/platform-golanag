#include "IPAllocator.h"
#include "network.h"
IPAllocator *IPAllocator::instance_ = nullptr;

IPAllocator *IPAllocator::get_instance() {
  if (instance_ == nullptr) {
    instance_ = new IPAllocator();
  }
  return instance_;
}

void IPAllocator::init_ipv4_pool(const std::string &network, int n) {
  ipv4_pool_ = network::get_ipv4_in_cidr_network(network, n);
}
void IPAllocator::init_ipv6_pool(const std::string &network, int n) {
  ipv6_pool_ = network::get_ipv6_in_cidr_network(network, n);
}

std::string IPAllocator::allocate_ipv4() {
  if (ipv4_pool_.empty()) {
    return "0.0.0.0";
  }

  const auto array_size = ipv4_pool_.size();

  const auto idx = array_size - (allocated_ipv4_num_++ % array_size) - 1;

  return ipv4_pool_[idx];
}
std::string IPAllocator::allocate_ipv6() {
  if (ipv6_pool_.empty()) {
    return "::";
  }

  const auto array_size = ipv6_pool_.size();

  const auto idx = array_size - (allocated_ipv6_num_++ % array_size) - 1;

  return ipv6_pool_[idx];
}
