package main

import (
	"log"
	"vibe-storm/internal/interfaces/http"
	"vibe-storm/pkg/config"
	"vibe-storm/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Echo instance with middleware
	e := middleware.InitializeEcho()

	// Setup routes
	http.SetupRoutes(e)

	// Start server
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
