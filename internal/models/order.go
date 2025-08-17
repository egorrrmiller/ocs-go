package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Status    int
	Created   time.Time
	IsDeleted bool
	Lines     []*Line
}
