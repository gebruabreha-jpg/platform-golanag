import type { IHousingService } from "@/services/interfaces/index.js";
import type { Request, Response } from "express";
import type { PaginatedResult } from "@/types/index.js";
import { AppError } from "@/errors/AppError.js";

export interface HousingControllerDeps {
  readonly housingService: IHousingService;
}

export class HousingController {
  constructor(private readonly services: HousingControllerDeps) {
    this.createListing = this.createListing.bind(this);
    this.getListings   = this.getListings.bind(this);
    this.getListing    = this.getListing.bind(this);
    this.updateListing = this.updateListing.bind(this);
    this.deleteListing = this.deleteListing.bind(this);
  }

  /* ── POST /housing/listings ──────────────────────────────────────────────── */
  async createListing(req: Request, res: Response): Promise<void> {
    const { housingService } = this.services;
    try {
      const listing = await housingService.createListing(req.body, req.user!.sub);
      res.status(201).json({ success: true, data: listing });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Listing creation failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  /* ── GET /housing/listings ───────────────────────────────────────────────── */
  async getListings(req: Request, res: Response): Promise<void> {
    const { housingService } = this.services;
    try {
      const { city, propertyType, minRent, maxRent, bedrooms, page = "1", limit = "20" } = req.query;
      const result: PaginatedResult<any> = await housingService.getListings({
        city:          city as string | undefined,
        propertyType:  propertyType as any,
        minRent:       minRent  ? +minRent : undefined,
        maxRent:       maxRent  ? +maxRent : undefined,
        bedrooms:      bedrooms ? +bedrooms: undefined,
        page:          +page,
        limit:         +limit,
      });
      res.status(200).json(paginatedResponse(result.items, result.meta));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Listing search failed." } }); }
  }

  /* ── GET /housing/listings/:id ───────────────────────────────────────────── */
  async getListing(req: Request, res: Response): Promise<void> {
    const { housingService } = this.services;
    try {
      const listing = await housingService.getListing(req.params.id!);
      if (!listing) return res.status(404).json({ success: false, error: { code: "LISTING_NOT_FOUND", message: "Listing not found." } });
      res.status(200).json(successResponse(listing));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to fetch listing." } }); }
  }

  /* ── PATCH /housing/listings/:id ─────────────────────────────────────────── */
  async updateListing(req: Request, res: Response): Promise<void> {
    const { housingService } = this.services;
    try {
      const listing = await housingService.updateListing(req.params.id!, req.body);
      res.status(200).json({ success: true, data: listing });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("NOT_FOUND", "Listing not found.", 404);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  /* ── DELETE /housing/listings/:id ────────────────────────────────────────── */
  async deleteListing(req: Request, res: Response): Promise<void> {
    try {
      await this.services.housingService.deleteListing(req.params.id!);
      res.status(200).json({ success: true, data: null, meta: { message: "Listing deleted." } });
    } catch { res.status(500).json(successResponse(null)); }
  }

  /* ── POST /housing/listings/:id/applications ─────────────────────────────── */
  async createApplication(req: Request, res: Response): Promise<void> {
    const { housingService } = this.services;
    try {
      const app = await housingService.createApplication(req.params.id!, req.user!.sub, req.body);
      res.status(201).json({ success: true, data: app });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Application failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  /* ── GET /housing/applications/:id ───────────────────────────────────────── */
  async getApplication(req: Request, res: Response): Promise<void> {
    try {
      const app = await this.services.housingService.getApplicationById(req.params.id!);
      if (!app) return res.status(404).json({ success: false, error: { code: "NOT_FOUND", message: "Application not found." } });
      res.status(200).json({ success: true, data: app });
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR" } }); }
  }

  /* ── PATCH /housing/applications/:id ─────────────────────────────────────── */
  async updateApplication(req: Request, res: Response): Promise<void> {
    try {
      const app = await this.services.housingService.updateApplication(req.params.id!, req.body);
      res.status(200).json({ success: true, data: app });
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("NOT_FOUND", "Application not found.", 404);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }
}
