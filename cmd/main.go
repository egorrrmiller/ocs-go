package main

import (
	"ocs-go/config"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	if err := Run("8080", cfg); err != nil {
		logrus.Fatal(err.Error())
	}
}
