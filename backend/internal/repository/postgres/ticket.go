package postgres

import "github.com/kldd0/travel-hack-2024/internal/pkg/postgres"

type TicketRepository struct {
	pg *postgres.Postgres
}

func NewTicketRepository(pg *postgres.Postgres) *TicketRepository {
	return &TicketRepository{pg}
}
