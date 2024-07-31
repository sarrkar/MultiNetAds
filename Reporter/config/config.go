package config

import (
	"os"
	"time"
)

type config struct {
	Client clientConfig
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
				Client: clientConfig{
					PanelApi: "http://panel-webserver:9001/api/ad/all_ads",
					Period:   10 * time.Minute,
					Retry:    5 * time.Second,
				},
			}
		} else {
			cfg = &config{
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
