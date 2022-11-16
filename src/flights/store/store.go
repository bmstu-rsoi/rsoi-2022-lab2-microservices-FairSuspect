package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ..
)

type Store struct {
	config            *Config
	db                *sql.DB
	airportRepository *AirportRepository
	flightRepository  *FlightRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {

	log.Default().Println("Connecting to db... with " + s.config.DatabaseURL)

	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	// defer db.Close()
	return nil

}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Airport() *AirportRepository {
	if s.airportRepository != nil {
		return s.airportRepository
	}
	s.airportRepository = &AirportRepository{store: s}

	return s.airportRepository
}
func (s *Store) Flight() *FlightRepository {
	if s.flightRepository != nil {
		return s.flightRepository
	}
	s.flightRepository = &FlightRepository{store: s}

	return s.flightRepository
}
