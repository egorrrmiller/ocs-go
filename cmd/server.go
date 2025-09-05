package main

import (
	_ "ocs-go/docs"
	"ocs-go/internal/handler"
	"ocs-go/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(productService *service.ProductService, orderService *service.OrderService) *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productHandlers := handler.NewProductHandler(productService)
	productHandlers.ConfigureRoutes(engine)

	orderHandlers := handler.NewOrderHandler(orderService)
	orderHandlers.ConfigureRoutes(engine)

	return engine
}
