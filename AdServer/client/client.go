package client

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/sarrkar/chan-ta-net/ad-server/config"
)

type Ad struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	ImageUrl     string `json:"image_url"`
	RedirectUrl  string `json:"redirect_url"`
	BID          uint   `json:"bid"`
	Active       bool   `json:"active"`
	Impression   uint   `json:"impression"`
	Click        uint   `json:"click"`
	AdvertiserID uint   `json:"advertiser_id"`
}

var ads []Ad = []Ad{}
var mu sync.Mutex = sync.Mutex{}

func GetBestAds() Ad {
	mu.Lock()
	defer mu.Unlock()
	if len(ads) >= 1 {
		return ads[0]
	} else {
		return Ad{}
	}
}

func UpdateAdsCache() {
	for {
		if err := getAds(); err != nil {
			time.Sleep(config.Config().Client.Retry)
			continue
		}
		time.Sleep(config.Config().Client.Period)
	}
}

func getAds() error {
	mu.Lock()
	defer mu.Unlock()
	resp, err := http.Get(config.Config().Client.PanelApi)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &ads)
	resp.Body.Close()

	sort.Slice(ads, func(i, j int) bool {
		return ads[i].BID > ads[j].BID
	})
	return nil
}
