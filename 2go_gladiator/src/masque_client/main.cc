#include "CliOptionParser.h"
#include "Config.h"
#include "IPAllocator.h"
#include "MasqueClient.h"
#include "PM.h"
#include "Worker.h"
#include "debug.h"
#include "network.h"
#include "tls_client_context_quictls.h"
#include "util.h"
#include <chrono>
#include <cstdio>
#include <cstdlib>
#include <cxxabi.h>
#include <dlfcn.h>
#include <execinfo.h>
#include <fstream>
#include <iostream>
#include <memory>
#include <signal.h>
#include <sys/sysinfo.h>
#include <thread>
#include <vector>

std::ofstream keylog_file;

namespace {
bool force_quit_g = false;
void sigint_handler(int signal) {
  printf("Received SIGINT (Ctrl+C).\n");
  force_quit_g = true;
}

void print_calltrace(int sig) {
  static constexpr int STACK_DEEP = 20;
  void *array[STACK_DEEP];
  size_t size;

  // Get pointers to all frames on the stack
  size = backtrace(array, STACK_DEEP);

  // Print out all the frames to stderr with demangled symbols
  fprintf(stderr, "Error: signal %d:\n", sig);
  for (size_t i = 0; i < size; i++) {
    Dl_info info;
    if (dladdr(array[i], &info) && info.dli_sname) {
      int status;
      char *demangled_symbol =
          abi::__cxa_demangle(info.dli_sname, 0, 0, &status);
      if (status == 0) {
        fprintf(stderr, "%lu: %s\n", i, demangled_symbol);
        std::free(
            demangled_symbol); // Don't forget to free the allocated memory
      } else {
        fprintf(stderr, "%lu: Demangling failed: %s\n", i, info.dli_sname);
      }
    } else {
      fprintf(stderr, "%lu: Unknown symbol\n", i);
    }
  }
  exit(1);
}

MasqueClient *create_client(gladiator::Worker *worker, int net_family,
                            const network::Address &local_addr,
                            const network::Address &proxy_address,
                            const network::Address &pvd_server_address,
                            const Config &app_config, PM &pm) {

  auto client = new MasqueClient(worker, local_addr, &app_config, &pm);

  for (const auto &ele : app_config.target_servers_to_requests) {
    auto &[host, port] = ele.first;
    network::Address target_address =
        network::get_remote_address(host.c_str(), port.c_str(), net_family);

    if (!app_config.quiet) {
      ngtcp2::debug::log_printf(
          nullptr, "Client [%s] local %s target: %s %s",
          client->get_id().c_str(), local_addr.to_string().c_str(),
          target_address.to_string().c_str(),
          local_addr.get_family() == AF_INET ? "IPv4" : "IPv6");

      for (auto &req : ele.second) {
        std::cerr << req << std::endl;
      }
    }

    client->add_target_server(target_address, host, port);
  }

  if (app_config.need_query_pvd_server()) {

    client->set_pvd_server_address(pvd_server_address,
                                   app_config.pvd_server.c_str());
  } else if (app_config.is_proxy_mode()) {

    client->set_proxy_server_address(proxy_address,
                                     app_config.proxy_server.c_str());
  }
  return client;
}

void print_counters(Config &app_config, PM &pm, double elapsed_seconds) {
  static int64_t prev_rx_bytes = 0;
  static int64_t prev_tx_bytes = 0;

  int64_t cur_rx_bytes = pm.get_udp_rx_bytes();
  int64_t cur_tx_bytes = pm.get_udp_tx_bytes();

  double rx_speed = ((cur_rx_bytes - prev_rx_bytes) << 3) / elapsed_seconds;
  double tx_speed = ((cur_tx_bytes - prev_tx_bytes) << 3) / elapsed_seconds;

  auto [rx_unit, rx_value] =
      ngtcp2::util::convert_data_with_suitable_unit(rx_speed);
  auto [tx_unit, tx_value] =
      ngtcp2::util::convert_data_with_suitable_unit(tx_speed);

  auto [rx_byte_unit, rx_byte_value] =
      ngtcp2::util::convert_data_with_suitable_unit(cur_rx_bytes, false);
  auto [tx_byte_unit, tx_byte_value] =
      ngtcp2::util::convert_data_with_suitable_unit(cur_tx_bytes, false);

  if (app_config.is_proxy_mode() ||
      (app_config.need_query_pvd_server() && not app_config.pvd_only)) {

    fprintf(stdout, "%8.2f%c%8.2f%c%8.2f%c%8.2f%c%8ld%8ld%8ld%8ld%8ld%8ld\n",
            rx_value, rx_unit, tx_value, tx_unit, rx_byte_value, rx_byte_unit,
            tx_byte_value, tx_byte_unit, pm.get_created_conns(),
            pm.get_closed_conns(), pm.get_opened_streams(),
            pm.get_closed_streams(), pm.get_tunnel_created_conns(),
            pm.get_tunnel_closed_conns());

  } else {

    fprintf(stdout, "%8.2f%c%8.2f%c%8.2f%c%8.2f%c%8ld%8ld%8ld%8ld\n", rx_value,
            rx_unit, tx_value, tx_unit, rx_byte_value, rx_byte_unit,
            tx_byte_value, tx_byte_unit, pm.get_created_conns(),
            pm.get_closed_conns(), pm.get_opened_streams(),
            pm.get_closed_streams());
  }
  prev_rx_bytes = cur_rx_bytes;
  prev_tx_bytes = cur_tx_bytes;
}

} // namespace

