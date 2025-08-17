package dto

import (
	"time"

	"github.com/google/uuid"
)

// swagger:model OrderRequestDto
type OrderRequestDto struct {
	Lines []Line `json:"lines"`
}

// swagger:model OrderResponseDto
type OrderResponseDto struct {
	// ID заказа
	// example: a8f5f167-fa8d-4c1d-9c12-aa3b77b9e5f1
	Id uuid.UUID `json:"id"`

	// Статус заказа
	// example: 1
	Status int `json:"status"`

	// Время создания
	// example: 2023-10-05T12:00:00Z
	Created time.Time `json:"created"`

	// Список товаров
	Lines []Line `json:"lines"`
}

// swagger:model Line
type Line struct {
	// ID товара
	// example: b5f5f167-fa8d-4c1d-9c12-aa3b77b9e5f2
	LineId uuid.UUID `json:"id"`

	// Количество товара
	// example: 5
	Quality int `json:"qty"`
}
