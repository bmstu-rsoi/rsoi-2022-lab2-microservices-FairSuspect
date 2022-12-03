package objects

import (
	_ "encoding/json"
)

type TicketPurchaseRequest struct {
	FlightNumber    string `json:"flightNumber"`
	Price           int    `json:"price"`
	PaidFromBalance bool   `json:"paidFromBalance"`
}

type TicketPurchaseResponse struct {
	TicketUid     string             `json:"ticketUid"`
	FlightNumber  string             `json:"flightNumber"`
	FromAirport   string             `json:"fromAirport"`
	ToAirport     string             `json:"toAirport"`
	Date          string             `json:"date"`
	Status        string             `json:"status"`
	Price         int                `json:"price"`
	PaidByMoney   int                `json:"paidByMoney"`
	PaidByBonuses int                `json:"paidByBonuses"`
	Privilege     PrivilegeShortInfo `json:"privilege"`
}

func NewTicketPurchaseResponse(flight *FlightResponse, ticket *TicketCreateResponse, privilege *AddHistoryResponce) *TicketPurchaseResponse {
	return &TicketPurchaseResponse{
		TicketUid:     ticket.TicketUid,
		FlightNumber:  flight.FlightNumber,
		FromAirport:   flight.FromAirport,
		ToAirport:     flight.ToAirport,
		Date:          flight.Date,
		Status:        ticket.Status,
		Price:         flight.Price,
		PaidByMoney:   privilege.PaidByMoney,
		PaidByBonuses: privilege.PaidByBonuses,
		Privilege:     privilege.Privilege,
	}
}

type TicketCreateRequest struct {
	FlightNumber string `json:"flightNumber" gorm:"not null"`
	Price        int    `json:"price" gorm:"not null"`
}

type TicketCreateResponse struct {
	TicketUid    string `json:"ticketUid"`
	Username     string `json:"username"`
	FlightNumber string `json:"flightNumber"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
}

type Ticket struct {
	TicketUid    string `json:"ticketUid"`
	Username     string `json:"username"`
	FlightNumber string `json:"flightNumber"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
}
type TicketArr []Ticket

type TicketResponse struct {
	TicketUid    string `json:"ticketUid"`
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
}

func MakeTicketResponseArr(tickets []Ticket, flights []FlightResponse) []TicketResponse {
	flight_map := make(map[string]FlightResponse)
	for _, v := range flights {
		flight_map[v.FlightNumber] = v
	}

	data := make([]TicketResponse, len(tickets))
	for k, v := range tickets {
		flight := flight_map[v.FlightNumber]

		data[k] = TicketResponse{
			TicketUid:    v.TicketUid,
			FlightNumber: v.FlightNumber,
			FromAirport:  flight.FromAirport,
			ToAirport:    flight.ToAirport,
			Date:         flight.Date,
			Price:        v.Price,
			Status:       v.Status,
		}
	}
	return data
}

func ToTicketResponce(ticket *Ticket, flight *FlightResponse) *TicketResponse {
	return &TicketResponse{
		TicketUid:    ticket.TicketUid,
		FlightNumber: ticket.FlightNumber,
		FromAirport:  flight.FromAirport,
		ToAirport:    flight.ToAirport,
		Date:         flight.Date,
		Price:        flight.Price,
		Status:       ticket.Status,
	}
}
