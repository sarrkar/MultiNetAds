package models

type Publisher struct {
	ID                uint   `gorm:"column:id; primary_key; not null" json:"id"`
	Name              string `json:"name"`
	Balance           int    `json:"balance"`
	CommissionPercent int    `json:"Commission_percent"`
	Category          string `json:"category"`
}
