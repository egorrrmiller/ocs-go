package kafka

import (
	"context"
	"encoding/json"
	"ocs-go/config"
	model "ocs-go/internal/model/db"
	"ocs-go/internal/service"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (p *Product) MapToModel() model.Product {
	return model.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	}
}

type ProductConsumer struct {
	productService *service.ProductService
}

func NewProductConsumer(productService *service.ProductService) *ProductConsumer {
	return &ProductConsumer{
		productService: productService,
	}
}

func (consumer *ProductConsumer) Start(ctx context.Context, cfg config.KafkaConfig) {
	c := NewConsumer(cfg, ProductTopic)

	c.Start(ctx, func(msg kafka.Message) error {
		var product Product

		err := json.Unmarshal(msg.Value, &product)
		if err != nil {
			return err
		}

		consumer.productService.AddProduct(product.MapToModel())

		logrus.Info("Топик products обработан.")

		return nil
	})
}
