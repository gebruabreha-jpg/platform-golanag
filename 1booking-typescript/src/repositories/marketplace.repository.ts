import type { IMarketplaceRepository } from "@/repositories/interfaces/IMarketplaceRepository.js";
import type { MarketplaceItem } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";

export class MarketplaceRepository implements IMarketplaceRepository {
  async findById(id: string): Promise<MarketplaceItem | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE id = $1`, [id]);
    return result.rows[0] as MarketplaceItem ?? null;
  }

  async findAll(): Promise<MarketplaceItem[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE is_active = TRUE AND is_sold = FALSE`);
    return result.rows as MarketplaceItem[];
  }

  async findBySeller(sellerId: string): Promise<MarketplaceItem[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE seller_id = $1`, [sellerId]);
    return result.rows as MarketplaceItem[];
  }

  async findByCategory(category: string): Promise<MarketplaceItem[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE category = $1`, [category]);
    return result.rows as MarketplaceItem[];
  }

  async search(query: Record<string, unknown>): Promise<MarketplaceItem[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE is_active = TRUE AND is_sold = FALSE`);
    return result.rows as MarketplaceItem[];
  }
}