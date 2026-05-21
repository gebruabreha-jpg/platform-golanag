import { Router } from "express";
import authRoutes from "./auth.routes.js";
import userRoutes from "./user.routes.js";
import communityRoutes from "./community.routes.js";
import housingRoutes from "./housing.routes.js";
import marketplaceRoutes from "./marketplace.routes.js";

const router = Router();

router.use("/auth", authRoutes);
router.use("/users", userRoutes);
router.use("/communities", communityRoutes);
router.use("/housing", housingRoutes);
router.use("/marketplace", marketplaceRoutes);

export default router;