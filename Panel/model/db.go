package model

import "sync"

type Ad struct {
	ID          int    `json:"id"`
	Headline    string `json:"headline"`
	PhotoURL    string `json:"photo_url"`
	Bid         int    `json:"bid"`
	Active      bool   `json:"active"`
	Spent       int    `json:"spent"`
	Impressions int    `json:"impressions"`
	Clicks      int    `json:"clicks"`
	Link        string `json:"link"`
}

type Advertiser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type Publisher struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Balance     int    `json:"balance"`
	Income      int    `json:"income"`
	Impressions int    `json:"impressions"`
	Clicks      int    `json:"clicks"`
}

var Ads = make(map[int]Ad)
var Advertisers = make(map[int]Advertiser)
var Publishers = make(map[int]Publisher)

var mu sync.Mutex
var NextAdID, NextAdvertiserID, NextPublisherID = 1, 1, 1

func Lock() {
	mu.Lock()
}

func Unlock() {
	mu.Unlock()
}
