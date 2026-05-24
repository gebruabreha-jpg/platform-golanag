package identity

import (
	"connectme/backend/internal/domains/identity/handler"
	"connectme/backend/internal/domains/identity/persistence/postgres"
	"connectme/backend/internal/domains/identity/repository"
	"connectme/backend/internal/domains/identity/service"
	"connectme/internal/database"
	"connectme/internal/platform/middleware"
)

// Module represents the identity module.
type Module struct {
	router       *gin.Engine
	authHandler  *handler.AuthHandler
	userHandler  *handler.ProfileHandler
	travellerHandler *handler.TravellerHandler
	sessionHandler *handler.SessionHandler
	securityHandler *handler.SecurityHandler
	preferenceHandler *handler.PreferenceHandler
	adminHandler *handler.AdminUserHandler
}

// NewModule creates a new identity module.
func NewModule(router *gin.Engine, db *pgxpool.Pool) *Module {
	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	travellerRepo := postgres.NewTravellerRepository(db)
	sessionRepo := postgres.NewSessionRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	travellerService := service.NewTravellerService(travellerRepo)
	sessionService := service.NewSessionService(sessionRepo)
	securityService := service.NewSecurityService(userRepo, sessionRepo)
	preferenceService := service.NewPreferenceService(userRepo)
	adminService := service.NewAdminUserService(userRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewProfileHandler(userService)
	travellerHandler := handler.NewTravellerHandler(travellerService)
	sessionHandler := handler.NewSessionHandler(sessionService)
	securityHandler := handler.NewSecurityHandler(securityService)
	preferenceHandler := handler.NewPreferenceHandler(preferenceService)
	adminHandler := handler.NewAdminUserHandler(adminService)

	return &Module{
		router:       router,
		authHandler:  authHandler,
		userHandler:  userHandler,
		travellerHandler: travellerHandler,
		sessionHandler: sessionHandler,
		securityHandler: securityHandler,
		preferenceHandler: preferenceHandler,
		adminHandler: adminHandler,
	}
}

// RegisterRoutes registers all routes for the identity module.
func (m *Module) RegisterRoutes() {
	// Auth routes
	authGroup := m.router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", m.authHandler.Register)
		authGroup.POST("/login", m.authHandler.Login)
		authGroup.POST("/logout", middleware.Auth(), m.authHandler.Logout)
		authGroup.POST("/refresh", m.authHandler.RefreshToken)
		authGroup.POST("/forgot-password", m.authHandler.ForgotPassword)
		authGroup.POST("/reset-password", m.authHandler.ResetPassword)
		authGroup.POST("/verify-email", m.authHandler.VerifyEmail)
	}

	// Profile routes
	profileGroup := m.router.Group("/api/v1/profile")
	{
		profileGroup.Use(middleware.Auth())
		profileGroup.GET("/", m.userHandler.GetProfile)
		profileGroup.PATCH("/", m.userHandler.UpdateProfile)
		profileGroup.DELETE("/", m.userHandler.DeleteAccount)
	}

	// Traveller routes
	travellerGroup := m.router.Group("/api/v1/travellers")
	{
		travellerGroup.Use(middleware.Auth())
		travellerGroup.GET("/", m.travellerHandler.GetTravellers)
		travellerGroup.POST("/", m.travellerHandler.CreateTraveller)
		travellerGroup.PATCH("/:id", m.travellerHandler.UpdateTraveller)
		travellerGroup.DELETE("/:id", m.travellerHandler.DeleteTraveller)
	}

	// Security routes
	securityGroup := m.router.Group("/api/v1/security")
	{
		securityGroup.Use(middleware.Auth())
		securityGroup.GET("/sessions", m.sessionHandler.GetSessions)
		securityGroup.DELETE("/sessions/:id", m.sessionHandler.RevokeSession)
		securityGroup.PATCH("/password", m.securityHandler.ChangePassword)
	}

	// Admin routes
	adminGroup := m.router.Group("/api/v1/admin")
	{
		adminGroup.Use(middleware.Auth(), middleware.Role("admin"))
		adminGroup.GET("/users", m.adminHandler.GetUsers)
		adminGroup.PATCH("/users/:id/block", m.adminHandler.BlockUser)
	}
}