package handler

import (
	"connectme/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommunityHandler struct {
	communityService *service.CommunityService
}

func NewCommunityHandler(communityService *service.CommunityService) *CommunityHandler {
	return &CommunityHandler{communityService: communityService}
}

// ListCommunities handles GET /api/v1/communities
func (h *CommunityHandler) ListCommunities(c *gin.Context) {
	category := c.Query("category")
	location := c.Query("location")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	communities, total, err := h.communityService.ListCommunities(service.ListCommunitiesFilter{
		Category: category,
		Location: location,
		Limit:    limit,
		Offset:   offset,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"communities": communities,
		"total":       total,
		"limit":       limit,
		"offset":      offset,
	})
}

// GetCommunity handles GET /api/v1/communities/:id
func (h *CommunityHandler) GetCommunity(c *gin.Context) {
	id := c.Param("id")
	community, err := h.communityService.GetCommunity(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Community not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"community": community})
}

// CreateCommunity handles POST /api/v1/communities
func (h *CommunityHandler) CreateCommunity(c *gin.Context) {
	var req service.CreateCommunityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set moderator ID from authenticated user
	req.ModeratorID = c.GetString("userID")

	community, err := h.communityService.CreateCommunity(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Community created",
		"community":  community,
	})
}

// JoinCommunity handles POST /api/v1/communities/:id/join
func (h *CommunityHandler) JoinCommunity(c *gin.Context) {
	communityID := c.Param("id")
	userID := c.GetString("userID")

	if err := h.communityService.JoinCommunity(communityID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined community"})
}

// LeaveCommunity handles POST /api/v1/communities/:id/leave
func (h *CommunityHandler) LeaveCommunity(c *gin.Context) {
	communityID := c.Param("id")
	userID := c.GetString("userID")

	if err := h.communityService.LeaveCommunity(communityID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully left community"})
}

// GetPosts handles GET /api/v1/communities/:id/posts
func (h *CommunityHandler) GetPosts(c *gin.Context) {
	communityID := c.Param("id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	posts, err := h.communityService.GetPosts(communityID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// CreatePost handles POST /api/v1/communities/:id/posts
func (h *CommunityHandler) CreatePost(c *gin.Context) {
	communityID := c.Param("id")
	userID := c.GetString("userID")

	var req service.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.CommunityID = communityID
	req.UserID = userID

	post, err := h.communityService.CreatePost(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created",
		"post":    post,
	})
}
