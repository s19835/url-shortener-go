package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/config"
	"github.com/s19835/url-shortener-go/internal/handlers"
	"github.com/s19835/url-shortener-go/internal/repositories"
	"github.com/s19835/url-shortener-go/internal/services"
)

func main() {
	// load env variagle
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//Initialize repository
	repo, err := repositories.NewURLRepository(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize services
	urlService := services.NewURLService(repo, cfg.Redis)

	// Initialize handlers
	urlHandler := handlers.NewURLHandler(urlService)

	// step a router
	r := gin.Default()

	// routes
	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:shortCode", urlHandler.RedirectURL)

	// starting a server
	log.Printf("Server running on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
