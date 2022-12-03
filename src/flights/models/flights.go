package models

import (
	"flights/errors"
	"flights/objects"
	"flights/repository"
)

type FlightsM struct {
	rep repository.FlightsRep
}

func NewFlightsM(rep repository.FlightsRep) *FlightsM {
	return &FlightsM{rep}
}

func (model *FlightsM) GetAll(page int, PageSize int) []objects.Flight {
	return model.rep.GetAll(page, PageSize)
}

func (model *FlightsM) Find(flight_number string) (*objects.Flight, error) {
	flight, err := model.rep.Find(flight_number)
	if err != nil {
		return nil, errors.RecordNotFound
	} else {
		return flight, nil
	}
}
