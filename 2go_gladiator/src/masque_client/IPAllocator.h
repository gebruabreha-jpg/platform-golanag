#ifndef IPALLOCATOR_H
#define IPALLOCATOR_H

#include "network.h"
#pragma once
#include <string>
#include <tuple>
#include <vector>

// singleton
class IPAllocator {
public:
  ~IPAllocator() = default;

  IPAllocator(const IPAllocator &) = delete;
  IPAllocator &operator=(const IPAllocator &) = delete;

  void init_ipv4_pool(const std::string &network, int n); // not thread-safe
  void init_ipv6_pool(const std::string &network, int n); // not thread-safe

  std::string allocate_ipv4(); // not thread-safe
  std::string allocate_ipv6(); // not thread-safe

  static IPAllocator *get_instance();

private:
  IPAllocator() = default;

  std::vector<std::string> ipv4_pool_;
  int allocated_ipv4_num_ = 0;
  std::vector<std::string> ipv6_pool_;
  int allocated_ipv6_num_ = 0;

  static IPAllocator *instance_;
};

#endif