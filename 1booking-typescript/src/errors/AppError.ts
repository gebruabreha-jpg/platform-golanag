import { Request, Response, NextFunction } from "express";
import { ZodIssue } from "zod";
import {
  ErrorCodes,
  errorCodeToStatus,
} from "@/errors/errorCodes.js";

/**
 * Custom application error.
 *
 * Operational errors (expected errors, `isOperational = true`) are
 * trusted and safe to expose to the client. Programming errors
 * (`isOperational = false`) must be hidden behind 500 to avoid
 * leaking stack traces.
 */
export class AppError extends Error {
  public readonly statusCode: number;
  public readonly isOperational: boolean;
  public readonly code: ErrorCodes;
  public readonly details?: string[];

  constructor(
    code: ErrorCodes = ErrorCodes.INTERNAL_ERROR,
    message: string,
    statusOrStatusCode: number = 500,
    isOperational: boolean = true,
    details?: string[],
  ) {
    // Allow callers to pass a ZodIssue[] and convert automatically
    const zodIssues: ZodIssue[] = (message as unknown as { _zod?: { issues?: ZodIssue[] } })?._zod
      ?.issues ?? [];
    super(zodIssues.length
      ? zodIssues.map((i) => `${i.path.join(".")}: ${i.message}`).join(", ")
      : message
    );

    this.name = code;
    this.code  = code;
    // If the caller didn't explicitly set a code-derived status, use the helper
    this.statusCode   =
      statusOrStatusCode >= 400 && statusOrStatusCode < 600
        ? statusOrStatusCode
        : errorCodeToStatus(code);
    this.isOperational = isOperational;
    this.details       = details;

    // Stack trace
    Error.captureStackTrace(this, this.constructor);
  }
}

/** Convenience factory helpers. */
export const UnauthorizedError = (
  msg: string = "Authentication required",
  code: ErrorCodes = ErrorCodes.UNAUTHORIZED,
) => new AppError(code, msg, 401);

export const ForbiddenError = (
  msg: string = "Access denied",
  code: ErrorCodes = ErrorCodes.FORBIDDEN,
) => new AppError(code, msg, 403);

export const NotFoundError = (
  msg: string = "Resource not found",
  code: ErrorCodes = ErrorCodes.NOT_FOUND,
) => new AppError(code, msg, 404);

export const ValidationError = (
  message: string | string[],
  code: ErrorCodes = ErrorCodes.VALIDATION_ERROR,
  details?: string[],
) => {
  const msg = Array.isArray(message)
    ? message.join(", ")
    : message;
  return new AppError(code, msg, 400, true, details);
};

export const ConflictError = (
  msg: string = "Conflict",
  code: ErrorCodes = ErrorCodes.CONFLICT,
) => new AppError(code, msg, 409);

export const InternalError = (
  message: string,
  isOperational = false,
) => new AppError(ErrorCodes.INTERNAL_ERROR, message, 500, isOperational);

export const RateLimitedError = (
  retryAfter?: number,
) => new AppError(ErrorCodes.RATE_LIMITED, "Rate limit exceeded", 429, true, retryAfter ? [`Retry after ${retryAfter}s`] : undefined);

/** Type guard – true only for AppError instances. */
export function isAppError(err: unknown): err is AppError {
  return err instanceof AppError;
}
