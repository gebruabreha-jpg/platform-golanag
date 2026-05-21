import type { User } from "@/types/index.js";

/**
 * Crud operations shared by every repository.
 */
export interface IBaseRepository<T> {
  findById(id: string):         Promise<T | null>;
  findByIds(ids: string[]):     Promise<T[]>;
  findAll(filter?: Record<string, unknown>):  Promise<T[]>;
  create(data: Partial<T>):     Promise<T>;
  update(id: string, data: Partial<T>): Promise<T | null>;
  delete(id: string):           Promise<boolean>;
  count(filter?: Record<string, unknown>):   Promise<number>;
  existsById(id: string):       Promise<boolean>;
}

/**
 * User-specific repository operations.
 */
export interface IUserRepository extends IBaseRepository<User> {
  findByEmail(email: string):                              Promise<User | null>;
  findByVerificationToken(token: string):                  Promise<User | null>;
  findByRefreshTokenHash(tokenHash: string):               Promise<User | null>;
  updatePassword(id: string, passwordHash: string):        Promise<void>;
  updateTrustScore(id: string, score: number):              Promise<void>;
  updateVerificationLevel(id: string, level: number):       Promise<void>;
  updateIsVerified(id: string, isVerified: boolean):        Promise<void>;
  findActiveUsers(limit: number, offset: number):           Promise<User[]>;
}

// Re-export every interface from a single entry point
export * from "@/repositories/interfaces/ICommunityRepository.js";
