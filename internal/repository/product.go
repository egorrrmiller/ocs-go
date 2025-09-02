package repository

import (
	"ocs-go/internal/models/database"

	"gorm.io/gorm"
)

type Product interface {
	AddProduct(product database.Product)
}

type ProductRepository struct {
	db *gorm.DB
}

func (r ProductRepository) AddProduct(product database.Product) {
	r.db.Create(&product)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
