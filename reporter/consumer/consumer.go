package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

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
	Ads                map[uint]*models.Ad
	Advs               map[uint]*models.Advertiser
	Pubs               map[uint]*models.Publisher
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
		make(map[uint]*models.Ad),
		make(map[uint]*models.Advertiser),
		make(map[uint]*models.Publisher),
	}
}

func (c *Consumer) GetClick() {
	for {
		msg, err := c.ClickConsumer.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("consume CLICK at topic: %v partition: %v offset: %v value: %s \n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		var km KafkaMessage
		json.Unmarshal(msg.Value, &km)
		c.addClick(km)
	}

	if err := c.ClickConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) addClick(m KafkaMessage) {
	c.Ads[m.AdID].Click++
	c.Advs[m.AdvID].Balance -= c.Ads[m.AdID].BID
	c.Pubs[m.PubID].Balance += (c.Pubs[m.PubID].CommissionPercent * c.Ads[m.AdID].BID) / 100
	c.Pubs[m.PubID].Click++
}

func (c *Consumer) GetImpression() {
	for {
		msg, err := c.ImpressionConsumer.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("consume IMPRRESION at topic: %v partition: %v offset: %v value: %s \n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		var km KafkaMessage
		json.Unmarshal(msg.Value, &km)
		c.addImpression(km)
	}

	if err := c.ImpressionConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) addImpression(m KafkaMessage) {
	c.Ads[m.AdID].Impression++
	c.Pubs[m.PubID].Impression++
}
