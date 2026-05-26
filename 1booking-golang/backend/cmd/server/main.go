package main

import (
	systemhttp "1booking-golang/backend/internal/system/http"

	"1booking-golang/backend/internal/platform/config"
	"1booking-golang/backend/internal/platform/database"

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

	// register route
	mux.HandleFunc(
		"/health",
		healthHandler.Health,
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