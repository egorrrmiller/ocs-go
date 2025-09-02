package repository

import (
	"ocs-go/internal/models/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order interface {
	GetOrder(id uuid.UUID) database.Order
	AddOrder(order database.Order) error
	UpdateOrder(order database.Order)
	DeleteOrder(id uuid.UUID)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrder(id uuid.UUID) (order database.Order) {
	order = database.Order{Id: id}
	r.db.First(&order)
	return
}

func (r *OrderRepository) AddOrder(order database.Order) error {
	result := r.db.Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *OrderRepository) UpdateOrder(order database.Order) {}

func (r *OrderRepository) DeleteOrder(id uuid.UUID) {
	order := database.Order{
		Id:        id,
		IsDeleted: true,
	}

	r.db.Updates(&order)
}
