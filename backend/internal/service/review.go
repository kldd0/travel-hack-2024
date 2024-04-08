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

func (s *ReviewService) GetAllByTourId(ctx context.Context, tourId int) ([]entity.DTOReview, error) {
	reviews, err := s.reviewRepository.GetAllByTourId(ctx, tourId)

	dtoReviews := make([]entity.DTOReview, 0, len(reviews))
	for _, review := range reviews {
		dtoReviews = append(dtoReviews, review.ToDTOModel())
	}

	return dtoReviews, err
}
