package controllers

import (
	"gateway/controllers/responses"
	"gateway/models"

	"net/http"

	"github.com/gorilla/mux"
)

type privilegeCtrl struct {
	privileges *models.PrivilegesM
}

func InitPrivileges(r *mux.Router, privileges *models.PrivilegesM) {
	ctrl := &privilegeCtrl{privileges}
	r.HandleFunc("/privilege", ctrl.fetch).Methods(http.MethodGet)
}

func (ctrl *privilegeCtrl) fetch(w http.ResponseWriter, r *http.Request) {
	data := ctrl.privileges.Fetch(r.Header.Get("X-User-Name"))
	responses.JsonSuccess(w, data)
}
