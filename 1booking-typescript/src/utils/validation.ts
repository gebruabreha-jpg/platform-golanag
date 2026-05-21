import { Request, Response, NextFunction } from "express";
import { z } from "zod";

/**
 * Wrap `value` in an express ValidationResult-compatible object
 * (plain `z.infer` result produces `.some()` / `.map()` automatically via
 *  the express-validator `body()` / `query()` helpers).
 */
export type ValidatedBody<T extends z.ZodType> = z.infer<T>;

/** Schema-based validation middleware factory. */
export function validateBody<T extends z.ZodType>(schema: T) {
  return (req: Request, _res: Response, next: NextFunction): void | NextFunction {
    const result = schema.safeParse(req.body);
    if (!result.success) {
      const errors = result.error.issues.map((i) => `${i.path.join(".")}: ${i.message}`);
      return next(new AppError(ValidationError, "Request validation failed", 400, undefined, errors));
    }
    next();
  };
}

export function validateQuery<T extends z.ZodType>(schema: T) {
  return (req: Request, _res: Response, next: NextFunction): void | NextFunction {
    const result = schema.safeParse(req.query);
    if (!result.success) {
      const errors = result.error.issues.map((i) => `${i.path.join(".")}: ${i.message}`);
      return next(new AppError(ValidationError, "Request validation failed", 400, undefined, errors));
    }
    next();
  };
}

/**
 * Sanitise a string by removing control characters.
 */
export function sanitiseString(str: unknown): string | undefined {
  if (typeof str !== "string") return undefined;
  return str.replace(/[^\x20-\x7E]/g, "");
}

export function validateEmail(email: unknown): string | undefined {
  if (typeof email !== "string") return undefined;
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return re.test(email) ? email.toLowerCase().trim() : undefined;
}
