import { z } from "zod";

/* ── Auth ───────────────────────────────────────────────────────────────────── */

export const registerSchema = z.object({
  email:    z.string().email("A valid email is required."),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long.")
    .regex(/[A-Z]/,              "Password must contain at least one uppercase letter.")
    .regex(/[a-z]/,              "Password must contain at least one lower-case letter.")
    .regex(/[0-9]/,              "Password must contain at least one digit."),
  firstName: z.string().min(1, "First name is required."               ).max(80),
  lastName:  z.string().min(1, "Last name is required."                ).max(80),
  phone:     z.string().regex(/^\+?[1-9]\d{1,14}$/, "Invalid phone number.").optional().nullable(),
  role:      z.enum(["DIASPORA", "LOCAL", "MERCHANT", "ADMIN"]).optional(),
  country:   z.string().max(80).optional().nullable(),
  city:      z.string().max(80).optional().nullable(),
});

export const loginSchema = z.object({
  email:    z.string().email("A valid email is required."),
  password: z.string().min(1, "Password is required."),
});

export const refreshTokenSchema = z.object({
  refreshToken: z.string().min(1, "Refresh token is required."),
});

export const changePasswordSchema = z.object({
  currentPassword: z.string().min(1, "Current password is required."),
  newPassword:     z.string()
    .min(8, "New password must be at least 8 characters.")
    .regex(/[A-Z]/, "New password must contain at least one uppercase letter.")
    .regex(/[a-z]/, "New password must contain at least one lower-case letter.")
    .regex(/[0-9]/, "New password must contain at least one digit."),
});

/* ── User ───────────────────────────────────────────────────────────────────── */

export const updateProfileSchema = z.object({
  firstName:    z.string().min(1).max(80).optional(),
  lastName:     z.string().min(1).max(80).optional(),
  phone:        z.string().regex(/^\+?[1-9]\d{1,14}$/).optional().nullable(),
  bio:          z.string().max(500).optional().nullable(),
  avatarUrl:    z.string().url().optional().nullable(),
  location:     z.string().max(120).optional().nullable(),
  country:      z.string().max(80).optional().nullable(),
  city:         z.string().max(80).optional().nullable(),
  trustScore:   z.number().min(0).max(10000).optional(),
});

export const updateUserSchema = updateProfileSchema.extend({
  role:                 z.enum(["DIASPORA", "LOCAL", "MERCHANT", "ADMIN"]).optional(),
  isVerified:           z.boolean().optional(),
  verificationLevel:    z.number().int().min(0).max(2).optional(),
  isActive:             z.boolean().optional(),
});

/* ── Pagination / list ──────────────────────────────────────────────────────── */

export const paginationSchema = z.object({
  page:       z.coerce.number().int().min(1).default(1),
  limit:      z.coerce.number().int().min(1).max(100).default(20),
  sortBy:     z.string().max(64).optional(),
  sortOrder:  z.enum(["asc", "desc"]).optional().default("desc"),
});

export const communityListSchema = paginationSchema.extend({
  category: z.enum(["SHIPPING", "HOUSING", "MARKETPLACE", "JOBS", "SCHOLARSHIPS", "BUSINESS"]).optional(),
  isPrivate: z.coerce.boolean().optional(),
  country:  z.string().max(80).optional(),
});

/* ── Community ──────────────────────────────────────────────────────────────── */

export const createCommunitySchema = z.object({
  name:        z.string().min(3, "Community name must be at least 3 characters.").max(100),
  description: z.string().max(1000).optional().nullable(),
  category:    z.enum(["SHIPPING", "HOUSING", "MARKETPLACE", "JOBS", "SCHOLARSHIPS", "BUSINESS"]),
  country:     z.string().max(80).optional().nullable(),
  isPrivate:   z.boolean().optional().default(false),
  imageUrl:    z.string().url().optional().nullable(),
  rules:       z.string().max(2000).optional().nullable(),
  tags:        z.array(z.string().max(40)).max(10).optional(),
});

/* ── Post ───────────────────────────────────────────────────────────────────── */

export const createPostSchema = z.object({
  communityId: z.string().uuid("communityId must be a valid UUID."),
  type:        z.enum(["OFFER", "REQUEST", "INFO", "DISCUSSION"]).default("DISCUSSION"),
  title:       z.string().min(5, "Title must be at least 5 characters.").max(120),
  content:     z.string().min(10, "Content must be at least 10 characters."),
  mediaUrls:   z.array(z.string().url()).max(10).optional(),
  isPinned:    z.boolean().optional().default(false),
});

/* ── Housing ────────────────────────────────────────────────────────────────── */

