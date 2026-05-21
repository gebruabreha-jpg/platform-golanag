import type { Request, Response, NextFunction } from "express";
import { v4 as uuidv4 } from "uuid";

/**
 * Attach a unique `req.requestId` (UUID v4) to every incoming request.
 * Useful for log correlation across services.
 */
export function requestId(req: Request, _res: Response, next: NextFunction): void {
  req.requestId = req.header("X-Request-ID") ?? uuidv4();
  next();
}
