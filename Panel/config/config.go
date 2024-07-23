package config

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
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "12345678",
				DbName:   "test",
			},
		}
	}

	return cfg
}
