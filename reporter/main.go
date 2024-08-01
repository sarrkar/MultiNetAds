package main

import (
	"log"
	"time"

	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/common/models"
	"github.com/sarrkar/chan-ta-net/reporter/consumer"
)

func main() {
	err := database.InitDb()
	defer database.CloseDb()
	if err != nil {
		log.Fatalln(err)
	}
	database.Migrate()

	db := database.GetDb()
	c := consumer.NewConsumer()
	go c.GetClick()
	go c.GetImpression()
	for {
		var ads []models.Ad
		var advs []models.Advertiser
		var pubs []models.Publisher

		if result := db.Find(&ads); result.Error != nil {
			log.Fatalln(err)
		}
		if result := db.Find(&advs); result.Error != nil {
			log.Fatalln(err)
		}
		if result := db.Find(&pubs); result.Error != nil {
			log.Fatalln(err)
		}

		c.Lock.Lock()
		for i := 0; i < len(ads); i++ {
			c.Ads[ads[i].ID] = &ads[i]
		}
		for i := 0; i < len(advs); i++ {
			c.Advs[advs[i].ID] = &advs[i]
		}
		for i := 0; i < len(pubs); i++ {
			c.Pubs[pubs[i].ID] = &pubs[i]
		}
		c.Lock.Unlock()
		time.Sleep(1 * time.Minute)
		db.Save(&ads)
		db.Save(&advs)
		db.Save(&pubs)
		for k := range c.Ads {
			delete(c.Ads, k)
		}
		for k := range c.Advs {
			delete(c.Advs, k)
		}
		for k := range c.Pubs {
			delete(c.Pubs, k)
		}
	}
}
