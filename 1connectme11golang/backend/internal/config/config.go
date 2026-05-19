package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	ServerAddr  string

	DatabaseURL     string
	MigrationsURL   string
	RedisAddr       string
	RedisPassword   string
	NatsURL         string
	NatsClusterID   string

	JWTSecret       string
	JWTExpiry       time.Duration
	RefreshSecret   string

	StripeSecretKey string
	StripeWebhookSecret string

	AWSRegion       string
	AWSAccessKey    string
	AWSSecretKey    string
	S3Bucket        string

	AI ServiceURL   string
	AI APIKey       string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	// Default values
	viper.SetDefault("environment", "development")
	viper.SetDefault("server.addr", ":8080")
	viper.SetDefault("database.url", "postgres://user:pass@localhost:5432/connectme?sslmode=disable")
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("nats.url", "nats://localhost:4222")
	viper.SetDefault("jwt.expiry", "24h")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Config file error: %v", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return &cfg
}
