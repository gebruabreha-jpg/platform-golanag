#ifndef REQUEST_H
#define REQUEST_H

#pragma once

#include <ostream>
#include <string>
#include <fstream>

class Request {
public:
  Request() = default;
  Request(const Request& other);
  ~Request() ;

  bool parse_uri(const char *uri);

  friend std::ostream &operator<<(std::ostream &out, const Request &req);

  std::string_view method;
  std::string_view scheme;
  std::string authority;
  std::string path;
  std::string protocol;
  std::string host_name;
  std::string port = "443";
  std::ofstream output;
  
};

#endif