package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name       string  `json:"name"`
	Balance    int     `json:"balance"`
	Commission float64 `json:"Commission"`
}
