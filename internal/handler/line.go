package handler

import "github.com/gin-gonic/gin"

// AddLine godoc
// @Summary Добавить новый товар
// @Description Создает новую запись товара в системе (независимо от заказов)
// @Tags lines
// @Accept json
// @Produce json
// @Param input body dto.LineRequestDto true "Данные товара"
// @Success 200 {object} dto.LineRequestDto
// @Failure 400
// @Failure 500
// @Router /lines [post]
func (h *Handler) AddLine(e *gin.Context) {

}
