import { Router } from "express";
import { HousingController } from "@/controllers/housing.controller.js";
import { HousingService } from "@/services/housing.service.js";
import { authenticate } from "@/middleware/auth.js";
import { verifyAccessToken } from "@/utils/crypto.js";

const router = Router();
const housingService = new HousingService();
const housingController = new HousingController({ housingService });

router.post("/", authenticate(verifyAccessToken, verifyAccessToken), housingController.createListing);
router.get("/", housingController.getListings);
router.get("/:id", housingController.getListing);
router.patch("/:id", authenticate(verifyAccessToken, verifyAccessToken), housingController.updateListing);
router.delete("/:id", authenticate(verifyAccessToken, verifyAccessToken), housingController.deleteListing);
router.post("/:id/applications", authenticate(verifyAccessToken, verifyAccessToken), housingController.createApplication);
router.get("/applications/:id", authenticate(verifyAccessToken, verifyAccessToken), housingController.getApplication);
router.patch("/applications/:id", authenticate(verifyAccessToken, verifyAccessToken), housingController.updateApplication);

export default router;