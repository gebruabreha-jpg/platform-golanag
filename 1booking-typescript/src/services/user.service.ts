import type { IUserService, UpdateProfileInput } from "@/services/interfaces/index.js";
import type { UserPublic } from "@/types/index.js";
import { AppError } from "@/errors/AppError.js";
import { ErrorCodes } from "@/errors/errorCodes.js";
import { getPostgresPool } from "@/clients/postgres.client.js";
import { hashPassword } from "@/utils/crypto.js";

export class UserService implements IUserService {
  async getProfile(id: string): Promise<UserPublic | null> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `SELECT id, email, first_name as "firstName", last_name as "lastName", 
              phone, bio, avatar_url as "avatarUrl", location, country, city, 
              role, is_verified as "isVerified", verification_level as "verificationLevel",
              trust_score as "trustScore", total_transactions as "totalTransactions",
              created_at as "createdAt", updated_at as "updatedAt"
       FROM users WHERE id = $1 AND deleted_at IS NULL`,
      [id],
    );
    if (result.rows.length === 0) return null;
    return result.rows[0] as UserPublic;
  }

  async getUserByEmail(email: string): Promise<UserPublic | null> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `SELECT id, email, first_name as "firstName", last_name as "lastName",
              phone, bio, avatar_url as "avatarUrl", location, country, city,
              role, is_verified as "isVerified", verification_level as "verificationLevel",
              trust_score as "trustScore", total_transactions as "totalTransactions",
              created_at as "createdAt", updated_at as "updatedAt"
       FROM users WHERE email = $1 AND deleted_at IS NULL`,
      [email],
    );
    if (result.rows.length === 0) return null;
    return result.rows[0] as UserPublic;
  }

  async updateProfile(id: string, data: UpdateProfileInput): Promise<UserPublic | null> {
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
      `UPDATE users SET ${fields.join(", ")}, updated_at = NOW()
       WHERE id = $${idx} AND deleted_at IS NULL
       RETURNING id, email, first_name as "firstName", last_name as "lastName",
                 phone, bio, avatar_url as "avatarUrl", location, country, city,
                 role, is_verified as "isVerified", verification_level as "verificationLevel",
                 trust_score as "trustScore", total_transactions as "totalTransactions",
                 created_at as "createdAt", updated_at as "updatedAt"`,
      values,
    );
    return result.rows[0] as UserPublic ?? null;
  }

  async getAllUsers(query: { page: number; limit: number; role?: string }): Promise<UserPublic[]> {
    const pool = getPostgresPool();
    const offset = (query.page - 1) * query.limit;
    let where = "deleted_at IS NULL";
    const values: unknown[] = [];
    if (query.role) {
      values.push(query.role);
      where += ` AND role = $${values.length}`;
    }
    values.push(offset, query.limit);
    const result = await pool.query(
      `SELECT id, email, first_name as "firstName", last_name as "lastName",
              phone, bio, avatar_url as "avatarUrl", location, country, city,
              role, is_verified as "isVerified", verification_level as "verificationLevel",
              trust_score as "trustScore", total_transactions as "totalTransactions",
              created_at as "createdAt", updated_at as "updatedAt"
       FROM users WHERE ${where} ORDER BY created_at DESC LIMIT $${values.length - 1} OFFSET $${values.length - 2}`,
      values,
    );
    return result.rows as UserPublic[];
  }

  async setUserRole(id: string, role: string): Promise<UserPublic | null> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `UPDATE users SET role = $1, updated_at = NOW()
       WHERE id = $2 AND deleted_at IS NULL
       RETURNING id, email, first_name as "firstName", last_name as "lastName",
                 phone, bio, avatar_url as "avatarUrl", location, country, city,
                 role, is_verified as "isVerified", verification_level as "verificationLevel",
                 trust_score as "trustScore", total_transactions as "totalTransactions",
                 created_at as "createdAt", updated_at as "updatedAt"`,
      [role, id],
    );
    return result.rows[0] as UserPublic ?? null;
  }

  async deactivateUser(id: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`UPDATE users SET deleted_at = NOW() WHERE id = $1`, [id]);
  }
}