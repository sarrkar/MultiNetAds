package models

import "gorm.io/gorm"

type Advertiser struct {
	gorm.Model
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	Ads     []Ad   `json:"ads"`
}
