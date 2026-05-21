#include "FlatBuffer.h"
#include <cassert>
#include <cstddef>
#include <cstdint>
#include <cstdio>
#include <cstring>

FlatBuffer::FlatBuffer(ssize_t size) : size_(size) {
  data_ = new uint8_t[size];
}

FlatBuffer::FlatBuffer(const FlatBuffer& other):size_(other.size()) {
 data_ = new uint8_t[size_];
 memcpy(data_, other.data_, size_);
 len_ = other.len();
}

FlatBuffer::~FlatBuffer() { delete [] data_; }

uint8_t *FlatBuffer::data() { return data_; }

ssize_t FlatBuffer::size() const { return size_; }

ssize_t FlatBuffer::len() const { return len_; }

void FlatBuffer::set_data_len(ssize_t len) {
  assert(len <= size_);
  len_ = len;
}