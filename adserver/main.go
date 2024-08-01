package main

import (
	"github.com/sarrkar/chan-ta-net/adserver/api"
	"github.com/sarrkar/chan-ta-net/adserver/client"
)

func main() {
	go client.UpdateAdsCache()
	api.InitServer()
}
