#ifndef MASQUECLIENT_H
#define MASQUECLIENT_H

#include "FdWatcher.h"
#include "HttpRequest.h"
#include "TimerWatcher.h"
#include <atomic>
#include <memory>
#include <ngtcp2/ngtcp2.h>
#include <string>
#pragma once
#include "BaseJob.h"
#include "EventHandler.h"
#include "network.h"
#include <chrono>
#include <cstdint>
#include <cstdlib>
#include <map>
#include <vector>
namespace gladiator {
class Worker;
}
class UdpLayer;
class InnerLayer;
class OuterLayer;
class TimerWatcher;
class PM;
class PvdClient;

namespace network {
struct Address;
}

class TransportLayer;
class MasqueStreamInfo;
class QuicConnection;
class FlatBuffer;

class MasqueClient : EventHandler {
public:
  MasqueClient() = delete;
  MasqueClient(const MasqueClient &other) = delete;

  MasqueClient(gladiator::Worker *worker,
               const struct network::Address &local_addr,
               const struct Config *config, class PM *pm = nullptr);

  virtual ~MasqueClient();

  enum class Status : int { INIT, RUNNING, STOPPED, TIMEOUT, CANCELLING };

  Status get_status() const { return status_.load(); }
  void set_status(Status status) { status_.store(status); }

  void async_start(double delay = 0.0);
  void start();

  void suicide();

  // stop the client actively
  void async_stop();
  void stop();

  struct ev_loop *get_ev_loop() const;

  void on_quic_connection_closed(TransportLayer *layer);
  void on_quic_timeout(bool handshake_completed, TransportLayer *layer);

  void on_http3_closed(TransportLayer *layer);
  void on_stop();

  // interfaces of EventHandler
  void on_timeout(TimerWatcher *watcher) final;
  void on_read_ready(FdWatcher *watcher) final{}; // do not care, no fd wather

  void on_tunnel_setup_ready_async(int64_t stream_id, MasqueStreamInfo &info);
  void on_tunnel_setup_ready(int64_t stream_id, MasqueStreamInfo &info);

  void print_stats();

  void add_target_server(const network::Address &target_server_addr,
                         const std::string &host_name, const std::string &port);

  void set_proxy_server_address(const network::Address &proxy_server_addr,
                                const char *host_name);

  void set_pvd_server_address(const network::Address &pvd_server_addr,
                              const char *host_name);

  void send_data();

  const Config *get_config() const;

  OuterLayer *get_outer_layer() const;
  PvdClient *get_pvd_client() const;
  std::string get_id() const { return std::to_string(id_); }

  PM *get_pm();

  void reset_idle_timer();

  double get_duration_ms() const {
    std::chrono::duration<double, std::milli> duration_ms =
        stop_time_ - start_time_;
    return duration_ms.count();
  }

  static double get_total_duration_ms();

  const network::Address &get_local_addr() const { return local_addr_; }

  gladiator::Worker *get_worker() const { return worker_; }

  void cancel_start_job();
  void async_cancel_start_job();

  void set_inner_layers_stopped();

  void reschedule_write_datagram(QuicConnection *conn, FlatBuffer *buffer);

private:
  void start_protocol_layers();
  void start_pvd_client();
  void start_outer_layer();
  void start_inner_layers();

  void stop_udp_layers();

  gladiator::Worker *worker_ = nullptr;
  const Config *config_ = nullptr;

  const struct network::Address local_addr_;

  std::map<std::tuple<std::string, std::string>, struct network::Address>
      host_port_to_address_map_;

  struct network::Address proxy_server_addr_;
  const char *proxy_server_name_ = nullptr;

  struct network::Address pvd_server_addr_;
  const char *pvd_server_name_ = nullptr;

  // represent the connections from this user to the target servers
  std::vector<InnerLayer *> inner_layers_;

  OuterLayer *outer_layer_ = nullptr;
  PvdClient *pvd_client_ = nullptr;

  // in proxy mode, there is only one udp layer
  std::vector<UdpLayer *> udp_layers_;

  TimerWatcher *idle_timer_watcher_ = nullptr;

  TimerWatcher *start_timer_watcher_ = nullptr;

  std::atomic<Status> status_{Status::INIT};

  int id_ = 0;

  PM *pm_ = nullptr;

  static int id_cnt_; // for client id allocation. not thread-safe

  double start_delay_ = 0; // delay in seconds. 0.001 means 1 ms.

  std::atomic_bool start_job_cancelled_ = false;

  std::chrono::time_point<std::chrono::high_resolution_clock> start_time_ =
      std::chrono::high_resolution_clock::now();
  std::chrono::time_point<std::chrono::high_resolution_clock> stop_time_ =
      std::chrono::high_resolution_clock::now();

  static std::atomic<
      std::chrono::time_point<std::chrono::high_resolution_clock>>
      earliest_start_time_;
  static std::atomic<
      std::chrono::time_point<std::chrono::high_resolution_clock>>
      latest_stop_time_;

  // private sub classes
  class StartJob : public gladiator::BaseJob {
  public:
    StartJob(MasqueClient *client) : client_(client){};
    void execute() override { client_->start(); }

  private:
    MasqueClient *client_ = nullptr;
  };

  class CancelStartJob : public gladiator::BaseJob {
  public:
    CancelStartJob(MasqueClient *client) : client_(client){};
    void execute() override { client_->cancel_start_job(); }

  private:
    MasqueClient *client_ = nullptr;
  };

  class DeleteClientJob : public gladiator::BaseJob {
  public:
    DeleteClientJob(MasqueClient *client) : client_(client){};
    void execute() override {
      client_->stop();
      delete client_;
    }

  private:
    MasqueClient *client_ = nullptr;
  };

  class StopJob : public gladiator::BaseJob {
  public:
    StopJob(MasqueClient *client) : client_(client){};
    void execute() override { client_->stop(); }

  private:
    MasqueClient *client_ = nullptr;
  };

  class InnerLayerStartJob : public gladiator::BaseJob {
  public:
    InnerLayerStartJob(MasqueClient *client, MasqueStreamInfo &info, int64_t id)
        : client_(client), stream_info_(info), stream_id_(id){};
    void execute() override {
      client_->on_tunnel_setup_ready(stream_id_, stream_info_);
    }

  private:
    MasqueClient *client_ = nullptr;
    MasqueStreamInfo &stream_info_;
    int64_t stream_id_;
  };

  class SendDatagramJob : public gladiator::BaseJob {
  public:
    SendDatagramJob(QuicConnection *conn, FlatBuffer *buffer);

    ~SendDatagramJob() override;

    void execute() override;

  private:
    QuicConnection *conn_ = nullptr;
    FlatBuffer *buffer_ = nullptr;
  };
};

#endif
