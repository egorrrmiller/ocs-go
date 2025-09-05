package service

import (
	model "ocs-go/internal/model/db"
	"ocs-go/internal/repository"
)

type ProductService struct {
	repository repository.Product
}

func (s *ProductService) AddProduct(model model.Product) {
	s.repository.AddProduct(model)
}

func NewProductService(repository repository.Product) *ProductService {
	return &ProductService{
		repository: repository,
	}
}
