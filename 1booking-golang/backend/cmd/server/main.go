package main

import (
	"1booking-golang/internal/platform/config"
	"1booking-golang/internal/platform/database"
)

func main() {
	cfg := config.LoadDev()

	db := database.NewPostgres(cfg)
	defer database.Close(db)
}
