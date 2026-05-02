#include "Worker.h"
#include <cstdint>
#include <cstdio>
#include <ev.h>
#include <iostream>
#include <stdexcept>
#include <sys/types.h>
#include <unistd.h>

namespace gladiator {

void Worker::static_entry(Worker *worker) { worker->entry(); }

static void handle_evnet_fd_ready_to_read(struct ev_loop *loop, ev_io *w,
                                          int revents) {
  Worker *worker = static_cast<Worker *>(w->data);
  worker->handle_jobs();
}

Worker::Worker(const char *name) {
  event_loop_ = ev_loop_new(EVBACKEND_EPOLL | EVFLAG_NOENV);
  event_fd_ = eventfd(0, EFD_NONBLOCK);

  // spawn thread at the last
  thread_ = std::make_unique<std::thread>(Worker::static_entry, this);
  const pthread_t native_handle = thread_->native_handle();
  pthread_setname_np(native_handle, name);
}

void Worker::bind_cpu(uint8_t cpu_id) {
  const pthread_t native_handle = thread_->native_handle();

  cpu_set_t cpuset;
  CPU_ZERO(&cpuset);
  CPU_SET(cpu_id, &cpuset);

  if (pthread_setaffinity_np(native_handle, sizeof(cpu_set_t), &cpuset) != 0) {
    fprintf(stderr, "Error: pthread_setaffinity_np");
  }
}

Worker::~Worker() {
  thread_->join();
  ev_loop_destroy(event_loop_);
}

void Worker::entry() {
  // monitor event_fd
  ev_io event_fd_watcher;
  event_fd_watcher.data = this;

  ev_io_init(&event_fd_watcher, handle_evnet_fd_ready_to_read, event_fd_,
             EV_READ);
  ev_io_start(event_loop_, &event_fd_watcher);

  stopped_ = true;

  ev_run(event_loop_, 0);

  stopped_ = false;
}

void Worker::add_job(BaseJob *job) {
  job->worker_ = this;
  if (job_queue_.push(job)) {
    // write event_fd to awake consumer
    const uint64_t n = 1;
    if (auto nbytes = write(event_fd_, &n, sizeof(uint64_t)); nbytes < 0) {
      // do nothing
    }
  } else {
    fprintf(stderr, "Job Queue is full!");
    exit(0);
  }
}

void Worker::handle_jobs() {
  // read event_fd
  uint64_t n = 0;
  if (auto nbyes = read(event_fd_, &n, sizeof(uint64_t)); nbyes < 0) {
    // do nothing
  }
  while (!job_queue_.is_emtpy()) {
    auto *job = job_queue_.pop();
    if (job->worker_ != this) {
      throw std::runtime_error("can not execture job does not belong me!");
    }
    job->execute();
    delete job;
  }
}

struct ev_loop *Worker::get_event_loop() const {
  return event_loop_;
}

void Worker::StopJob::execute() { ev_break(worker_->event_loop_, EVBREAK_ALL); }

void Worker::stop() { add_job(new Worker::StopJob(this)); }

} // namespace gladiator