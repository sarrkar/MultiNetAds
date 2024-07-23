package main

import (
	"github.com/sarrkar/chan-ta-net/Panel/api"
)

func main() {

	//err := database.InitDb()
	//defer database.CloseDb()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	api.InitServer()
}
