import { Router } from "express";
import { CommunityController } from "@/controllers/community.controller.js";
import { CommunityService } from "@/services/community.service.js";
import { authenticate } from "@/middleware/auth.js";
import { verifyAccessToken } from "@/utils/crypto.js";

const router = Router();
const communityService = new CommunityService();
const communityController = new CommunityController(communityService);

router.use(authenticate(verifyAccessToken, verifyAccessToken));
router.post("/", communityController.createCommunity);
router.get("/", communityController.getCommunities);
router.get("/:id", communityController.getCommunity);
router.post("/:id/posts", communityController.addPost);
router.get("/:id/posts", communityController.getPosts);
router.post("/:id/join", communityController.joinCommunity);
router.post("/:id/leave", communityController.leaveCommunity);
router.patch("/:id", communityController.updateCommunity);
router.delete("/:id", communityController.deleteCommunity);

export default router;