int main(int argc, char *argv[]) {
  // signal(SIGSEGV, print_calltrace);
  signal(SIGABRT, print_calltrace);
  signal(SIGINT, sigint_handler);
  Config app_config{};
  PM pm;

  CliOptionParser cli_option(app_config);
  if (!cli_option.parse_cli(argc, const_cast<const char **>(argv))) {
    exit(EXIT_FAILURE);
  }

  int family = AF_INET;

  IPAllocator *ip_allocator = IPAllocator::get_instance();

  if (!app_config.ipv4_pool.empty()) {
    ip_allocator->init_ipv4_pool(app_config.ipv4_pool, app_config.clients_num);
    family = AF_INET;
  }

  if (!app_config.ipv6_pool.empty()) {
    ip_allocator->init_ipv6_pool(app_config.ipv6_pool, app_config.clients_num);
    family = AF_INET6;
  }

  // TLS setting
  TLSClientContext tls_ctx(app_config);
  if (tls_ctx.init(nullptr, nullptr) != 0) {
    exit(EXIT_FAILURE);
  }

  app_config.tls_context = &tls_ctx;

  // for wireshark decode usage
  auto keylog_filename = getenv("SSLKEYLOGFILE");
  if (keylog_filename) {
    keylog_file.open(keylog_filename, std::ios_base::app);
    if (keylog_file) {
      tls_ctx.enable_keylog();
    }
  }

  if (ngtcp2::util::generate_secret(app_config.static_secret.data(),
                                    app_config.static_secret.size()) != 0) {
    std::cerr << "Unable to generate static secret" << std::endl;
    exit(EXIT_FAILURE);
  }

  const int num_cpus = get_nprocs();

  // create worker threads
  std::vector<gladiator::Worker *> worker_pool;
  for (int i = 0; i < app_config.workers_num; i++) {
    char worker_name[32];
    snprintf(worker_name, sizeof(worker_name), "worker[%d]", i);
    auto worker = new gladiator::Worker(worker_name);
    worker->bind_cpu(i % num_cpus);
    worker_pool.emplace_back(worker);
  }

  std::vector<MasqueClient *> clients;
  network::Address proxy_address;
  network::Address pvd_server_address;

  if (app_config.need_query_pvd_server()) {
    pvd_server_address =
        network::get_remote_address(app_config.pvd_server.c_str(),
                                    app_config.pvd_server_port.c_str(), family);
  } else if (app_config.is_proxy_mode()) {
    proxy_address = network::get_remote_address(
        app_config.proxy_server.c_str(), app_config.proxy_server_port.c_str(),
        family);
  }

  for (int i = 0; i < app_config.clients_num; i++) {

    auto local_ip_str = family == AF_INET ? ip_allocator->allocate_ipv4()
                                          : ip_allocator->allocate_ipv6();
    auto local_addr = network::Address(family, local_ip_str.c_str(), htons(0));

    auto client = create_client(worker_pool[i % app_config.workers_num], family,
                                local_addr, proxy_address, pvd_server_address,
                                app_config, pm);

    clients.emplace_back(client);

    const double delay = 1.0 / app_config.new_conn_per_second;
    client->async_start(delay);
  }

  ngtcp2::debug::log_printf(nullptr, "Wait all clients close.\n");

  const char *fg_green_color = "\u001b[32;1m";
  const char *fg_red_color = "\u001b[31;1m";
  const char *reset_color = "\u001b[0m";
  //const char *bold = "\u001b[1m";
  const char *reversed_color = "\u001b[7m";

  if (app_config.is_proxy_mode() ||
      (app_config.need_query_pvd_server() && !app_config.pvd_only)) {

    fprintf(stdout, "%s%16s%16s%16s%16s%16s%s\n", reversed_color,
            "UDP Througput", "UDP bytes", "    INNER CONNECTIONS",
            "   QUIC STREAMS  ", " OUTER CONNECTIONS ", reset_color);
    fprintf(stderr, "%s%8s%8s%8s%8s%8s%8s%8s%8s%8s%8s%s\n",
            fg_green_color, // reversed_color,
            "    RX ", "    TX ", "    RX ", "    TX ", "    CREATED",
            "   CLOSED", "   CREATED", "  CLOSED", "  CREATED", "  CLOSED  ",
            reset_color);

  } else {

    fprintf(stdout, "%s%16s%16s%16s%16s%s\n", reversed_color, "UDP Througput",
            "UDP bytes", "    QUIC CONNECTIONS", "    QUIC STREAMS  ",
            reset_color);
    fprintf(stdout, "%s%8s%8s%8s%8s%8s%8s%8s%8s%s\n",
            fg_green_color, // reversed_color,
            "     RX ", "     TX ", "     RX ", "     TX ", "    CREATED",
            "   CLOSED", "  CREATED", "   CLOSED", reset_color);
  }

  bool all_clients_stopped = false;
  int total_seconds = 0;
  bool recreate_client = app_config.duration > 0;

  struct timespec prev_ts;
  clock_gettime(CLOCK_MONOTONIC, &prev_ts);

  do {

    usleep(1E6); // 1 second sleep
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);

    double elapsed_seconds = (ts.tv_sec - prev_ts.tv_sec) +
                             (double)(ts.tv_nsec - prev_ts.tv_nsec) / 1E9;

    print_counters(app_config, pm, elapsed_seconds);

    if (recreate_client) {
      // if there is a client stopped or timeout then delete it and create a new
      // one to replace it.
      for (size_t i = 0; i < clients.size(); i++) {
        auto client = clients[i];
        if (client->get_status() == MasqueClient::Status::STOPPED ||
            client->get_status() == MasqueClient::Status::TIMEOUT) {
          auto new_client = create_client(
              client->get_worker(), family, client->get_local_addr(),
              proxy_address, pvd_server_address, app_config, pm);
          // replace slot with the new client
          clients[i] = new_client;
          // start the new client to replace the old one
          // it's async call and executed in the worker thread.
          new_client->async_start();
          client->suicide();
        }
      }
    }

    all_clients_stopped = true;
    for (auto *client : clients) {
      const auto stopped =
          client->get_status() == MasqueClient::Status::STOPPED ||
          client->get_status() == MasqueClient::Status::TIMEOUT ||
          client->get_status() == MasqueClient::Status::CANCELLING;
      all_clients_stopped &= stopped;
    }

    if (recreate_client and ++total_seconds > app_config.duration) {
      recreate_client = false;
      for (auto client : clients) {
        if (client->get_status() == MasqueClient::Status::INIT) {
          client->async_cancel_start_job();
        } else if (client->get_status() == MasqueClient::Status::RUNNING) {
          client->async_stop();
        }
      }
    }

    clock_gettime(CLOCK_MONOTONIC, &prev_ts);

  } while (!force_quit_g && all_clients_stopped == false);

  auto [unit, value] = ngtcp2::util::convert_data_with_suitable_unit(
      pm.get_user_rx_bytes(), false);

  ngtcp2::debug::log_printf(nullptr,
                            "Total of %s%.1f%c%s bytes of "
                            "user payload data received in %s%.3f(ms)%s",
                            fg_green_color, value, unit, reset_color,
                            fg_green_color,
                            MasqueClient::get_total_duration_ms(), reset_color);

  if (app_config.is_proxy_mode() || app_config.need_query_pvd_server()) {

    fprintf(stdout, "QUIC-Tunnel packets Rx: %s%ld%s Tx: %s%ld%s\n",
            fg_green_color, pm.get_quic_tunnel_rx_packets(), reset_color,
            fg_green_color, pm.get_quic_tunnel_tx_packets(), reset_color);
    if (app_config.request_quic_forwarding) {
      fprintf(stdout, "QUIC-Forward packets Rx: %s%ld%s Tx: %s%ld%s\n",
              fg_green_color, pm.get_quic_forwarding_rx_packets(), reset_color,
              fg_green_color, pm.get_quic_forwarding_tx_packets(), reset_color);
    }
  }

  auto tunnel_handshake_timeout_times = pm.get_tunnel_quic_timeout(false);
  if (tunnel_handshake_timeout_times) {

    fprintf(stdout,
            "There are %s%ld%s QUIC connections fail to setup with the MASQUE "
            "Proxy.\n",
            fg_red_color, tunnel_handshake_timeout_times, reset_color);
  }

  auto tunnel_idle_timeout_times = pm.get_tunnel_quic_timeout(true);
  if (tunnel_idle_timeout_times) {

    fprintf(stdout,
            "There are %s%ld%s QUIC connections idle timeout with the MASQUE "
            "Proxy.\n",
            fg_red_color, tunnel_idle_timeout_times, reset_color);
  }

  auto inner_handshake_timeout_times = pm.get_inner_quic_timeout(false);
  if (inner_handshake_timeout_times) {

    fprintf(stdout,
            "There are %s%ld%s QUIC connections fail to setup with the Target "
            "Server.\n",
            fg_red_color, inner_handshake_timeout_times, reset_color);
  }

  auto inner_idle_timeout_times = pm.get_inner_quic_timeout(true);
  if (inner_idle_timeout_times) {

    fprintf(stdout,
            "There are %s%ld%s QUIC connections idle timeout with the Target "
            "Server.\n",
            fg_red_color, inner_idle_timeout_times, reset_color);
  }

  for (auto worker : worker_pool) {
    worker->stop();
  }

  // wait all worker stop
  while (true) {
    bool all_workers_stopped = true;
    for (auto worker : worker_pool) {
      if (worker->is_stopped()) {
        all_workers_stopped = false;
      }
    }

    if (all_workers_stopped) {
      break;
    }
  }

  for (auto client : clients) {
    if (!app_config.quiet) {
      client->print_stats();
    }
    delete client;
  }

  for (auto worker : worker_pool) {
    delete worker;
  }
  return 0;
}
