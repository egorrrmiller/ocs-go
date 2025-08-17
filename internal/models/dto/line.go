package dto

import "github.com/google/uuid"

// swagger:model LineRequestDto
type LineRequestDto struct {
	// ID товара
	// example: c7f5f167-fa8d-4c1d-9c12-aa3b77b9e5f3
	Id uuid.UUID `json:"id"`
}
