package sql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"studentRecordsApp/internal/service"
)

var _ service.StudentDb = (*Storage)(nil)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}
