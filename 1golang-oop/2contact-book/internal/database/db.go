// Package database provides PostgreSQL connection pooling.
// It is a separate package so the connection logic can be reused
// without importing the config or handler packages.
package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Connect opens a PostgreSQL connection pool using the given DSN.
// It sets pool limits and verifies the connection with a Ping.
// Returns an error if the DSN is empty or the database is unreachable.
func Connect(dsn string) (*sql.DB, error) {
	// Guard: if no database URL is provided, fail immediately.
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not configured")
	}

	// sql.Open creates a connection pool but does not connect yet.
	// The actual connection happens on the first query or Ping.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	// Pool tuning: adjust these values based on your database server's capacity.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Ping verifies the connection is alive.
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}
