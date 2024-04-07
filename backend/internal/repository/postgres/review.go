package postgres

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

type ReviewRepository struct {
	*postgres.Postgres
}

func NewReviewRepository(pg *postgres.Postgres) *ReviewRepository {
	return &ReviewRepository{pg}
}

func (r *ReviewRepository) GetAllByTourId(ctx context.Context, tourId int) ([]entity.Review, error) {
	return []entity.Review{}, nil
}
