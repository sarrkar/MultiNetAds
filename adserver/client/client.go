package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/sarrkar/chan-ta-net/adserver/config"
	"github.com/sarrkar/chan-ta-net/common/models"
)

var ads []models.Ad = []models.Ad{}
var pubs []models.Publisher = []models.Publisher{}
var advs []models.Advertiser = []models.Advertiser{}
var mu sync.Mutex = sync.Mutex{}

func UpdateCache() {
	for {
		if err := getAds(); err != nil {
			time.Sleep(config.Config().Client.Retry)
			continue
		}
		if err := getPubs(); err != nil {
			time.Sleep(config.Config().Client.Retry)
			continue
		}
		if err := getAdvs(); err != nil {
			time.Sleep(config.Config().Client.Retry)
			continue
		}
		time.Sleep(config.Config().Client.Period)
	}
}

func getAds() error {
	mu.Lock()
	defer mu.Unlock()
	log.Println("try to update ads from panel api")
	resp, err := http.Get(config.Config().Client.PanelApi + "/all_ads")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &ads)
	resp.Body.Close()
	log.Printf("Fetched %d ads\n", len(ads))

	sort.Slice(ads, func(i, j int) bool {
		if ads[i].Impression != 0 && ads[j].Impression != 0 {
			return uint(ads[i].BID)*(ads[i].Click/ads[i].Impression) > uint(ads[j].BID)*(ads[j].Click/ads[j].Impression)
		}
		return ads[i].BID > ads[j].BID
	})
	return nil
}

func getPubs() error {
	mu.Lock()
	defer mu.Unlock()
	log.Println("try to update publishers from panel api")
	resp, err := http.Get(config.Config().Client.PanelApi + "/all_publishers")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &pubs)
	resp.Body.Close()
	log.Printf("Fetched %d publishers\n", len(pubs))
	return nil
}

func getAdvs() error {
	mu.Lock()
	defer mu.Unlock()
	log.Println("try to update advertisers from panel api")
	resp, err := http.Get(config.Config().Client.PanelApi + "/all_advertisers")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &advs)
	resp.Body.Close()
	log.Printf("Fetched %d advertisers\n", len(advs))
	return nil
}

func getPublisherByID(publisherID uint) *models.Publisher {
	for _, publisher := range pubs {
		if publisher.ID == publisherID {
			return &publisher
		}
	}
	return nil
}

func GetBestAds(publisherID uint, title string) ([]models.Ad, models.Publisher) {
	mu.Lock()
	defer mu.Unlock()

	results := make([]models.Ad, 0)

	publisher := getPublisherByID(publisherID)
	if publisher == nil {
		log.Printf("No publisher found for ID %d\n", publisherID)
		return []models.Ad{}, models.Publisher{}
	}
	log.Printf("Publisher ID: %d, Category: %s\n", publisher.ID, publisher.Category)

	titleTokens := strings.Split(title, " ")

	for _, ad := range ads {
		adTokens := strings.Split(ad.Title, " ")
		for _, token1 := range titleTokens {
			for _, token2 := range adTokens {
				if token1 == token2 && len(token1) > 2 {
					log.Printf("Matched Ad ID: %d, Token: %s\n", ad.ID, token1)
					results = append(results, ad)
				}
			}
		}
	}

	for _, ad := range ads {
		if ad.Category == publisher.Category {
			log.Printf("Matched Ad ID: %d, Category: %s\n", ad.ID, ad.Category)
			results = append(results, ad)
		}
	}

	return results, *publisher
}
