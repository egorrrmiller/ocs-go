package db

import (
	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`

	// связи
	OrderItems []OrderProducts `gorm:"foreignKey:ProductId"`
}
