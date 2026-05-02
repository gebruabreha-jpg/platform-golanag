/*
 * ngtcp2
 *
 * Copyright (c) 2017 ngtcp2 contributors
 * Copyright (c) 2012 nghttp2 contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 * LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
#include "util.h"

#include <arpa/inet.h>
#include <fcntl.h>
#include <net/if.h>
#include <netdb.h>
#include <netinet/in.h>
#include <stdexcept>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

#include <algorithm>
#include <array>
#include <cassert>
#include <charconv>
#include <chrono>
#include <cstring>
#include <fstream>
#include <iostream>
#include <limits>
#include <regex>

#include "template.h"

using namespace std::literals;

namespace ngtcp2 {

namespace util {

std::optional<std::string> read_pem(const std::string_view &filename,
                                    const std::string_view &name,
                                    const std::string_view &type);

int write_pem(const std::string_view &filename, const std::string_view &name,
              const std::string_view &type, const uint8_t *data,
              size_t datalen);

namespace {
int64_t convert_value_with_unit(
    const std::string &value_str, const std::string &regex_str,
    const std::unordered_map<std::string, uint64_t> &unit_multipliers) {
  std::smatch match;
  std::regex regex(regex_str);
  // Perform regex search
  if (std::regex_match(value_str, match, regex)) {
    std::string numeric_part = match[1].str();
    std::string unit = match[2].str();
    double value = std::stod(numeric_part);
    uint64_t multiplier = unit_multipliers.at(unit);

    // Calculate the final value in bits per second
    return static_cast<int64_t>(std::round(value * multiplier));
  }

  // Return -1 if the input string does not match the expected format
  return -1;
}
} // namespace

namespace {
constexpr char LOWER_XDIGITS[] = "0123456789abcdef";
} // namespace

std::string format_hex(uint8_t c) {
  std::string s;
  s.resize(2);

  s[0] = LOWER_XDIGITS[c >> 4];
  s[1] = LOWER_XDIGITS[c & 0xf];

  return s;
}

std::string format_hex(const uint8_t *s, size_t len) {
  std::string res;
  res.resize(len * 2);

  for (size_t i = 0; i < len; ++i) {
    auto c = s[i];

    res[i * 2] = LOWER_XDIGITS[c >> 4];
    res[i * 2 + 1] = LOWER_XDIGITS[c & 0x0f];
  }
  return res;
}

std::string format_hex(const std::string_view &s) {
  return format_hex(reinterpret_cast<const uint8_t *>(s.data()), s.size());
}

std::string decode_hex(const std::string_view &s) {
  assert(s.size() % 2 == 0);
  std::string res(s.size() / 2, '0');
  auto p = std::begin(res);
  for (auto it = std::begin(s); it != std::end(s); it += 2) {
    *p++ = (hex_to_uint(*it) << 4) | hex_to_uint(*(it + 1));
  }
  return res;
}

namespace {
// format_fraction2 formats |n| as fraction part of integer.  |n| is
// considered as fraction, and its precision is 3 digits.  The last
// digit is ignored.  The precision of the resulting fraction is 2
// digits.
std::string format_fraction2(uint32_t n) {
  n /= 10;

  if (n < 10) {
    return {'.', '0', static_cast<char>('0' + n)};
  }
  return {'.', static_cast<char>('0' + n / 10),
          static_cast<char>('0' + (n % 10))};
}
} // namespace

namespace {
// round2even rounds the last digit of |n| so that the n / 10 becomes
// even.
uint64_t round2even(uint64_t n) {
  if (n % 10 == 5) {
    if ((n / 10) & 1) {
      n += 10;
    }
  } else {
    n += 5;
  }
  return n;
}
} // namespace

std::string format_durationf(uint64_t ns) {
  static constexpr const std::string_view units[] = {"us"sv, "ms"sv, "s"sv};
  if (ns < 1000) {
    return format_uint(ns) + "ns";
  }
  auto unit = 0;
  if (ns < 1000000) {
    // do nothing
  } else if (ns < 1000000000) {
    ns /= 1000;
    unit = 1;
  } else {
    ns /= 1000000;
    unit = 2;
  }

  ns = round2even(ns);

  if (ns / 1000 >= 1000 && unit < 2) {
    ns /= 1000;
    ++unit;
  }

  auto res = format_uint(ns / 1000);
  res += format_fraction2(ns % 1000);
  res += units[unit];

  return res;
}

std::mt19937 make_mt19937() {
  std::random_device rd;
  return std::mt19937(rd());
}

ngtcp2_tstamp timestamp() {
  return std::chrono::duration_cast<std::chrono::nanoseconds>(
             std::chrono::steady_clock::now().time_since_epoch())
      .count();
}

bool numeric_host(const char *hostname) {
  return numeric_host(hostname, AF_INET) || numeric_host(hostname, AF_INET6);
}

bool numeric_host(const char *hostname, int family) {
  int rv;
  std::array<uint8_t, sizeof(struct in6_addr)> dst;

  rv = inet_pton(family, hostname, dst.data());

  return rv == 1;
}

namespace {
void hexdump8(FILE *out, const uint8_t *first, const uint8_t *last) {
  auto stop = std::min(first + 8, last);
  for (auto k = first; k != stop; ++k) {
    fprintf(out, "%02x ", *k);
  }
  // each byte needs 3 spaces (2 hex value and space)
  for (; stop != first + 8; ++stop) {
    fputs("   ", out);
  }
  // we have extra space after 8 bytes
  fputc(' ', out);
}
} // namespace

void hexdump(FILE *out, const uint8_t *src, size_t len) {
  if (len == 0) {
    return;
  }
  size_t buflen = 0;
  auto repeated = false;
  std::array<uint8_t, 16> buf{};
  auto end = src + len;
  auto i = src;
  for (;;) {
    auto nextlen =
        std::min(static_cast<size_t>(16), static_cast<size_t>(end - i));
    if (nextlen == buflen &&
        std::equal(std::begin(buf), std::begin(buf) + buflen, i)) {
      // as long as adjacent 16 bytes block are the same, we just
      // print single '*'.
      if (!repeated) {
        repeated = true;
        fputs("*\n", out);
      }
      i += nextlen;
      continue;
    }
    repeated = false;
    fprintf(out, "%08lx", static_cast<unsigned long>(i - src));
    if (i == end) {
      fputc('\n', out);
      break;
    }
    fputs("  ", out);
    hexdump8(out, i, end);
    hexdump8(out, i + 8, std::max(i + 8, end));
    fputc('|', out);
    auto stop = std::min(i + 16, end);
    buflen = stop - i;
    auto p = buf.data();
    for (; i != stop; ++i) {
      *p++ = *i;
      if (0x20 <= *i && *i <= 0x7e) {
        fputc(*i, out);
      } else {
        fputc('.', out);
      }
    }
    fputs("|\n", out);
  }
}

std::string make_cid_key(const ngtcp2_cid *cid) {
  return std::string(cid->data, cid->data + cid->datalen);
}

std::string make_cid_key(const uint8_t *cid, size_t cidlen) {
  return std::string(cid, cid + cidlen);
}

std::string straddr(const sockaddr *sa, socklen_t salen) {
  std::array<char, NI_MAXHOST> host;
  std::array<char, NI_MAXSERV> port;

  auto rv = getnameinfo(sa, salen, host.data(), host.size(), port.data(),
                        port.size(), NI_NUMERICHOST | NI_NUMERICSERV);
  if (rv != 0) {
    std::cerr << "getnameinfo: " << gai_strerror(rv) << std::endl;
    return "";
  }
  std::string res = "[";
  res.append(host.data(), strlen(host.data()));
  res += "]:";
  res.append(port.data(), strlen(port.data()));
  return res;
}

std::string_view strccalgo(ngtcp2_cc_algo cc_algo) {
  switch (cc_algo) {
  case NGTCP2_CC_ALGO_RENO:
    return "reno"sv;
  case NGTCP2_CC_ALGO_CUBIC:
    return "cubic"sv;
  case NGTCP2_CC_ALGO_BBR:
    return "bbr"sv;
  default:
    assert(0);
    abort();
  }
}

namespace {
constexpr bool rws(char c) { return c == '\t' || c == ' '; }
} // namespace

std::optional<std::unordered_map<std::string, std::string>>
read_mime_types(const std::string_view &filename) {
  std::ifstream f(filename.data());
  if (!f) {
    return {};
  }

  std::unordered_map<std::string, std::string> dest;

  std::string line;
  while (std::getline(f, line)) {
    if (line.empty() || line[0] == '#') {
      continue;
    }

    auto p = std::find_if(std::begin(line), std::end(line), rws);
    if (p == std::begin(line) || p == std::end(line)) {
      continue;
    }

    auto media_type = std::string{std::begin(line), p};
    for (;;) {
      auto ext = std::find_if_not(p, std::end(line), rws);
      if (ext == std::end(line)) {
        break;
      }

      p = std::find_if(ext, std::end(line), rws);
      dest.emplace(std::string{ext, p}, media_type);
    }
  }

  return dest;
}

std::string format_duration(ngtcp2_duration n) {
  if (n >= 3600 * NGTCP2_SECONDS && (n % (3600 * NGTCP2_SECONDS)) == 0) {
    return format_uint(n / (3600 * NGTCP2_SECONDS)) + 'h';
  }
  if (n >= 60 * NGTCP2_SECONDS && (n % (60 * NGTCP2_SECONDS)) == 0) {
    return format_uint(n / (60 * NGTCP2_SECONDS)) + 'm';
  }
  if (n >= NGTCP2_SECONDS && (n % NGTCP2_SECONDS) == 0) {
    return format_uint(n / NGTCP2_SECONDS) + 's';
  }
  if (n >= NGTCP2_MILLISECONDS && (n % NGTCP2_MILLISECONDS) == 0) {
    return format_uint(n / NGTCP2_MILLISECONDS) + "ms";
  }
  if (n >= NGTCP2_MICROSECONDS && (n % NGTCP2_MICROSECONDS) == 0) {
    return format_uint(n / NGTCP2_MICROSECONDS) + "us";
  }
  return format_uint(n) + "ns";
}

namespace {
std::optional<std::pair<uint64_t, size_t>>
parse_uint_internal(const std::string_view &s) {
  uint64_t res = 0;

  if (s.empty()) {
    return {};
  }

  for (size_t i = 0; i < s.size(); ++i) {
    auto c = s[i];
    if (c < '0' || '9' < c) {
      return {{res, i}};
    }

    auto d = c - '0';
    if (res > (std::numeric_limits<uint64_t>::max() - d) / 10) {
      return {};
    }

    res *= 10;
    res += d;
  }

  return {{res, s.size()}};
}
} // namespace

std::optional<uint64_t> parse_uint(const std::string_view &s) {
  auto o = parse_uint_internal(s);
  if (!o) {
    return {};
  }
  auto [res, idx] = *o;
  if (idx != s.size()) {
    return {};
  }
  return res;
}

std::optional<uint64_t> parse_uint_iec(const std::string_view &s) {
  auto o = parse_uint_internal(s);
  if (!o) {
    return {};
  }
  auto [res, idx] = *o;
  if (idx == s.size()) {
    return res;
  }
  if (idx + 1 != s.size()) {
    return {};
  }

  uint64_t m;
  switch (s[idx]) {
  case 'G':
  case 'g':
    m = 1 << 30;
    break;
  case 'M':
  case 'm':
    m = 1 << 20;
    break;
  case 'K':
  case 'k':
    m = 1 << 10;
    break;
  default:
    return {};
  }

  if (res > std::numeric_limits<uint64_t>::max() / m) {
    return {};
  }

  return res * m;
}

std::optional<uint64_t> parse_duration(const std::string_view &s) {
  auto o = parse_uint_internal(s);
  if (!o) {
    return {};
  }
  auto [res, idx] = *o;
  if (idx == s.size()) {
    return res * NGTCP2_SECONDS;
  }

  uint64_t m;
  if (idx + 1 == s.size()) {
    switch (s[idx]) {
    case 'H':
    case 'h':
      m = 3600 * NGTCP2_SECONDS;
      break;
    case 'M':
    case 'm':
      m = 60 * NGTCP2_SECONDS;
      break;
    case 'S':
    case 's':
      m = NGTCP2_SECONDS;
      break;
    default:
      return {};
    }
  } else if (idx + 2 == s.size() && (s[idx + 1] == 's' || s[idx + 1] == 'S')) {
    switch (s[idx]) {
    case 'M':
    case 'm':
      m = NGTCP2_MILLISECONDS;
      break;
    case 'U':
    case 'u':
      m = NGTCP2_MICROSECONDS;
      break;
    case 'N':
    case 'n':
      return res;
    default:
      return {};
    }
  } else {
    return {};
  }

  if (res > std::numeric_limits<uint64_t>::max() / m) {
    return {};
  }

  return res * m;
}

namespace {
template <typename InputIt> InputIt eat_file(InputIt first, InputIt last) {
  if (first == last) {
    *first++ = '/';
    return first;
  }

  if (*(last - 1) == '/') {
    return last;
  }

  auto p = last;
  for (; p != first && *(p - 1) != '/'; --p)
    ;
  if (p == first) {
    // this should not happened in normal case, where we expect path
    // starts with '/'
    *first++ = '/';
    return first;
  }

  return p;
}
} // namespace

namespace {
template <typename InputIt> InputIt eat_dir(InputIt first, InputIt last) {
  auto p = eat_file(first, last);

  --p;

  assert(*p == '/');

  return eat_file(first, p);
}
} // namespace

std::string normalize_path(const std::string_view &path) {
  assert(path.size() <= 1024);
  assert(path.size() > 0);
  assert(path[0] == '/');

  std::array<char, 1024> res;
  auto p = res.data();

  auto first = std::begin(path);
  auto last = std::end(path);

  *p++ = '/';
  ++first;
  for (; first != last && *first == '/'; ++first)
    ;

  for (; first != last;) {
    if (*first == '.') {
      if (first + 1 == last) {
        break;
      }
      if (*(first + 1) == '/') {
        first += 2;
        continue;
      }
      if (*(first + 1) == '.') {
        if (first + 2 == last) {
          p = eat_dir(res.data(), p);
          break;
        }
        if (*(first + 2) == '/') {
          p = eat_dir(res.data(), p);
          first += 3;
          continue;
        }
      }
    }
    if (*(p - 1) != '/') {
      p = eat_file(res.data(), p);
    }
    auto slash = std::find(first, last, '/');
    if (slash == last) {
      p = std::copy(first, last, p);
      break;
    }
    p = std::copy(first, slash + 1, p);
    first = slash + 1;
    for (; first != last && *first == '/'; ++first)
      ;
  }
  return std::string{res.data(), p};
}

std::vector<std::string_view> split_str(const std::string_view &s, char delim) {
  size_t len = 1;
  auto last = std::end(s);
  std::string_view::const_iterator d;
  for (auto first = std::begin(s); (d = std::find(first, last, delim)) != last;
       ++len, first = d + 1)
    ;

  auto list = std::vector<std::string_view>(len);

  len = 0;
  for (auto first = std::begin(s);; ++len) {
    auto stop = std::find(first, last, delim);
    // xcode clang does not understand std::string_view{first, stop}.
    list[len] = std::string_view{first, static_cast<size_t>(stop - first)};
    if (stop == last) {
      break;
    }
    first = stop + 1;
  }
  return list;
}

std::optional<uint32_t> parse_version(const std::string_view &s) {
  auto k = s;
  if (!util::istarts_with(k, "0x"sv)) {
    return {};
  }
  k = k.substr(2);
  uint32_t v;
  auto rv = std::from_chars(k.data(), k.data() + k.size(), v, 16);
  if (rv.ptr != k.data() + k.size() || rv.ec != std::errc{}) {
    return {};
  }

  return v;
}

std::optional<std::string> read_token(const std::string_view &filename) {
  return read_pem(filename, "token", "QUIC TOKEN");
}

int write_token(const std::string_view &filename, const uint8_t *token,
                size_t tokenlen) {
  return write_pem(filename, "token", "QUIC TOKEN", token, tokenlen);
}

std::optional<std::string>
read_transport_params(const std::string_view &filename) {
  return read_pem(filename, "transport parameters",
                  "QUIC TRANSPORT PARAMETERS");
}

int write_transport_params(const std::string_view &filename,
                           const uint8_t *data, size_t datalen) {
  return write_pem(filename, "transport parameters",
                   "QUIC TRANSPORT PARAMETERS", data, datalen);
}

int compare_ipv6(const in6_addr &a, const in6_addr &b) {
  for (int i = 0; i < 16; i++) {
    if (a.s6_addr[i] > b.s6_addr[i]) {
      return 1;
    } else if (a.s6_addr[i] < b.s6_addr[i]) {
      return -1;
    }
  }
  return 0;
}

std::tuple<char, double> convert_data_with_suitable_unit(double value,
                                                         bool speed) {

  if (speed) {
    if (value < 1E3) {
      return {' ', value};
    } else if (value < 1E6) {
      return {'K', value / 1E3};
    } else if (value < 1E9) {
      return {'M', value / 1E6};
    }
    return {'G', value / 1E9};
  } else {
    if (value < 1_k) {
      return {' ', value};
    } else if (value < 1_m) {
      return {'K', value / 1_k};
    } else if (value < 1_g) {
      return {'M', value / 1_m};
    }
    return {'G', value / 1_g};
  }
}

int encode_var_len_integer(uint64_t val, uint8_t *buff, int buf_size) {
  int length = 0;
  uint64_t MASK = 0xFF;
  if (val < (1L << 6)) {
    length = 1;
    assert(buf_size >> length);
    buff[0] = val;

  } else if (val < (1L << 14)) {

    length = 2;
    assert(buf_size >> length);
    buff[0] = ((val & (MASK << 8)) >> 8) | (1 << 6);
    buff[1] = val & MASK;

  } else if (val < (1L << 30)) {
    length = 4;
    assert(buf_size >> length);
    buff[0] = ((val & (MASK << 24)) >> 24) | (2 << 6);

    for (int i = 1; i < 4; i++) {
      auto shift_bits = 8 * (3 - i);
      buff[i] = (val & (MASK << shift_bits)) >> shift_bits;
    }

  } else if (val < (1L << 62)) {
    length = 8;
    assert(buf_size >> length);
    buff[0] = ((val & (MASK << 56)) >> 56) | (3 << 6);

    for (int i = 1; i < 8; i++) {
      auto shift_bits = 8 * (7 - i);
      buff[i] = (val & (MASK << shift_bits)) >> shift_bits;
    }
  }
  return length;
}

int decode_var_len_integer(uint64_t &val, const uint8_t *buff, int buf_size) {

  auto v = buff[0];
  auto prefix = (v & 0xC0) >> 6;
  auto length = 1 << prefix;
  // fprintf(stderr,"length=%d buf_size:%d\n",length,buf_size);
  if (length > buf_size or length > 8) {
    fprintf(stderr, "decode_var_len_integer() length=%d buf_size:%d\n", length,
            buf_size);
    // throw std::runtime_error("length > buf_size");
    return 0;
  }
  // assert(length <= buf_size);

  val = v & 0x3f;

  for (int i = 1; i < length; i++) {
    val = (val << 8) + buff[i];
  }

  return length;
}

int64_t parse_bitrate(const std::string &bitrate_str) {
  const std::string regex_str = R"(([\d.]+)\s*([GMKgmk])?)";
  const std::unordered_map<std::string, uint64_t> unit_multipliers = {
      {"G", 1000000000UL}, // Gigabit
      {"M", 1000000UL},    // Megabit
      {"K", 1000UL},       // Kilobit
      {"g", 1000000000UL}, // Gigabit
      {"m", 1000000UL},    // Megabit
      {"k", 1000UL},       // Kilobit
      {"", 1UL}            // No unit, assume bits
  };
  return convert_value_with_unit(bitrate_str, regex_str, unit_multipliers);
}

int64_t parse_bytes(const std::string& byte_str){
  const std::string regex_str = R"(([\d.]+)\s*([GMKgmk])?)";
  const std::unordered_map<std::string, uint64_t> unit_multipliers = {
      {"G", 1UL << 30}, // Giga byte
      {"M", 1UL << 20},    // Mega byte
      {"K", 1UL << 10},       // Kilo byte
      {"g", 1UL << 30}, // Giga byte
      {"m", 1UL << 20},    // Mega byte
      {"k", 1UL << 10},       // Kilo byte
      {"", 1UL}            // No unit, assume byte
  };
  return convert_value_with_unit(byte_str, regex_str, unit_multipliers);
}

int parse_duration2(const std::string &duration_str) {
  std::string regex_str = R"(^([\d.]+)\s*([HMShms])?$)";
  std::unordered_map<std::string, uint64_t> unit_multipliers = {
      {"H", 3600UL}, // hour
      {"M", 60UL},   // minute
      {"S", 1UL},    // second
      {"h", 3600UL}, // hour
      {"m", 60UL},   // minute
      {"s", 1UL},    // second
      {"", 1UL}      // No unit, assume second
  };
  return convert_value_with_unit(duration_str, regex_str, unit_multipliers);
}

} // namespace util

std::ostream &operator<<(std::ostream &os, const ngtcp2_cid &cid) {
  return os << "0x" << util::format_hex(cid.data, cid.datalen);
}
} // namespace ngtcp2
