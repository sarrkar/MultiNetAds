package consumer

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/common/models"
	"github.com/sarrkar/chan-ta-net/reporter/config"
	"github.com/segmentio/kafka-go"
)

type KafkaMessage struct {
	AdID  uint `json:"ad_id"`
	AdvID uint `json:"adv_id"`
	PubID uint `json:"pub_id"`
}

type Consumer struct {
	ClickConsumer      *kafka.Reader
	ImpressionConsumer *kafka.Reader
	Lock               sync.Mutex
	AdsBid             map[uint]int
	AdsClick           map[uint]uint
	AdsImpression      map[uint]uint
	AdvsExpense        map[uint]int
	PubsClick          map[uint]uint
	PubsImpression     map[uint]uint
	PubsIncome         map[uint]int
}

func NewConsumer() *Consumer {
	return &Consumer{
		kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.Config().Client.KafkaUrl},
			GroupID: "reporter",
			Topic:   "click_events",
		}),
		kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.Config().Client.KafkaUrl},
			GroupID: "reporter",
			Topic:   "impression_events",
		}),
		sync.Mutex{},
		make(map[uint]int),
		make(map[uint]uint),
		make(map[uint]uint),
		make(map[uint]int),
		make(map[uint]uint),
		make(map[uint]uint),
		make(map[uint]int),
	}
}

func (c *Consumer) GetClick() {
	for {
		msg, err := c.ClickConsumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
			time.Sleep(config.Config().Client.Retry)
			break
		}
		var km KafkaMessage
		json.Unmarshal(msg.Value, &km)
		log.Printf("consume CLICK at topic: %v partition: %v offset: %v value: %v \n", msg.Topic, msg.Partition, msg.Offset, km)
		c.addClick(km)
	}

	if err := c.ClickConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) addClick(m KafkaMessage) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.AdsClick[m.AdID]++
	c.PubsClick[m.AdID]++
	c.AdvsExpense[m.AdvID] += c.AdsBid[m.AdID]
	c.PubsIncome[m.AdvID] += c.AdsBid[m.AdID]
}

func (c *Consumer) GetImpression() {
	for {
		msg, err := c.ImpressionConsumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
			time.Sleep(config.Config().Client.Retry)
			break
		}
		var km KafkaMessage
		json.Unmarshal(msg.Value, &km)
		log.Printf("consume IMPRRESION at topic: %v partition: %v offset: %v value: %v \n", msg.Topic, msg.Partition, msg.Offset, km)
		c.addImpression(km)
	}

	if err := c.ImpressionConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) addImpression(m KafkaMessage) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.AdsImpression[m.AdID]++
	c.PubsImpression[m.AdID]++
}

func (c *Consumer) GetBidCatch() {
	db := database.GetDb()
	var ads []models.Ad
	if result := db.Find(&ads); result.Error != nil {
		log.Fatalln(result.Error)
	}
	log.Println("Get Ads BID from db")
	c.Lock.Lock()
	defer c.Lock.Unlock()
	for i := 0; i < len(ads); i++ {
		c.AdsBid[ads[i].ID] = ads[i].BID
	}
}

func (c *Consumer) Save() {
	db := database.GetDb()
	var ads []models.Ad
	var advs []models.Advertiser
	var pubs []models.Publisher
	if result := db.Find(&ads); result.Error != nil {
		log.Fatalln(result.Error)
	}
	if result := db.Find(&advs); result.Error != nil {
		log.Fatalln(result.Error)
	}
	if result := db.Find(&pubs); result.Error != nil {
		log.Fatalln(result.Error)
	}
	c.Lock.Lock()
	defer c.Lock.Unlock()
	for i := 0; i < len(ads); i++ {
		ads[i].Click += c.AdsClick[ads[i].ID]
		ads[i].Impression += c.AdsImpression[ads[i].ID]
	}
	for i := 0; i < len(advs); i++ {
		advs[i].Balance -= c.AdvsExpense[advs[i].ID]
	}
	for i := 0; i < len(pubs); i++ {
		pubs[i].Balance += c.PubsIncome[pubs[i].ID] * pubs[i].CommissionPercent / 100
	}

	db.Save(&ads)
	db.Save(&advs)
	db.Save(&pubs)

	log.Println("save changes in db successfully")

}

func (c *Consumer) ResetCatch() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.AdsBid = make(map[uint]int)
	c.AdsClick = make(map[uint]uint)
	c.AdsImpression = make(map[uint]uint)
	c.AdvsExpense = make(map[uint]int)
	c.PubsClick = make(map[uint]uint)
	c.PubsImpression = make(map[uint]uint)
	c.PubsIncome = make(map[uint]int)
}
