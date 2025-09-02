package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ocs-go/config"
	kafkaModel "ocs-go/internal/model/kafka"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	go func() {
		kafkaMock(cfg.Kafka)
	}()

	if err := Run(cfg.Port, cfg); err != nil {
		logrus.Fatal(err.Error())
	}
}

// мок сервиса, отвечающего за шаринг "товаров" в системе
func kafkaMock(config config.KafkaConfig) {

	ctx := context.Background()

	idProduct, _ := uuid.NewUUID()
	data, _ := json.Marshal(kafkaModel.Product{
		Id:          idProduct,
		Name:        fmt.Sprintf("name %s", idProduct),
		Description: fmt.Sprintf("description %s", idProduct),
	})

	w := &kafka.Writer{
		Addr:  kafka.TCP(config.Brokers...),
		Topic: "products",
	}

	for i := 0; i < 250; i++ {
		err := w.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(idProduct.String()),
				Value: data,
			})

		if err != nil {
			logrus.Fatal(err.Error())
		}

		logrus.Println("Сообщение в топик producers отправлено")

		time.Sleep(5 * time.Second)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
