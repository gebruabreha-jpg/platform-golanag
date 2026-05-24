package main

import (
	"connectme/internal/config"
	"connectme/internal/database"
)

func main() {
	cfg := config.LoadProd()

	db := database.NewPostgres(cfg)
	defer database.Close(db)
}
