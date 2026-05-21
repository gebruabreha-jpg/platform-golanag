import { z } from "zod";

export const createPostSchema = z.object({
  communityId: z.string().uuid("communityId must be a valid UUID."),
  type:        z.enum(["OFFER", "REQUEST", "INFO", "DISCUSSION"]).default("DISCUSSION"),
  title:       z.string().min(5, "Title must be at least 5 characters.").max(120),
  content:     z.string().min(10, "Content must be at least 10 characters."),
  mediaUrls:   z.array(z.string().url()).max(10).optional(),
  isPinned:    z.boolean().optional().default(false),
});

export type CreatePost = z.infer<typeof createPostSchema>;
