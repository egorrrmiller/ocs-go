package service

import (
	"ocs-go/internal/models"
	"ocs-go/internal/repository"

	"github.com/google/uuid"
)

type Order interface {
	GetOrder(id uuid.UUID) models.Order
	AddOrder(order models.Order) error
	UpdateOrder(order models.Order)
	DeleteOrder(id uuid.UUID)
}

type OrderService struct {
	repository repository.Order
}

func NewOrderService(repository repository.Order) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) GetOrder(id uuid.UUID) models.Order {
	return s.repository.GetOrder(id)
}

func (s *OrderService) AddOrder(order models.Order) error {
	return s.repository.AddOrder(order)
}

func (s *OrderService) UpdateOrder(order models.Order) {
	s.repository.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uuid.UUID) {
	s.repository.DeleteOrder(id)
}
