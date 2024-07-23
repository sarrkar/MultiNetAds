package main

import (
	"log"

	"github.com/sarrkar/Chan-ta-net/Panel/api"
	"github.com/sarrkar/Chan-ta-net/Panel/config"
	"github.com/sarrkar/Chan-ta-net/Panel/database"
)

func main() {
	cfg := config.GetConfig()

	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}

	api.InitServer(cfg)
}
