package config

import (
	"os"
	"time"
)

type config struct {
	Client consumerConfig
}

type consumerConfig struct {
	KafkaUrl string
	Period   time.Duration
	Retry    time.Duration
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Client: consumerConfig{
					KafkaUrl: "kafka:9092",
					Retry:    5 * time.Second,
					Period:   1 * time.Minute,
				},
			}
		} else {
			cfg = &config{
				Client: consumerConfig{
					KafkaUrl: "localhost:29092",
					Retry:    5 * time.Second,
					Period:   1 * time.Minute,
				},
			}
		}
	}

	return cfg
}
