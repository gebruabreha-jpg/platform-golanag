#ifndef WORKER_H
#define WORKER_H
#include "ev.h"
#include <atomic>
#include <memory>
#include <sys/eventfd.h>
#include <thread>

#include "BaseJob.h"
#include "RingQueue.h"

#pragma once

namespace gladiator {

class Worker {
public:
  Worker(const char *name = "anonym");
  ~Worker();
  void add_job(BaseJob *job);
  void handle_jobs();
  struct ev_loop *get_event_loop() const;
  void bind_cpu(uint8_t cpu_id);
  void stop();

  bool is_stopped() const { return stopped_.load(); }

private:
  class StopJob : public BaseJob {
  public:
    StopJob(Worker *worker){};
    void execute() override;
  };

  void entry();
  static void static_entry(Worker *worker);
  // make sure the woker can handle concurrent startup of 30,000 users
  // that's why we set size of the job queue as 32K
  RingQueue<BaseJob *, (1 << 15)> job_queue_;
  int event_fd_;
  std::unique_ptr<std::thread> thread_ = nullptr;
  struct ev_loop *event_loop_ = nullptr;

  std::atomic_bool stopped_ = true;
};

} // namespace gladiator

#endif