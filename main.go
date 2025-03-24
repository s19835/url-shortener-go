package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a router
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		log.Println("Connected")
	})

	r.Run()
}
