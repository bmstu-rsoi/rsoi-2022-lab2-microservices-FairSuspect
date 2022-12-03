package objects

import (
	_ "encoding/json"
)

type FlightResponse struct {
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
	Items         []FlightResponse `json:"items"`
}
