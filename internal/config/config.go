package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database string
	LogLevel string
}

func LoadConfig() (*Config, error) {
	// Load .env if present; ignore error if file missing so local dev works without it
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		Port:     port,
		Database: os.Getenv("DATABASE_URL"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}
