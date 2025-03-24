package repository

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

func SaveURL(shortCode string, longURL string) error {
	return redisClient.Set(context.Background(), shortCode, longURL, 0).Err()
}

func GetURL(shortCode string) (string, error) {
	url, err := redisClient.Get(context.Background(), shortCode).Result()

	if err != nil {
		return "", errors.New("URL not found")
	}

	return url, err
}
