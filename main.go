package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/config"
)

func main() {
	// load env variagle
	config.Load()

	// get neccessary variables
	DB_URL := config.GetEnv("DB_URL", "optional")
	log.Println(DB_URL)

	// create a router
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		log.Println("Connected")
	})

	r.Run()
}
