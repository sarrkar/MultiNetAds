package main

import (
	"log"

	"github.com/sarrkar/Chan-ta-net/Panel/api"
	"github.com/sarrkar/Chan-ta-net/Panel/database"
)

func main() {

	err := database.InitDb()
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}

	api.InitServer()
}
