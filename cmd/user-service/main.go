package main

import (
    "log"

    "go-api-backend/internal/config"
    "go-api-backend/internal/user"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    srv := user.NewServer()
    log.Printf("starting user-service on %s", cfg.Port)
    if err := srv.Run(cfg.Port); err != nil {
        log.Fatalf("user-service stopped: %v", err)
    }
}
