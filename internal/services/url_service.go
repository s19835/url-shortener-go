package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/s19835/url-shortener-go/internal/models"
	"github.com/s19835/url-shortener-go/internal/repositories"
	"github.com/s19835/url-shortener-go/pkg/utils"
)

type URLService struct {
	repo  repositories.URLRepository
	redis *redis.Client
}

func NewURLService(repo repositories.URLRepository, redisCfg models.RedisConfig) *URLService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	return &URLService{
		repo:  repo,
		redis: rdb,
	}
}

func (s *URLService) ShortenURL(ctx context.Context, originalURL string, expiry time.Duration) (string, error) {
	//generate short code
	shortCode, err := utils.GenerateShortCode(originalURL)
	if err != nil {
		return "", err
	}

	// create URL record
	url := &models.URL{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(expiry),
	}

	// stores in Database
	if err := s.repo.Create(ctx, url); err != nil {
		return "", fmt.Errorf("fail to save URL: %w", err)
	}

	// cache in redis
	if err := s.cacheURL(ctx, url); err != nil {
		fmt.Printf("Warning: failed to cache URL: %v\n", err)
	}

	return shortCode, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	// Try to get from cache file
	cachedURL, err := s.redis.Get(ctx, shortCode).Result()
	if err == nil {
		return cachedURL, err
	}

	// fall back to database
	url, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("Error: Url not found : %w", err)
	}

	// cache the result for feature request
	if err := s.cacheURL(ctx, url); err != nil {
		fmt.Printf("Warning: failed to cache URL: %v\n", err)
	}

	return url.OriginalURL, nil
}

func (s *URLService) cacheURL(ctx context.Context, url *models.URL) error {
	ttl := time.Until(url.ExpiresAt)

	if ttl <= 0 {
		return errors.New("URL already expried")
	}

	return s.redis.Set(ctx, url.ShortCode, url.OriginalURL, ttl).Err()
}
