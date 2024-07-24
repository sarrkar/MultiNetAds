package main

import (
	"github.com/sarrkar/Chan-ta-net/AdServer/api"
	"github.com/sarrkar/Chan-ta-net/AdServer/client"
)

func main() {
	go client.UpdateAdsCache()
	api.InitServer()
}
