#include "CliOptionParser.h"
#include <cstdio>
#include <filesystem>
#include <getopt.h>

#include "Config.h"
#include "util.h"
#include <sys/stat.h>

namespace {
bool is_dir_existed(const std::string &path) {
  struct stat info;

  // Check if the path exists
  if (stat(path.c_str(), &info) != 0)
    return false;

  // Check if it's a directory
  if (info.st_mode & S_IFDIR)
    return true;
  else
    return false;
}
} // namespace

CliOptionParser::CliOptionParser(Config &config) : config_(config) {}

void CliOptionParser::print_usage(const std::string &app_name) const {
  fprintf(stderr, "Usage: %s [OPTIONS] <URI> [<URI>,...]\n", app_name.c_str());
}
void CliOptionParser::print_help_information(
    const std::string &app_name) const {

  static constexpr char help_info[] = R"(
  <URI>  [<URI>,...]     One or more URIs with the same domain name.
Options:
  -o, --output=<DIR> save the download files to the directory.
  -q, --quiet         disable debug output
  --connect-udp=<proxy-server:port>
              If specified, the client enter MASQUE client mode.
              The client sends an HTTP3 CONNECT request to the proxy server to
              establish a tunnel to the target server.
  --pvd-server=<pvd-server:port>
              If specified, the client connects the PVD server to get
              proxy-server's information.
              Note: --pvd and --connect-udp are mutually exclusive.

  --pvd-only
              Indicate it only connects the PVD server. After get the proxy-server's
              information, it close the connection and do not connect masque proxy-server.

  --clients-num=<N>
              The number of clients. Every client setup a `connection` to
              the http server. The default value is 1.

  --ipv4-pool=<network/mask>
              If specified, the client use the local IPv4 allocated from
              this pool.

  --ipv6-pool=<network/mask>
              If specified, the client use the local IPv6 allocated from
              this pool.

  --timeout=<N>
              The QUIC connection idle timeout(seconds).

  --connect-speed=<N>
              In a multiple clients scenario, it defines the rate at which new
              connections are established per second. The default value is 10.
              It will be ignored in single client scenario.

  --mtu=<N>
              Maximum Transmission Unit. If not set, it will be get automatically.

  --quic-forwarding
              Indicate the MASQUE client requests QUIC-Aware forwarding service.
              It's ignored if option --connect-udp is not provided.

  --request-dup-factor
              Indicate the duplication times of the requests sent in a single connection.
              Since the QUIC steam is created per request, it will increase the number of streams.
              Useful to test the performance of the server. The default value is 1.

  --duration=<N>
              If this option is specified, traffic will continue to operate until the elapsed time
              exceeds the specified duration. Support postfix unit 'H','M','S', 'h','m','s' or no postfix. 
              e.g. 1.5h, 72H, 30m, 45s, 120 etc. 

  -h, --help  Display this help and exit.
)";
  fprintf(stderr, help_info);
}

