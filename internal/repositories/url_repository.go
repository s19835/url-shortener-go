package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/s19835/url-shortener-go/internal/models"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL)
	FindByShortCode(ctx context.Context, shortCode string) (*models.URL, error)
}

type postgresURLRepository struct {
	db *pgxpool.Pool
}

// Create implements URLRepository.
func (p *postgresURLRepository) Create(ctx context.Context, url *models.URL) {
	panic("unimplemented")
}

// FindByShortCode implements URLRepository.
func (p *postgresURLRepository) FindByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	panic("unimplemented")
}

func NewURLRepository(db_url string) (URLRepository, error) {
	pool, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &postgresURLRepository{db: pool}, nil
}
