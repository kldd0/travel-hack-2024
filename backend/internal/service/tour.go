package service

import (
	"context"
	"sort"
	"time"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type TourService struct {
	tourRepository repository.Tour
}

func NewTourService(tourRepository repository.Tour) *TourService {
	return &TourService{tourRepository: tourRepository}
}

func (s *TourService) GetById(ctx context.Context, id int) (entity.DTOTour, error) {
	tour, err := s.tourRepository.GetById(ctx, id)
	return tour.ToDTOModel(), err
}

func (s *TourService) GetMany(ctx context.Context, filters map[string]interface{}) ([]entity.SimplifiedTourView, error) {
	tours, err := s.tourRepository.GetMany(ctx, filters)

	simplifiedTours := make([]entity.SimplifiedTourView, 0, len(tours))
	for _, tour := range tours {
		simplifiedTours = append(simplifiedTours, tour.ToSimplifiedTour())
	}

	return simplifiedTours, err
}

func (s *TourService) GetHotMany(ctx context.Context, filters map[string]interface{}) ([]entity.SimplifiedTourView, error) {
	tours, err := s.tourRepository.GetMany(ctx, filters)

	simplifiedTours := make([]entity.SimplifiedTourView, 0, len(tours))
	for _, tour := range tours {
		// check if tour has no dates
		if len(tour.Dates) < 1 {
			continue
		}

		// sort tour dates to get the nearest date to compare with today date
		// to find out if tour is hot (start in or less than 72 hours)
		sort.Slice(tour.Dates, func(i, j int) bool { return tour.Dates[i].Before(tour.Dates[j]) })
		if !(time.Until(tour.Dates[0]) <= time.Hour*72) {
			continue
		}

		simplifiedTours = append(simplifiedTours, tour.ToSimplifiedTour())
	}

	return simplifiedTours, err
}
