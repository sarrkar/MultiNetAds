package models

type Publisher struct {
	ID         int     `gorm:"column:id; primary_key; not null" json:"id"`
	Name       string  `json:"name"`
	Balance    int     `json:"balance"`
	Commission float64 `json:"Commission"`
}
