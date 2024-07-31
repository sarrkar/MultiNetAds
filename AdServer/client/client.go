package client

import (
	"encoding/json"
	"fmt"
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

type Advertiser struct{
	ID 		uint 	`gorm:"column:id; primary_key; not null" json:"id"`
	Name 	string 	`json:"name"`
	Balance int 	`json:"balance"`
	Ads 	[]Ad 	`json:"ads"`
	}	

var ads []Ad = []Ad{}
var mu sync.Mutex = sync.Mutex{}

var advertiser []Advertiser = []Advertiser{}
var mu_avertiser sync.Mutex = sync.Mutex{}



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
	fmt.Println("try to update ads from panel api")
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
// =================================================================


// func createAdvertiserMap(advertisers []Advertiser) map[uint]int {
// 	advertiserMap := make(map[uint]int)
// 	for _, advertiser := range advertisers {
// 		advertiserMap[advertiser.ID] = advertiser.Balance
// 	}
// 	return advertiserMap
// }

// advertiserMap := createAdvertiserMap()


// func checkBalanceAdvertiser(ads []Ad, advertiserMap map[uint]int){
// 	for i, ad := range ads {
// 		if balance, exists := advertiserMap[ad]; exists && balance < 100{
// 			ads[i].Active = false
// 		}
// 	}
// }