import { z } from "zod";

export const updateProfileSchema = z.object({
  firstName:  z.string().min(1).max(80).optional(),
  lastName:   z.string().min(1).max(80).optional(),
  phone:      z.string().regex(/^\+?[1-9]\d{1,14}$/).optional().nullable(),
  bio:        z.string().max(500).optional().nullable(),
  avatarUrl:  z.string().url().optional().nullable(),
  location:   z.string().max(120).optional().nullable(),
  country:    z.string().max(80).optional().nullable(),
  city:       z.string().max(80).optional().nullable(),
});

export const updateUserSchema = updateProfileSchema.extend({
  role:               z.enum(["DIASPORA", "LOCAL", "MERCHANT", "ADMIN"]).optional(),
  isVerified:         z.boolean().optional(),
  verificationLevel:  z.number().min(0).max(2).optional(),
  isActive:           z.boolean().optional(),
});

export type UpdateProfileInput = z.infer<typeof updateProfileSchema>;
