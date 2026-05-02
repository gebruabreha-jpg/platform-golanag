#include "TimerWatcher.h"
#include "EventHandler.h"
#include "debug.h"
#include <ev.h>
namespace {
void timeout_callback(struct ev_loop *loop, ev_timer *watcher, int revents) {
  TimerWatcher *timer_watcher = reinterpret_cast<TimerWatcher *>(watcher->data);
  timer_watcher->get_handler()->on_timeout(timer_watcher);
}
} // namespace

TimerWatcher::TimerWatcher(struct ev_loop *event_loop, EventHandler *handler)
    : event_loop_(event_loop), handler_(handler) {
  timer_watcher_.data = this;
  ev_timer_init(&timer_watcher_, timeout_callback, 0.0, 0.0);
}

TimerWatcher::~TimerWatcher() {}

void TimerWatcher::start() {
  ev_timer_start(event_loop_, &timer_watcher_);
  stopped_ = false;
};
void TimerWatcher::stop() {
  if (not stopped_) {
    ev_timer_stop(event_loop_, &timer_watcher_);
    stopped_ = true;
  }
}

void TimerWatcher::reset_timer(double seconds) {
  timer_watcher_.repeat = seconds;
  ev_timer_again(event_loop_, &timer_watcher_);
  stopped_ = false;
  // the .after has been reset, the timer will repeat after .repeat
};
void TimerWatcher::awake_timer() {
  ev_feed_event(event_loop_, &timer_watcher_, EV_TIMER);
};

EventHandler *TimerWatcher::get_handler() const { return handler_; }