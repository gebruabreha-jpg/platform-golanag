import { Router } from "express";
import { UserController } from "@/controllers/user.controller.js";
import { UserService } from "@/services/user.service.js";
import { authenticate } from "@/middleware/auth.js";
import { verifyAccessToken } from "@/utils/crypto.js";

const router = Router();
const userService = new UserService();
const userController = new UserController(userService);

router.use(authenticate(verifyAccessToken, verifyAccessToken));
router.get("/profile", userController.getProfile);
router.patch("/profile", userController.updateProfile);
router.get("/", userController.getAllUsers);
router.delete("/:id", userController.deleteUser);

export default router;