package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s19835/url-shortener-go/internal/config"
	"github.com/s19835/url-shortener-go/pkg/utils"
)

func main() {
	// load env variagle
	config.Load()

	// get neccessary variables
	DB_URL := config.GetEnv("DB_URL", "optional")
	log.Println(DB_URL)

	// check utils short code generations
	s, err := utils.GenerateShortCode(DB_URL)
	if err != nil {
		log.Println("Unable generate shrot code!")
	}
	log.Println("short code: ", s)

	// create a router
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		log.Println("Connected")
	})

	r.Run()
}
