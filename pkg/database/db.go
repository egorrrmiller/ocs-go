package database

import (
	"fmt"
	"ocs-go/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
