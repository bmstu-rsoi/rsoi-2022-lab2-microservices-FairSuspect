package models

import (
	"flights/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Flights *FlightsM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Flights = NewFlightsM(repository.NewPGFlightsRep(db))

	return models
}
