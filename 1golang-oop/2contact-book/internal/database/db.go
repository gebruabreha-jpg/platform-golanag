package database

import (
	"database/sql"
	"fmt"
	"time"

	"contact-book-api/internal/config"

	_ "github.com/lib/pq"
)

// Connect opens a PostgreSQL connection pool using the configured DATABASE_URL.
func Connect(cfg *config.Config) (*sql.DB, error) {
	if cfg == nil || !cfg.UsePostgres() {
		return nil, fmt.Errorf("DATABASE_URL is not configured")
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}