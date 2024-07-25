package config

import "os"

type config struct {
	Server   serverConfig
	Postgres postgresConfig
}

type serverConfig struct {
	Port    string
	RunMode string
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
		cfg = &config{
			Server: serverConfig{
				Port:    "8080",
				RunMode: "debug",
			},
			Postgres: postgresConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     "5432",
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DbName:   os.Getenv("POSTGRES_DB"),
			},
		}
	}

	return cfg
}
