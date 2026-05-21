#ifndef FLATBUFFER_H
#define FLATBUFFER_H

#pragma once
#include <cstdint>
#include <cstdlib>


class FlatBuffer {
public:
  FlatBuffer() = delete;
  FlatBuffer(ssize_t size);
  FlatBuffer(const FlatBuffer& other);
  ~FlatBuffer();
  uint8_t* data();

  // the capacity for write
  ssize_t size() const;

  // the bytes in buffer to read
  ssize_t len() const;

  void set_data_len(ssize_t len);

private:
  uint8_t *data_ = nullptr;
  const ssize_t size_;
  ssize_t len_ = 0;
};

#endif