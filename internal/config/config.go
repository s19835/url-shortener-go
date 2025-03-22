package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// load the enviornment variables using godotenv
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("warning: No .env file found!\n%s", err)
	}
}

// get the enviornment variables or return default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
