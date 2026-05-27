package main

import (
	systemhttp "1booking-golang/backend/internal/system/http"

	"1booking-golang/backend/internal/platform/config"
	"1booking-golang/backend/internal/platform/database"

	"1booking-golang/backend/internal/domains/ota/identity/handler"
	"1booking-golang/backend/internal/domains/ota/identity/service"
	"1booking-golang/backend/internal/domains/ota/identity/persistence/postgres"

	"log"
	"net/http"
)

func main() {

	// load config
	cfg := config.LoadDev()

	// create database connection
	db := database.NewPostgres(cfg)

	// close DB when app shuts down
	defer database.Close(db)

	// create router
	mux := http.NewServeMux()

	// create handler
	healthHandler := systemhttp.NewHealthHandler()

	// create repositories
	userRepo := postgres.NewUserRepository(db)

	// create services
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// register health route
	mux.HandleFunc(
		"/health",
		healthHandler.Health,
	)

	// register auth routes
	mux.HandleFunc(
		"/api/auth/register",
		authHandler.Register,
	)

	// start server
	log.Println("server running on :8080")

	err := http.ListenAndServe(
		":8080",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}