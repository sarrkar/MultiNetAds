package main

import (
//	"log"

	"github.com/sarrkar/chan-ta-net/panel/api"
//	"github.com/sarrkar/chan-ta-net/panel/database"
)

func main() {
/*
	err := database.InitDb()
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}
*/
	api.InitServer()
}
