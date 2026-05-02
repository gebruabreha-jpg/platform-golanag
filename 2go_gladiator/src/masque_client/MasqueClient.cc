#include "MasqueClient.h"
#include "Config.h"
#include "Http3Connection.h"
#include "InnerLayer.h"
#include "MasqueStreamInfo.h"
#include "OuterLayer.h"
#include "PvdClient.h"
#include "QuicConnection.h"
#include "TimerWatcher.h"
#include "UdpLayer.h"
#include "Worker.h"
#include "debug.h"
#include <atomic>

#include "PM.h"
#include <chrono>
#include <random>

int MasqueClient::id_cnt_ = 0;
std::atomic<std::chrono::time_point<std::chrono::high_resolution_clock>>
    MasqueClient::earliest_start_time_{
        std::chrono::high_resolution_clock::now() + std::chrono::seconds(10)};
std::atomic<std::chrono::time_point<std::chrono::high_resolution_clock>>
    MasqueClient::latest_stop_time_{std::chrono::high_resolution_clock::now() -
                                    std::chrono::seconds(10)};

MasqueClient::MasqueClient(gladiator::Worker *worker,
                           const struct network::Address &local_addr,
                           const struct Config *config, PM *pm)
    : worker_(worker), config_(config), local_addr_(local_addr), pm_(pm) {

  id_ = ++MasqueClient::id_cnt_; // Monotonically increasing
}

MasqueClient::~MasqueClient() {

  delete outer_layer_;

  for (auto &inner_layer : inner_layers_) {
    delete inner_layer;
  }
  inner_layers_.clear();

  delete pvd_client_;
  delete idle_timer_watcher_;
  delete start_timer_watcher_;

  for (auto &udp_layer : udp_layers_) {
    delete udp_layer;
  }
  udp_layers_.clear();
}

const Config *MasqueClient::get_config() const { return config_; }

OuterLayer *MasqueClient::get_outer_layer() const { return outer_layer_; }
PvdClient *MasqueClient::get_pvd_client() const { return pvd_client_; }

struct ev_loop *MasqueClient::get_ev_loop() const {
  return worker_->get_event_loop();
}

void MasqueClient::async_start(double delay) {
  start_delay_ = delay;
  worker_->add_job(new StartJob(this));
}

void MasqueClient::suicide() { worker_->add_job(new DeleteClientJob(this)); }

void MasqueClient::start() {

  if (start_job_cancelled_.load()) {
    status_ = Status::STOPPED;
    return;
  }

  if (status_.load() != Status::INIT) {
    return;
  }

  start_timer_watcher_ = new TimerWatcher(get_ev_loop(), this);
  if (start_delay_) {
    auto delay = start_delay_ * id_;
    start_timer_watcher_->reset_timer(delay);

  } else {
    if (!config_->quiet) {
      ngtcp2::debug::log_printf(nullptr, "client[%0d] start with no delay ...",
                                id_);
    }
    start_protocol_layers();
  }
}

void MasqueClient::async_cancel_start_job() {
  worker_->add_job(new CancelStartJob(this));
}

void MasqueClient::cancel_start_job() {
  if (status_ == Status::INIT) {
    start_job_cancelled_ = true;
    status_ = Status::CANCELLING;
    start_timer_watcher_->stop();
  }
}

void MasqueClient::start_pvd_client() {
  if (pvd_client_) {
    return;
  }

  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%d] start PVD Client ...", id_);
  }
  // create the UDP and PvdClient
  auto udp = new UdpLayer(local_addr_.get_family(), this);
  udp->set_local_address(local_addr_);
  udp->set_remote_address(pvd_server_addr_);
  udp_layers_.push_back(udp);
  udp->start();

  pvd_client_ =
      new PvdClient(this, pvd_server_name_, local_addr_, pvd_server_addr_,
                    udp->get_max_udp_payload_size());

  pvd_client_->set_udp_layer(udp);
  udp->set_upper_layer(pvd_client_);

  pvd_client_->start();

  status_.store(Status::RUNNING);
}

void MasqueClient::start_outer_layer() {
  if (outer_layer_) {
    return;
  }

  if (!config_->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%d] start Outer layer ...", id_);
  }

  // create the UDP and OuterLayer
  auto udp = new UdpLayer(local_addr_.get_family(), this);
  udp->set_local_address(local_addr_);
  udp->set_remote_address(proxy_server_addr_);
  udp_layers_.push_back(udp);
  udp->start();

  outer_layer_ =
      new OuterLayer(this, proxy_server_name_, local_addr_, proxy_server_addr_,
                     udp->get_max_udp_payload_size());

  outer_layer_->set_udp_layer(udp);
  udp->set_upper_layer(outer_layer_);

  outer_layer_->start();
}

