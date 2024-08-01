package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/sarrkar/chan-ta-net/adserver/config"
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
	Category     string `json:"category"`
}

type Publisher struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Balance           int    `json:"balance"`
	CommissionPercent int    `json:"commission_percent"`
	Category          string `json:"category"`
}

var ads []Ad = []Ad{}
var publishers []Publisher = []Publisher{}
var mu sync.Mutex = sync.Mutex{}

func GetBestAds(publisherID uint, title string) ([]Ad, Publisher) {
	mu.Lock()
	defer mu.Unlock()

	results := make([]Ad, 0)

	publisher := getPublisherByID(publisherID)
	if publisher == nil {
		fmt.Printf("No publisher found for ID %d\n", publisherID)
		return []Ad{}, Publisher{}
	}
	fmt.Printf("Publisher ID: %d, Category: %s\n", publisher.ID, publisher.Category)

	titleTokens := strings.Split(title, " ")

	for _, ad := range ads {
		adTokens := strings.Split(ad.Title, " ")
		for _, token1 := range titleTokens {
			for _, token2 := range adTokens {
				if token1 == token2 && len(token1) > 2 {
					fmt.Printf("Matched Ad ID: %d, Token: %s\n", ad.ID, token1)
					results = append(results, ad)
				}
			}
		}
	}

	for _, ad := range ads {
		if ad.Category == publisher.Category {
			fmt.Printf("Matched Ad ID: %d, Category: %s\n", ad.ID, ad.Category)
			results = append(results, ad)
		}
	}

	return results, *publisher
}

func getPublisherByID(publisherID uint) *Publisher {
	for _, publisher := range publishers {
		if publisher.ID == publisherID {
			return &publisher
		}
	}
	return nil
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
		if ads[i].Click != 0 && ads[j].Click != 0 {
			return ads[i].BID*(ads[i].Click/ads[i].Impression) > ads[j].BID*(ads[j].Click/ads[j].Impression)
		}
		return ads[i].BID > ads[j].BID
	})
	return nil
}

func UpdatePublishersCache() {
	for {
		if err := getPublishers(); err != nil {
			time.Sleep(config.Config().Client.Retry)
			continue
		}
		time.Sleep(config.Config().Client.Period)
	}
}

func getPublishers() error {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("try to update publishers from panel api")
	resp, err := http.Get(config.Config().Client.PublisherApi)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &publishers)
	resp.Body.Close()
	fmt.Printf("Fetched %d publishers\n", len(publishers))
	return nil
}
