import { IRoleService } from "@/services/interfaces/index.js";
import type { ICommunityService } from "@/services/interfaces/index.js";
import type { Request, Response } from "express";
import { AppError } from "@/errors/AppError.js";

export class CommunityController {
  constructor(private readonly communities: ICommunityService) {
    this.createCommunity     = this.createCommunity.bind(this);
    this.getCommunities      = this.getCommunities.bind(this);
    this.getCommunity        = this.getCommunity.bind(this);
    this.addPost             = this.addPost.bind(this);
    this.getPosts            = this.getPosts.bind(this);
    this.joinCommunity       = this.joinCommunity.bind(this);
    this.leaveCommunity      = this.leaveCommunity.bind(this);
    this.updateCommunity     = this.updateCommunity.bind(this);
    this.deleteCommunity     = this.deleteCommunity.bind(this);
  }

  async createCommunity(req: Request, res: Response): Promise<void> {
    try {
      const community = await this.communities.create(req.body);
      res.status(201).json(community);
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Community creation failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async getCommunities(req: Request, res: Response): Promise<void> {
    try {
      const result = await this.communities.getCommunities(req.query as any);
      res.status(200).json(result);
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to list communities." } }); }
  }

  async getCommunity(req: Request, res: Response): Promise<void> {
    try {
      const community = await this.communities.getCommunityById(req.params.id);
      if (!community) return res.status(404).json({ success: false, error: { code: "COMMUNITY_NOT_FOUND", message: "Community not found." } });
      res.status(200).json({ success: true, data: community });
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to fetch community." } }); }
  }

  async addPost(req: Request, res: Response): Promise<void> {
    try {
      const post = await this.communities.addPost(req.params.id, { ...req.body, userId: req.user!.sub });
      res.status(201).json({ success: true, data: post });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Failed to add post.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async getPosts(req: Request, res: Response): Promise<void> {
    try {
      const posts = await this.communities.getPosts(req.params.id);
      res.status(200).json({ success: true, data: posts });
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to fetch posts." } }); }
  }

  async joinCommunity(req: Request, res: Response): Promise<void> {
    try {
      await this.communities.connectUserToCommunity(req.params.id, req.user!.sub);
      res.status(200).json({ success: true, data: null, meta: { message: "Joined community." } });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Failed to join.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async leaveCommunity(req: Request, res: Response): Promise<void> {
    try {
      await this.communities.disconnectUserFromCommunity(req.params.id, req.user!.sub);
      res.status(200).json({ success: true, data: null, meta: { message: "Left community." } });
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to leave." } }); }
  }

  async updateCommunity(req: Request, res: Response): Promise<void> {
    try {
      const community = await this.communities.update(req.params.id, req.body);
      res.status(200).json({ success: true, data: community });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("NOT_FOUND", "Community not found or cannot be updated.", 404);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async deleteCommunity(req: Request, res: Response): Promise<void> {
    try {
      await this.communities.delete(req.params.id);
      res.status(200).json({ success: true, data: null, meta: { message: "Community deleted." } });
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Delete failed." } }); }
  }
}
