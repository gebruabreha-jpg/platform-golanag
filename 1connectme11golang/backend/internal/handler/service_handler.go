package handler

import (
	"net/http"
	"strconv"
	"time"

	"connectme/internal/service"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	serviceService *service.ServiceService
}

func NewServiceHandler(serviceService *service.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService}
}

// ListLawyers handles GET /api/v1/services/lawyers
func (h *ServiceHandler) ListLawyers(c *gin.Context) {
	specialization := c.Query("specialization")
	country := c.Query("country")
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	lawyers, total, err := h.serviceService.ListLawyers(service.ListLawyersFilter{
		Specialization: specialization,
		Country:        country,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lawyers": lawyers,
		"total":   total,
		"limit":   limit,
		"offset":  offset,
	})
}

// BookLawyer handles POST /api/v1/services/lawyers/:id/book
func (h *ServiceHandler) BookLawyer(c *gin.Context) {
	lawyerID := c.Param("id")
	userID := c.GetString("userID")

	var req service.BookLawyerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.LawyerID = lawyerID
	req.UserID = userID

	booking, err := h.serviceService.BookLawyer(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created",
		"booking": booking,
	})
}

// GetCurrencyRates handles GET /api/v1/services/currency-rates
func (h *ServiceHandler) GetCurrencyRates(c *gin.Context) {
	from := c.DefaultQuery("from", "USD")
	to := c.DefaultQuery("to", "ETB")

	rates, err := h.serviceService.GetCurrencyRates(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rates": rates,
		"timestamp": time.Now(),
	})
}