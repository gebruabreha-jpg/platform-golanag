import { Router } from "express";
import { AuthController } from "@/controllers/auth.controller.js";
import { authenticate } from "@/middleware/auth.js";
import { verifyAccessToken } from "@/utils/crypto.js";

const router = Router();

// Public routes
router.post("/register", AuthController.register);
router.post("/login", AuthController.login);
router.post("/refresh", AuthController.refresh);

// Protected routes
router.use(authenticate(verifyAccessToken, verifyAccessToken));
router.post("/logout", AuthController.logout);
router.post("/change-password", AuthController.changePassword);

export default router;