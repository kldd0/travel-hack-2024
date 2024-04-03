package repository

import (
	"context"
	"errors"

	"github.com/kldd0/travel-hack-2024/internal/entity"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type User interface {
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (entity.User, error)
}

type Account interface {
	CreateAccount(ctx context.Context) (int, error)
	GetAccountById(ctx context.Context, id int) (entity.Account, error)
}

type Tour interface {
	GetTourById(ctx context.Context, id int) (entity.Tour, error)
	GetAllTours(ctx context.Context /* params for filtering */) ([]entity.Tour, error)
}

type Review interface {
	GetAllReviewsByTourId(ctx context.Context /* params for filtering */) ([]entity.Review, error)
}

type Repositories struct {
	User
	Account
	Review
	Tour
	// Reservation
}

func NewRepositories() *Repositories {
	return &Repositories{}
}
