package postgres

import "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"

// сущность отзыва из бд
type ReviewRepository struct {
	pg *postgres.Postgres
}

func NewReviewRepository(pg *postgres.Postgres) *ReviewRepository {
	return &ReviewRepository{pg}
}
