package main

import (
	"ocs-go/config"
	"ocs-go/internal/handler"
	"ocs-go/internal/repository"
	"ocs-go/internal/service"
	"ocs-go/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(port string, config config.Config) error {
	db, err := database.InitDb(config)
	if err != nil {
		logrus.Fatal(err)
	}

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	router := gin.New()

	newHandler.ConfigureRoutes(router)

	return router.Run(":" + port)
}
