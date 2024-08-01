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
	PanelApi     string
	PublisherApi string
	Period       time.Duration
	Retry        time.Duration
}

var cfg *config

func Config() *config {
	if cfg == nil {
		host := os.Getenv("EVENTSERVERHOST")
		port := os.Getenv("EVENTSERVERPORT")
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:                   "5002",
					RunMode:                "debug",
					EventSeverExternalHost: host + ":" + port,
					OTLlength:              10,
				},
				Client: clientConfig{
					PanelApi:     "http://panel-webserver:9001/api/ad/all_ads",
					PublisherApi: "http://panel-webserver:9001/api/publisher/all_publishers",
					Period:       10 * time.Minute,
					Retry:        5 * time.Second,
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
					PanelApi:     "http://localhost:5001/api/ad/all_ads",
					PublisherApi: "http://localhost:5001/api/publisher/all_publishers",
					Period:       1 * time.Minute,
					Retry:        5 * time.Second,
				},
			}
		}
	}

	return cfg
}
