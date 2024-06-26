package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type Tour interface {
	GetById(ctx context.Context, id int) (entity.DTOTour, error)
	GetMany(ctx context.Context, filters map[string]interface{}) ([]entity.SimplifiedTourView, error)
	GetHotMany(ctx context.Context, filters map[string]interface{}) ([]entity.SimplifiedTourView, error)
}

type Review interface {
	GetAllByTourId(ctx context.Context, tourId int) ([]entity.DTOReview, error)
}

type Order interface {
	GetOrderById(ctx context.Context, id int /* params for filtering */) (entity.Order, error)
}

type City interface {
	GetMany(ctx context.Context, prefix string, limit int) ([]entity.City, error)
}

type Services struct {
	// Auth        Auth
	// Account     Account

	Tour   Tour
	Order  Order
	Review Review
	City   City
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Tour:   NewTourService(deps.Repos.Tour),
		Review: NewReviewService(deps.Repos.Review),
		City:   NewCityService(deps.Repos.City),
	}
}
