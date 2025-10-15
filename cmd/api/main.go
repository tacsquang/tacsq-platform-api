package main

import (
	"log"
	"net/http"

	"go-api-backend/internal/api/routes"
	"go-api-backend/internal/config"
	"go-api-backend/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Try to connect to the database, but keep running if unavailable (useful for local dev)
	if cfg.Database != "" {
		db, err := database.ConnectDB(cfg.Database)
		if err != nil {
			log.Printf("warning: could not connect to database: %v; continuing without DB", err)
		} else {
			defer db.Close()
		}
	}

	// Set up router (routes.SetupRoutes accepts a gin.Engine)
	router := routes.SetupRoutes()

	// Start the HTTP server
	bindAddr := cfg.Port
	// If cfg.Port starts with ':' assume it's just the port and bind to 0.0.0.0
	if len(bindAddr) > 0 && bindAddr[0] == ':' {
		bindAddr = "0.0.0.0" + bindAddr
	}

	log.Printf("Starting server on %s", bindAddr)
	if err := http.ListenAndServe(bindAddr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
