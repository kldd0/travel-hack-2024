package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type TourService struct {
	tourRepository repository.Tour
}

func NewTourService(tourRepository repository.Tour) *TourService {
	return &TourService{tourRepository: tourRepository}
}

func (s *TourService) GetTourById(ctx context.Context, id int) (entity.Tour, error) {
	return s.tourRepository.GetTourById(ctx, id)
}

func (s *TourService) GetAllTours(ctx context.Context /* filtering params */) ([]entity.Tour, error) {
	return s.tourRepository.GetAllTours(ctx)
}
