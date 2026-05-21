package handler

import (
	"net/http"
	"strconv"

	"connectme/internal/service"

	"github.com/gin-gonic/gin"
)

type HousingHandler struct {
	housingService *service.HousingService
}

func NewHousingHandler(housingService *service.HousingService) *HousingHandler {
	return &HousingHandler{housingService: housingService}
}

// ListListings handles GET /api/v1/housing/listings
func (h *HousingHandler) ListListings(c *gin.Context) {
	city := c.Query("city")
	country := c.Query("country")
	propertyType := c.Query("property_type")
	roomType := c.Query("room_type")
	minRent, _ := strconv.ParseFloat(c.Query("min_rent"), 64)
	maxRent, _ := strconv.ParseFloat(c.Query("max_rent"), 64)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	listings, total, err := h.housingService.ListListings(service.ListListingsFilter{
		City:         city,
		Country:      country,
		PropertyType: propertyType,
		RoomType:     roomType,
		MinRent:      minRent,
		MaxRent:      maxRent,
		Limit:        limit,
		Offset:       offset,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"total":    total,
		"limit":    limit,
		"offset":   offset,
	})
}

// GetListing handles GET /api/v1/housing/listings/:id
func (h *HousingHandler) GetListing(c *gin.Context) {
	id := c.Param("id")
	listing, err := h.housingService.GetListing(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listing": listing})
}

// CreateListing handles POST /api/v1/housing/listings
func (h *HousingHandler) CreateListing(c *gin.Context) {
	userID := c.GetString("userID")
	var req service.CreateListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.LandlordID = userID

	listing, err := h.housingService.CreateListing(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Listing created",
		"listing":  listing,
	})
}

// SubmitApplication handles POST /api/v1/housing/applications
func (h *HousingHandler) SubmitApplication(c *gin.Context) {
	userID := c.GetString("userID")
	var req struct {
		ListingID string `json:"listing_id"`
		Message   string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.housingService.SubmitApplication(req.ListingID, userID, req.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted"})
}