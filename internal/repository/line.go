package repository

import "gorm.io/gorm"

type LineRepository struct {
	db *gorm.DB
}

func (r LineRepository) AddLine() {
	//TODO implement me
	panic("implement me")
}

func NewLineRepository(db *gorm.DB) *LineRepository {
	return &LineRepository{db: db}
}
