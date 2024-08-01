package config

import "os"

type config struct {
	Server serverConfig
	Client clientConfig
}

type serverConfig struct {
	Port    string
	RunMode string
}

type clientConfig struct {
	PanelApi string
	KafkaUrl string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:    "9003",
					RunMode: "release",
				},
				Client: clientConfig{
					PanelApi: "http://panel-webserver:9001/api/ad",
					KafkaUrl: "kafka:9092",
				},
			}
		} else {
			cfg = &config{
				Server: serverConfig{
					Port:    "5003",
					RunMode: "debug",
				},
				Client: clientConfig{
					PanelApi: "http://localhost:5001/api/ad",
					KafkaUrl: "localhost:29092",
				},
			}
		}
	}

	return cfg
}
