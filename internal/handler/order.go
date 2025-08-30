package handler

import (
	"net/http"
	"ocs-go/internal/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logrus.Error("Invalid Id")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result := h.service.GetOrder(id)

	c.JSON(http.StatusOK, result)
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
	var requestDto dto.OrderRequestDto

	if err := c.ShouldBindJSON(&requestDto); err != nil {
		logrus.Error("Invalid Order")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	order := requestDto.MapToModel()

	err := h.service.AddOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
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
	var requestDto dto.OrderRequestDto

	if err := c.ShouldBindJSON(&requestDto); err != nil {
		logrus.Error("Invalid Order")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	order := requestDto.MapToModel()
	h.service.UpdateOrder(order)
	c.JSON(http.StatusOK, nil)
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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logrus.Error("Invalid Id")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	h.service.DeleteOrder(id)

	c.JSON(http.StatusOK, nil)
}
