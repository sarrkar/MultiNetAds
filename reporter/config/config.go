package config

import (
	"os"
)

type config struct {
	Client consumerConfig
}

type consumerConfig struct {
	KafkaUrl string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Client: consumerConfig{
					KafkaUrl: "kafka:9092",
				},
			}
		} else {
			cfg = &config{
				Client: consumerConfig{
					KafkaUrl: "localhost:29092",
				},
			}
		}
	}

	return cfg
}
