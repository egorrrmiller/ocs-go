package service

import "ocs-go/internal/repository"

type ProductService struct {
	repository repository.Product
}

func (s *ProductService) AddProduct() {
	//TODO implement me
	panic("implement me")
}

func NewProductService(repository repository.Product) *ProductService {
	return &ProductService{
		repository: repository,
	}
}
