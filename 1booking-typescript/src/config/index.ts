import dotenv from "dotenv";
import { z } from "zod";
import path    from "path";
import { fileURLToPath } from "url";

dotenv.config({ path: path.join(process.cwd(), ".env") });
const __dirname = path.dirname(fileURLToPath(import.meta.url));

const envSchema = z.object({
  NODE_ENV:                  z.enum(["development", "test", "production"]).default("development"),
  PORT:                      z.coerce.number().default(3100),
  DATABASE_URL:              z.string().min(1, "DATABASE_URL is required"),
  REDIS_URL:                 z.string().min(1, "REDIS_URL is required"),
  JWT_SECRET:                z.string().min(32, "JWT_SECRET must be at least 32 characters"),
  JWT_REFRESH_SECRET:        z.string().min(32, "JWT_REFRESH_SECRET must be at least 32 characters"),
  JWT_ACCESS_EXPIRY:         z.string().default("15m"),
  JWT_REFRESH_EXPIRY:        z.string().default("7d"),
  BCRYPT_ROUNDS:             z.coerce.number().int().min(4).max(16).default(12),
  RATE_LIMIT_WINDOW_MS:      z.coerce.number().min(1000).default(60000),
  RATE_LIMIT_MAX_REQUESTS:   z.coerce.number().min(1).default(100),
  CORS_ORIGIN:               z.string().default("http://localhost:3001"),
  LOG_LEVEL:                 z.enum(["trace", "debug", "info", "warn", "error", "fatal"]).default("info"),
  STRIPE_SECRET_KEY:         z.string().optional().or(z.literal("")),
  STRIPE_WEBHOOK_SECRET:     z.string().optional().or(z.literal("")),
  S3_ENDPOINT:               z.string().optional().or(z.literal("")),
  S3_ACCESS_KEY:             z.string().optional().or(z.literal("")),
  S3_SECRET_KEY:             z.string().optional().or(z.literal("")),
  S3_BUCKET:                 z.string().optional().or(z.literal("")),
  AI_SERVICE_URL:            z.string().optional().or(z.literal("")),
  OPENAI_API_KEY:            z.string().optional().or(z.literal("")),
});
export type EnvConfig = z.infer<typeof envSchema>;

let _env: EnvConfig | null = null;
export function getEnv(): EnvConfig {
  if (!_env) {
    const parsed = envSchema.safeParse(process.env);
    if (!parsed.success) {
      const issues = parsed.error.issues.map((i) => `${i.path.join(".")}: ${i.message}`).join("\n");
      throw new Error(`Invalid environment configuration:\n${issues}`);
    }
    _env = parsed.data;
  }
  return _env;
}
export function get isProduction(): boolean { return getEnv().NODE_ENV === "production"; }