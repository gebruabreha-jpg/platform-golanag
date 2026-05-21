import type { Request, Response } from "express";
import { IUserService } from "@/services/interfaces/index.js";
import { successResponse, paginatedResponse } from "@/utils/response.js";
import { AppError } from "@/errors/AppError.js";

export class UserController {
  constructor(private readonly users: IUserService) {
    this.getProfile    = this.getProfile.bind(this);
    this.updateProfile = this.updateProfile.bind(this);
    this.getAllUsers   = this.getAllUsers.bind(this);
    this.deleteUser    = this.deleteUser.bind(this);
  }

  /* ── GET /users/profile ─────────────────────────────────────────────────── */
  async getProfile(req: Request, res: Response): Promise<void> {
    try {
      const user = await this.users.getProfile(req.user!.sub);
      if (!user) { res.status(404).json({ success:false, error:{code:"NOT_FOUND",message:"User not found."} }); return; }
      res.status(200).json(successResponse(user));
    } catch { res.status(500).json({ success:false, error:{code:"INTERNAL_ERROR",message:"Failed to fetch profile."} }); }
  }

  /* ── PATCH /users/profile ────────────────────────────────────────────────── */
  async updateProfile(req: Request, res: Response): Promise<void> {
    try {
      const user = await this.users.updateProfile(req.user!.sub, req.body);
      res.status(200).json(successResponse(user));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Profile update failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  /* ── GET /users ──────────────────────────────────────────────────────────── */
  async getAllUsers(req: Request, res: Response): Promise<void> {
    try {
      const { page = "1", limit = "20", role } = req.query as Record<string, string>;
      const users  = await this.users.getAllUsers({ page: +page, limit: +limit, role });
      const total  = await this.users.count();
      res.status(200).json(paginatedResponse(users, { page: +page, limit: +limit, total, totalPages: 1 }));
    } catch {
      res.status(500).json({ success:false, error:{code:"INTERNAL_ERROR",message:"Failed to fetch users."} });
    }
  }

  /* ── DELETE /users/:id ───────────────────────────────────────────────────── */
  async deleteUser(req: Request, res: Response): Promise<void> {
    try {
      await this.users.deactivateUser(req.params.id);
      res.status(200).json(successResponse(null, { message: "User deactivated." }));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("NOT_FOUND", "User not found.", 404);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }
}
