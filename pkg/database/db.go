package database

import (
	"fmt"
	"ocs-go/config"
	"ocs-go/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err := db.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderProducts{}); err != nil {
		logrus.Fatal(err.Error())
	}

	return db, err
}