bool CliOptionParser::parse_cli(int argc, const char **argv) {
  // define options
  static int flag = 0;
  constexpr static option long_opts[] = {
      {"help", no_argument, nullptr, 'h'},
      {"quiet", no_argument, nullptr, 'q'},
      {"output", required_argument, nullptr, 'o'},
      {"connect-udp", required_argument, &flag, 1},
      {"clients-num", required_argument, &flag, 2},
      {"connect-speed", required_argument, &flag, 3},
      {"ipv4-pool", required_argument, &flag, 4},
      {"ipv6-pool", required_argument, &flag, 5},
      {"timeout", required_argument, &flag, 6},
      {"mtu", required_argument, &flag, 7},
      {"quic-forwarding", no_argument, &flag, 8},
      {"pvd-server", required_argument, &flag, 9},
      {"pvd-only", no_argument, &flag, 10},
      {"request-dup-factor", required_argument, &flag, 11},
      {"duration", required_argument, &flag, 12},
  };

  auto app_name = std::filesystem::path(argv[0]).filename();

  // read options
  for (;;) {
    int optidx = 0;

    auto c = getopt_long(argc, const_cast<char **>(argv), "hqo:", long_opts,
                         &optidx);
    if (c == -1) {
      break;
    }
    switch (c) {
    case 'h':
      // --help
      print_usage(app_name);
      print_help_information(app_name);
      exit(EXIT_SUCCESS);

    case '?':
      print_usage(app_name);
      exit(EXIT_FAILURE);

    case 'o': {
      config_.output_dir = optarg;
    } break;

    case 'q': {
      config_.quiet = true;
    } break;

    case 0:
      switch (flag) {
      case 1: {
        auto proxy_server_and_port = ngtcp2::util::split_str(optarg, ':');
        config_.proxy_server = proxy_server_and_port[0];
        config_.proxy_server_port = proxy_server_and_port[1];
      } break;

      case 2: {
        if (auto n = ngtcp2::util::parse_uint(optarg); !n) {
          fprintf(stderr, "clients-num: invalid argument\n");
          return false;
        } else if (*n > INT16_MAX) {
          fprintf(stderr, "clients-num: must not exceed (1 << 15) - 1\n");

          return false;
        } else {
          config_.clients_num = static_cast<uint16_t>(*n);
          config_.workers_num = config_.clients_num;
        }

      } break;

      case 3: {
        if (auto n = ngtcp2::util::parse_uint(optarg); !n) {
          fprintf(stderr, "connects-speed: invalid argument\n");
          return false;
        } else if (*n > 1000 or *n == 0) {
          fprintf(stderr, "connects-speed: must be in range [1,1000]\n");

          return false;
        } else {
          config_.new_conn_per_second = static_cast<uint16_t>(*n);
        }
      } break;

      case 4: {
        config_.ipv4_pool = optarg;
      } break;

      case 5: {
        config_.ipv6_pool = optarg;
      } break;

      case 6: {
        if (auto n = ngtcp2::util::parse_uint(optarg); !n) {
          fprintf(stderr, "timeout: invalid argument\n");

          return false;
        } else if (*n > 60) {
          fprintf(stderr, "timeout: must not exceed 60\n");

          return false;
        } else {
          config_.timeout = static_cast<uint8_t>(*n) * 1E9;
        }

      } break;

      case 7: {
        if (auto n = ngtcp2::util::parse_uint(optarg); !n) {
          fprintf(stderr, "mtu: invalid argument\n");

          return false;
        } else if (*n < 1280) {
          fprintf(stderr, "mtu: must not less than 1280\n");

          return false;
        } else {
          config_.max_trans_unit = static_cast<uint16_t>(*n);
        }
      } break;

      case 8: {

        config_.request_quic_forwarding = true;

      } break;

      case 9: {
        auto pvd_server_and_port = ngtcp2::util::split_str(optarg, ':');
        config_.pvd_server = pvd_server_and_port[0];
        config_.pvd_server_port = pvd_server_and_port[1];
        printf("pvd_server: %s, pvd_server_port: %s\n",
               config_.pvd_server.c_str(), config_.pvd_server_port.c_str());
      } break;

      case 10: {
        config_.pvd_only = true;
      } break;

      case 11: {
        if (auto n = ngtcp2::util::parse_uint(optarg); !n) {
          fprintf(stderr, "request-dup-factor: invalid argument\n");

          return false;
        } else if (*n > 50 or *n == 0) {
          fprintf(stderr, "request-dup-factor: must be in range [1,50]\n");

          return false;
        } else {
          config_.request_duplication_factor = static_cast<int>(*n);
        }
      } break;

      case 12: {
        if (auto n = ngtcp2::util::parse_duration2(optarg); n < 1) {
          fprintf(stderr, "duration: invalid argument\n");
          return false;
        } else {
          config_.duration = n;
        }
      } break;

        // end of numeric option
      }
      break;

    default:
      break;
    }
  }

  // URL is mandatory argument if it's not pvd_only mode
  if (!config_.pvd_only && argc - optind < 1) {
    print_usage(app_name);
    exit(EXIT_FAILURE);
  }

  // if pvd-only is true, the pvd-server must be specified
  if (config_.pvd_only && config_.pvd_server.empty()) {
    fprintf(stderr, "option --pvd-only must be used with --pvd-server\n");
    print_usage(app_name);
    exit(EXIT_FAILURE);
  }

  // output directory must exist if it's not empty
  if (!config_.output_dir.empty() && !is_dir_existed(config_.output_dir)) {
    fprintf(stderr, "Neither '%s' exist nor it's a directory!\n",
            config_.output_dir.c_str());
    exit(EXIT_FAILURE);
  }

  // mutual exclusion check
  if (!config_.ipv4_pool.empty() && !config_.ipv6_pool.empty()) {
    fprintf(stderr,
            "option --ipv4-pool and --ipv6-pool are mutual exclusive!\n");
    print_usage(app_name);
    exit(EXIT_FAILURE);
  }

  if (!config_.pvd_server.empty() && !config_.proxy_server.empty()) {
    fprintf(stderr,
            "option --pvd-server and --connect-udp are mutual exclusive!\n");
    print_usage(app_name);
    exit(EXIT_FAILURE);
  }

  if (parse_requests(&argv[optind], argc - optind) != 0) {

    return false;
  }
  return true;
}

int CliOptionParser::parse_requests(const char **argv, size_t argvlen) {
  for (size_t i = 0; i < argvlen; ++i) {
    auto uri = argv[i];
    HttpRequest req;
    req.method = config_.http_method;
    if (!req.parse_uri(uri)) {
      fprintf(stderr, "Could not parse URI: %s\n", uri);
      return -1;
    }
    for (int i = 0; i < config_.request_duplication_factor; i++) {
      config_.target_servers_to_requests[std::tuple(req.host_name, req.port)]
          .push_back(req);
    }
  }

  return 0;
}
