#ifndef REQUEST_H
#define REQUEST_H

#pragma once

#include <cstdint>
#include <fstream>
#include <ostream>
#include <string>

static constexpr int MAX_HTTP3_BODY_DATA_LEN = 2048;

class HttpRequest {
public:
  HttpRequest();
  HttpRequest(const HttpRequest &other);
  ~HttpRequest();

  bool parse_uri(const char *uri);

  friend std::ostream &operator<<(std::ostream &out, const HttpRequest &req);

  std::string method;
  std::string scheme;
  std::string authority;
  std::string path;
  std::string accept;
  std::string protocol;
  std::string host_name;
  std::string port = "443";
  std::ofstream *output = nullptr;
  uint8_t data_body[MAX_HTTP3_BODY_DATA_LEN];
  int32_t data_len = 0;
  int32_t acked_len = 0;
  int64_t stream_id = -1;
  bool data_end = false;
};

#endif