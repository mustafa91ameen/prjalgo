package main

import (
	"context"
	"log"

	"github.com/mustafa91ameen/prjalgo/backend/internal/config"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
	"github.com/mustafa91ameen/prjalgo/backend/internal/db"
	"github.com/mustafa91ameen/prjalgo/backend/internal/logger"
	"github.com/mustafa91ameen/prjalgo/backend/internal/server"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load config from .env
	cfg := config.Load()

	// Connect to database with retry logic
	ctx := context.Background()
	pool, err := db.Connect(ctx, cfg.DSN(), cfg.DBMaxRetries, cfg.DBRetryDelay)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Println("Connected to database")

	// Build dependencies
	c := container.New(pool, cfg)

	// Create and start server
	srv := server.NewServer(cfg.ServerPort, c, cfg)
	srv.RegisterRoutes()

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
