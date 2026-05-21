import { NextFunction, Request, Response } from "express";
import logger from "@/utils/logger.js";
import { isAppError } from "@/errors/AppError.js";
import { errorResponse } from "@/utils/response.js";

const DEFAULT_PRODUCTION_MESSAGE = "An unexpected error occurred. Please try again later.";

/**
 * Global Express error-handling middleware.
 *
 * Called after every route handler / middleware throws or passes an error
 * to `next(err)`. Converts known `AppError` instances to a uniform
 * JSON error envelope and hides unexpected server errors behind 500.
 */
export function errorHandler(
  err: unknown,
  req: Request,
  res: Response,
  _next: NextFunction,
): Response | void {
  const path    = req.originalUrl ?? req.url;
  const method  = req.method;
  const ip      = req.ip ?? req.socket?.remoteAddress ?? "unknown";

  // Log every error with full context
  if (isAppError(err)) {
    logger.warn(
      {
        code:          err.code ?? "UNKNOWN",
        statusCode:    err.statusCode,
        isOperational: err.isOperational,
        details:       err.details,
        method,
        path,
        ip,
      },
      err.message,
    );

    const payload = {
      code: err.code ?? "INTERNAL_ERROR",
      message: err.message,
      ...(err.details?.length ? { details: err.details } : {}),
    };

    return res.status(err.statusCode).json({
      ...errorResponse(err.code ?? "INTERNAL_ERROR", err.message, err.details),
      meta: { timestamp: new Date().toISOString(), method, path, ip },
    });
  }

  // Unexpected / programming error
  logger.error(
    { method, path, ip, name: err instanceof Error ? err.name : undefined },
    err instanceof Error ? err.message : String(err),
  );

  const response = errorResponse(
    "INTERNAL_ERROR",
    process.env.NODE_ENV === "production" ? DEFAULT_PRODUCTION_MESSAGE : (err as Error).message,
  );

  return res.status(500).json({
    ...response,
    meta: { timestamp: new Date().toISOString(), method, path, ip },
  });
}
