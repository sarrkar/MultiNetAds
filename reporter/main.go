package main

import (
	"log"
	"time"

	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/reporter/config"
	"github.com/sarrkar/chan-ta-net/reporter/consumer"
)

func main() {
	err := database.InitDb()
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}
	database.Migrate()

	c := consumer.NewConsumer()
	go c.GetClick()
	go c.GetImpression()
	for {
		c.GetBidCatch()
		time.Sleep(config.Config().Client.Period)
		c.Save()
		c.ResetCatch()
	}
}
