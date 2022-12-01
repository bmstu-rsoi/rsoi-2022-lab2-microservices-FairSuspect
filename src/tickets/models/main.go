package models

import (
	"tickets/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Tickets *TicketsM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Tickets = NewTicketsM(repository.NewPGTicketsRep(db))

	return models
}
