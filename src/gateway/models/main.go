package models

import "net/http"

type Models struct {
	Flights    *FlightsM
	Privileges *PrivilegesM
	Tickets    *TicketsM
}

func InitModels() *Models {
	models := new(Models)
	client := &http.Client{}

	models.Flights = NewFlightsM(client)
	models.Privileges = NewPrivilegesM(client)
	models.Tickets = NewTicketsM(client, models.Flights)

	return models
}
