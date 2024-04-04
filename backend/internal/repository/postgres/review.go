package postgres

import "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"

type ReviewRepository struct {
	pg *postgres.Postgres
}

func NewReviewRepository(pg *postgres.Postgres) *ReviewRepository {
	return &ReviewRepository{pg}
}
