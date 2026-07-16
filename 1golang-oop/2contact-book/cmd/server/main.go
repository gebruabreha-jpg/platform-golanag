package main

import (
	"fmt"
	"os"

	"contact-book-api/internal/config"
	"contact-book-api/internal/database"

	_ "github.com/joho/godotenv/autoload"
	"log/slog"
)

func main() {
	cfg := config.Load()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.Connect(cfg)
	if err != nil {
		log.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	stats := db.Stats()
	log.Info("database connected",
		"max_open", stats.MaxOpenConnections,
		"open", stats.OpenConnections,
		"in_use", stats.InUse,
		"idle", stats.Idle,
		"wait_count", stats.WaitCount,
		"wait_duration", stats.WaitDuration,
	)

	fmt.Println("OK: database connection successful")
}