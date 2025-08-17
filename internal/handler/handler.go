package handler

import (
	"ocs-go/internal/service"

	_ "ocs-go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) ConfigureRoutes(e *gin.Engine) {

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	line := e.Group("/lines")
	{
		line.POST("/", h.AddLine)
	}

	order := e.Group("/orders")
	{
		order.GET("/:id", h.GetOrder)
		order.POST("/", h.AddOrder)
		order.PUT("/:id", h.UpdateOrder)
		order.DELETE("/:id", h.DeleteOrder)
	}
}
