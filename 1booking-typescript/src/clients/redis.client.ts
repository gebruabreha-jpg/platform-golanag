import { createRedisClient } from "@/config/redis.config.js";

let redis: ReturnType<typeof createRedisClient> | null = null;

export function getRedisClient(): ReturnType<typeof createRedisClient> {
  if (!redis) redis = createRedisClient();
  return redis;
}

export async function initRedis(): Promise<void> {
  const client = getRedisClient();
  // Eagerly connect so we fail at startup if Redis is unavailable.
  await client.connect();
}

export async function closeRedis(): Promise<void> {
  if (redis) { await redis.quit(); redis = null; }
}