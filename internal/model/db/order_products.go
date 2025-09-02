package db

import "github.com/google/uuid"

type OrderProducts struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	OrderId   uuid.UUID
	ProductId uuid.UUID
	Quantity  int

	// связи
	Product Product `gorm:"foreignKey:ProductId"`
	Order   Order   `gorm:"foreignKey:OrderId"`
}
