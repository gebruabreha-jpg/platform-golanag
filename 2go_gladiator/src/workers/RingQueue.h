/*
 * Author: boris.wang
 *
 * RingQueue can be used in producer-consumer mode.
 * It does not need lock for single producer push() and single consumer pop()
 */
#ifndef RING_QUEUE_H
#define RING_QUEUE_H

namespace gladiator {

constexpr bool is_power_of_two(const unsigned x) { return x && !(x & (x - 1)); }

template <typename T, unsigned size> class RingQueue {
  static_assert(is_power_of_two(size), "error: size is not power of 2");

public:
  RingQueue() : head_(0), tail_(0) {}
  ~RingQueue() = default;

  bool push(const T &element) {
    if (is_full()) {
      return false;
    }

    queue_[tail_] = element;
    tail_ = (tail_ + 1) & mask_;

    return true;
  }

  T &pop() {
    if (is_emtpy()) {
      throw(this);
    }
    T &element = queue_[head_];
    head_ = (head_ + 1) & mask_;
    return element;
  }

  T &front() {
    if (is_emtpy()) {
      throw(this);
    }
    return queue_[head_];
  }

  constexpr unsigned get_size() { return size - 1; }

  bool is_full() { return ((tail_ + 1) & mask_) == head_; }

  bool is_emtpy() { return tail_ == head_; }

private:
  T queue_[size];
  const int mask_ = size - 1;
  int head_;
  int tail_;
};

} // namespace gladiator

#endif