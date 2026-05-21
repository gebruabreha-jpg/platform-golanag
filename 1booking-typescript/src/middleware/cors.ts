import { Request, Response, NextFunction } from "express";
import cors from "cors";

const defaultOrigins = process.env.CORS_ORIGIN
  ? process.env.CORS_ORIGIN.split(",").map((s) => s.trim())
  : ["http://localhost:3001"];

/**
 * CORS configuration.
 */
export const corsConfig = cors({
  origin: (origin, callback) => {
    if (!origin || defaultOrigins.includes(origin)) return callback(null, true);
    callback(new Error("Not allowed by CORS."));
  },
  credentials: true,
  methods: ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"],
  allowedHeaders: [
    "Authorization",
    "Content-Type",
    "Accept",
    "X-Request-ID",
    "X-CSRF-Token",
  ],
  exposedHeaders: ["X-Request-ID", "X-RateLimit-Limit", "X-RateLimit-Remaining"],
  maxAge: 86400,
  preflightContinue: false,
});

/**
 * Express middleware that handles CORS preflight explicitly before the cors
 * package handles the actual request to avoid double-processing.
 */
export function corsPreflight(req: Request, res: Response, next: NextFunction): void {
  if (req.method === "OPTIONS") {
    res.status(204).send("");
  } else {
    next();
  }
}
