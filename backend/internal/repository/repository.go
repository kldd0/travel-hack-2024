package repository

import (
	"context"
	"errors"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository/postgres"

	pgdb "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
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
	GetMany(ctx context.Context /* params for filtering */) ([]entity.Tour, error)
}

type Review interface {
	GetMany(ctx context.Context /* params for filtering */) ([]entity.Review, error)
}

type Repositories struct {
	User
	Account
	Review
	Tour

	// Reservation
}

func NewRepositories() *Repositories {
	return &Repositories{
		Tour: postgres.NewTourRepository((*pgdb.Postgres)(nil)),
	}
}
