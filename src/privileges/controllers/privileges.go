package controllers

import (
	"encoding/json"
	"privileges/controllers/responses"
	"privileges/models"
	"privileges/objects"

	"net/http"

	"github.com/gorilla/mux"
)

type privilegesCtrl struct {
	privileges *models.PrivilegesM
}

func InitPrivileges(r *mux.Router, privileges *models.PrivilegesM) {
	ctrl := &privilegesCtrl{privileges}
	r.HandleFunc("/privilege", ctrl.get).Methods(http.MethodGet)
	r.HandleFunc("/history", ctrl.addTicket).Methods(http.MethodPost)
	r.HandleFunc("/history/{ticketUid}", ctrl.deleteTicket).Methods(http.MethodDelete)
}

func (ctrl *privilegesCtrl) get(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Name")

	privilege, history, _ := ctrl.privileges.Find(username)

	data := objects.ToPrivilegeInfoResponse(privilege, history)
	responses.JsonSuccess(w, data)
}

func (ctrl *privilegesCtrl) addTicket(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.AddTicketRequest)
	err := json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	username := r.Header.Get("X-User-Name")

	data, err := ctrl.privileges.AddTicket(username, req_body)
	switch err {
	case nil:
		responses.JsonSuccess(w, data)
	default:
		responses.BadRequest(w, err.Error())
	}
}

func (ctrl *privilegesCtrl) deleteTicket(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	ticket_uid := urlParams["ticketUid"]
	username := r.Header.Get("X-User-Name")

	err := ctrl.privileges.DeleteTicket(username, ticket_uid)
	switch err {
	case nil:
		responses.SuccessTicketDeletion(w)
	default:
		responses.BadRequest(w, err.Error())
	}
}
