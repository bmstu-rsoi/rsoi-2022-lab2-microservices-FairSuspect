package models

import (
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type FlightsM struct {
	client *http.Client
}

func (model *FlightsM) Find(FlightNumber string) (*objects.FlightResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/flights/%s", utils.Config.FlightsEndpoint, FlightNumber), nil)
	response, err := model.client.Do(request)
	if err != nil {
		return nil, err
	} else {
		data := &objects.FlightResponse{}
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *FlightsM) Fetch(page int, PageSize int) *objects.PaginationResponse {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/flights", utils.Config.FlightsEndpoint), nil)
	q := request.URL.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("size", fmt.Sprintf("%d", PageSize))
	request.URL.RawQuery = q.Encode()
	response, err := model.client.Do(request)
	if err != nil {
		panic("Http request failed\n")
	}

	data := &objects.PaginationResponse{}
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, data)
	return data
}
func NewFlightsM(client *http.Client) *FlightsM {
	return &FlightsM{client: client}
}
