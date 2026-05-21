#ifndef BASEJOB_H
#define BASEJOB_H
#pragma once

namespace gladiator {

class Worker;

class BaseJob {
public:
  friend class Worker;
  BaseJob() = default;
  virtual ~BaseJob() = default;
  virtual void execute() = 0;
  Worker *get_worker() { return worker_; }

private:
  Worker *worker_ = nullptr;
};

} // namespace gladiator

#endif