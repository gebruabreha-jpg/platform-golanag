import type { MarketplaceItem } from "@/types/index.js";

export interface IMarketplaceRepository {
  findByCategory(category: string):             Promise<MarketplaceItem[]>;
  findWithFilters(filters: MarketplaceFilterInput): Promise<MarketplaceItem[]>;
  search(query: MarketplaceSearchInput):        Promise<PaginatedResult<MarketplaceItem>>;
  incrementViewCount(id: string):               Promise<void>;
  incrementInterestCount(id: string):           Promise<void>;
}

export interface MarketplaceFilterInput {
  category?:         string;
  condition?:        string;
  minPrice?:         number;
  maxPrice?:         number;
  country?:          string;
  sellerId?:         string;
  isSold?:           boolean;
}

export interface MarketplaceSearchInput {
  query:            string;
  category?:        string;
  minPrice?:        number;
  maxPrice?:        number;
  page:             number;
  limit:            number;
}

import type { PaginatedResult } from "@/types/index.js";
