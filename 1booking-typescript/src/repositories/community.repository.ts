import type { ICommunityRepository } from "@/repositories/interfaces/ICommunityRepository.js";
import type { Community } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";

export class CommunityRepository implements ICommunityRepository {
  async findById(id: string): Promise<Community | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM communities WHERE id = $1`, [id]);
    return result.rows[0] as Community ?? null;
  }

  async findAll(): Promise<Community[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM communities`);
    return result.rows as Community[];
  }

  async findByCategory(category: string): Promise<Community[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM communities WHERE category = $1`, [category]);
    return result.rows as Community[];
  }

  async findByName(name: string): Promise<Community | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM communities WHERE name = $1`, [name]);
    return result.rows[0] as Community ?? null;
  }

  async getPosts(communityId: string, limit: number, offset: number): Promise<any[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM posts WHERE community_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, [communityId, limit, offset]);
    return result.rows;
  }

  async addMember(communityId: string, userId: string, role: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`INSERT INTO community_members (id, community_id, user_id, role) VALUES (gen_random_uuid(), $1, $2, $3) ON CONFLICT DO NOTHING`, [communityId, userId, role]);
    await pool.query(`UPDATE communities SET member_count = member_count + 1 WHERE id = $1`, [communityId]);
  }

  async removeMember(communityId: string, userId: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`DELETE FROM community_members WHERE community_id = $1 AND user_id = $2`, [communityId, userId]);
    await pool.query(`UPDATE communities SET member_count = GREATEST(member_count - 1, 0) WHERE id = $1`, [communityId]);
  }

  async isMember(communityId: string, userId: string): Promise<boolean> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT 1 FROM community_members WHERE community_id = $1 AND user_id = $2`, [communityId, userId]);
    return result.rowCount > 0;
  }
}