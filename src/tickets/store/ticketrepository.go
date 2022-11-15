package store

import (
	"http-rest-api/internal/app/apiserver/model"
)

type TicketRepository struct {
	store *Store
}

func (r *TicketRepository) Create(t *model.Ticket) (int, error) {
	id := -1
	err := r.store.db.QueryRow("INSERT INTO Tickets (flight_number, price) VALUES ($1, $2) RETURNING Id", t.FlightNumber, t.Price).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
func (r *TicketRepository) GetAll() ([]*model.Ticket, error) {
	Tickets := []*model.Ticket{}
	query, err := r.store.db.Query("SELECT * from Tickets")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		t := &model.Ticket{}
		if err := query.Scan(&t.Id, &t.TicketUid, &t.UserName, &t.FlightNumber, &t.Price, &t.Status); err != nil {
			return nil, err
		}
		Tickets = append(Tickets, t)
	}
	return Tickets, nil
}

func (r *TicketRepository) GetById(id int) (*model.Ticket, error) {
	t := &model.Ticket{}
	if err := r.store.db.QueryRow("SELECT * from Tickets WHERE Id = $1", id).Scan(&t.Id, &t.TicketUid, &t.UserName, &t.FlightNumber, &t.Price, &t.Status); err != nil {
		return nil, err
	}
	return t, nil
}
func (r *TicketRepository) DeleteById(id int) (*model.Ticket, error) {
	p := &model.Ticket{}
	if _, err := r.store.db.Query("DELETE from Tickets WHERE Id = $1", id); err != nil {
		return nil, err
	}
	return p, nil
}
