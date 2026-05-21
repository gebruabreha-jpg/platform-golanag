import { Router } from "express";
import { MarketplaceController } from "@/controllers/marketplace.controller.js";
import { MarketplaceService } from "@/services/marketplace.service.js";
import { authenticate } from "@/middleware/auth.js";
import { verifyAccessToken } from "@/utils/crypto.js";

const router = Router();
const marketplaceService = new MarketplaceService();
const marketplaceController = new MarketplaceController(marketplaceService);

router.post("/", authenticate(verifyAccessToken, verifyAccessToken), marketplaceController.createItem);
router.get("/", marketplaceController.getItems);
router.get("/:id", marketplaceController.getItem);
router.patch("/:id", authenticate(verifyAccessToken, verifyAccessToken), marketplaceController.updateItem);
router.delete("/:id", authenticate(verifyAccessToken, verifyAccessToken), marketplaceController.deleteItem);
router.post("/:id/sold", authenticate(verifyAccessToken, verifyAccessToken), marketplaceController.markItemSold);

export default router;