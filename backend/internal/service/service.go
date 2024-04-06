package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type Tour interface {
	GetById(ctx context.Context, id int) (entity.Tour, error)
	GetMany(ctx context.Context /* ...filtering params */) ([]entity.Tour, error)
}

type Review interface {
	GetAllById(ctx context.Context /* параметры фильтра*/) ([]entity.Review, error)
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
		Tour: deps.Repos.Tour,
		City: deps.Repos.City,
	}
}
