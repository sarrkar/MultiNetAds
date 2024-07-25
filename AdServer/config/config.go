package config

import "time"

type config struct {
	Server serverConfig
	Client clientConfig
}

type serverConfig struct {
	Port           string
	RunMode        string
	EventSeverHost string
	OTLlength      int
}

type clientConfig struct {
	PanelApi string
	Period   time.Duration
}

var cfg *config

func Config() *config {
	if cfg == nil {
		cfg = &config{
			Server: serverConfig{
				Port:           "9000",
				RunMode:        "debug",
				EventSeverHost: "http://event-server-webserver:7000",
				OTLlength:      10,
			},
			Client: clientConfig{
				PanelApi: "http://panel-webserver:8080/api/ad/all_ads",
				Period:   1 * time.Minute,
			},
		}
	}
	return cfg
}
