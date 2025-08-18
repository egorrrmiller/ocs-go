package repository

import "gorm.io/gorm"

type Order interface {
}

type Product interface {
	AddProduct()
}

type Repository struct {
	Order
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order:   NewOrderRepository(db),
		Product: NewProductRepository(db),
	}
}
