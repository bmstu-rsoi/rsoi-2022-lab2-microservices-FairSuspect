package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/errors"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type TicketsM struct {
	client *http.Client

	flights    *FlightsM
	privileges *PrivilegesM
}

func NewTicketsM(client *http.Client, flights *FlightsM) *TicketsM {
	return &TicketsM{
		client:  client,
		flights: flights,
	}
}

func (model *TicketsM) FetchUser(username string) (*objects.UserInfoResponse, error) {
	data := new(objects.UserInfoResponse)
	tickets, err := model.fetch(username)
	if err != nil {
		return nil, err
	}
	flights := model.flights.Fetch(1, 100).Items
	data.Tickets = objects.MakeTicketResponseArr(tickets, flights)

	privilege := model.privileges.Fetch(username)
	data.Privilege = objects.PrivilegeShortInfo{
		Balance: privilege.Balance,
		Status:  privilege.Status,
	}
	return data, nil
}

func (model *TicketsM) fetch(username string) (objects.TicketArr, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/tickets", utils.Config.TicketsEndpoint), nil)
	if username != "" {
		req.Header.Set("X-User-Name", username)
	}
	resp, err := model.client.Do(req)
	if err != nil {
		return nil, err
	}

	data := new(objects.TicketArr)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, data)
	return *data, nil
}

func (model *TicketsM) Fetch() ([]objects.TicketResponse, error) {
	tickets, err := model.fetch("")
	if err != nil {
		return nil, err
	}

	flights := model.flights.Fetch(1, 100).Items
	return objects.MakeTicketResponseArr(tickets, flights), nil
}

func (model *TicketsM) create(flight_number string, price int, username string) (*objects.TicketCreateResponse, error) {
	req_body, _ := json.Marshal(&objects.TicketCreateRequest{FlightNumber: flight_number, Price: price})
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/tickets", utils.Config.TicketsEndpoint), bytes.NewBuffer(req_body))
	req.Header.Add("X-User-Name", username)

	if resp, err := model.client.Do(req); err != nil {
		return nil, err
	} else {
		data := &objects.TicketCreateResponse{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *TicketsM) Create(flight_number string, username string, price int, from_balance bool) (*objects.TicketPurchaseResponse, error) {
	flight, err := model.flights.Find(flight_number)
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	ticket, err := model.create(flight_number, price, username)
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	privilege, err := model.privileges.AddTicket(username, &objects.AddHistoryRequest{
		TicketUID:       ticket.TicketUid,
		Price:           flight.Price,
		PaidFromBalance: from_balance,
	})
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	return objects.NewTicketPurchaseResponse(flight, ticket, privilege), nil
}

func (model *TicketsM) find(ticket_uid string) (*objects.Ticket, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/tickets/%s", utils.Config.TicketsEndpoint, ticket_uid), nil)
	resp, err := model.client.Do(req)
	if err != nil {
		return nil, err
	} else {
		data := &objects.Ticket{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *TicketsM) Find(ticket_uid string, username string) (*objects.TicketResponse, error) {
	ticket, err := model.find(ticket_uid)
	if err != nil {
		return nil, err
	} else if username != ticket.Username {
		return nil, errors.ForbiddenTicket
	}

	flight, err := model.flights.Find(ticket.FlightNumber)
	if err != nil {
		return nil, err
	} else {
		return objects.ToTicketResponce(ticket, flight), nil
	}
}

func (model *TicketsM) delete(ticket_uid string) error {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/tickets/%s", utils.Config.TicketsEndpoint, ticket_uid), nil)
	_, err := model.client.Do(req)
	return err
}

func (model *TicketsM) Delete(ticket_uid string, username string) error {
	ticket, err := model.find(ticket_uid)
	if err != nil {
		return err
	} else if username != ticket.Username {
		return errors.ForbiddenTicket
	}

	if err = model.delete(ticket_uid); err != nil {
		return err
	}

	return model.privileges.DeleteTicket(username, ticket_uid)
}
