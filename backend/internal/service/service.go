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
	Process(ctx context.Context /* params for filtering */) error // обработка заявки
}

type Services struct {
	// Auth        Auth
	// Account     Account
	Order  Order
	Review Review
	Tour   Tour
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Tour: deps.Repos.Tour,
	}
}
