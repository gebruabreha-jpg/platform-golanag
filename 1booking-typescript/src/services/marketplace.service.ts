import type { IMarketplaceService, CreateMarketplaceItemInput, MarketPlaceSearchInput, CreateTransactionInput } from "@/services/interfaces/index.js";
import type { MarketplaceItem, Transaction } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";
import { randomUUID } from "crypto";

export class MarketplaceService implements IMarketplaceService {
  async createItem(data: CreateMarketplaceItemInput, sellerId: string): Promise<MarketplaceItem> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO marketplace_items (id, seller_id, title, description, category, subcategory, price, currency,
              condition, location, country, shipping_available, shipping_cost, image_urls)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
       RETURNING *`,
      [randomUUID(), sellerId, data.title, data.description, data.category, data.subcategory, data.price, data.currency ?? "USD",
       data.condition, data.location, data.country, data.shippingAvailable ?? false, data.shippingCost ?? 0, data.imageUrls ?? []],
    );
    return result.rows[0] as MarketplaceItem;
  }

  async updateItem(id: string, data: Partial<CreateMarketplaceItemInput>): Promise<MarketplaceItem> {
    const pool = getPostgresPool();
    const fields: string[] = [];
    const values: unknown[] = [];
    let idx = 1;
    for (const [key, value] of Object.entries(data)) {
      if (value !== undefined) {
        fields.push(`${key} = $${idx++}`);
        values.push(value);
      }
    }
    values.push(id);
    const result = await pool.query(
      `UPDATE marketplace_items SET ${fields.join(", ")}, updated_at = NOW() WHERE id = $${idx} RETURNING *`,
      values,
    );
    return result.rows[0] as MarketplaceItem;
  }

  async getItemById(id: string): Promise<MarketplaceItem | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM marketplace_items WHERE id = $1`, [id]);
    return result.rows[0] as MarketplaceItem ?? null;
  }

  async searchItems(query: MarketPlaceSearchInput): Promise<{ items: MarketplaceItem[]; meta: { page: number; limit: number; total: number; totalPages: number } }> {
    const pool = getPostgresPool();
    const offset = (query.page - 1) * query.limit;
    let where = "is_active = TRUE AND is_sold = FALSE";
    const values: unknown[] = [];
    if (query.category) { values.push(query.category); where += ` AND category = $${values.length}`; }
    if (query.condition) { values.push(query.condition); where += ` AND condition = $${values.length}`; }
    if (query.minPrice) { values.push(query.minPrice); where += ` AND price >= $${values.length}`; }
    if (query.maxPrice) { values.push(query.maxPrice); where += ` AND price <= $${values.length}`; }
    if (query.country) { values.push(query.country); where += ` AND country = $${values.length}`; }
    if (query.search) { values.push(`%${query.search}%`); where += ` AND title ILIKE $${values.length}`; }
    values.push(query.limit, offset);
    const result = await pool.query(
      `SELECT * FROM marketplace_items WHERE ${where} ORDER BY created_at DESC LIMIT $${values.length - 1} OFFSET $${values.length}`,
      values,
    );
    const totalResult = await pool.query(`SELECT COUNT(*) FROM marketplace_items WHERE ${where.split("LIMIT")[0].trim()}`, values.slice(0, -2));
    const total = parseInt(totalResult.rows[0].count, 10);
    return { items: result.rows as MarketplaceItem[], meta: { page: query.page, limit: query.limit, total, totalPages: Math.ceil(total / query.limit) } };
  }

  async markItemSold(itemId: string): Promise<void> {
    await getPostgresPool().query(`UPDATE marketplace_items SET is_sold = TRUE WHERE id = $1`, [itemId]);
  }

  async createTransaction(data: CreateTransactionInput): Promise<Transaction> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO transactions (id, buyer_id, seller_id, item_id, type, amount, currency, status, payment_method)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
       RETURNING *`,
      [randomUUID(), data.buyerId, data.sellerId, data.itemId, data.type, data.amount, data.currency ?? "USD", "PENDING", data.paymentMethod],
    );
    return result.rows[0] as Transaction;
  }

  async getTransaction(id: string): Promise<Transaction | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM transactions WHERE id = $1`, [id]);
    return result.rows[0] as Transaction ?? null;
  }

  async getSellerTransactions(sellerId: string): Promise<Transaction[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM transactions WHERE seller_id = $1 ORDER BY created_at DESC`, [sellerId]);
    return result.rows as Transaction[];
  }
}