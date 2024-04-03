package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type Tour interface {
	GetTourById(ctx context.Context, id int) (entity.Tour, error)
	GetAllTours(ctx context.Context /* ...filtering params */) ([]entity.Tour, error)
}

type Review interface {
	GetAllReviewsByTourId(ctx context.Context /* параметры фильтра*/) ([]entity.Review, error)
}

type Services struct {
	// Auth        Auth
	// Account     Account
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
