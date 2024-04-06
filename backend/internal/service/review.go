package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type ReviewService struct {
	reviewRepository repository.Review
}

func NewReviewService(reviewRepository repository.Review) *ReviewService {
	return &ReviewService{reviewRepository: reviewRepository}
}

// вернем отзывы по нужному туру
func (s *ReviewService) GetAllReviewsByTourId(ctx context.Context) ([]entity.Review, error) {
	return s.reviewRepository.GetAllById(ctx)
}
