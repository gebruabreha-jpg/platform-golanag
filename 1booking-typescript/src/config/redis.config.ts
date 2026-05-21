import Redis from "ioredis";

const url = process.env.REDIS_URL ?? "redis://localhost:6379";

export function createRedisClient(): Redis {
  return new Redis(url, { lazyConnect: true, maxRetriesPerRequest: null, retryStrategy(times) { return times > 3 ? null : Math.min(times * 200, 2000); } });
}