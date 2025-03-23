package main

import (
	"log"

	"github.com/s19835/url-shortener-go/internal/config"
)

func main() {
	config.LoadEnv()
	log.Println("loaded environment variables")
}
