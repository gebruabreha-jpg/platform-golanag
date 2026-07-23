// Package config reads runtime configuration from environment variables.
// It centralizes all config loading so main.go doesn't need to know where values come from.
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all runtime configuration sourced from environment variables.
// Add new fields here when you need more configuration values.
type Config struct {
	// Port is the HTTP server listen port.
	Port int

	// DatabaseURL is the PostgreSQL connection string.
	// Empty string means no database (use in-memory store).
	DatabaseURL string

	// ShutdownTimeoutSeconds is the maximum time to wait for in-flight
	// requests to finish during graceful shutdown.
	ShutdownTimeoutSeconds int
}

// Load reads configuration from the environment, applying sensible defaults.
// It reads PORT (default 8080), DATABASE_URL (default empty), and
// SHUTDOWN_TIMEOUT_SECONDS (default 10).
func Load() *Config {
	return &Config{
		Port:                   getEnvInt("PORT", 8080),
		DatabaseURL:            getEnv("DATABASE_URL", ""),
		ShutdownTimeoutSeconds: getEnvInt("SHUTDOWN_TIMEOUT_SECONDS", 10),
	}
}

// Validate checks that the configuration values are within acceptable ranges.
// Returns an error describing the first validation failure found.
func (c *Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535, got %d", c.Port)
	}
	return nil
}

// UsePostgres reports whether a database connection string was provided.
// Used by main.go to decide whether to connect to Postgres or skip it.
func (c *Config) UsePostgres() bool {
	return c.DatabaseURL != ""
}

// Addr returns the listen address for the HTTP server (e.g. ":8080").
func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}

// getEnv returns the value of the environment variable key, or fallback if unset or empty.
// This is a helper to avoid repeating the os.LookupEnv pattern everywhere.
func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

// getEnvInt returns the integer value of the environment variable key, or fallback if unset or invalid.
// It silently returns fallback on parse errors (e.g. PORT=abc) — this is intentional for config loading.
func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}