import bcrypt from "bcrypt";
import jwt, { SignOptions, JwtPayload } from "jsonwebtoken";
import { randomUUID } from "crypto";
import type { JWTPayload } from "@/types/index.js";
import type logger from "@/utils/logger.js";

const BCRYPT_ROUNDS = parseInt(process.env.BCRYPT_ROUNDS ?? "12", 10);
const JWT_SECRET = process.env.JWT_SECRET;
const JWT_REFRESH_SECRET = process.env.JWT_REFRESH_SECRET;
const JWT_ACCESS_EXPIRY = process.env.JWT_ACCESS_EXPIRY ?? "15m";
const JWT_REFRESH_EXPIRY = process.env.JWT_REFRESH_EXPIRY ?? "7d";

const signOptions: SignOptions = {
  algorithm: "HS256",
  expiresIn: JWT_ACCESS_EXPIRY,
};

const signRefreshOptions: SignOptions = {
  algorithm: "HS256",
  expiresIn: JWT_REFRESH_EXPIRY,
};

/* ── Password ───────────────────────────────────────────────────────────────── */

/**
 * Hash a plain-text password using bcrypt.
 * @param password — The plain text password.
 */
export async function hashPassword(password: string): Promise<string> {
  return bcrypt.hash(password, BCRYPT_ROUNDS);
}

/**
 * Validate a plain-text password against a bcrypt hash.
 */
export async function comparePassword(
  password: string,
  hash: string,
): Promise<boolean> {
  return bcrypt.compare(password, hash);
}

/* ── Tokens ─────────────────────────────────────────────────────────────────── */

/**
 * Sign a short-lived access JWT for a user.
 */
export function signAccessToken(payload: JWTPayload): string {
  if (!JWT_SECRET) throw new Error("JWT_SECRET is not set");
  return jwt.sign(payload, JWT_SECRET, signOptions);
}

/**
 * Sign a long-lived refresh JWT (stored hashed in DB/Redis).
 */
export function signRefreshToken(payload: JWTPayload): string {
  if (!JWT_REFRESH_SECRET) throw new Error("JWT_REFRESH_SECRET is not set");
  return jwt.sign(payload, JWT_REFRESH_SECRET, signRefreshOptions);
}

/**
 * Verify an access JWT from the Authorization header.
 */
export function verifyAccessToken(token: string): JWTPayload | null {
  if (!JWT_SECRET) return null;
  try {
    return jwt.verify(token, JWT_SECRET, { algorithms: ["HS256"] }) as JWTPayload;
  } catch {
    return null;
  }
}

/**
 * Verify a refresh JWT.
 */
export function verifyRefreshToken(token: string): JWTPayload | null {
  if (!JWT_REFRESH_SECRET) return null;
  try {
    return jwt.verify(token, JWT_REFRESH_SECRET, { algorithms: ["HS256"] }) as JWTPayload;
  } catch {
    return null;
  }
}

/* ── Token Pair ─────────────────────────────────────────────────────────────── */

/**
 * Generate a fresh pair of access + refresh tokens for a user session.
 */
export function generateTokenPair(userId: string, role: string): {
  accessToken: string;
  refreshToken: string;
  expiresIn: string;
} {
  const payload: JWTPayload = {
    sub: userId,
    role: role as JWTPayload["role"],
    type: "access",
  };
  return {
    accessToken: signAccessToken(payload),
    refreshToken: signRefreshToken({ ...payload, type: "refresh" }),
    expiresIn: JWT_ACCESS_EXPIRY,
  };
}

/* ── Token hashing ──────────────────────────────────────────────────────────── */

/**
 * One-way hash a raw refresh token before persisting.
 * We store `bcrypt.hash(token)` so the plain token can never be recovered.
 */
export async function hashToken(token: string): Promise<string> {
  return bcrypt.hash(token, BCRYPT_ROUNDS);
}

/**
 * Compare a raw refresh-token string against its stored bcrypt hash.
 */
export async function compareTokenHash(
  token: string,
  hash: string,
): Promise<boolean> {
  return bcrypt.compare(token, hash);
}
