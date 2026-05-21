import type { IHousingService, CreateListingInput, ListingSearchInput, CreateApplicationInput } from "@/services/interfaces/index.js";
import type { HousingListing, HousingApplication } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";
import { randomUUID } from "crypto";

export class HousingService implements IHousingService {
  async createListing(data: CreateListingInput & { landlordId: string }): Promise<HousingListing> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO housing_listings (id, landlord_id, title, description, property_type, room_type,
              bedrooms, bathrooms, monthly_rent, deposit, address, city, country, latitude, longitude,
              available_from, lease_term, furnished, includes_utilities, image_urls)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
       RETURNING *`,
      [randomUUID(), data.landlordId, data.title, data.description, data.propertyType, data.roomType,
       data.bedrooms, data.bathrooms, data.monthlyRent, data.deposit, data.address, data.city, data.country,
       data.latitude, data.longitude, data.availableFrom, data.leaseTerm, data.furnished, data.includesUtilities, data.imageUrls],
    );
    return result.rows[0] as HousingListing;
  }

  async updateListing(id: string, data: Partial<CreateListingInput>): Promise<HousingListing> {
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
      `UPDATE housing_listings SET ${fields.join(", ")}, updated_at = NOW() WHERE id = $${idx} RETURNING *`,
      values,
    );
    return result.rows[0] as HousingListing;
  }

  async getListing(id: string): Promise<HousingListing | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE id = $1`, [id]);
    return result.rows[0] as HousingListing ?? null;
  }

  async getListings(query: ListingSearchInput): Promise<{ items: HousingListing[]; meta: { page: number; limit: number; total: number; totalPages: number } }> {
    const pool = getPostgresPool();
    const offset = (query.page - 1) * query.limit;
    let where = "is_active = TRUE";
    const values: unknown[] = [];
    if (query.city) { values.push(query.city); where += ` AND city = $${values.length}`; }
    if (query.propertyType) { values.push(query.propertyType); where += ` AND property_type = $${values.length}`; }
    if (query.minRent) { values.push(query.minRent); where += ` AND monthly_rent >= $${values.length}`; }
    if (query.maxRent) { values.push(query.maxRent); where += ` AND monthly_rent <= $${values.length}`; }
    if (query.bedrooms) { values.push(query.bedrooms); where += ` AND bedrooms >= $${values.length}`; }
    values.push(query.limit, offset);
    const result = await pool.query(
      `SELECT * FROM housing_listings WHERE ${where} ORDER BY created_at DESC LIMIT $${values.length - 1} OFFSET $${values.length}`,
      values,
    );
    const totalResult = await pool.query(`SELECT COUNT(*) FROM housing_listings WHERE ${where.split("LIMIT")[0].trim()}`, values.slice(0, -2));
    const total = parseInt(totalResult.rows[0].count, 10);
    return { items: result.rows as HousingListing[], meta: { page: query.page, limit: query.limit, total, totalPages: Math.ceil(total / query.limit) } };
  }

  async createApplication(listingId: string, userId: string, data: CreateApplicationInput): Promise<HousingApplication> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO housing_applications (id, listing_id, user_id, message, move_in_date, proposed_rent)
       VALUES ($1, $2, $3, $4, $5, $6)
       RETURNING *`,
      [randomUUID(), listingId, userId, data.message, data.moveInDate, data.proposedRent],
    );
    await pool.query(`UPDATE housing_listings SET application_count = application_count + 1 WHERE id = $1`, [listingId]);
    return result.rows[0] as HousingApplication;
  }

  async getApplications(listingId: string): Promise<HousingApplication[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_applications WHERE listing_id = $1`, [listingId]);
    return result.rows as HousingApplication[];
  }

  async approveApplication(applicationId: string): Promise<void> {
    await getPostgresPool().query(`UPDATE housing_applications SET status = 'APPROVED' WHERE id = $1`, [applicationId]);
  }

  async rejectApplication(applicationId: string): Promise<void> {
    await getPostgresPool().query(`UPDATE housing_applications SET status = 'REJECTED' WHERE id = $1`, [applicationId]);
  }

  async deleteListing(id: string): Promise<void> {
    await getPostgresPool().query(`UPDATE housing_listings SET is_active = FALSE WHERE id = $1`, [id]);
  }

  async getApplicationById(id: string): Promise<HousingApplication | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_applications WHERE id = $1`, [id]);
    return result.rows[0] as HousingApplication ?? null;
  }

  async updateApplication(id: string, data: Partial<HousingApplication>): Promise<HousingApplication> {
    const pool = getPostgresPool();
    const fields: string[] = [];
    const values: unknown[] = [];
    let idx = 1;
    for (const [key, value] of Object.entries(data)) {
      if (value !== undefined) {
        const snakeKey = key.replace(/[A-Z]/g, (m) => "_" + m.toLowerCase());
        fields.push(`${snakeKey} = $${idx++}`);
        values.push(value);
      }
    }
    values.push(id);
    const result = await pool.query(
      `UPDATE housing_applications SET ${fields.join(", ")}, updated_at = NOW() WHERE id = $${idx} RETURNING *`,
      values,
    );
    return result.rows[0] as HousingApplication;
  }
}