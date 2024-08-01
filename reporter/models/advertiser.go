package models

type Advertiser struct {
	ID      uint   `gorm:"column:id; primary_key; not null" json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	Ads     []Ad   `json:"ads"`
}
