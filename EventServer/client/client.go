package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sarrkar/chan-ta-net/event-server/config"
)

func AddClick(adID, advID, pubID string) {
	fmt.Println("add click")
	resp, err := http.Get(config.Config().Client.PanelApi + adID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func AdImperession(adID, advID, pubID string) {
	fmt.Println("add impression")
	resp, err := http.Get(config.Config().Client.PanelApi + adID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
