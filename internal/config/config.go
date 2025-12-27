package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	DatabaseURL string
	BaseURL     string
	SentryDSN   string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		AppPort:     "8080",
		DatabaseURL: os.Getenv("DATABASE_URL"),
		BaseURL:     os.Getenv("BASE_URL"),
		SentryDSN:   os.Getenv("SENTRY_DSN"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "http://localhost:8080"
	}

	return cfg
}
