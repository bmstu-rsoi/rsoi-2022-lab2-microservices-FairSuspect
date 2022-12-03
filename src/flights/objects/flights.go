package objects

import (
	"fmt"
)

type Airport struct {
	Id      int    `json:"id" gorm:"primary_key;index"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func (Airport) TableName() string {
	return "airport"
}

type Flight struct {
	Id            int     `json:"id" gorm:"primary_key;index"`
	FlightNumber  string  `json:"flightNumber"`
	Datetime      string  `json:"datetime"`
	FromAirport   Airport `json:"fromAirport" gorm:"foreignKey:FromAirportID"`
	ToAirport     Airport `json:"toAirport" gorm:"foreignKey:ToAirportID"`
	FromAirportID int     `gorm:"index"`
	ToAirportID   int     `gorm:"index"`
	Price         int     `json:"price"`
}

func (Flight) TableName() string {
	return "flight"
}

func (flight *Flight) ToFilghtResponse() *FilghtResponse {
	return &FilghtResponse{
		flight.FlightNumber,
		fmt.Sprintf("%s %s", flight.FromAirport.City, flight.FromAirport.Name),
		fmt.Sprintf("%s %s", flight.ToAirport.City, flight.ToAirport.Name),
		flight.Datetime,
		flight.Price,
	}
}

func ToFilghtResponses(flights []Flight) []FilghtResponse {
	resps := make([]FilghtResponse, len(flights))
	for k, v := range flights {
		resps[k] = *v.ToFilghtResponse()
	}
	return resps
}

type FilghtResponse struct {
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
}

type PaginationResponse struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	TotalElements int              `json:"totalElements"`
	Items         []FilghtResponse `json:"items"`
}
