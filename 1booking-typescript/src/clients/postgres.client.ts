import { Pool } from "pg";
import { loadDatabaseConfig } from "@/config/database.config.js";

let pool: Pool | null = null;

export function getPostgresPool(): Pool {
  if (!pool) {
    const { url, pool: poolCfg } = loadDatabaseConfig();
    pool = new Pool({ connectionString: url, ...poolCfg });
    pool.on("connect",  () => console.log("[db] connection established"));
    pool.on("error",   (err) => { console.error("[db] pool error:", err); });
  }
  return pool;
}

export async function closePostgresPool(): Promise<void> {
  if (pool) { await pool.end(); pool = null; }
}