export const createListingSchema = z.object({
  title:        z.string().min(3).max(120),
  description:  z.string().max(2000).optional().nullable(),
  address:      z.string().min(5, "Address is required."),
  city:         z.string().min(1).max(80),
  country:      z.string().min(2).max(80),
  latitude:     z.number().min(-90).max(90),
  longitude:    z.number().min(-180).max(180),
  monthlyRent:  z.number().positive("Rent must be positive."),
  deposit:      z.number().min(0).optional().default(0),
  bedrooms:     z.number().int().min(0).default(0),
  bathrooms:    z.number().int().min(0).default(0),
  propertyType: z.enum(["APARTMENT", "HOUSE", "ROOM", "STUDIO"]).default("APARTMENT"),
  roomType:     z.enum(["PRIVATE", "SHARED", "MASTER"]).optional(),
  leaseTerm:    z.enum(["MONTHLY", "SHORT_TERM", "LONG_TERM"]).default("MONTHLY"),
  furnished:    z.boolean().optional().default(false),
  includesUtilities: z.boolean().optional().default(false),
  availableFrom: z.string().datetime().transform((s) => new Date(s)),
  imageUrls:    z.array(z.string().url()).max(20).optional(),
});

export const updateListingSchema = createListingSchema.partial();

export const createListingAppSchema = z.object({
  listingId:   z.string().uuid(),
  message:     z.string().max(1000).optional().nullable(),
  moveInDate:  z.string().datetime().optional().nullable().transform((v) => (v ? new Date(v) : null)),
  proposedRent: z.number().positive().optional(),
});

/* ── Marketplace ────────────────────────────────────────────────────────────── */

export const createItemSchema = z.object({
  title:              z.string().min(3).max(120),
  description:        z.string().max(2000).optional().nullable(),
  category:           z.string().max(80),
  subcategory:        z.string().max(80).optional().nullable(),
  price:              z.number().positive("Price must be greater than 0."),
  currency:           z.string().length(3).default("USD"),
  condition:          z.enum(["NEW", "LIKE_NEW", "GOOD", "FAIR"]),
  location:           z.string().max(120).optional().nullable(),
  country:            z.string().max(80).optional().nullable(),
  shippingAvailable:  z.boolean().optional().default(false),
  shippingCost:       z.number().min(0).optional().default(0),
  imageUrls:          z.array(z.string().url()).max(20).optional(),
});

export const updateItemSchema = createItemSchema.partial();

export const itemSearchSchema = z.object({
  category:  z.string().max(80).optional(),
  condition: z.enum(["NEW", "LIKE_NEW", "GOOD", "FAIR"]).optional(),
  minPrice:  z.coerce.number().min(0).optional(),
  maxPrice:  z.coerce.number().min(0).optional(),
  country:   z.string().max(80).optional(),
  search:    z.string().max(120).optional(),
  page:      z.coerce.number().int().min(1).default(1),
  limit:     z.coerce.number().int().min(1).max(100).default(20),
});

/* ── Services ───────────────────────────────────────────────────────────────── */

export const createServiceSchema = z.object({
  name:             z.string().min(3).max(120),
  description:      z.string().max(2000).optional().nullable(),
  category:         z.string().max(60),
  subcategory:      z.string().max(80).optional().nullable(),
  hourlyRate:       z.number().positive().optional(),
  currency:         z.string().length(3).default("USD"),
  location:         z.string().max(120).optional().nullable(),
  availableRemotely: z.boolean().optional().default(false),
  imageUrl:         z.string().url().optional().nullable(),
  skills:           z.array(z.string().max(60)).max(20).optional(),
});

export const bookServiceSchema = z.object({
  serviceListingId: z.string().uuid("A valid service listing ID is required."),
  scheduledAt:      z.string().datetime("Invalid scheduled date-time.").transform((s) => new Date(s)),
  durationMin:      z.number().int().min(15).max(1440),
  notes:            z.string().max(2000).optional().nullable(),
});

/* ── Payment ────────────────────────────────────────────────────────────────── */

export const createEscrowSchema = z.object({
  itemId:       z.string().uuid("A valid item ID is required."),
  buyerId:      z.string().uuid("A valid buyer ID is required."),
  sellerId:     z.string().uuid("A valid seller ID is required."),
  amount:       z.number().positive("Escrow amount must be greater than 0."),
  currency:     z.string().length(3).default("USD"),
  description:  z.string().max(500).optional().nullable(),
});

export const releaseEscrowSchema = z.object({
  transactionId: z.string().uuid(),
});

export const refundSchema = z.object({
  transactionId:   z.string().uuid(),
  reason:          z.string().max(500).optional().nullable(),
});

/* ── Shipping ───────────────────────────────────────────────────────────────── */

