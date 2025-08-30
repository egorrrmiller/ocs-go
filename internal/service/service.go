package service

import (
	"ocs-go/internal/repository"
)

type Product interface {
	AddProduct()
}

type Service struct {
	Order
	Product
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(r.Product),
		Order:   NewOrderService(r.Order),
	}
}
