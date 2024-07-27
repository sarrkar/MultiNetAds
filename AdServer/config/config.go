package config

import (
	"os"
	"time"
)

type config struct {
	Server serverConfig
	Client clientConfig
}

type serverConfig struct {
	Port                   string
	RunMode                string
	EventSeverExternalHost string
	OTLlength              int
}

type clientConfig struct {
	PanelApi string
	Period   time.Duration
	Retry    time.Duration
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:                   "9002",
					RunMode:                "debug",
					EventSeverExternalHost: "http://localhost:5003",
					OTLlength:              10,
				},
				Client: clientConfig{
					PanelApi: "http://panel-webserver:9001/api/ad/all_ads",
					Period:   10 * time.Minute,
					Retry:    5 * time.Second,
				},
			}
		} else {
			cfg = &config{
				Server: serverConfig{
					Port:                   "5002",
					RunMode:                "debug",
					EventSeverExternalHost: "http://localhost:5003",
					OTLlength:              10,
				},
				Client: clientConfig{
					PanelApi: "http://localhost:5001/api/ad/all_ads",
					Period:   1 * time.Minute,
					Retry:    5 * time.Second,
				},
			}
		}
	}

	return cfg
}
