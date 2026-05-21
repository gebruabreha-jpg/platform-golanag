import type { HousingListing, HousingApplication } from "@/types/index.js";

export interface IHousingRepository {
  findByLocation(city: string, country?: string): Promise<HousingListing[]>;
  findByPriceRange(min: number, max: number):   Promise<HousingListing[]>;
  findWithApplications():                       Promise<HousingListing[]>;
  search(query: HousingSearchInput):            Promise<HousingListing[]>;
  createApplication(data: CreateHousingAppInput): Promise<HousingApplication>;
  getApplications(listingId: string):           Promise<HousingApplication[]>;
}

export interface HousingSearchInput {
  city?:         string;
  propertyType?: string;
  minRent?:      number;
  maxRent?:      number;
  bedrooms?:     number;
  limit:         number;
  offset:        number;
}

export interface CreateHousingAppInput {
  listingId:    string;
  userId:       string;
  message?:     string | null;
  moveInDate?:  Date | null;
  proposedRent?: number | null;
}
