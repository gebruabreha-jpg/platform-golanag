package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDev() *Config {
	baseDir, _ := os.LookupEnv("APP_ENV")
	if baseDir == "" {
		baseDir, _ = os.Getwd()
	}

	candidates := []string{
		baseDir + "/.env",
		baseDir + "/../.env",
		".env",
	}

	for _, path := range candidates {
		if err := godotenv.Load(path); err == nil {
			log.Printf("env loaded: %s", path)
			break
		}
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
