package database

import (
	"fmt"
	"ocs-go/config"
	model "ocs-go/internal/model/db"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err := db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderProducts{}); err != nil {
		logrus.Fatal(err.Error())
	}

	return db, err
}
