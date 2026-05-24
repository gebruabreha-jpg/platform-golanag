package main

import (
	"1booking-golang/internal/platform/config"
	"1booking-golang/internal/platform/database"
)

func main() {
	cfg := config.LoadProd()

	db := database.NewPostgres(cfg)
	defer database.Close(db)
}
