package controllers

import (
	"flights/models"
	"flights/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	api1_r := r.PathPrefix("/api/v1/").Subrouter()

	InitFlights(api1_r, m.Flights)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	models := models.InitModels(db)

	initControllers(router, models)
	return router
}

func RunRouter(r *mux.Router, port uint16) error {
	c := cors.New(cors.Options{})
	handler := c.Handler(r)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler)
}
