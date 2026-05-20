package handler

import (
	"connectme/internal/config"
	"connectme/internal/middleware"
	"connectme/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	userService *service.UserService
	cfg         *config.Config
}

func NewUserHandler(userService *service.UserService, cfg *config.Config) *UserHandler {
	return &UserHandler{userService: userService, cfg: cfg}
}

// Register handles POST /api/v1/auth/register
func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles POST /api/v1/auth/login
func (h *UserHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, err := middleware.GenerateToken(user.ID, user.Email, user.Role, h.cfg.JWTSecret, 24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	refreshToken, err := middleware.GenerateToken(user.ID, user.Email, user.Role, h.cfg.RefreshSecret, 168)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"user":         user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// GetProfile handles GET /api/v1/users/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")
	user, err := h.userService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateProfile handles PUT /api/v1/users/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	_ = c.GetString("userID")
	// Bind and update profile
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

// RefreshToken handles POST /api/v1/auth/refresh
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.cfg.RefreshSecret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	accessToken, err := middleware.GenerateToken(user.ID, user.Email, user.Role, h.cfg.JWTSecret, 24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	newRefresh, err := middleware.GenerateToken(user.ID, user.Email, user.Role, h.cfg.RefreshSecret, 168)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": newRefresh,
	})
}

// SubmitVerification handles POST /api/v1/users/verification
func (h *UserHandler) SubmitVerification(c *gin.Context) {
	// Upload ID, proof of address, etc.
	c.JSON(http.StatusOK, gin.H{"message": "Verification submitted"})
}
