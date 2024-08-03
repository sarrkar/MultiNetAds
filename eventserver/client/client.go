package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sarrkar/chan-ta-net/eventserver/config"
	"github.com/segmentio/kafka-go"
)

type Client interface {
	AddClick(adID, advID, pubID string)
	AddImperession(adID, advID, pubID string)
}

type HttpClient struct{}

func (h HttpClient) AddClick(adID, advID, pubID string) {
	log.Printf("add click %s %s %s \n", adID, advID, pubID)
	resp, err := http.Get(config.Config().Client.PanelApi + "/inc_click/" + adID + "/" + advID + "/" + pubID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func (h HttpClient) AddImperession(adID, advID, pubID string) {
	log.Printf("add impression %s %s %s \n", adID, advID, pubID)
	resp, err := http.Get(config.Config().Client.PanelApi + "/inc_impression/" + adID + "/" + advID + "/" + pubID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

type KafkaClient struct {
	ClickProducer      *kafka.Writer
	ImpressionProducer *kafka.Writer
}

func NewKafkaClinet() *KafkaClient {
	return &KafkaClient{
		kafka.NewWriter(kafka.WriterConfig{Topic: "click_events", Brokers: []string{config.Config().Client.KafkaUrl}}),
		kafka.NewWriter(kafka.WriterConfig{Topic: "impression_events", Brokers: []string{config.Config().Client.KafkaUrl}}),
	}
}

type KafkaMessage struct {
	AdID  uint `json:"ad_id"`
	AdvID uint `json:"adv_id"`
	PubID uint `json:"pub_id"`
}

func (c KafkaClient) AddClick(adID, advID, pubID string) {
	log.Printf("add click KAFKA %s %s %s \n", adID, advID, pubID)
	adIDi, _ := strconv.Atoi(adID)
	advIDi, _ := strconv.Atoi(advID)
	pubIDi, _ := strconv.Atoi(pubID)
	data := KafkaMessage{uint(adIDi), uint(advIDi), uint(pubIDi)}

	b, _ := json.Marshal(data)
	err := c.ClickProducer.WriteMessages(context.Background(), kafka.Message{Value: b})
	if err != nil {
		log.Fatalln(err)
	}
}

func (c KafkaClient) AddImperession(adID, advID, pubID string) {
	log.Printf("add impression KAFKA %s %s %s \n", adID, advID, pubID)
	adIDi, _ := strconv.Atoi(adID)
	advIDi, _ := strconv.Atoi(advID)
	pubIDi, _ := strconv.Atoi(pubID)
	data := KafkaMessage{uint(adIDi), uint(advIDi), uint(pubIDi)}

	b, _ := json.Marshal(data)
	err := c.ImpressionProducer.WriteMessages(context.Background(), kafka.Message{Value: b})
	if err != nil {
		log.Fatalln(err)
	}
}
