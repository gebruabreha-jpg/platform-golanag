import type { ServiceProfessional } from "@/types/index.js";

export interface IServiceRepository {
  findByCategory(category: string):           Promise<ServiceProfessional[]>;
  findBySpecialization(specialization: string): Promise<ServiceProfessional[]>;
  search(query: ServiceSearchInput):         Promise<PaginatedResult<ServiceProfessional>>;
  updateRating(id: string, rating: number):   Promise<void>;
}

export interface ServiceSearchInput {
  category?:         string;
  location?:         string;
  country?:          string;
  remoteOnly?:       boolean;
  page:              number;
  limit:             number;
}

export interface PaginatedResult<T> {
  items: T[];
  meta:  { page: number; limit: number; total: number; totalPages: number };
}
