package sql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) *Storage {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return &Storage{db}
}
