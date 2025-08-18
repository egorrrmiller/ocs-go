package repository

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

func (r ProductRepository) AddProduct() {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
