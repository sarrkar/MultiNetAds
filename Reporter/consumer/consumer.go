package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sarrkar/chan-ta-net/Reporter/config"
	"github.com/sarrkar/chan-ta-net/Reporter/database"
	"github.com/sarrkar/chan-ta-net/Reporter/models"
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
		addClick(km)
	}

	if err := c.ClickConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func addClick(m KafkaMessage) {
	var ad models.Ad
	var adv models.Advertiser
	var pub models.Publisher

	DB := database.GetDb()
	DB.First(&ad, m.AdID)
	ad.Click++
	DB.Save(&ad)
	DB.First(&adv, m.AdvID)
	adv.Balance -= ad.BID
	DB.Save(&adv)
	DB.First(&pub, m.PubID)
	pub.Balance += (pub.CommissionPercent * ad.BID) / 100
	// pub.Click++
	DB.Save(&pub)
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
		addImpression(km)
	}

	if err := c.ImpressionConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func addImpression(m KafkaMessage) {
	var ad models.Ad
	// var pub models.Publisher

	DB := database.GetDb()
	DB.First(&ad, m.AdID)
	ad.Impression++
	DB.Save(&ad)
	// DB.First(&pub, m.PubID)
	// pub.Impression++
	// DB.Save(&pub)
}
