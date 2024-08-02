package config

import "os"

type config struct {
	Server serverConfig
}

type serverConfig struct {
	Port               string
	RunMode            string
	TemplateDir        string
	StaticDir          string
	AdSeverExternalAPI string
	PanelExternalHost  string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		asserverHost := os.Getenv("ADSERVERHOST")
		panelHost := os.Getenv("PANELHOST")
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:               "9001",
					RunMode:            "release",
					TemplateDir:        "/app/panel/api/templates/*",
					StaticDir:          "/app/panel/api/static",
					AdSeverExternalAPI: asserverHost + "/api/ad",
					PanelExternalHost:  panelHost,
				},
			}
		} else {
			cfg = &config{
				Server: serverConfig{
					Port:               "8080",
					RunMode:            "debug",
					TemplateDir:        "./api/templates/*",
					StaticDir:          "./api/static",
					AdSeverExternalAPI: "http://localhost:5002/api/ad",
					PanelExternalHost:  "http://localhost:8080",
				},
			}

		}
	}

	return cfg
}
