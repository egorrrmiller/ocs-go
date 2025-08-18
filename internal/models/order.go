package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Created   time.Time `gorm:"autoCreateTime"`
	IsDeleted bool
	Items     []*OrderItem
}
