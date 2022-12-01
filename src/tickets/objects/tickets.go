package objects

type Ticket struct {
	Id           int    `gorm:"primary_key;index"`
	TicketUid    string `json:"ticketUid" gorm:"not null;unique"`
	Username     string `json:"username" gorm:"not null"`
	FlightNumber string `json:"flightNumber" gorm:"not null"`
	Price        int    `json:"price" gorm:"not null"`
	Status       string `json:"status" gorm:"not null"`
}

func (Ticket) TableName() string {
	return "ticket"
}

type CreateRequest struct {
	FlightNumber string `json:"flightNumber" gorm:"not null"`
	Price        int    `json:"price" gorm:"not null"`
}
