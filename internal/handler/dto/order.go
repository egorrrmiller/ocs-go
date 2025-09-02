package dto

import (
	"ocs-go/internal/model/db"
	"time"

	"github.com/google/uuid"
)

// swagger:model OrderRequestDto
type OrderRequestDto struct {
	Products []Product `json:"products"`
}

// swagger:model OrderResponseDto
type OrderResponseDto struct {
	// ID заказа
	// example: a8f5f167-fa8d-4c1d-9c12-aa3b77b9e5f1
	Id uuid.UUID `json:"id"`

	// Время создания
	// example: 2023-10-05T12:00:00Z
	Created time.Time `json:"created"`

	// Список товаров
	Products []Product `json:"products"`
}

// swagger:model Product
type Product struct {
	// ID товара
	// example: b5f5f167-fa8d-4c1d-9c12-aa3b77b9e5f2
	ProductId uuid.UUID `json:"id"`

	// Количество товара
	// example: 5
	Quantity int `json:"qty"`
}

func (o *OrderRequestDto) MapToModel() db.Order {
	products := make([]db.OrderProducts, 0)
	orderId := uuid.New()

	for _, product := range o.Products {
		products = append(products, db.OrderProducts{
			OrderId:   orderId,
			ProductId: product.ProductId,
			Quantity:  product.Quantity,
		})
	}

	return db.Order{
		Id:       orderId,
		Products: products,
	}
}
