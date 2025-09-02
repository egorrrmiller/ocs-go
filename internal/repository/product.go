package repository

import (
	"ocs-go/internal/model/db"

	"gorm.io/gorm"
)

type Product interface {
	AddProduct(product db.Product)
}

type ProductRepository struct {
	db *gorm.DB
}

func (r ProductRepository) AddProduct(product db.Product) {
	r.db.Create(&product)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
