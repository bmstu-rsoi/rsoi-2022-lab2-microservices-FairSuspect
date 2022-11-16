package store

import (
	"http-rest-api/internal/app/apiserver/model"
)

type FlightRepository struct {
	store *Store
}

func (r *FlightRepository) GetAll() ([]*model.Flight, error) {
	Flights := []*model.Flight{}
	query, err := r.store.db.Query("SELECT * from Flights")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		f := &model.Flight{}
		if err := query.Scan(&f.ID, &f.FlightNumber, &f.DateTime, &f.FromAirportID, &f.ToAirportID, &f.FromAirportID, &f.Price); err != nil {
			return nil, err
		}
		Flights = append(Flights, f)
	}
	return Flights, nil
}

func (r *FlightRepository) GetById(id int) (*model.Flight, error) {
	f := &model.Flight{}
	if err := r.store.db.QueryRow("SELECT * from Flights WHERE Id = $1", id).Scan(&f.ID, &f.FlightNumber, &f.DateTime, &f.FromAirportID, &f.ToAirportID, &f.FromAirportID, &f.Price); err != nil {
		return nil, err
	}
	return f, nil
}
