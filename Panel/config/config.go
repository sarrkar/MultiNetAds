package config

import "os"

type config struct {
	Server   serverConfig
	Postgres postgresConfig
}

type serverConfig struct {
	Port        string
	RunMode     string
	TemplateDir string
	StaticDir   string
}

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	TimeZone string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &config{
				Server: serverConfig{
					Port:        "9001",
					RunMode:     "release",
					TemplateDir: "/app/api/templates/*",
					StaticDir:   "app/api/static",
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
				Server: serverConfig{
					Port:        "5001",
					RunMode:     "debug",
					TemplateDir: "api/templates/*",
					StaticDir:   "api/static",
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
