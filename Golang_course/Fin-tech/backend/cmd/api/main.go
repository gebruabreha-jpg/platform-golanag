package main

import (
	"log"
	"time"

	"fintech-backend/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("Error loading configs/.env file: ", err)
	}

	cfg := LoadConfig()

	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer database.Close()
	log.Println("Database connected")

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api/health", func(c *gin.Context) {
		if err := database.Pool.Ping(c.Request.Context()); err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "database unreachable"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "database": "connected"})
	})

	router.Run(":8080")
}
