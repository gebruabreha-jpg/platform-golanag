import type { IMarketplaceService, MarketPlaceSearchInput } from "@/services/interfaces/index.js";
import type { Request, Response } from "express";
import { AppError } from "@/errors/AppError.js";
import { successResponse, paginatedResponse } from "@/utils/response.js";

export class MarketplaceController {
  constructor(private readonly marketplace: IMarketplaceService) {
    this.createItem = this.createItem.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getItems = this.getItems.bind(this);
    this.updateItem = this.updateItem.bind(this);
    this.deleteItem = this.deleteItem.bind(this);
    this.markItemSold = this.markItemSold.bind(this);
    this.createTransaction = this.createTransaction.bind(this);
    this.getTransaction = this.getTransaction.bind(this);
  }

  async createItem(req: Request, res: Response): Promise<void> {
    try {
      const item = await this.marketplace.createItem(req.body, req.user!.sub);
      res.status(201).json(successResponse(item));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Item creation failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async getItem(req: Request, res: Response): Promise<void> {
    try {
      const item = await this.marketplace.getItemById(req.params.id!);
      if (!item) return res.status(404).json({ success: false, error: { code: "NOT_FOUND", message: "Item not found." } });
      res.status(200).json(successResponse(item));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to fetch item." } }); }
  }

  async getItems(req: Request, res: Response): Promise<void> {
    try {
      const { page = "1", limit = "20", category, condition, minPrice, maxPrice, country, search } = req.query;
      const result = await this.marketplace.searchItems({
        category,
        condition: condition as MarketPlaceSearchInput["condition"],
        minPrice: minPrice ? +minPrice : undefined,
        maxPrice: maxPrice ? +maxPrice : undefined,
        country: country as string | undefined,
        search: search as string | undefined,
        page: +page,
        limit: +limit,
      });
      res.status(200).json(paginatedResponse(result.items, result.meta));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Search failed." } }); }
  }

  async updateItem(req: Request, res: Response): Promise<void> {
    try {
      const item = await this.marketplace.updateItem(req.params.id!, req.body);
      res.status(200).json(successResponse(item));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("NOT_FOUND", "Item not found.", 404);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async deleteItem(req: Request, res: Response): Promise<void> {
    try {
      await this.marketplace.markItemSold(req.params.id!);
      res.status(200).json(successResponse(null, { message: "Item deleted." }));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Delete failed." } }); }
  }

  async markItemSold(req: Request, res: Response): Promise<void> {
    try {
      await this.marketplace.markItemSold(req.params.id!);
      res.status(200).json(successResponse(null, { message: "Item marked as sold." }));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Operation failed." } }); }
  }

  async createTransaction(req: Request, res: Response): Promise<void> {
    try {
      const tx = await this.marketplace.createTransaction({ ...req.body, buyerId: req.user!.sub });
      res.status(201).json(successResponse(tx));
    } catch (err: unknown) {
      const e = err instanceof AppError ? err : new AppError("INTERNAL_ERROR", "Transaction failed.", 500);
      res.status(e.statusCode).json({ success: false, error: { code: e.code, message: e.message } });
    }
  }

  async getTransaction(req: Request, res: Response): Promise<void> {
    try {
      const tx = await this.marketplace.getTransaction(req.params.id!);
      if (!tx) return res.status(404).json({ success: false, error: { code: "NOT_FOUND", message: "Transaction not found." } });
      res.status(200).json(successResponse(tx));
    } catch { res.status(500).json({ success: false, error: { code: "INTERNAL_ERROR", message: "Failed to fetch transaction." } }); }
  }
}