void MasqueClient::start_inner_layers() {
  if (!inner_layers_.empty()) {
    return;
  }

  // create the UDP layers
  for (auto &ele : host_port_to_address_map_) {

    auto &[host, port] = ele.first;
    auto &target_addr = ele.second;

    auto udp = new UdpLayer(local_addr_.get_family(), this);
    udp->set_local_address(local_addr_);
    udp->set_remote_address(target_addr);
    udp_layers_.push_back(udp);
    udp->start();

    const auto &requests =
        get_config()->target_servers_to_requests.at(ele.first);

    auto inner = new InnerLayer(this, host.c_str(), local_addr_, target_addr,
                                &requests, udp->get_max_udp_payload_size());

    inner_layers_.push_back(inner);

    inner->set_udp_layer(udp);
    udp->set_upper_layer(inner);

    inner->start();
  }
}

void MasqueClient::start_protocol_layers() {

  idle_timer_watcher_ = new TimerWatcher(get_ev_loop(), this);

  if (config_->need_query_pvd_server()) {
    start_pvd_client();
  } else if (config_->is_proxy_mode()) {
    start_outer_layer();
  } else {
    start_inner_layers();
  }

  status_.store(Status::RUNNING);

  start_time_ = std::chrono::high_resolution_clock::now();

  if (start_time_ < earliest_start_time_.load()) {
    earliest_start_time_.store(start_time_);
  }

  send_data();
}

void MasqueClient::async_stop() {
  if (status_.load() == Status::RUNNING || status_.load() == Status::TIMEOUT) {
    worker_->add_job(new StopJob(this));
  };
}
void MasqueClient::stop() {
  if (status_ != Status::STOPPED) {
    for (auto layer : inner_layers_) {
      if (!layer->is_stopped()) {
        layer->get_quic()->disconnect(false);
      }
    }

    if (outer_layer_ && !outer_layer_->is_stopped()) {
      outer_layer_->get_quic()->disconnect(false);
    }

    idle_timer_watcher_->stop();
    if (!config_->quiet) {
      ngtcp2::debug::log_printf(nullptr, "Client[%d] stop.", id_);
    }
    stop_time_ = std::chrono::high_resolution_clock::now();

    if (stop_time_ > latest_stop_time_.load()) {
      latest_stop_time_.store(stop_time_);
    }
  }
}

void MasqueClient::on_timeout(TimerWatcher *watcher) {

  if (watcher == idle_timer_watcher_) {
    watcher->stop();
    set_status(Status::TIMEOUT);
    stop_time_ = std::chrono::high_resolution_clock::now();
    if (stop_time_ > latest_stop_time_.load()) {
      latest_stop_time_.store(stop_time_);
    }
    async_stop();
  } else if (watcher == start_timer_watcher_) {
    watcher->stop();
    if (!start_job_cancelled_.load()) {
      start_protocol_layers();
    } else {
      status_ = Status::STOPPED;
    }
  }
}

void MasqueClient::print_stats() {

  // if (inner_layer_) {
  //   fprintf(stderr, "Client[%d],Inner layer\n", id_);
  // }

  // if (outer_layer_) {
  //   fprintf(stderr, "Client[%d],Tunnel layer\n", id_);
  // }

  for (auto &udp_layer : udp_layers_) {
    auto udp_stats = udp_layer->get_stats();
    fprintf(stderr,
            "UDP (%s - %s) stats:\n\ttx_pkts: %lu\n\ttx_bytes: %lu\n\trx_pkts: "
            "%lu\n\trx_bytes: %lu\n",
            udp_layer->get_local_address().to_string().c_str(),
            udp_layer->get_remote_address().to_string().c_str(),
            udp_stats.total_tx_udp_pkts, udp_stats.total_tx_bytes,
            udp_stats.total_rx_udp_pkts, udp_stats.total_rx_bytes);
  }
}

void MasqueClient::add_target_server(const network::Address &target_server_addr,
                                     const std::string &host_name,
                                     const std::string &port) {
  host_port_to_address_map_[std::make_tuple(host_name, port)] =
      target_server_addr;
}

void MasqueClient::set_proxy_server_address(
    const network::Address &proxy_server_addr, const char *host_name) {
  proxy_server_addr_ = proxy_server_addr;
  proxy_server_name_ = host_name;
}

void MasqueClient::set_pvd_server_address(
    const network::Address &pvd_server_addr, const char *host_name) {
  pvd_server_addr_ = pvd_server_addr;
  pvd_server_name_ = host_name;
}

void MasqueClient::send_data() {

  if (pvd_client_) {
    pvd_client_->send_data();
  }

  // the data will flow down to low layer by order
  for (auto &inner_layer : inner_layers_) {
    inner_layer->send_data();
  }
  if (outer_layer_) {
    outer_layer_->send_data();
  }
}

void MasqueClient::on_tunnel_setup_ready_async(int64_t stream_id,
                                               MasqueStreamInfo &info) {
  worker_->add_job(new InnerLayerStartJob(this, info, stream_id));
}

