package kafka

import (
	"context"
	"ocs-go/config"
	"ocs-go/internal/service"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

const (
	ProductTopic = "products"
)

type Consumer struct {
	reader *kafka.Reader
}

type Services struct {
	product *service.ProductService
}

func NewConsumer(cfg config.KafkaConfig, topic string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   cfg.Brokers,
		GroupID:   cfg.ConsumerGroup,
		Partition: 0,
		Topic:     topic,
		MaxBytes:  10e6, // 10MB
	})

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Start(ctx context.Context, handler func(msg kafka.Message) error) {
	go func() {
		defer func(reader *kafka.Reader) {
			err := reader.Close()
			if err != nil {
				logrus.Error(err.Error())
			}
		}(c.reader)

		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					logrus.Error(err.Error())
					return
				}
				logrus.Error(err.Error())
				continue
			}

			if err := handler(m); err != nil {
				logrus.Error(err.Error())
			}
		}

	}()
}
