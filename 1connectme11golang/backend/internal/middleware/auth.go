package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		role := claims["role"].(string)

		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}
}

func AuthRequired(secret string) gin.HandlerFunc {
	return AuthMiddleware(secret)
}

func CORS() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

// RequireRole enforces RBAC — at least one of the listed roles must
// match the role declared in the JWT claims.
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	roleSet := make(map[string]struct{}, len(allowedRoles))
	for _, r := range allowedRoles {
		roleSet[r] = struct{}{}
	}

	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		if _, ok := roleSet[role.(string)]; !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// IsAdmin is a convenience check for ADMIN-level operations.
func IsAdmin(c *gin.Context) bool {
	role, ok := c.Get("role")
	return ok && role.(string) == "ADMIN"
}

// RateLimit protects routes with a Redis-backed sliding-window counter.
// Default: 100 requests per minute per IP+endpoint key.
func RateLimit(redisClient *redis.Client) gin.HandlerFunc {
	const (
		limit  = 100 // max requests
		window = 60  // seconds
	)

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		key := "ratelimit:" + c.ClientIP() + ":" + c.Request.URL.Path

		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			// Fail open — allow the request if Redis is down
			c.Next()
			return
		}

		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.Header("Retry-After", strconv.Itoa(window))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Limit", strconv.Itoa(limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(limit-int(count)))

		c.Next()
	}
}

// JWTClaims represents the JWT token claims
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID, email, role string, secret string, expiryHours int) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiryHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
