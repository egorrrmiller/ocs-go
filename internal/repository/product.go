package repository

import (
	"ocs-go/internal/models"

	"gorm.io/gorm"
)

type Product interface {
	AddProduct(product models.Product)
}

type ProductRepository struct {
	db *gorm.DB
}

func (r ProductRepository) AddProduct(product models.Product) {
	r.db.Create(&product)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
