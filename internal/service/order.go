package service

import (
	"ocs-go/internal/repository"
)

type OrderService struct {
	repository *repository.Order
}

func NewOrderService(repository *repository.Order) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) GetOrder()    {}
func (s *OrderService) AddOrder()    {}
func (s *OrderService) UpdateOrder() {}
func (s *OrderService) DeleteOrder() {

}
