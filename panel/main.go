package main

import (
	"log"

	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/panel/api"
)

func main() {

	err := database.InitDb()
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}
	database.Migrate()

	api.InitServer()
}
