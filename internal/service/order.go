package service

import (
	"ocs-go/internal/model/db"
	"ocs-go/internal/repository"

	"github.com/google/uuid"
)

type OrderService struct {
	repository repository.Order
}

func NewOrderService(repository repository.Order) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) GetOrder(id uuid.UUID) db.Order {
	return s.repository.GetOrder(id)
}

func (s *OrderService) AddOrder(order db.Order) error {
	return s.repository.AddOrder(order)
}

func (s *OrderService) UpdateOrder(order db.Order) {
	s.repository.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uuid.UUID) {
	s.repository.DeleteOrder(id)
}
