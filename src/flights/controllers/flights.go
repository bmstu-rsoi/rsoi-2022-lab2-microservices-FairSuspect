package controllers

import (
	"flights/controllers/responses"
	"flights/errors"
	"flights/models"
	"flights/objects"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type filghtCtrl struct {
	model *models.FlightsM
}

func InitFlights(r *mux.Router, model *models.FlightsM) {
	ctrl := &filghtCtrl{model}
	r.HandleFunc("/flights", ctrl.getAll).Methods(http.MethodGet)
	r.HandleFunc("/flights/{flightNumber}", ctrl.get).Methods(http.MethodGet)
}

func (ctrl *filghtCtrl) getAll(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	PageSize, _ := strconv.Atoi(queryParams.Get("size"))
	items := ctrl.model.GetAll(page, PageSize)

	data := &objects.PaginationResponse{
		Page:          page,
		PageSize:      PageSize,
		TotalElements: len(items),
		Items:         objects.ToFilghtResponses(items),
	}

	responses.JsonSuccess(w, data)
}

func (ctrl *filghtCtrl) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	flight_number := urlParams["flightNumber"]

	data, err := ctrl.model.Find(flight_number)
	switch err {
	case nil:
		responses.JsonSuccess(w, data.ToFilghtResponse())
	case errors.RecordNotFound:
		responses.RecordNotFound(w, flight_number)
	default:
		responses.InternalError(w)
	}
}
