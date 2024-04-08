package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type CityService struct {
	cityRepository repository.City
}

func NewCityService(cityRepository repository.City) *CityService {
	return &CityService{cityRepository: cityRepository}
}

func (s *CityService) GetMany(ctx context.Context, prefix string, limit int) ([]entity.City, error) {
	return s.cityRepository.GetMany(ctx, prefix, limit)
}
