package models

import "github.com/google/uuid"

type OrderItem struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	ProductId uuid.UUID
	Quantity  int

	// связи
	Product Product `gorm:"foreignKey:ProductId"`
	Order   Order   `gorm:"foreignKey:OrderId"`
}
