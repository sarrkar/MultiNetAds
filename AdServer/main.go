package main

import (
	"github.com/sarrkar/chan-ta-net/ad-server/api"
	"github.com/sarrkar/chan-ta-net/ad-server/client"
)

func main() {
	go client.UpdateAdsCache()
	api.InitServer()
}
