package database

import (
	"context"
	"log"
	"time"

	"your-module-name/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(cfg *config.Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	log.Println("Database connected")

	return db
}