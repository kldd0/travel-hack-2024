package postgres

import "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"

type OrderRepository struct {
	*postgres.Postgres
}

func NewOrderRepository(pg *postgres.Postgres) *OrderRepository {
	return &OrderRepository{pg}
}
