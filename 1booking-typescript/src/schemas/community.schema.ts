import { z } from "zod";

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

export const communitySearchSchema = z.object({
  category:  z.enum(["SHIPPING", "HOUSING", "MARKETPLACE", "JOBS", "SCHOLARSHIPS", "BUSINESS"]).optional(),
  isPrivate: z.coerce.boolean().optional(),
  country:   z.string().max(80).optional(),
  search:    z.string().max(120).optional(),
  page:      z.coerce.number().int().min(1).default(1),
  limit:     z.coerce.number().int().min(1).max(100).default(20),
});

export type CreateCommunity    = z.infer<typeof createCommunitySchema>;
export type SearchCommunityInput = z.infer<typeof communitySearchSchema>;
