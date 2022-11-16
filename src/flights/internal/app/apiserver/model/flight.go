package model

import "time"

type Flight struct {
	ID            int       `json:"id" gorm:"primary_key"`
	FlightNumber  string    `json:"flight_number" gorm:"not null"`
	DateTime      time.Time `json:"datetime" gorm:"not null"`
	FromAirportID int       `json:"from_airport_id" gorm:"references:AirportID"`
	ToAirportID   int       `json:"to_airport_id" gorm:"references:AirportID"`
	Price         int       `json:"price" gorm:"not null"`
}
