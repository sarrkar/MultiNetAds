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
}

var cfg *config

func Config() *config {
	if cfg == nil {
		host := os.Getenv("ADSERVERHOST")
		port := os.Getenv("ADSERVERPORT")
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:               "9001",
					RunMode:            "release",
					TemplateDir:        "/app/panel/api/templates/*",
					StaticDir:          "/app/panel/api/static",
					AdSeverExternalAPI: host + ":" + port + "/api/ad",
				},
			}
		} else {
			cfg = &config{
				Server: serverConfig{
					Port:               "8080",
					RunMode:            "debug",
					TemplateDir:        "api/templates/*",
					StaticDir:          "api/static",
					AdSeverExternalAPI: "http://localhost:5002/api/ad",
				},
			}

		}
	}

	return cfg
}
