package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Println("Retrive the env variables")
		return value
	}

	return defaultValue
}
