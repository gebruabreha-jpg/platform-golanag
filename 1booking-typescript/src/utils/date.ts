import { v4 as uuidv4 } from "uuid";

/**
 * Format a Date (or ISO string) as `YYYY-MM-DD`.
 */
export function formatDate(value: Date | string): string {
  const d = value instanceof Date ? value : new Date(value);
  return d.toISOString().split("T")[0];
}

/**
 * Return a Date shifted forward by `days` days.
 */
export function addDays(value: Date | string, days: number): Date {
  const d = value instanceof Date ? value : new Date(value);
  d.setDate(d.getDate() + days);
  return d;
}

/**
 * Check whether a plain number or ISO-string timestamp is in the past.
 */
export function isExpired(value: number | string): boolean {
  return Date.now() > new Date(value).getTime();
}

/**
 * Human-readable relative string (e.g. "3 minutes ago").
 */
export function timeAgo(value: Date | string | number): string {
  const seconds = Math.floor((Date.now() - new Date(value).getTime()) / 1000);

  const intervals: [number, string][] = [
    [31536000, "year"],
    [2592000,   "month"],
    [604800,    "week"],
    [86400,     "day"],
    [3600,      "hour"],
    [60,        "minute"],
    [1,         "second"],
  ];

  for (const [secondsPerUnit, unit] of intervals) {
    const count = Math.floor(seconds / secondsPerUnit);
    if (count >= 1) return `${count} ${unit}${count !== 1 ? "s" : ""} ago`;
  }
  return "just now";
}

/**
 * Generate a compact sequential short-id (cryptographically random).
 */
export function generateCode(length = 8): string {
  const bytes = crypto.getRandomValues(new Uint8Array(length));
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("")
    .slice(0, length)
    .toUpperCase();
}

export { uuidv4 as generateUUID };
