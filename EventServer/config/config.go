package config

type config struct {
	Server serverConfig
	Client clientConfig
}

type serverConfig struct {
	Port    string
	RunMode string
}

type clientConfig struct {
	PanelApi string
}

var cfg *config

func Config() *config {
	if cfg == nil {
		cfg = &config{
			Server: serverConfig{
				Port:    "7000",
				RunMode: "debug",
			},
			Client: clientConfig{
				PanelApi: "http://panel-webserver:8080/api/ad",
			},
		}
	}
	return cfg
}
