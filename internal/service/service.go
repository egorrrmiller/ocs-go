package service

import "ocs-go/internal/repository"

type Order interface {
	GetOrder()
	AddOrder()
	UpdateOrder()
	DeleteOrder()
}

type Line interface {
	AddLine()
}

type Service struct {
	Order
	Line
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Line:  NewLineService(&r.Line),
		Order: NewOrderService(&r.Order),
	}
}
