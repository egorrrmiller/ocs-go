package repository

import "gorm.io/gorm"

type Order interface {
}

type Line interface {
	AddLine()
}

type Repository struct {
	Order
	Line
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order: NewOrderRepository(db),
		Line:  NewLineRepository(db),
	}
}
