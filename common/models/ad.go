package models

type Ad struct {
	ID           uint   `gorm:"column:id; primary_key; not null" json:"id"`
	Title        string `json:"title"`
	ImageUrl     string `json:"image_url"`
	RedirectUrl  string `json:"redirect_url"`
	BID          int    `gorm:";column:bid" json:"bid"`
	Active       bool   `json:"active"`
	Impression   uint   `json:"impression"`
	Click        uint   `json:"click"`
	AdvertiserID uint   `json:"advertiser_id"`
	Category     string `json:"category"`
}
