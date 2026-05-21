import type { IServiceRepository } from "@/repositories/interfaces/IServiceRepository.js";
import type { ServiceProfessional } from "@/types/index.js";
import { getPostgresPool } from "@/clients/postgres.client.js";

export class ServiceRepository implements IServiceRepository {
  async findById(id: string): Promise<ServiceProfessional | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM service_professionals WHERE id = $1`, [id]);
    return result.rows[0] as ServiceProfessional ?? null;
  }

  async findAll(): Promise<ServiceProfessional[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM service_professionals`);
    return result.rows as ServiceProfessional[];
  }

  async findByUserId(userId: string): Promise<ServiceProfessional | null> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM service_professionals WHERE user_id = $1`, [userId]);
    return result.rows[0] as ServiceProfessional ?? null;
  }

  async findBySpecialization(specialization: string): Promise<ServiceProfessional[]> {
    const pool = getPostgresPool();
    const result = await pool.query(`SELECT * FROM service_professionals WHERE specialization = $1`, [specialization]);
    return result.rows as ServiceProfessional[];
  }
}