package repository

import (
	"flights/errors"
	"flights/objects"

	"github.com/jinzhu/gorm"
)

type FlightsRep interface {
	GetAll(page int, PageSize int) []objects.Flight
	Find(flight_number string) (*objects.Flight, error)
}

type FlightsRepImpl struct {
	db *gorm.DB
}

func NewPGFlightsRep(db *gorm.DB) *FlightsRepImpl {
	return &FlightsRepImpl{db}
}

func paginate(page int, PageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offet := (page - 1) * PageSize
		return db.Offset(offet).Limit(PageSize)
	}
}

func (rep *FlightsRepImpl) GetAll(page int, PageSize int) []objects.Flight {
	flights := []objects.Flight{}
	rep.db.
		Scopes(paginate(page, PageSize)).
		Model(&objects.Flight{}).
		Preload("FromAirport").
		Preload("ToAirport").
		Find(&flights)

	return flights
}

func (rep *FlightsRepImpl) Find(FlightNumber string) (*objects.Flight, error) {
	temp := new(objects.Flight)
	err := rep.db.
		Where(&objects.Flight{FlightNumber: FlightNumber}).
		Preload("FromAirport").
		Preload("ToAirport").
		First(temp).
		Error
	switch err {
	case gorm.ErrRecordNotFound:
		temp, err = nil, errors.RecordNotFound
	case nil:
		break
	default:
		temp, err = nil, errors.UnknownError
	}

	return temp, err
}
