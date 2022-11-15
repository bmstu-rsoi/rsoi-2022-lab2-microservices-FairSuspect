package model

import (
	"github.com/google/uuid"
)

type Ticket struct {
	Id           int       `json:"id" gorm:"primary_key"`
	TicketUid    uuid.UUID `json:"ticket_uid" gorm:"not null"`
	UserName     string    `json:"username" gorm:"not null"`
	FlightNumber string    `json:"flight_number" gorm:"not null"`
	Price        int       `json:"price" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null"`
}
