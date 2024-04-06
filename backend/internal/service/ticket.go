package service

import (
	"context"

	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type TicketService struct {
	ticketRepository repository.Ticket
}

func NewTicketService(ticketRepository repository.Ticket) *TicketService {
	return &TicketService{ticketRepository: ticketRepository}
}

func (s *TicketService) GetTicketById(ctx context.Context, id int) (entity.Ticket, error) {
	return s.ticketRepository.GetTicketById(ctx, id)
}

func (s *TicketService) Process(ctx context.Context) error {
	return s.ticketRepository.Process(ctx) //обработка
}
