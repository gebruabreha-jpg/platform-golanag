import { z } from "zod";

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

export type CreateItemInput = z.infer<typeof createItemSchema>;
export type UpdateItemInput = z.infer<typeof updateItemSchema>;
export type ItemSearchInput = z.infer<typeof itemSearchSchema>;
