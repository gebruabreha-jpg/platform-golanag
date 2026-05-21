import { Request, Response } from "express";
import { errorResponse } from "@/utils/response.js";

/**
 * 404 handler – catches any request that was not matched by a route.
 */
export function notFoundHandler(_req: Request, res: Response): Response {
  return res.status(404).json(
    errorResponse("NOT_FOUND", "Endpoint not found."),
  );
}
