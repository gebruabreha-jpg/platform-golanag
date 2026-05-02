#ifndef EVENTHANDLER_H
#define EVENTHANDLER_H

#pragma once

class TimerWatcher;
class FdWatcher;

class EventHandler {
public:
  EventHandler() = default;
  ~EventHandler() = default;
  virtual void on_timeout(TimerWatcher *watcher) = 0;
  virtual void on_read_ready(FdWatcher *watcher) = 0;

private:
};

#endif