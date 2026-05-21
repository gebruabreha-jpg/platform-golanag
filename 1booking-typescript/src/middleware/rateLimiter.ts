import { NextFunction, Request, Response } from "express";
import type IORedis from "ioredis";

/**
 * Build a rate-limiter middleware factory keyed by `ip + method + path`.
 */
export function rateLimiter(redis: IORedis) {
  const WINDOW_MS   = parseInt(process.env.RATE_LIMIT_WINDOW_MS   ?? "60000", 10);
  const MAX_REQUESTS = parseInt(process.env.RATE_LIMIT_MAX_REQUESTS ?? "100",  10);

  return async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const ip   = req.ip ?? "unknown";
    const key  = `rl:${req.method}:${req.originalUrl}:${ip}`;

    try {
      const current = await redis.incr(key);
      if (current === 1) await redis.pexpire(key, WINDOW_MS);
      const remaining = Math.max(MAX_REQUESTS - current, 0);
      res.setHeader("X-RateLimit-Limit",     String(MAX_REQUESTS        ));
      res.setHeader("X-RateLimit-Remaining", String(remaining           ));
      res.setHeader("X-RateLimit-Reset",     String(Math.ceil(Date.now() / 1000 + WINDOW_MS / 1000)));

      if (current > MAX_REQUESTS) {
        const retryAfter = (await redis.pttl(key)) / 1000;
        res.setHeader("Retry-After", String(Math.max(1, Math.ceil(retryAfter))));
        return next(new Error(`RATE_LIMITED:Too many requests. Retry in ${Math.ceil(retryAfter)}s.`));
      }
      return next();
    } catch {
      // If Redis is down, pass through – do not block legitimate traffic.
      return next();
    }
  };
}
