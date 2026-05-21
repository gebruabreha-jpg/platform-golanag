/**
 * Application error code enum.
 * Keep codes uppercase and machine-readable so they can be mapped
 * directly to HTTP status codes without string comparisons.
 */
export enum ErrorCodes {
  /* ── HTTP client errors ────────────────────────────────────────────────────── */
  UNAUTHORIZED            = "UNAUTHORIZED",
  FORBIDDEN               = "FORBIDDEN",
  NOT_FOUND               = "NOT_FOUND",
  BAD_REQUEST             = "BAD_REQUEST",
  CONFLICT                = "CONFLICT",
  VALIDATION_ERROR        = "VALIDATION_ERROR",
  RATE_LIMITED            = "RATE_LIMITED",

  /* ── HTTP server errors ────────────────────────────────────────────────────── */
  INTERNAL_ERROR          = "INTERNAL_ERROR",
  DATABASE_ERROR          = "DATABASE_ERROR",
  REDIS_ERROR             = "REDIS_ERROR",
  EXTERNAL_SERVICE_ERROR  = "EXTERNAL_SERVICE_ERROR",

  /* ── Token-related ────────────────────────────────────────────────────────── */
  TOKEN_EXPIRED           = "TOKEN_EXPIRED",
  TOKEN_INVALID           = "TOKEN_INVALID",
  TOKEN_REVOKED           = "TOKEN_REVOKED",
  REFRESH_TOKEN_EXPIRED   = "REFRESH_TOKEN_EXPIRED",
  REFRESH_TOKEN_REVOKED   = "REFRESH_TOKEN_REVOKED",
  REFRESH_TOKEN_NOT_FOUND = "REFRESH_TOKEN_NOT_FOUND",
  TOKEN_MISSING           = "TOKEN_MISSING",

  /* ── Auth / user ───────────────────────────────────────────────────────────── */
  INVALID_CREDENTIALS     = "INVALID_CREDENTIALS",
  EMAIL_EXISTS            = "EMAIL_EXISTS",
  PASSWORD_TOO_WEAK       = "PASSWORD_TOO_WEAK",
  USER_NOT_FOUND          = "USER_NOT_FOUND",
  ACCOUNT_LOCKED          = "ACCOUNT_LOCKED",
  ACCOUNT_DISABLED        = "ACCOUNT_DISABLED",
  EMAIL_NOT_VERIFIED      = "EMAIL_NOT_VERIFIED",
  PHONE_NOT_VERIFIED      = "PHONE_NOT_VERIFIED",

  /* ── Business domain ───────────────────────────────────────────────────────── */
  LISTING_NOT_FOUND       = "LISTING_NOT_FOUND",
  COMMUNITY_NOT_FOUND     = "COMMUNITY_NOT_FOUND",
  POST_NOT_FOUND          = "POST_NOT_FOUND",
  BOOKING_CONFLICT        = "BOOKING_CONFLICT",
  BOOKING_NOT_FOUND       = "BOOKING_NOT_FOUND",
  PAYMENT_FAILED          = "PAYMENT_FAILED",
  INSUFFICIENT_FUNDS      = "INSUFFICIENT_FUNDS",
  ESCROW_NOT_RELEASED     = "ESCROW_NOT_RELEASED",
  SERVICE_NOT_FOUND       = "SERVICE_NOT_FOUND",
  REVIEW_NOT_FOUND        = "REVIEW_NOT_FOUND",
  UNAUTHORIZED_ACTION     = "UNAUTHORIZED_ACTION",
  INSUFFICIENT_PRIVILEGES = "INSUFFICIENT_PRIVILEGES",
}

/**
 * Map ErrorCode → HTTP status.
 */
export function errorCodeToStatus(code: ErrorCodes): number {
  switch (code) {
    case ErrorCodes.UNAUTHORIZED:
    case ErrorCodes.TOKEN_MISSING:
    case ErrorCodes.INVALID_CREDENTIALS:
      return 401;

    case ErrorCodes.TOKEN_EXPIRED:
    case ErrorCodes.REFRESH_TOKEN_EXPIRED:
      return 401;

    case ErrorCodes.TOKEN_INVALID:
    case ErrorCodes.TOKEN_REVOKED:
    case ErrorCodes.REFRESH_TOKEN_REVOKED:
    case ErrorCodes.REFRESH_TOKEN_NOT_FOUND:
      return 401;

    case ErrorCodes.FORBIDDEN:
    case ErrorCodes.INSUFFICIENT_PRIVILEGES:
    case ErrorCodes.ACCOUNT_LOCKED:
    case ErrorCodes.ACCOUNT_DISABLED:
      return 403;

    case ErrorCodes.NOT_FOUND:
    case ErrorCodes.USER_NOT_FOUND:
    case ErrorCodes.LISTING_NOT_FOUND:
    case ErrorCodes.COMMUNITY_NOT_FOUND:
    case ErrorCodes.POST_NOT_FOUND:
    case ErrorCodes.SERVICE_NOT_FOUND:
    case ErrorCodes.REVIEW_NOT_FOUND:
    case ErrorCodes.BOOKING_NOT_FOUND:
      return 404;

    case ErrorCodes.CONFLICT:
    case ErrorCodes.BOOKING_CONFLICT:
    case ErrorCodes.EMAIL_EXISTS:
      return 409;

    case ErrorCodes.BAD_REQUEST:
    case ErrorCodes.VALIDATION_ERROR:
    case ErrorCodes.INVALID_CREDENTIALS:
    case ErrorCodes.PASSWORD_TOO_WEAK:
    case ErrorCodes.EMAIL_NOT_VERIFIED:
    case ErrorCodes.PHONE_NOT_VERIFIED:
    case ErrorCodes.INSUFFICIENT_FUNDS:
    case ErrorCodes.ESCROW_NOT_RELEASED:
    case ErrorCodes.UNAUTHORIZED_ACTION:
      return 400;

    case ErrorCodes.RATE_LIMITED:
      return 429;

    case ErrorCodes.PAYMENT_FAILED:
      return 402;

    case ErrorCodes.INTERNAL_ERROR:
    case ErrorCodes.DATABASE_ERROR:
    case ErrorCodes.REDIS_ERROR:
    case ErrorCodes.EXTERNAL_SERVICE_ERROR:
    default:
      return 500;
  }
}

/** HTTP 402 – Payment Required (reserved for payment endpoint). */
export const PaymentRequired = 402;
