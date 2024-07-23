package config

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	TimeZone string
}

func GetConfig() *Config {
	cfg := &Config{
		Server: ServerConfig{
			Port:    "8080",
			RunMode: "debug",
		},
		Postgres: PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DbName:   "test",
		},
	}

	return cfg
}
