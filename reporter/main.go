package main

import (
	"github.com/sarrkar/chan-ta-net/reporter/consumer"
	"github.com/sarrkar/chan-ta-net/reporter/database"
)

func main() {
	database.InitDb()
	c := consumer.NewConsumer()
	go c.GetClick()
	c.GetImpression()
	// tick := time.Tick(1 * time.Minute)
	// for range tick {

	// }
}
