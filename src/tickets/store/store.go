package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ..
)

type Store struct {
	config           *Config
	db               *sql.DB
	TicketRepository *TicketRepository
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

func (s *Store) Ticket() *TicketRepository {
	if s.TicketRepository != nil {
		return s.TicketRepository
	}
	s.TicketRepository = &TicketRepository{store: s}

	return s.TicketRepository
}
