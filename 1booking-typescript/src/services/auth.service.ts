import type { IAuthService, RegisterInput, LoginResult, RefreshResult } from "@/services/interfaces/index.js";
import type { User } from "@/types/index.js";
import { AppError } from "@/errors/AppError.js";
import { ErrorCodes } from "@/errors/errorCodes.js";
import { getPostgresPool } from "@/clients/postgres.client.js";
import { getRedisClient } from "@/clients/redis.client.js";
import { hashPassword, comparePassword, signAccessToken, signRefreshToken, verifyRefreshToken as verifyRT, hashToken, compareTokenHash } from "@/utils/crypto.js";
import { randomUUID } from "crypto";

const TOKEN_BLACKLIST_PREFIX = "bl:";

export class AuthService implements IAuthService {
  async register(data: RegisterInput): Promise<User> {
    const pool = getPostgresPool();
    const existing = await pool.query(`SELECT id FROM users WHERE email = $1`, [data.email]);
    if (existing.rows.length > 0) {
      throw new AppError(ErrorCodes.EMAIL_EXISTS, "Email already registered");
    }
    const passwordHash = await hashPassword(data.password);
    const result = await pool.query(
      `INSERT INTO users (id, email, password_hash, first_name, last_name, phone, role, country, city)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
       RETURNING *`,
      [randomUUID(), data.email, passwordHash, data.firstName, data.lastName, data.phone, data.role ?? "DIASPORA", data.country, data.city],
    );
    return result.rows[0] as User;
  }

  async login(email: string, password: string): Promise<LoginResult> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`, [email]);
    if (result.rows.length === 0) {
      throw new AppError(ErrorCodes.INVALID_CREDENTIALS, "Invalid email or password");
    }
    const user = result.rows[0] as User;
    const valid = await comparePassword(password, user.password_hash);
    if (!valid) {
      throw new AppError(ErrorCodes.INVALID_CREDENTIALS, "Invalid email or password");
    }
    const accessToken = signAccessToken({ sub: user.id, role: user.role, type: "access" });
    const refreshToken = signRefreshToken({ sub: user.id, role: user.role, type: "refresh" });
    const tokenHash = await hashToken(refreshToken);
    await pool.query(
      `INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, NOW() + INTERVAL '7 days')`,
      [randomUUID(), user.id, tokenHash],
    );
    return { accessToken, refreshToken, expiresIn: process.env.JWT_ACCESS_EXPIRY ?? "15m" };
  }

  async logout(userId: string, refreshToken: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE refresh_tokens SET revoked_at = NOW() WHERE user_id = $1`, [userId]);
    const redis = getRedisClient();
    await redis.setex(TOKEN_BLACKLIST_PREFIX + refreshToken, 604800, "revoked");
  }

  async validateToken(token: string): Promise<User | null> {
    // This is handled by verifyAccessToken in crypto.ts
    return null;
  }

  async changePassword(userId: string, currentPassword: string, newPassword: string): Promise<void> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT password_hash FROM users WHERE id = $1`, [userId]);
    if (result.rows.length === 0) throw new AppError(ErrorCodes.USER_NOT_FOUND, "User not found");
    const valid = await comparePassword(currentPassword, result.rows[0].password_hash);
    if (!valid) throw new AppError(ErrorCodes.INVALID_CREDENTIALS, "Current password is incorrect");
    const newHash = await hashPassword(newPassword);
    await pool.query(`UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`, [newHash, userId]);
  }

  async refresh(refreshToken: string): Promise<RefreshResult> {
    const payload = verifyRT(refreshToken);
    if (!payload) throw new AppError(ErrorCodes.TOKEN_INVALID, "Invalid refresh token");
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM refresh_tokens WHERE user_id = $1 AND revoked_at IS NULL`, [payload.sub]);
    if (result.rows.length === 0) throw new AppError(ErrorCodes.REFRESH_TOKEN_REVOKED, "Refresh token not found");
    const stored = result.rows[0];
    const valid = await compareTokenHash(refreshToken, stored.token_hash);
    if (!valid) throw new AppError(ErrorCodes.REFRESH_TOKEN_INVALID, "Invalid refresh token");
    const accessToken = signAccessToken({ sub: payload.sub, role: payload.role, type: "access" });
    return { accessToken, expiresIn: process.env.JWT_ACCESS_EXPIRY ?? "15m" };
  }

  async revokeRefreshToken(tokenHash: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE refresh_tokens SET revoked_at = NOW() WHERE token_hash = $1`, [tokenHash]);
  }
}