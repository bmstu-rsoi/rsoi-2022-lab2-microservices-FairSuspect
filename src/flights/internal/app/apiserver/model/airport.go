package model

type Airport struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}
