import { z } from "zod";

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

export type CreateServiceInput = z.infer<typeof createServiceSchema>;
export type BookServiceInput   = z.infer<typeof bookServiceSchema>;
