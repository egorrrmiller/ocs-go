package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Consumer struct {
	Name string `mapstructure:"name"`
}

type Producer struct {
	Name string `mapstructure:"name"`
}

type KafkaConfig struct {
	Brokers       []string   `mapstructure:"brokers"`
	Consumers     []Consumer `mapstructure:"consumers"`
	Producers     []Producer `mapstructure:"producers"`
	ConsumerGroup string     `mapstructure:"consumer_group"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Config struct {
	Port     string      `mapstructure:"port"`
	Kafka    KafkaConfig `mapstructure:"kafka"`
	Database DBConfig    `mapstructure:"database"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	// Устанавливаем значения по умолчанию
	viper.SetDefault("port", "8080")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("kafka.consumer_group", "ocs-go")

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return
}
