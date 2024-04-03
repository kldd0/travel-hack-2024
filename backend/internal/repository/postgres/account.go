package postgres

import (
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

type AccountRepository struct {
	pg *postgres.Postgres
}

func NewAccountRepository(pg *postgres.Postgres) *AccountRepository {
	return &AccountRepository{pg}
}
