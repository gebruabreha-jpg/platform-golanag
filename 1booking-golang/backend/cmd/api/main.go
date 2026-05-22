package main

import (
	"connectme/internal/config"
	"connectme/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg)
	defer database.Close(db)
}
