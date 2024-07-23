package models

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Title        string `json:"title"`
	ImageUrl     string `json:"image_url"`
	RedirectUrl  string `json:"redirect_url"`
	BID          uint   `json:"bid"`
	Active       bool   `json:"active"`
	Impression   uint   `json:"impression"`
	Click        uint   `json:"click"`
	AdvertiserID uint   `json:"advertiser_id"`
}
