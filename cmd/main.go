package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"ocs-go/config"
	kafkaModel "ocs-go/internal/model/kafka"
	"ocs-go/pkg/database"
	consumer "ocs-go/pkg/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Shutting down everything...")
		cancel()
	}()

	// загрузка конфига
	cfg, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// инициализация бд
	db, err := database.InitDb(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	// контейнер зависимостей
	container := NewContainer(cfg, db)

	// мок внешнего сервиса
	go kafkaMock(ctx, container.Config.Kafka)

	// регистрация консюмеров
	consumerRegistry := consumer.NewRegistry()
	consumerRegistry.RegisterConsumer(consumer.NewProductConsumer(container.ProductService))
	consumerRegistry.StartAll(ctx, container.Config.Kafka)

	// gin
	engine := NewRouter(container.ProductService, container.OrderService)
	server := &http.Server{
		Addr:    ":" + container.Config.Port,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Error(err.Error())
		}
	}()

	<-ctx.Done()

	// закрываем сервер сразу, без таймаута
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Error(err.Error())
	}

	log.Println("Все сервисы остановлены")
}

// мок сервиса, отвечающего за шаринг "товаров" в системе
func kafkaMock(ctx context.Context, config config.KafkaConfig) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(config.Brokers...),
		Topic:                  "products",
		AllowAutoTopicCreation: true,
	}
	defer w.Close()

	for i := 0; i < 250; i++ {

		select {
		case <-ctx.Done():
			return
		default:
		}

		idProduct := uuid.New()
		data, _ := json.Marshal(kafkaModel.Product{
			Id:          idProduct,
			Name:        fmt.Sprintf("name %s", idProduct),
			Description: fmt.Sprintf("description %s", idProduct),
		})

		err := w.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(idProduct.String()),
				Value: data,
			})

		if err != nil {
			logrus.Error(err.Error())
			continue
		}

		logrus.Println("Сообщение в топик producers отправлено")

		sleep := rand.Intn(10)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}
