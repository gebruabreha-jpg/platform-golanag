package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all runtime configuration sourced from environment variables.
type Config struct {
	// HTTP server
	Port int

	// Postgres connection
	DatabaseURL string
}

// Load reads configuration from the environment, applying sensible defaults.
func Load() *Config {
	return &Config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

// UsePostgres reports whether a database connection string was provided.
func (c *Config) UsePostgres() bool {
	return c.DatabaseURL != ""
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

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

// Addr returns the listen address for the HTTP server.
func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}
