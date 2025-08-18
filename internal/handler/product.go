package handler

import "github.com/gin-gonic/gin"

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
func (h *Handler) AddProduct(e *gin.Context) {

}
