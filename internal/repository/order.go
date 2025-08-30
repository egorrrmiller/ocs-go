package repository

import (
	"ocs-go/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order interface {
	GetOrder(id uuid.UUID) models.Order
	AddOrder(order models.Order) error
	UpdateOrder(order models.Order)
	DeleteOrder(id uuid.UUID)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrder(id uuid.UUID) (order models.Order) {
	order = models.Order{Id: id}
	r.db.First(&order)
	return
}

func (r *OrderRepository) AddOrder(order models.Order) error {
	result := r.db.Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *OrderRepository) UpdateOrder(order models.Order) {}

func (r *OrderRepository) DeleteOrder(id uuid.UUID) {
	order := models.Order{
		Id:        id,
		IsDeleted: true,
	}

	r.db.Updates(&order)
}
