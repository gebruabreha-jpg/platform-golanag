import type { IAuthService, RegisterInput, RefreshResult } from "@/services/interfaces/index.js";
import type { Request, Response } from "express";
import { AppError } from "@/errors/AppError.js";
import { successResponse } from "@/utils/response.js";
import { AuthService } from "@/services/auth.service.js";

export class AuthController {
  private static authService = new AuthService();

  static async register(req: Request, res: Response): Promise<void> {
    try {
      const user = await AuthController.authService.register(req.body as RegisterInput);
      res.status(201).json(successResponse(user));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Registration failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  static async login(req: Request, res: Response): Promise<void> {
    try {
      const { email, password } = req.body;
      const result = await AuthController.authService.login(email, password);
      res.status(200).json(successResponse(result));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Login failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  static async refresh(req: Request, res: Response): Promise<void> {
    try {
      const { refreshToken } = req.body;
      const result = await AuthController.authService.refresh(refreshToken);
      res.status(200).json(successResponse(result));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("TOKEN_INVALID", "Invalid refresh token.", 401);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  static async logout(req: Request, res: Response): Promise<void> {
    try {
      await AuthController.authService.logout(req.user!.sub, req.body.refreshToken);
      res.status(200).json(successResponse(null, { message: "Logged out successfully." }));
    } catch {
      res.status(200).json(successResponse(null, { message: "Logged out successfully." }));
    }
  }

  static async changePassword(req: Request, res: Response): Promise<void> {
    try {
      await AuthController.authService.changePassword(req.user!.sub, req.body.currentPassword, req.body.newPassword);
      res.status(200).json(successResponse(null, { message: "Password updated successfully." }));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Password change failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }
}