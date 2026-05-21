#ifndef FDWATCHER_H
#define FDWATCHER_H

#include "EventHandler.h"
#include <ev.h>
#pragma once

class FdWatcher {
public:
  FdWatcher() = delete;
  FdWatcher(struct ev_loop *event_loop, int fd, EventHandler *handler);
  ~FdWatcher();

  void start();
  void stop();
  EventHandler* get_handler() const;

private:
  struct ev_loop *const event_loop_ = nullptr;
  int fd_ = -1;
  EventHandler *const handler_ = nullptr;
  ev_io fd_watcher_;
};

#endif