export const createShippingRequestSchema = z.object({
  originCountry:    z.string().min(2).max(80),
  originCity:       z.string().min(1).max(80),
  destinationCountry: z.string().min(2).max(80),
  destinationCity:  z.string().min(1).max(80),
  itemDescription:  z.string().min(5).max(1000),
  weightKg:         z.number().positive().optional(),
  dimensionsCm: z.object({
    length: z.number().positive("Length must be > 0."),
    width:  z.number().positive("Width must be > 0."),
    height: z.number().positive("Height must be > 0."),
  }),
  offeredPrice:     z.number().positive("Offered price must be greater than 0."),
  currency:         z.string().length(3).default("USD"),
});

export const createShippingOfferSchema = z.object({
  requestId:        z.string().uuid(),
  offeredPrice:     z.number().positive("Offered price must be greater than 0."),
  availableDate:    z.string().datetime().transform((s) => new Date(s)),
  notes:            z.string().max(1000).optional().nullable(),
});

/* ── Scholarship ────────────────────────────────────────────────────────────── */

export const createScholarshipSchema = z.object({
  title:         z.string().min(5).max(200),
  description:   z.string().max(5000),
  provider:      z.string().min(2).max(120),
  providerType:  z.enum(["GOVERNMENT", "UNIVERSITY", "FOUNDATION"]),
  country:       z.string().max(80),
  city:          z.string().max(80),
  level:         z.enum(["UNDERGRADUATE", "GRADUATE", "PHD", "POSTDOC"]),
  field:         z.string().max(120).optional().nullable(),
  amount:        z.number().nonnegative().optional().default(0),
  currency:      z.string().length(3).default("USD"),
  covers:        z.array(z.string()).nonempty().max(20).optional().default([]),
  deadline:      z.string().datetime().optional().nullable().transform((v) => (v ? new Date(v) : null)),
  eligibility:   z.string().max(2000).optional().nullable(),
  requirements:  z.string().max(2000).optional().nullable(),
  applicationUrl: z.string().url().optional().nullable(),
  isFeatured:    z.boolean().optional().default(false),
});

/* ── Job ────────────────────────────────────────────────────────────────────── */

export const createJobSchema = z.object({
  title:            z.string().min(5).max(200),
  description:      z.string().max(5000),
  jobType:          z.enum(["FULL_TIME", "PART_TIME", "CONTRACT", "FREELANCE"]),
  remote:           z.boolean().optional().default(false),
  location:         z.string().max(120).optional().nullable(),
  country:          z.string().max(80).optional().nullable(),
  salaryMin:        z.number().min(0).optional().nullable(),
  salaryMax:        z.number().min(0).optional().nullable(),
  currency:         z.string().length(3).default("USD"),
  industry:         z.string().max(80).optional().nullable(),
  skills:           z.array(z.string()).max(20).optional().default([]),
  benefits:         z.array(z.string()).max(20).optional().default([]),
  applicationUrl:   z.string().url().optional().nullable(),
  expiresAt:        z.string().datetime().optional().nullable().transform((v) => (v ? new Date(v) : null)),
});

export const updateJobSchema = createJobSchema.partial();

/* ── Currency rate ─────────────────────────────────────────────────────────── */

export const currencyRateSchema = z.object({
  fromCurrency: z.string().length(3),
  toCurrency:   z.string().length(3),
  rate:         z.number().positive(),
});

/** Infer TypeScript types from zod schemas for reuse across the app. */
export type RegisterUserInput  = z.infer<typeof registerSchema>;
export type LoginInput         = z.infer<typeof loginSchema>;
export type RefreshTokenInput  = z.infer<typeof refreshTokenSchema>;
export type ChangePasswordInput= z.infer<typeof changePasswordSchema>;
export type UpdateProfileInput = z.infer<typeof updateProfileSchema>;
export type CreateCommunity    = z.infer<typeof createCommunitySchema>;
export type CreatePost         = z.infer<typeof createPostSchema>;
export type CreateListing      = z.infer<typeof createListingSchema>;
export type UpdateListing      = z.infer<typeof updateListingSchema>;
export type CreateListingApp   = z.infer<typeof createListingAppSchema>;
export type CreateItem         = z.infer<typeof createItemSchema>;
export type UpdateItem         = z.infer<typeof updateItemSchema>;
export type ItemSearch         = z.infer<typeof itemSearchSchema>;
export type CreateService      = z.infer<typeof createServiceSchema>;
export type BookService        = z.infer<typeof bookServiceSchema>;
export type CreateEscrow       = z.infer<typeof createEscrowSchema>;
export type ReleaseEscrowInput = z.infer<typeof releaseEscrowSchema>;
export type RefundInput        = z.infer<typeof refundSchema>;
export type CreateShippingRequest = z.infer<typeof createShippingRequestSchema>;
export type CreateShippingOffer   = z.infer<typeof createShippingOfferSchema>;
export type CreateScholarship  = z.infer<typeof createScholarshipSchema>;
export type CreateJob          = z.infer<typeof createJobSchema>;
