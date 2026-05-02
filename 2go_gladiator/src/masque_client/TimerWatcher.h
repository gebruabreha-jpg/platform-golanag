#ifndef TIMERWATCHER_H
#define TIMERWATCHER_H

#include <atomic>
#pragma once
#include <ev.h>
class EventHandler;

class TimerWatcher {
public:
  TimerWatcher() = delete;
  TimerWatcher(struct ev_loop *event_loop, EventHandler *handler);
  ~TimerWatcher();

  void start();
  void stop();

  void reset_timer(double seconds);
  void awake_timer();

  EventHandler* get_handler() const;

private:
  struct ev_loop *const event_loop_ = nullptr;
  EventHandler *const handler_ = nullptr;
  ev_timer timer_watcher_;
  std::atomic_bool stopped_ = false;
};

#endif