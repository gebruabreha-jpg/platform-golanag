import { z } from "zod";

/* ── Core ───────────────────────────────────────────────────────────────────── */

export interface PaginationMeta {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

export interface ApiResponse<T = unknown> {
  success: boolean;
  data?: T;
  error?: {
    code: string;
    message: string;
    details?: string[];
  };
  meta?: {
    timestamp: string;
    path: string;
    [key: string]: unknown;
  };
}

/* ── Auth ───────────────────────────────────────────────────────────────────── */

export interface JWTPayload {
  sub: string;
  role: string;
  type: "access" | "refresh";
  iat?: number;
  exp?: number;
}

export interface LoginResponse {
  accessToken:  string;
  refreshToken: string;
  expiresIn:    string;
}

export interface RegisterRequest {
  email:    string;
  password: string;
  firstName: string;
  lastName:  string;
  phone?:    string | null;
  role:      "DIASPORA" | "LOCAL" | "MERCHANT" | "ADMIN";
  country?:  string | null;
  city?:     string | null;
}

export interface RefreshTokenRequest {
  refreshToken: string;
}

/* ── User ───────────────────────────────────────────────────────────────────── */

export type UserRole = "DIASPORA" | "LOCAL" | "MERCHANT" | "ADMIN";

export type VerificationLevel = 0 | 1 | 2;

export interface User {
  id:                string;
  email:             string;
  phone?:            string | null;
  firstName:         string;
  lastName:          string;
  dateOfBirth?:      Date | null;
  bio?:              string | null;
  avatarUrl?:        string | null;
  location?:         string | null;
  country?:          string | null;
  city?:             string | null;
  role:              UserRole;
  isVerified:        boolean;
  verificationLevel: VerificationLevel;
  trustScore:        number;
  totalTransactions: number;
  passwordHash:      string;
  createdAt:         Date;
  updatedAt:         Date;
  deletedAt?:        Date | null;
}

export interface UserPublic {
  id:                string;
  email?:            string | null;
  firstName:         string;
  lastName:          string;
  location?:         string | null;
  country?:          string | null;
  city?:             string | null;
  role:              UserRole;
  isVerified:        boolean;
  verificationLevel: VerificationLevel;
  trustScore:        number;
  totalTransactions: number;
  avatarUrl?:        string | null;
  createdAt:         Date;
  updatedAt:         Date;
}

/* ── Community ──────────────────────────────────────────────────────────────── */

export type CommunityCategory =
  | "SHIPPING"
  | "HOUSING"
  | "MARKETPLACE"
  | "JOBS"
  | "SCHOLARSHIPS"
  | "BUSINESS";

export type PostType = "OFFER" | "REQUEST" | "INFO" | "DISCUSSION";

export interface Community {
  id:           string;
  name:         string;
  description?: string | null;
  category:     CommunityCategory;
  location?:    string | null;
  country?:     string | null;
  isPrivate:    boolean;
  memberCount:  number;
  moderatorId:  string;
  rules?:       string | null;
  imageUrl?:    string | null;
  tags?:        ReadonlyArray<string> | null;
  createdAt:    Date;
  updatedAt:    Date;
}

export interface Post {
  id:            string;
  communityId:   string;
  userId:        string;
  type:          PostType;
  title:         string;
  content:       string;
  mediaUrls?:    ReadonlyArray<string> | null;
  isPinned:      boolean;
  isClosed:      boolean;
  replyCount:    number;
  viewCount:     number;
  createdAt:     Date;
  updatedAt:     Date;
}

/* ── Shipping ──────────────────────────────────────────────────────────────── */

export type PaymentStatus = "PENDING" | "COMPLETED" | "CANCELLED" | "DISPUTED";

export interface Transaction {
  id:            string;
  buyerId:       string;
  sellerId:      string;
  itemId?:       string | null;
  type:          string;
  amount:        number;
  currency:      string;
  status:        PaymentStatus;
  paymentMethod: string;
  escrowId?:     string | null;
  releasedAt?:   Date | null;
  completedAt?:  Date | null;
  createdAt:     Date;
  updatedAt:     Date;
}

/* ── Housing ───────────────────────────────────────────────────────────────── */

export type PropertyType = "APARTMENT" | "HOUSE" | "ROOM" | "STUDIO";
export type RoomType     = "PRIVATE" | "SHARED" | "MASTER";
export type LeaseTerm    = "MONTHLY" | "SHORT_TERM" | "LONG_TERM";

export interface HousingListing {
  id:               string;
  landlordId:       string;
  title:            string;
  description?:     string | null;
  propertyType:     PropertyType;
  roomType?:        RoomType | null;
  bedrooms:         number;
  bathrooms:        number;
  monthlyRent:      number;
  deposit?:         number;
  address:          string;
  city:             string;
  country:          string;
  latitude:         number;
  longitude:        number;
  availableFrom?:   Date | null;
  leaseTerm:        LeaseTerm;
  furnished:        boolean;
  includesUtilities: boolean;
  imageUrls?:       ReadonlyArray<string> | null;
  isActive:         boolean;
  applicationCount: number;
  createdAt:        Date;
  updatedAt:        Date;
}

export interface HousingApplication {
  id:            string;
  listingId:     string;
  userId:        string;
  message?:      string | null;
  moveInDate?:   Date | null;
  proposedRent?: number | null;
  status:        string;
  createdAt:     Date;
  updatedAt:     Date;
}

/* ── Marketplace ───────────────────────────────────────────────────────────── */

export type ItemCategory   = string;
export type ItemCondition  = "NEW" | "LIKE_NEW" | "GOOD" | "FAIR";

export interface MarketplaceItem {
  id:                  string;
  sellerId:            string;
  title:               string;
  description?:        string | null;
  category:            ItemCategory;
  subcategory?:        string | null;
  price:               number;
  currency:            string;
  condition:           ItemCondition;
  location?:           string | null;
  country?:            string | null;
  shippingAvailable:   boolean;
  shippingCost:        number;
  imageUrls?:          ReadonlyArray<string> | null;
  isActive:            boolean;
  isSold:              boolean;
  viewCount:           number;
  interestCount:       number;
  createdAt:           Date;
  updatedAt:           Date;
}

/* ── Services ──────────────────────────────────────────────────────────────── */

export type ServiceCategory =
  | "LEGAL"
  | "MEDICAL"
  | "CONSULTING"
  | "TECHNICAL"
  | "CREATIVE"
  | "TRADES"
  | "TRANSPORT"
  | "OTHER";

export type LawyerSpecialization =
  | "IMMIGRATION"
  | "BUSINESS"
  | "FAMILY"
  | "PROPERTY"
  | "CRIMINAL";

export interface ServiceProfessional {
  id:                 string;
  userId:             string;
  specialization:     string;
  name:               string;
  firm?:              string | null;
  yearsExperience:    number;
  location:           string;
  country:            string;
  consultationFee:    number;
  currency:           string;
  isVerified:         boolean;
  rating:             number;
  reviewCount:        number;
  createdAt:          Date;
  updatedAt:          Date;
}

export interface LawyerBooking {
  id:           string;
  lawyerId:     string;
  userId:       string;
  scheduledAt:  Date;
  duration:     number;
  type:         "IN_PERSON" | "REMOTE";
  status:       string;
  notes?:       string | null;
  createdAt:    Date;
  updatedAt:    Date;
}

/* ── Scholarship ───────────────────────────────────────────────────────────── */

export interface Scholarship {
  id:             string;
  title:          string;
  description:    string;
  provider:       string;
  providerType:   "GOVERNMENT" | "UNIVERSITY" | "FOUNDATION";
  country:        string;
  city?:          string | null;
  level:          string;
  field:          string;
  amount?:        number | null;
  currency?:      string | null;
  isActive:       boolean;
  isFeatured:     boolean;
  deadline?:      Date | null;
  createdAt:      Date;
  updatedAt:      Date;
}

/* ── Job ───────────────────────────────────────────────────────────────────── */

export type JobType = "FULL_TIME" | "PART_TIME" | "CONTRACT" | "FREELANCE";

export interface Job {
  id:               string;
  employerId:       string;
  title:            string;
  description:      string;
  jobType:          JobType;
  remote:           boolean;
  location?:        string | null;
  country?:         string | null;
  salaryMin?:       number | null;
  salaryMax?:       number | null;
  currency?:        string | null;
  industry?:        string | null;
  skills?:          ReadonlyArray<string> | null;
  benefits?:        ReadonlyArray<string> | null;
  applicationUrl?:  string | null;
  expiresAt?:       Date | null;
  isActive:         boolean;
  viewCount:        number;
  applicationCount: number;
  createdAt:         Date;
  updatedAt:         Date;
}

/* ── Review ────────────────────────────────────────────────────────────────── */

export type SubjectType = "USER" | "LISTING" | "COMMUNITY";

export interface TrustReview {
  id:           string;
  reviewerId:   string;
  subjectId:    string;
  subjectType:  SubjectType;
  rating:       number;
  comment?:     string | null;
  isVerified:   boolean;
  createdAt:    Date;
}

/* ── Paginated ──────────────────────────────────────────────────────────────── */

export interface PaginatedResult<T> {
  items: T[];
  meta:  PaginationMeta;
}
