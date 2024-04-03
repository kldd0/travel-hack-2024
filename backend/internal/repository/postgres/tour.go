package postgres

import (
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

type TourRepository struct {
	pg *postgres.Postgres
}

func NewTourRepository(pg *postgres.Postgres) *TourRepository {
	return &TourRepository{pg}
}
