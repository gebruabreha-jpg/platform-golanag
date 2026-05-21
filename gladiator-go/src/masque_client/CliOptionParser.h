#ifndef CLIOPTIONPARSER_H
#define CLIOPTIONPARSER_H
#include <cstdlib>
#include <string>
#pragma once

class Config;

class CliOptionParser {
public:
  CliOptionParser(Config &config);
  ~CliOptionParser() = default;

  bool parse_cli(int argc, const char **argv);

private:
  void print_usage(const std::string &app_name) const;
  void print_help_information(const std::string &app_name) const;
  int parse_requests(const char **argv, size_t argvlen);
  Config &config_;
};

#endif