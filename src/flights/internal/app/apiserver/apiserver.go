package apiserver

import (
	"encoding/json"
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
	s.logger.Info("starting api server")

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

	s.router.HandleFunc("/api/v1/flights", s.Flights())
	s.router.HandleFunc("/api/v1/flights/{id:[0-9]+}", s.handleFlightId())

}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err

	}
	s.store = st

	return nil
}
func (s *APIServer) Flights() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			persons, err := s.store.Flight().GetAll()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			j, err := json.MarshalIndent(persons, "", "\t")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			io.WriteString(w, string(j)+"\n")
			break

		default:
			refuseMethod(w)

		}

	}
}
func (s *APIServer) handleFlightId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":

			person, err := s.store.Flight().GetById(id)
			if err != nil {
				http.Error(w, err.Error(), 404)
				return
			}
			j, err := json.Marshal(person)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string(j)+"\n")
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
