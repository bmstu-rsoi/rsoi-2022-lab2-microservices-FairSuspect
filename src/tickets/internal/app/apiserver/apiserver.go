package apiserver

import (
	"encoding/json"
	"fmt"
	"http-rest-api/internal/app/apiserver/model"
	"http-rest-api/store"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}

// Start ..
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.router.Use(commonMiddleware)
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting ticket api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil

}

func (s *APIServer) configureRouter() {

	s.router.HandleFunc("/api/v1/tickets", s.handleTickets())
	s.router.HandleFunc("/api/v1/tickets/{id:[0-9]+}", s.handleTicketsId())

}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err

	}
	s.store = st

	return nil
}
func (s *APIServer) handleTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			Tickets, err := s.store.Ticket().GetAll()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			j, err := json.MarshalIndent(Tickets, "", "\t")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			io.WriteString(w, string(j)+"\n")
			break
		case http.MethodPost:
			var p *model.Ticket

			// Try to decode the request body into the struct. If there is an error,
			// respond to the client with the error message and a 400 status code.
			if r.Body == nil {
				http.Error(w, "Empty body", http.StatusBadRequest)
				return
			}
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				io.WriteString(w, "Failed to parse model: "+err.Error())
				return
			}
			id, err := s.store.Ticket().Create(p)
			if err != nil {
				io.WriteString(w, "Failed to create model \n"+err.Error())
				return
			}
			w.Header().Add("Location", fmt.Sprintf("/api/v1/Tickets/%d", id))
			w.WriteHeader(http.StatusCreated)

			break
		default:
			refuseMethod(w)

		}

	}
}
func (s *APIServer) handleTicketsId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:

			Ticket, err := s.store.Ticket().GetById(id)
			if err != nil {
				http.Error(w, err.Error(), 404)
				return
			}
			j, err := json.Marshal(Ticket)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string(j)+"\n")
			break
		case http.MethodDelete:

			_, err := s.store.Ticket().DeleteById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)

			break

		}

	}
}

func refuseMethod(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", 405)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
