package config

import "os"

type config struct {
	Server serverConfig
}

type serverConfig struct {
	Port        string
	RunMode     string
	TemplateDir string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:        "9004",
					RunMode:     "release",
					TemplateDir: "app/publisherwebsite/templates/*",
				},
			}
		} else {
			cfg = &config{
				Server: serverConfig{
					Port:        "80",
					RunMode:     "debug",
					TemplateDir: "templates/*",
				},
			}

		}
	}

	return cfg
}
