import type { IHousingRepository } from "@/repositories/interfaces/IHousingRepository.js";
import type { HousingListing } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";

export class HousingRepository implements IHousingRepository {
  async findById(id: string): Promise<HousingListing | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE id = $1`, [id]);
    return result.rows[0] as HousingListing ?? null;
  }

  async findAll(): Promise<HousingListing[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE is_active = TRUE`);
    return result.rows as HousingListing[];
  }

  async findByLandlord(landlordId: string): Promise<HousingListing[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE landlord_id = $1`, [landlordId]);
    return result.rows as HousingListing[];
  }

  async findByLocation(city: string, country: string): Promise<HousingListing[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE city = $1 AND country = $2`, [city, country]);
    return result.rows as HousingListing[];
  }

  async search(query: Record<string, unknown>): Promise<HousingListing[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM housing_listings WHERE is_active = TRUE`);
    return result.rows as HousingListing[];
  }
}