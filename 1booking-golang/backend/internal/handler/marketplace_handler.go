package handler

import (
	"connectme/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MarketplaceHandler struct {
	marketplaceService *service.MarketplaceService
}

func NewMarketplaceHandler(marketplaceService *service.MarketplaceService) *MarketplaceHandler {
	return &MarketplaceHandler{marketplaceService: marketplaceService}
}

// ListItems handles GET /api/v1/marketplace/items
func (h *MarketplaceHandler) ListItems(c *gin.Context) {
	category := c.Query("category")
	country := c.Query("country")
	minPrice, _ := strconv.ParseFloat(c.Query("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(c.Query("max_price"), 64)
	condition := c.Query("condition")
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "newest")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	items, total, err := h.marketplaceService.ListItems(service.ListItemsFilter{
		Category:  category,
		Country:   country,
		MinPrice:  minPrice,
		MaxPrice:  maxPrice,
		Condition: condition,
		Search:    search,
		SortBy:    sortBy,
		Limit:     limit,
		Offset:    offset,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
		"limit": limit,
		"offset": offset,
	})
}

// GetItem handles GET /api/v1/marketplace/items/:id
func (h *MarketplaceHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := h.marketplaceService.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// CreateItem handles POST /api/v1/marketplace/items
func (h *MarketplaceHandler) CreateItem(c *gin.Context) {
	userID := c.GetString("userID")
	var req service.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.SellerID = userID

	item, err := h.marketplaceService.CreateItem(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item created",
		"item":    item,
	})
}

// UpdateItem handles PUT /api/v1/marketplace/items/:id
func (h *MarketplaceHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	// Get existing item to verify ownership
	item, err := h.marketplaceService.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if item.SellerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to edit this item"})
		return
	}

	// Update item (implement update logic)
	c.JSON(http.StatusOK, gin.H{"message": "Item updated"})
}

// InitiateSecurePayment handles POST /api/v1/marketplace/transactions/:id/secure-payment
func (h *MarketplaceHandler) InitiateSecurePayment(c *gin.Context) {
	txID := c.Param("id")
	userID := c.GetString("userID")
	// TODO: load transaction, verify ownership, create Stripe intent / escrow hold
	c.JSON(http.StatusOK, gin.H{
		"message":   "Secure payment initiated",
		"tx_id":     txID,
		"user_id":   userID,
	})
}

func (h *MarketplaceHandler) ExpressInterest(c *gin.Context) {
	itemID := c.Param("id")
	userID := c.GetString("userID")

	var req service.ExpressInterestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ItemID = itemID
	req.BuyerID = userID

	if err := h.marketplaceService.ExpressInterest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Interest expressed. Seller will be notified."})
}
