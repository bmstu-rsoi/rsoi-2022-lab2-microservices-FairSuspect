package controllers

import (
	"gateway/controllers/responses"
	"gateway/errors"
	"gateway/models"
	"gateway/objects"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ticketsCtrl struct {
	tickets *models.TicketsM
}

func InitTickets(r *mux.Router, tickets *models.TicketsM) {
	ctrl := &ticketsCtrl{tickets: tickets}
	r.HandleFunc("/me", ctrl.me).Methods(http.MethodGet)
	r.HandleFunc("/tickets", ctrl.fetch).Methods(http.MethodGet)
	r.HandleFunc("/tickets", ctrl.post).Methods(http.MethodPost)
	r.HandleFunc("/tickets/{ticketUid}", ctrl.get).Methods(http.MethodGet)
	r.HandleFunc("/tickets/{ticketUid}", ctrl.delete).Methods(http.MethodDelete)
}

func (ctrl *ticketsCtrl) me(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Name")
	data, err := ctrl.tickets.FetchUser(username)
	if err != nil {
		responses.InternalError(w)
	} else {
		responses.JsonSuccess(w, data)
	}
}

func (ctrl *ticketsCtrl) fetch(w http.ResponseWriter, r *http.Request) {
	data, err := ctrl.tickets.Fetch()
	if err != nil {
		responses.InternalError(w)
	} else {
		responses.JsonSuccess(w, data)
	}
}

func (ctrl *ticketsCtrl) post(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.TicketPurchaseRequest)
	err := json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	data, err := ctrl.tickets.Create(req_body.FlightNumber, r.Header.Get("X-User-Name"), req_body.Price, req_body.PaidFromBalance)
	if err != nil {
		responses.InternalError(w)
	} else {
		responses.JsonSuccess(w, data)
	}
}

func (ctrl *ticketsCtrl) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	ticket_uid := urlParams["ticketUid"]
	username := r.Header.Get("X-User-Name")

	data, err := ctrl.tickets.Find(ticket_uid, username)
	switch err {
	case nil:
		responses.JsonSuccess(w, data)
	case errors.ForbiddenTicket:
		responses.Forbidden(w)
	default:
		responses.RecordNotFound(w, ticket_uid)
	}
}

func (ctrl *ticketsCtrl) delete(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	ticket_uid := urlParams["ticketUid"]
	username := r.Header.Get("X-User-Name")

	err := ctrl.tickets.Delete(ticket_uid, username)
	switch err {
	case nil:
		responses.SuccessTicketDeletion(w)
	case errors.ForbiddenTicket:
		responses.Forbidden(w)
	default:
		responses.RecordNotFound(w, ticket_uid)
	}
}
