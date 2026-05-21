import { z } from "zod";

export const createListingSchema = z.object({
  title:              z.string().min(3).max(120),
  description:        z.string().max(2000).optional().nullable(),
  address:            z.string().min(5, "Address is required."),
  city:               z.string().min(1).max(80),
  country:            z.string().min(2).max(80),
  latitude:           z.number().min(-90).max(90),
  longitude:          z.number().min(-180).max(180),
  monthlyRent:        z.number().positive("Rent must be positive."),
  deposit:            z.number().min(0).optional().default(0),
  bedrooms:           z.number().int().min(0).default(0),
  bathrooms:          z.number().int().min(0).default(0),
  propertyType:       z.enum(["APARTMENT", "HOUSE", "ROOM", "STUDIO"]).default("APARTMENT"),
  roomType:           z.enum(["PRIVATE", "SHARED", "MASTER"]).optional(),
  leaseTerm:          z.enum(["MONTHLY", "SHORT_TERM", "LONG_TERM"]).default("MONTHLY"),
  furnished:          z.boolean().optional().default(false),
  includesUtilities:  z.boolean().optional().default(false),
  availableFrom:      z.string().datetime().transform((s) => new Date(s)),
  imageUrls:          z.array(z.string().url()).max(20).optional(),
});

export const updateListingSchema = createListingSchema.partial();

export const listingSearchSchema = z.object({
  city:         z.string().max(80).optional(),
  propertyType: z.enum(["APARTMENT", "HOUSE", "ROOM", "STUDIO"]).optional(),
  minRent:      z.coerce.number().min(0).optional(),
  maxRent:      z.coerce.number().min(0).optional(),
  bedrooms:     z.coerce.number().int().min(0).optional(),
  page:         z.coerce.number().int().min(1).default(1),
  limit:        z.coerce.number().int().min(1).max(100).default(20),
});

export const createApplicationSchema = z.object({
  listingId:   z.string().uuid("A valid listing ID is required."),
  message:     z.string().max(1000).optional().nullable(),
  moveInDate:  z.string().datetime().optional().nullable().transform((v) => (v ? new Date(v) : null)),
  proposedRent: z.number().positive().optional(),
});

export type CreateListingInput       = z.infer<typeof createListingSchema>;
export type UpdateListingInput       = z.infer<typeof updateListingSchema>;
export type ListingSearchInput       = z.infer<typeof listingSearchSchema>;
export type CreateApplicationInput   = z.infer<typeof createApplicationSchema>;
