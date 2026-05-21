#include "FdWatcher.h"
#include "EventHandler.h"

#include <cstdio>

namespace {
void read_ready_callback(struct ev_loop *loop, ev_io *watcher, int revents) {

  FdWatcher *fd_watcher = reinterpret_cast<FdWatcher *>(watcher->data);
  fd_watcher->get_handler()->on_read_ready(fd_watcher);
}
} // namespace

FdWatcher::FdWatcher(struct ev_loop *event_loop, int fd, EventHandler *handler)
    : event_loop_(event_loop), fd_(fd), handler_(handler) {
  fd_watcher_.data = this;
  ev_io_init(&fd_watcher_, read_ready_callback, fd, EV_READ);
}

FdWatcher::~FdWatcher() {}

void FdWatcher::start() { ev_io_start(event_loop_, &fd_watcher_); }
void FdWatcher::stop() { ev_io_stop(event_loop_, &fd_watcher_); }

EventHandler *FdWatcher::get_handler() const { return handler_; }
