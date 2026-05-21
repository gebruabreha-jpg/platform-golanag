package routes

import (
	"fintech-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/create-checkout-session", handlers.CreateCheckout)
}
