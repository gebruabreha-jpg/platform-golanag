/* eslint-disable @typescript-eslint/no-empty-function */
import pino, { LoggerOptions, LogLevel } from "pino";

const level =
  (process.env.LOG_LEVEL as LogLevel) ?? (process.env.NODE_ENV === "production" ? "info" : "debug");

const options: LoggerOptions = {
  level,
  formatters: {
    level(label) {
      return { level: label };
    },
  },
  timestamp: pino.stdTimeFunctions.isoTime,
  redact: ["request.headers.authorization", "password", "token"],
};

const logger = pino(options);

export default logger;
