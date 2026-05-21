import type { PoolConfig } from "pg";

export interface DatabaseConfig {
  url: string;
  pool: PoolConfig;
}

export function loadDatabaseConfig(): DatabaseConfig {
  const url = process.env.DATABASE_URL ?? "postgresql://postgres:postgres@localhost:5432/onepagecommerce_db";
  return { url, pool: { max: 20, min: 5, idleTimeoutMillis: 30000, connectionTimeoutMillis: 5000 } };
}