void MasqueClient::on_tunnel_setup_ready(int64_t stream_id,
                                         MasqueStreamInfo &info) {

  // check outer layer's status again
  if (outer_layer_ && outer_layer_->is_stopped()) {
    ngtcp2::debug::log_printf(
        nullptr, "Tunnel has shutdonw, do not create any inner layer!");
    return;
  }

  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Tunnel for stream %ld is ready!",
                              stream_id);
  }

  auto host_port = std::make_tuple(info.host, info.port);

  const auto &requests = get_config()->target_servers_to_requests.at(host_port);

  auto &target_addr = host_port_to_address_map_.at(host_port);

  auto inner =
      new InnerLayer(this, info.host.c_str(), local_addr_, target_addr,
                     &requests, udp_layers_[0]->get_max_udp_payload_size());

  info.inner_layer = inner;

  // store the scid
  inner->get_quic()->get_scid(info.scid);

  inner->set_tunnel_mode(true);
  inner->set_masque_stream_id(stream_id);
  inner_layers_.push_back(inner);
  inner->start();
  inner->send_data();
}

PM *MasqueClient::get_pm() { return pm_; }

void MasqueClient::reset_idle_timer() {
  idle_timer_watcher_->reset_timer(config_->timeout / 1E09);
}

void MasqueClient::on_stop() {

  stop_time_ = std::chrono::high_resolution_clock::now();
  if (stop_time_ > latest_stop_time_.load()) {
    latest_stop_time_.store(stop_time_);
  }
  idle_timer_watcher_->stop();
  start_timer_watcher_->stop();
  status_.store(Status::STOPPED);
}

void MasqueClient::stop_udp_layers() {
  for (auto &udp : udp_layers_) {
    udp->stop();
  }
}

void MasqueClient::set_inner_layers_stopped() {
  for (auto inner : inner_layers_) {
    inner->set_stopped();
  }
}

void MasqueClient::reschedule_write_datagram(QuicConnection *conn,
                                             FlatBuffer *buffer) {
  worker_->add_job(new SendDatagramJob(conn, buffer));
}

void MasqueClient::on_http3_closed(TransportLayer *layer) {
  // disconnect the quic connection
  if (!get_config()->quiet) {
    ngtcp2::debug::log_printf(nullptr, "Client[%d] %s on_http3_closed", id_,
                              layer->get_layer_str());
  }

  layer->get_quic()->disconnect(false);
}

void MasqueClient::on_quic_connection_closed(TransportLayer *layer) {

  if (status_.load() == Status::STOPPED) {
    return;
  }

  if (layer->is_pvd_client()) {

    pvd_client_->stop();
    stop_udp_layers();

    if (!get_config()->quiet) {
      ngtcp2::debug::log_printf(nullptr, "Client[%d] PVD Client closed!", id_);
    }

    get_pm()->inc_closed_connects();

    if (!get_config()->pvd_only) {
      start_outer_layer();
      if (outer_layer_) {
        outer_layer_->send_data();
      }
    } else {
      on_stop();
    }

    return;
  }

  if (layer->is_outer_layer()) {
    if (layer->get_quic()->is_handshake_completed()) {
      get_pm()->inc_tunnel_closed_connects();
    }

    for (auto &inner : inner_layers_) {
      inner->set_stopped();
    }

    stop_udp_layers();
    on_stop();
    return;
  }

  if (layer->is_inner_layer()) {

    if (layer->get_quic()->is_handshake_completed()) {
      get_pm()->inc_closed_connects();
    }

    bool all_stopped = true;

    for (auto &inner : inner_layers_) {
      if (!inner->is_stopped()) {
        all_stopped = false;
        break;
      }
    }

    if (all_stopped) {

      if (outer_layer_) {
        outer_layer_->get_http3()->disconnect();
      } else {
        stop_udp_layers();
        on_stop();
      }
    }
  }
}

double MasqueClient::get_total_duration_ms() {

  std::chrono::duration<double, std::milli> duration_ms =
      latest_stop_time_.load() - earliest_start_time_.load();

  if (duration_ms.count() < 0) {
    return 0;
  }

  return duration_ms.count();
}

void MasqueClient::on_quic_timeout(bool handshake_completed,
                                   TransportLayer *layer) {

  if (layer->is_inner_layer() || layer->is_pvd_client()) {
    pm_->inc_inner_quic_timeout(handshake_completed);
  } else if (layer->is_outer_layer()) {
    pm_->inc_tunnel_quic_timeout(handshake_completed);
  }
}

MasqueClient::SendDatagramJob::SendDatagramJob(QuicConnection *conn,
                                               FlatBuffer *buffer)
    : conn_(conn) {
  buffer_ = new FlatBuffer(*buffer);
}

MasqueClient::SendDatagramJob::~SendDatagramJob() { delete buffer_; }

void MasqueClient::SendDatagramJob::execute() {
  conn_->write_datagram_data(buffer_);
}
