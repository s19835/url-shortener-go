package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/config"
)

func main() {
	config.LoadEnv()
	log.Println("loaded environment variables")

	// set up routers
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
