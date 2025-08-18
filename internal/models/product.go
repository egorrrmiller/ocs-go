package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Id uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	OrderItems []*OrderItem `gorm:"foreignKey:ProductId"`
}
