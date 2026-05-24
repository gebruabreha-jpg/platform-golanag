package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDev() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env not found, using system env")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
