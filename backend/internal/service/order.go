package service

import (
	"github.com/kldd0/travel-hack-2024/internal/repository"
)

type OrderService struct {
	orderRepository repository.Order
}

func NewOrderService(orderRepository repository.Order) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

/*
func (s *OrderService) GetOrderById(ctx context.Context, id int) (entity.Order, error) {
	return s.orderRepository.GetById(ctx, id)
}

func (s *OrderService) Process(ctx context.Context) error {
	return s.orderRepository.Process(ctx) //обработка
}

*/
