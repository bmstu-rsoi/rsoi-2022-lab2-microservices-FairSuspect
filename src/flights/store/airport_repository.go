package store

import (
	"http-rest-api/internal/app/apiserver/model"
)

type AirportRepository struct {
	store *Store
}

func (r *AirportRepository) GetAll() ([]*model.Airport, error) {
	airports := []*model.Airport{}
	query, err := r.store.db.Query("SELECT * from Airports")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		a := &model.Airport{}
		if err := query.Scan(&a.ID, &a.Name, &a.City, &a.Country); err != nil {
			return nil, err
		}
		airports = append(airports, a)
	}
	return airports, nil
}

func (r *AirportRepository) GetById(id int) (*model.Airport, error) {
	a := &model.Airport{}
	if err := r.store.db.QueryRow("SELECT * from Airports WHERE Id = $1", id).Scan(&a.ID, &a.Name, &a.City, &a.Country); err != nil {
		return nil, err
	}
	return a, nil
}
