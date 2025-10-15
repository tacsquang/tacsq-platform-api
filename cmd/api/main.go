package main

import (
	"log"
	"net/http"
	"os"

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

	// Connect DB (optional)
	if cfg.Database != "" {
		db, err := database.ConnectDB(cfg.Database)
		if err != nil {
			log.Printf("warning: could not connect to database: %v; continuing without DB", err)
		} else {
			defer db.Close()
		}
	}

	// Set up router
	router := routes.SetupRoutes()

	// ⚠️ Thêm đoạn này để Render hiểu port
	port := os.Getenv("PORT")
	if port == "" {
		// dùng port từ config (local dev)
		if cfg.Port != "" {
			port = cfg.Port
		} else {
			port = "8080"
		}
	}

	bindAddr := "0.0.0.0:" + port

	log.Printf("Starting server on %s", bindAddr)
	if err := http.ListenAndServe(bindAddr, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
