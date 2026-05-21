import { Response, NextFunction } from "express";
import type { verifyAccessToken, verifyRefreshToken } from "@/utils/crypto.js";

export interface AuthenticatedRequest {
  user?: { sub: string; role: string; type: string };
}

/**
 * Strict authentication middleware factory.
 */
export function authenticate(
  verifyAt: typeof verifyAccessToken,
  verifyRt: typeof verifyRefreshToken,
): (req: unknown, res: Response, next: NextFunction) => void {
  return (req: unknown, res: Response, next: NextFunction): void => {
    const r = req as AuthenticatedRequest;
    const header = r.headers.authorization;
    if (!header?.startsWith("Bearer ")) return sendUnauthorized(res);
    const token = header.slice("Bearer ".length);
    const payload = verifyAt(token);
    if (!payload) return sendUnauthorized(res, "Invalid or expired token.");
    r.user = payload;
    return next();
  };
}

/** Optional auth – adds req.user only when token is valid. */
export function optionalAuth(
  verifyAt: typeof verifyAccessToken,
): (req: unknown, res: Response, next: NextFunction) => void {
  return (req: unknown, _res: Response, next: NextFunction): void => {
    const r = req as AuthenticatedRequest;
    const header = r.headers.authorization;
    if (header?.startsWith("Bearer ")) {
      const payload = verifyAt(header.slice("Bearer ".length));
      if (payload) r.user = payload;
    }
    return next();
  };
}

/** RBAC guard – chain after authenticate(). */
export function requireRole(
  ...roles: string[]
): (req: AuthenticatedRequest, res: Response, next: NextFunction) => void {
  return (req: AuthenticatedRequest, res: Response, next: NextFunction): void => {
    if (!req.user) return sendUnauthorized(res, "Authentication required before role check.");
    if (roles.includes(req.user.role) || roles.includes("*")) return next();
    sendForbidden(res, `Role "${req.user.role}" is not authorized for this resource.`);
  };
}

/* helpers */
function sendUnauthorized(res: Response, message = "Authentication required"): void {
  const { errorResponse } = require("@/utils/response") as typeof import("@/utils/response");
  return res.status(401).json(errorResponse("UNAUTHORIZED", message));
}

function sendForbidden(res: Response, message = "Access denied"): void {
  const { errorResponse } = require("@/utils/response") as typeof import("@/utils/response");
  return res.status(403).json(errorResponse("FORBIDDEN", message));
}