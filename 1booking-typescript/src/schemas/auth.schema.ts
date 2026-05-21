import { z } from "zod";

export const registerSchema = z.object({
  email:    z.string().email("A valid email is required."),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long.")
    .regex(/[A-Z]/, "Password must contain at least one uppercase letter.")
    .regex(/[a-z]/, "Password must contain at least one lower-case letter.")
    .regex(/[0-9]/, "Password must contain at least one digit."),
  firstName: z.string().min(1, "First name is required."                ).max(80),
  lastName:  z.string().min(1, "Last name is required."                 ).max(80),
  phone:     z.string().regex(/^\+?[1-9]\d{1,14}$/).optional().nullable(),
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

export type RegisterUserInput  = z.infer<typeof registerSchema>;
export type LoginInput         = z.infer<typeof loginSchema>;
export type RefreshTokenInput  = z.infer<typeof refreshTokenSchema>;
export type ChangePasswordInput= z.infer<typeof changePasswordSchema>;
