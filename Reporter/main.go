package main

import (
	"github.com/sarrkar/chan-ta-net/Reporter/consumer"
)

func main() {
	c := consumer.NewConsumer()
	go c.GetClick()
	go c.GetImpression()
	// tick := time.Tick(1 * time.Minute)
	// for range tick {

	// }
}
