import type { PaginationMeta } from "./response.js";

/**
 * Slice an array into a page window.
 */
export function paginate<T>(
  items: T[],
  page = 1,
  limit = 20,
): { data: T[]; meta: PaginationMeta } {
  const safePage  = Math.max(1, page);
  const safeLimit = Math.max(1, limit);
  const total     = items.length;
  const totalPages = Math.ceil(total / safeLimit);
  const start = (safePage - 1) * safeLimit;
  const data  = items.slice(start, start + safeLimit);

  return {
    data,
    meta: {
      page:   safePage,
      limit:  safeLimit,
      total,
      totalPages,
    },
  };
}
