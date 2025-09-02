package repository

import (
	"ocs-go/internal/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order interface {
	GetOrder(id uuid.UUID) db.Order
	AddOrder(order db.Order) error
	UpdateOrder(order db.Order)
	DeleteOrder(id uuid.UUID)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrder(id uuid.UUID) (order db.Order) {
	order = db.Order{Id: id}
	r.db.First(&order)
	return
}

func (r *OrderRepository) AddOrder(order db.Order) error {
	result := r.db.Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *OrderRepository) UpdateOrder(order db.Order) {}

func (r *OrderRepository) DeleteOrder(id uuid.UUID) {
	order := db.Order{
		Id:        id,
		IsDeleted: true,
	}

	r.db.Updates(&order)
}
