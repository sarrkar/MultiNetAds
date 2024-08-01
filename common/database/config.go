package database

import "os"

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	TimeZone string
	SslMode  string
}

var cfg *postgresConfig

func PostgresConfig() *postgresConfig {
	if cfg == nil {
		if os.Getenv("APP_ENV") == "docker" {
			cfg = &postgresConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     "5432",
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DbName:   os.Getenv("POSTGRES_DB"),
				SslMode:  "disable",
				TimeZone: "Asia/Tehran",
			}
		} else {
			cfg = &postgresConfig{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "Mam@d123",
				DbName:   "panel",
				SslMode:  "disable",
				TimeZone: "Asia/Tehran",
			}
		}
	}
	return cfg
}
