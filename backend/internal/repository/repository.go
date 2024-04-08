package repository

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository/postgres"

	pgdb "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

type User interface {
	GetByUsernameAndPassword(ctx context.Context, username, password string) (entity.User, error)
}

type Account interface {
	Create(ctx context.Context) (int, error)
	GetById(ctx context.Context, id int) (entity.Account, error)
}

type Tour interface {
	GetById(ctx context.Context, id int) (entity.Tour, error)
	GetMany(ctx context.Context, filters map[string]interface{}) ([]entity.Tour, error)
}

type Review interface {
	GetAllByTourId(ctx context.Context, tourId int) ([]entity.Review, error)
}

type Order interface {
	GetOrderById(ctx context.Context, id int /* params for filtering */) (entity.Order, error)
	Process(ctx context.Context /* params for filtering */) error // обработка заявки
}

type City interface {
	GetMany(ctx context.Context, prefix string, limit int) ([]entity.City, error)
}

type Repositories struct {
	// User
	// Account

	Tour
	Order
	City
	Review

	// Reservation
}

func NewRepositories(pg *pgdb.Postgres) *Repositories {
	return &Repositories{
		Tour:   postgres.NewTourRepository(pg),
		Review: postgres.NewReviewRepository(pg),
		City:   postgres.NewCityRepository(pg),
	}
}
