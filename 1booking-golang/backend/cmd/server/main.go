package main

import (
	"connectme/internal/platform/config"
	"connectme/internal/platform/database"
)

func main() {
	cfg := config.LoadProd()

	db := database.NewPostgres(cfg)
	defer database.Close(db)
}
