#include "IPAllocator.h"
#include <gtest/gtest.h>

TEST(IPV4Pool, Test_01) {
  IPAllocator *ip_allocator = IPAllocator::get_instance();
  ip_allocator->init_ipv4_pool("172.0.0.5/32", 3);
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.5");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.5");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.5");
}


TEST(IPV4Pool, Test_02) {
  IPAllocator *ip_allocator = IPAllocator::get_instance();
  ip_allocator->init_ipv4_pool("172.0.0.1/24", 3);
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.1");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.2");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.3");

  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.1");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.2");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.3");

  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.1");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.2");
  ASSERT_EQ(ip_allocator->allocate_ipv4(), "172.0.0.3");
}