import type { UserPublic } from "@/types/index.js";

/* ── Auth service interface ─────────────────────────────────────────────────── */
export interface IAuthService {
  register(data: RegisterInput): Promise<UserPublic>;
  login(email: string, password: string): Promise<LoginResult>;
  logout(userId: string, refreshToken: string): Promise<void>;
  validateToken(token: string): Promise<UserPublic | null>;
  changePassword(userId: string, currentPassword: string, newPassword: string): Promise<void>;

  /* helpers */
  refresh(refreshToken: string): Promise<RefreshResult>;
  revokeRefreshToken(tokenHash: string): Promise<void>;
}

export type RegisterInput = {
  email:       string;
  password:    string;
  firstName:   string;
  lastName:    string;
  phone?:      string | null;
  role?:       "DIASPORA" | "LOCAL" | "MERCHANT" | "ADMIN";
  country?:    string | null;
  city?:       string | null;
};

/* ── User service interface ─────────────────────────────────────────────────── */
export interface IUserService {
  getProfile(id: string): Promise<UserPublic | null>;
  getUserByEmail(email: string): Promise<UserPublic | null>;
  updateProfile(id: string, data: UpdateProfileInput): Promise<UserPublic | null>;
  getAllUsers(query: { page: number; limit: number; role?: string }): Promise<UserPublic[]>;
  setUserRole(id: string, role: string): Promise<UserPublic | null>;
  deactivateUser(id: string): Promise<void>;
}

export type UpdateProfileInput = {
  firstName?: string;
  lastName?:  string;
  phone?:     string | null;
  bio?:       string | null;
  avatarUrl?: string | null;
};

/* ── Community service interface ────────────────────────────────────────────── */
export interface ICommunityService {
  create(data: CreateCommunityInput): Promise<unknown>;
  getById(id: string): Promise<unknown | null>;
  getCommunities(query: SearchCommunityInput): Promise<PaginatedResult<unknown>>;
  addPost(communityId: string, data: CreatePostInput): Promise<unknown>;
  addMember(communityId: string, userId: string): Promise<void>;
  removeMember(communityId: string, userId: string): Promise<void>;
}

export interface CreateCommunityInput {
  name: string;
  description?: string | null;
  category: "SHIPPING" | "HOUSING" | "MARKETPLACE" | "JOBS" | "SCHOLARSHIPS" | "BUSINESS";
  country?:     string | null;
  isPrivate?:   boolean;
  imageUrl?:    string | null;
  rules?:       string | null;
  tags?: string[];
}

export interface SearchCommunityInput {
  category?:  string;
  isPrivate?: boolean;
  country?:   string;
  search?:    string;
  page:       number;
  limit:      number;
}
export interface PaginatedResult<T> { items: T[]; meta: { page: number; limit: number; total: number; totalPages: number }; }

export interface CreatePostInput {
  communityId: string;
  type:        "OFFER" | "REQUEST" | "INFO" | "DISCUSSION";
  title:       string;
  content:     string;
  mediaUrls?:  string[];
  isPinned?:   boolean;
}

/* ── Housing service interface ──────────────────────────────────────────────── */
export interface IHousingService {
  createListing(data: CreateListingInput): Promise<unknown>;
  updateListing(id: string, data: Partial<CreateListingInput>): Promise<unknown>;
  getListing(id: string): Promise<unknown | null>;
  getListings(query: ListingSearchInput): Promise<PaginatedResult<unknown>>;
  createApplication(listingId: string, userId: string, data: CreateApplicationInput): Promise<unknown>;
  getApplications(listingId: string): Promise<unknown[]>;
  approveApplication(applicationId: string): Promise<void>;
  rejectApplication(applicationId: string): Promise<void>;
}

