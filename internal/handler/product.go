package handler

import (
	"ocs-go/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) ConfigureRoutes(e *gin.Engine) {
	product := e.Group("/products")
	{
		product.POST("/", h.AddProduct)
	}
}

// AddProduct godoc
// @Summary Добавить новый товар
// @Description Создает новую запись товара в системе (независимо от заказов)
// @Tags products
// @Accept json
// @Produce json
// @Param input body dto.ProductRequestDto true "Данные товара"
// @Success 200 {object} dto.ProductRequestDto
// @Failure 400
// @Failure 500
// @Router /products [post]
func (h *ProductHandler) AddProduct(e *gin.Context) {

}
