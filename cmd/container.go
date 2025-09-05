package main

import (
	"ocs-go/config"
	"ocs-go/internal/repository"
	"ocs-go/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	ProductService *service.ProductService
	OrderService   *service.OrderService

	ProductRepository *repository.ProductRepository
	OrderRepository   *repository.OrderRepository

	Config config.Config
}

func NewContainer(cfg config.Config, db *gorm.DB) *Container {

	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo)

	return &Container{
		ProductService:    productService,
		OrderService:      orderService,
		ProductRepository: productRepo,
		OrderRepository:   orderRepo,
		Config:            cfg,
	}
}
