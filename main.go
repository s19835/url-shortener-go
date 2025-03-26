package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/config"
)

func main() {
	// load env variagle
	config.Load()

	// Initialize repository
	// repo, err := repositories.NewURLRepository(DB_URL)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize repository: %v", err)
	// }

	// // Initialize services
	// urlService := services.NewURLService(repo, )

	// create a router
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		log.Println("Connected")
	})

	r.Run()
}
