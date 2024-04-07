package postgres

import (
	"context"
	"fmt"

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
	sql, args, _ := r.Builder.
		Select("*").
		From("reviews").
		Where("tour_id = ?", tourId).
		ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ReviewRepository.GetAllByTourId - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var reviews []entity.Review
	for rows.Next() {
		var review entity.Review
		err := rows.Scan(
			&review.Id,
			&review.TourId,
			&review.Liked,
			&review.Username,
			&review.PositiveComment,
			&review.NegativeComment,
			&review.LocalResident,
			&review.Type,
			&review.Frequency,
			&review.Video,
		)
		if err != nil {
			return nil, fmt.Errorf("ReviewRepository.GetAllByTourId - rows.Scan: %v", err)
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}
