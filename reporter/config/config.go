package config

import (
	"os"
)

type config struct {
	Client   consumerConfig
	Postgres postgresConfig
}

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	TimeZone string
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
				Postgres: postgresConfig{
					Host:     os.Getenv("POSTGRES_HOST"),
					Port:     "5432",
					User:     os.Getenv("POSTGRES_USER"),
					Password: os.Getenv("POSTGRES_PASSWORD"),
					DbName:   os.Getenv("POSTGRES_DB"),
				},
			}
		} else {
			cfg = &config{
				Client: consumerConfig{
					KafkaUrl: "localhost:29092",
				},
				Postgres: postgresConfig{
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "Mam@d123",
					DbName:   "panel",
				},
			}
		}
	}

	return cfg
}
