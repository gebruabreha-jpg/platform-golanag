package handlers

import (
	"net/http"

	"fintech-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateCheckout(c *gin.Context) {
	s, err := services.CreateCheckoutSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": s.URL,
	})
}
