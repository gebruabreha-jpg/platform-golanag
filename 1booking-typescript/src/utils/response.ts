/* eslint-disable @typescript-eslint/no-explicit-any */

/**
 * Generic API response envelope.
 * @template T
 */
export interface ApiResponse<T = unknown> {
  success: boolean;
  data?: T;
  error?: {
    code: string;
    message: string;
    details?: string[];
  };
  meta?: {
    timestamp: string;
    path: string;
    [key: string]: unknown;
  };
}

/**
 * Pagination metadata object returned alongside paginated data.
 */
export interface PaginationMeta {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

/**
 * Wrap a success payload into a uniform ApiResponse.
 * @template T
 */
export function successResponse<T>(
  data: T,
  meta?: { [key: string]: unknown },
): ApiResponse<T> {
  return {
    success: true,
    data,
    meta: {
      timestamp: new Date().toISOString(),
      ...meta,
    },
  };
}

/**
 * Wrap an error payload into a uniform ApiResponse.
 * @template T
 */
export function errorResponse<T = never>(
  code: string,
  message: string,
  details?: string[],
): ApiResponse<T> {
  return {
    success: false,
    error: {
      code,
      message,
      details,
    },
    meta: {
      timestamp: new Date().toISOString(),
    },
  };
}

/**
 * Wrap a list payload with pagination metadata.
 */
export function paginatedResponse<T>(
  data: T[],
  meta: PaginationMeta,
): ApiResponse<T[]> {
  return {
    success: true,
    data,
    meta: {
      ...meta,
      timestamp: new Date().toISOString(),
    },
  };
}
