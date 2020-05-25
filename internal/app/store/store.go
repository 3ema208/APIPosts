package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Config store
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig return new config
func NewConfig() *Config {
	return &Config{}
}

// New ..
func New(config *Config) *Store {
	return &Store{config: config}
}

// Store ..
type Store struct {
	config    *Config
	db        *sql.DB
	postStore *PostStore
}

// Open ..
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if errP := db.Ping(); errP != nil {
		return errP
	}
	s.db = db
	return nil
}

// Close ..
func (s *Store) Close() {
	s.db.Close()
}

// Post return object
func (s *Store) Post() *PostStore {
	if s.postStore != nil {
		return s.postStore
	}
	return &PostStore{
		store: s,
	}
}
