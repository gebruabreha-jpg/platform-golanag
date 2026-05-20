package main

import (
	"log"
	"net/http"

	"connectme/internal/config"
	"connectme/internal/handler"
	"connectme/internal/middleware"
	"connectme/internal/repository/postgres"
	"connectme/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := postgres.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	// Initialize repositories via postgres package;
	// concrete types satisfy the repository interfaces.
	userRepo          := postgres.NewUserRepository(db)
	communityRepo     := postgres.NewCommunityRepository(db)
	postRepo          := postgres.NewPostRepository(db)
	marketplaceRepo   := postgres.NewMarketplaceRepository(db)
	housingRepo       := postgres.NewHousingRepository(db)
	serviceRepo       := postgres.NewServiceRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	communityService := service.NewCommunityService(communityRepo, postRepo)
	marketplaceService := service.NewMarketplaceService(marketplaceRepo)
	housingService := service.NewHousingService(housingRepo)
	serviceService := service.NewServiceService(serviceRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, cfg)
	communityHandler := handler.NewCommunityHandler(communityService)
	marketplaceHandler := handler.NewMarketplaceHandler(marketplaceService)
	housingHandler := handler.NewHousingHandler(housingService)
	serviceHandler := handler.NewServiceHandler(serviceService)

	// Setup Gin router
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.RateLimit(rdb))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "connectme-backend"})
	})

	// API routes
	api := router.Group("/api/v1")

	// Auth routes
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/refresh", userHandler.RefreshToken)
	}

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Use(middleware.AuthRequired(cfg.JWTSecret))
	{
		userRoutes.GET("/profile", userHandler.GetProfile)
		userRoutes.PUT("/profile", userHandler.UpdateProfile)
		userRoutes.POST("/verification", userHandler.SubmitVerification)
	}

	// Community routes
	communityRoutes := api.Group("/communities")
	{
		communityRoutes.GET("", communityHandler.ListCommunities)
		communityRoutes.GET("/:id", communityHandler.GetCommunity)
		communityRoutes.POST("", communityHandler.CreateCommunity)
		communityRoutes.POST("/:id/join", communityHandler.JoinCommunity)
		communityRoutes.POST("/:id/leave", communityHandler.LeaveCommunity)
		communityRoutes.GET("/:id/posts", communityHandler.GetPosts)
		communityRoutes.POST("/:id/posts", communityHandler.CreatePost)
	}

	// Marketplace routes
	marketplaceRoutes := api.Group("/marketplace")
	{
		marketplaceRoutes.GET("/items", marketplaceHandler.ListItems)
		marketplaceRoutes.GET("/items/:id", marketplaceHandler.GetItem)
		marketplaceRoutes.POST("/items", marketplaceHandler.CreateItem)
		marketplaceRoutes.PUT("/items/:id", marketplaceHandler.UpdateItem)
		marketplaceRoutes.POST("/items/:id/interest", marketplaceHandler.ExpressInterest)
		marketplaceRoutes.POST("/transactions/:id/secure-payment", marketplaceHandler.InitiateSecurePayment)
	}

	// Housing routes
	housingRoutes := api.Group("/housing")
	{
		housingRoutes.GET("/listings", housingHandler.ListListings)
		housingRoutes.GET("/listings/:id", housingHandler.GetListing)
		housingRoutes.POST("/listings", housingHandler.CreateListing)
		housingRoutes.POST("/applications", housingHandler.SubmitApplication)
	}

	// Service routes
	serviceRoutes := api.Group("/services")
	{
		serviceRoutes.GET("/currency-rates", serviceHandler.GetCurrencyRates)
		serviceRoutes.GET("/lawyers", serviceHandler.ListLawyers)
		serviceRoutes.POST("/lawyers/:id/book", serviceHandler.BookLawyer)
	}

	// Start server
	log.Printf("ConnectMe API server starting on %s", cfg.ServerAddr)
	if err := router.Run(cfg.ServerAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
