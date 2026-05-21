import type { ICommunityService, CreateCommunityInput, SearchCommunityInput, CreatePostInput } from "@/services/interfaces/index.js";
import type { Community, Post } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";
import { randomUUID } from "crypto";

export class CommunityService implements ICommunityService {
  async getCommunityById(id: string): Promise<Community | null> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `SELECT id, name, description, category, location, country, is_private as "isPrivate",
              member_count as "memberCount", moderator_id as "moderatorId", rules, image_url as "imageUrl",
              tags, created_at as "createdAt", updated_at as "updatedAt"
       FROM communities WHERE id = $1`,
      [id],
    );
    return result.rows[0] as Community ?? null;
  }

  async getById(id: string): Promise<Community | null> {
    return this.getCommunityById(id);
  }

  async create(data: CreateCommunityInput & { moderatorId?: string }): Promise<Community> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO communities (id, name, description, category, country, is_private, moderator_id, image_url, rules, tags)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
       RETURNING id, name, description, category, location, country, is_private as "isPrivate",
                 member_count as "memberCount", moderator_id as "moderatorId", rules, image_url as "imageUrl",
                 tags, created_at as "createdAt", updated_at as "updatedAt"`,
      [randomUUID(), data.name, data.description, data.category, data.country, data.isPrivate ?? false, 
       data.moderatorId ?? "system", data.imageUrl, data.rules, data.tags ?? []],
    );
    return result.rows[0] as Community;
  }

  async getCommunities(query: SearchCommunityInput): Promise<{ items: Community[]; meta: { page: number; limit: number; total: number; totalPages: number } }> {
    const pool = getPostgresPool();
    const offset = (query.page - 1) * query.limit;
    let where = "TRUE";
    const values: unknown[] = [];
    if (query.category) { values.push(query.category); where += ` AND category = $${values.length}`; }
    if (query.country) { values.push(query.country); where += ` AND country = $${values.length}`; }
    if (query.isPrivate !== undefined) { values.push(query.isPrivate); where += ` AND is_private = $${values.length}`; }
    if (query.search) { values.push(`%${query.search}%`); where += ` AND name ILIKE $${values.length}`; }
    values.push(query.limit, offset);
    const result = await pool.query(
      `SELECT id, name, description, category, location, country, is_private as "isPrivate",
              member_count as "memberCount", moderator_id as "moderatorId", rules, image_url as "imageUrl",
              tags, created_at as "createdAt", updated_at as "updatedAt"
       FROM communities WHERE ${where} ORDER BY created_at DESC LIMIT $${values.length - 1} OFFSET $${values.length}`,
      values,
    );
    const totalResult = await pool.query(`SELECT COUNT(*) FROM communities WHERE ${where.split("LIMIT")[0].trim()}`, values.slice(0, -2));
    const total = parseInt(totalResult.rows[0].count, 10);
    return {
      items: result.rows as Community[],
      meta: { page: query.page, limit: query.limit, total, totalPages: Math.ceil(total / query.limit) },
    };
  }

  async addPost(communityId: string, data: CreatePostInput): Promise<Post> {
    const pool = getPostgresPool();
    const result = await pool.query(
      `INSERT INTO posts (id, community_id, user_id, type, title, content, media_urls, is_pinned)
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
       RETURNING id, community_id as "communityId", user_id as "userId", type, title, content,
                 media_urls as "mediaUrls", is_pinned as "isPinned", is_closed as "isClosed",
                 reply_count as "replyCount", view_count as "viewCount", created_at as "createdAt", updated_at as "updatedAt"`,
      [randomUUID(), communityId, data.userId, data.type, data.title, data.content, data.mediaUrls ?? [], data.isPinned ?? false],
    );
    return result.rows[0] as Post;
  }

  async addMember(communityId: string, userId: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`INSERT INTO community_members (id, community_id, user_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`, [randomUUID(), communityId, userId]);
    await pool.query(`UPDATE communities SET member_count = member_count + 1 WHERE id = $1`, [communityId]);
  }

  async connectUserToCommunity(communityId: string, userId: string): Promise<void> {
    await this.addMember(communityId, userId);
  }

  async disconnectUserFromCommunity(communityId: string, userId: string): Promise<void> {
    await this.removeMember(communityId, userId);
  }

  async removeMember(communityId: string, userId: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`DELETE FROM community_members WHERE community_id = $1 AND user_id = $2`, [communityId, userId]);
    await pool.query(`UPDATE communities SET member_count = GREATEST(member_count - 1, 0) WHERE id = $1`, [communityId]);
  }

  async update(id: string, data: Partial<CreateCommunityInput>): Promise<Community | null> {
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
      `UPDATE communities SET ${fields.join(", ")}, updated_at = NOW() WHERE id = $${idx}
       RETURNING id, name, description, category, location, country, is_private as "isPrivate",
                 member_count as "memberCount", moderator_id as "moderatorId", rules, image_url as "imageUrl",
                 tags, created_at as "createdAt", updated_at as "updatedAt"`,
      values,
    );
    return result.rows[0] as Community ?? null;
  }

  async delete(id: string): Promise<void> {
    const pool = getPostgresPool();
    await pool.query(`DELETE FROM communities WHERE id = $1`, [id]);
  }

  async getPosts(communityId: string): Promise<Post[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM posts WHERE community_id = $1 ORDER BY created_at DESC`, [communityId]);
    return result.rows as Post[];
  }
}