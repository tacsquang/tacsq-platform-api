package config

import (
	"os"
)

type Config struct {
	Port     string
	Database string
	LogLevel string
}

func LoadConfig() (*Config, error) {
	// Do not load .env automatically in production builds (Render provides env vars).
	// Rely on environment variables only.

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		// Render (and some platforms) provide PORT without a leading colon, e.g. "10000"
		// Ensure the port string starts with ':' for ListenAndServe
		if port[0] != ':' {
			port = ":" + port
		}
	}

	return &Config{
		Port:     port,
		Database: os.Getenv("DATABASE_URL"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}
