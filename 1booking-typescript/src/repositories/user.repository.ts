import { IUserRepository, IBaseRepository } from "@/repositories/interfaces/IUserRepository.js";
import type { User } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";

export class UserRepository implements IUserRepository {
  async findById(id: string): Promise<User | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`, [id]);
    return result.rows[0] as User ?? null;
  }

  async findByIds(ids: string[]): Promise<User[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM users WHERE id = ANY($1) AND deleted_at IS NULL`, [ids]);
    return result.rows as User[];
  }

  async findAll(filter: Record<string, unknown> = {}): Promise<User[]> {
    const pool = getPostgresPool();
    const where = Object.keys(filter).map((k, i) => `${k} = $${i + 1}`).join(" AND ");
    const values = Object.values(filter);
    const result = await pool.query(`SELECT * FROM users WHERE deleted_at IS NULL ${where ? " AND " + where : ""}`, values);
    return result.rows as User[];
  }

  async create(data: Partial<User>): Promise<User> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO users (id, email, password_hash, first_name, last_name, phone, role, country, city)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`,
      [data.id, data.email, data.passwordHash, data.firstName, data.lastName, data.phone, data.role, data.country, data.city],
    );
    return result.rows[0] as User;
  }

  async update(id: string, data: Partial<User>): Promise<User | null> {
    const pool = getPostgresPool();
    const fields = Object.keys(data).filter((k) => k !== "id").map((k, i) => `"${k}" = $${i + 1}`);
    const values = [...Object.values(data).filter((_, i) => Object.keys(data)[i] !== "id"), id];
    const result = await pool.query(`UPDATE users SET ${fields.join(", ")}, updated_at = NOW() WHERE id = $${values.length} RETURNING *`, values);
    return result.rows[0] as User ?? null;
  }

  async delete(id: string): Promise<boolean> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET deleted_at = NOW() WHERE id = $1`, [id]);
    return true;
  }

  async count(filter: Record<string, unknown> = {}): Promise<number> {
    const pool = getPostgresPool();
    const where = Object.keys(filter).map((k, i) => `${k} = $${i + 1}`).join(" AND ");
    const values = Object.values(filter);
    const result = await pool.query(`SELECT COUNT(*) FROM users WHERE deleted_at IS NULL ${where ? " AND " + where : ""}`, values);
    return parseInt(result.rows[0].count, 10);
  }

  async existsById(id: string): Promise<boolean> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT 1 FROM users WHERE id = $1 AND deleted_at IS NULL`, [id]);
    return result.rowCount > 0;
  }

  async findByEmail(email: string): Promise<User | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`, [email]);
    return result.rows[0] as User ?? null;
  }

  async findByVerificationToken(token: string): Promise<User | null> {
    return null;
  }

  async findByRefreshTokenHash(tokenHash: string): Promise<User | null> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `SELECT u.* FROM users u JOIN refresh_tokens rt ON u.id = rt.user_id WHERE rt.token_hash = $1 AND rt.revoked_at IS NULL`,
      [tokenHash],
    );
    return result.rows[0] as User ?? null;
  }

  async updatePassword(id: string, passwordHash: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`, [passwordHash, id]);
  }

  async updateTrustScore(id: string, score: number): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET trust_score = $1, updated_at = NOW() WHERE id = $2`, [score, id]);
  }

  async updateVerificationLevel(id: string, level: number): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET verification_level = $1, updated_at = NOW() WHERE id = $2`, [level, id]);
  }

  async updateIsVerified(id: string, isVerified: boolean): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET is_verified = $1, updated_at = NOW() WHERE id = $2`, [isVerified, id]);
  }

  async findActiveUsers(limit: number, offset: number): Promise<User[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`, [limit, offset]);
    return result.rows as User[];
  }
}