package handler

import "github.com/gin-gonic/gin"

// GetOrder godoc
// @Summary Получить заказ по ID
// @Description Возвращает информацию о заказе и его товарах
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "UUID заказа" Format(uuid)
// @Success 200 {object} dto.OrderResponseDto
// @Failure 400
// @Failure 500
// @Router /orders/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {

}

// AddOrder godoc
// @Summary Создать новый заказ
// @Description Создает новый заказ с указанными товарами
// @Tags orders
// @Accept json
// @Produce json
// @Param input body dto.OrderRequestDto true "Данные для создания заказа"
// @Success 200 {object} dto.OrderResponseDto
// @Failure 400
// @Failure 500
// @Router /orders [post]
func (h *Handler) AddOrder(c *gin.Context) {

}

// UpdateOrder godoc
// @Summary Обновить существующий заказ
// @Description Обновляет данные заказа и его товары
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "UUID заказа" Format(uuid)
// @Param input body dto.OrderRequestDto true "Обновленные данные заказа"
// @Success 200 {object} dto.OrderResponseDto
// @Failure 400
// @Failure 500
// @Router /orders/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {

}

// DeleteOrder godoc
// @Summary Удалить заказ
// @Description Удаляет заказ по его идентификатору
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "UUID заказа" Format(uuid)
// @Success 200
// @Failure 400
// @Failure 500
// @Router /orders/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {

}