export type CreateListingInput = {
  title: string;
  description?: string | null;
  address: string;
  city: string;
  country: string;
  latitude: number;
  longitude: number;
  monthlyRent: number;
  deposit?: number;
  bedrooms: number;
  bathrooms: number;
  propertyType: "APARTMENT" | "HOUSE" | "ROOM" | "STUDIO";
  roomType?: "PRIVATE" | "SHARED" | "MASTER" | null;
  leaseTerm?: "MONTHLY" | "SHORT_TERM" | "LONG_TERM";
  furnished?: boolean;
  includesUtilities?: boolean;
  availableFrom?: Date | null;
  imageUrls?: string[];
};

export type ListingSearchInput = {
  city?: string;
  propertyType?: "APARTMENT" | "HOUSE" | "ROOM" | "STUDIO";
  minRent?: number;
  maxRent?: number;
  bedrooms?: number;
  page:   number;
  limit:  number;
};

export type CreateApplicationInput = {
  message?:     string | null;
  moveInDate?:  Date | null;
  proposedRent?: number | null;
};

/* ── Marketplace service interface ─────────────────────────────────────────── */
export interface IMarketplaceService {
  createItem(data: CreateMarketplaceItemInput, sellerId: string): Promise<unknown>;
  updateItem(id: string, data: Partial<CreateMarketplaceItemInput>): Promise<unknown>;
  getItemById(id: string): Promise<unknown | null>;
  searchItems(query: MarketPlaceSearchInput): Promise<PaginatedResult<unknown>>;
  markItemSold(itemId: string): Promise<void>;
  createTransaction(data: CreateTransactionInput): Promise<unknown>;
  getTransaction(id: string): Promise<unknown | null>;
  getSellerTransactions(sellerId: string): Promise<unknown[]>;
}

export type CreateMarketplaceItemInput = {
  title: string;
  description?: string | null;
  category: string;
  subcategory?: string | null;
  price: number;
  currency?: string;
  condition: "NEW" | "LIKE_NEW" | "GOOD" | "FAIR";
  location?: string | null;
  country?:   string | null;
  shippingAvailable?: boolean;
  shippingCost?: number;
  imageUrls?: string[];
};

export type MarketPlaceSearchInput = {
  category?:  string;
  condition?: "NEW" | "LIKE_NEW" | "GOOD" | "FAIR";
  minPrice?:  number;
  maxPrice?:  number;
  country?:   string;
  search?:    string;
  page:       number;
  limit:      number;
};

export type CreateTransactionInput = {
  itemId:       string;
  buyerId:      string;
  sellerId:     string;
  amount:       number;
  currency?:    string;
  type:         string;
  paymentMethod: string;
};

/* ── Payment service interface ──────────────────────────────────────────────── */
export interface IPaymentService {
  createEscrow(data: CreateEscrowInput): Promise<TransactionRecord>;
  releaseEscrow(transactionId: string): Promise<void>;
  cancelEscrow(transactionId: string, reason?: string): Promise<void>;
  getTransaction(id: string): Promise<TransactionRecord | null>;
  getUserTransactions(userId: string): Promise<TransactionRecord[]>;
}

export interface CreateEscrowInput {
  itemId:       string;
  buyerId:      string;
  sellerId:     string;
  amount:       number;
  currency?:    string;
  description?: string | null;
}

export interface TransactionRecord {
  id:            string;
  buyerId:       string;
  sellerId:      string;
  itemId?:       string | null;
  type:          string;
  amount:        number;
  currency:      string;
  status:        "PENDING" | "COMPLETED" | "CANCELLED" | "DISPUTED";
  paymentMethod: string;
  escrowId?:     string | null;
  releasedAt?:   Date | null;
  completedAt?:  Date | null;
  createdAt:     Date;
  updatedAt:     Date;
}

/* ── External / AI service interface ───────────────────────────────────────── */
export interface IExternalService {
  callExternalAPI<T>(url: string, method?: "GET" | "POST" | "PUT" | "DELETE", body?: unknown, headers?: Record<string, string>): Promise<T>;
  callWithRetry<T>(url: string, method?: "GET", body?: unknown): Promise<T>;
}
