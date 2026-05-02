#include "Request.h"
#include "http_parser.h"
#include "util.h"
#include <cstring>

namespace {
std::string_view get_string(const char *uri, const http_parser_url &u,
                            http_parser_url_fields f) {
  auto p = &u.field_data[f];
  return {uri + p->off, p->len};
}
}


Request::Request(const Request& other){
  method = other.method;
  scheme = other.scheme;
  authority = other.authority;
  path = other.path;
  protocol = other.protocol;
  host_name = other.host_name;
  port = other.port;
  
}

bool  Request::parse_uri(const char *uri) {
  http_parser_url u;

  http_parser_url_init(&u);
  if (http_parser_parse_url(uri, strlen(uri), /* is_connect = */ 0, &u) != 0) {
    return false;
  }

  if (!(u.field_set & (1 << UF_SCHEMA)) || !(u.field_set & (1 << UF_HOST))) {
    return false;
  }

  scheme = get_string(uri, u, UF_SCHEMA);

  auto host = std::string(get_string(uri, u, UF_HOST));
  host_name = host;
  if (ngtcp2::util::numeric_host(host.c_str(), AF_INET6)) {
    authority = '[';
    authority += host;
    authority += ']';
  } else {
    authority = std::move(host);
  }

  if (u.field_set & (1 << UF_PORT)) {
    port = get_string(uri, u, UF_PORT);
    authority += ':';
    authority += port;
  }

  if (u.field_set & (1 << UF_PATH)) {
    path = get_string(uri, u, UF_PATH);
  } else {
    path = "/";
  }

  if (u.field_set & (1 << UF_QUERY)) {
    path += '?';
    path += get_string(uri, u, UF_QUERY);
  }

  return true;
}

std::ostream& operator << (std::ostream& out, const Request& req) {
    out << "method:    " << req.method << "\n";
    out << "scheme:    " << req.scheme << "\n";
    out << "authority: " << req.authority << "\n";
    out << "port:      " << req.port << "\n";
    out << "protocol:  " << req.protocol << "\n";
    out << "path:      " << req.path << "\n";
    
    return out;
}

Request::~Request(){
  if (output.is_open()){
    output.close();
  }
}