package kafka

import (
	"context"
	"ocs-go/config"
)

type ConsumerInterface interface {
	Start(ctx context.Context, cfg config.KafkaConfig)
}

type Registry struct {
	consumers []ConsumerInterface
}

func NewRegistry() *Registry {
	return &Registry{
		consumers: []ConsumerInterface{},
	}
}

func (r *Registry) RegisterConsumer(c ConsumerInterface) {
	r.consumers = append(r.consumers, c)
}

func (r *Registry) StartAll(ctx context.Context, cfg config.KafkaConfig) {
	for _, consumer := range r.consumers {
		consumer.Start(ctx, cfg)
	}
}
