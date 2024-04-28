package sql

import (
	"context"
	"log"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/pkg/password"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/service"
)

var (
	_ service.StudentDb     = (*Storage)(nil)
	_ service.DocumentDb    = (*Storage)(nil)
	_ service.PhoneNumberDb = (*Storage)(nil)
	_ service.UserDb        = (*Storage)(nil)
	_ service.ApplicationDb = (*Storage)(nil)
)

type Storage struct {
	db *sqlx.DB
}

func New(connectionString string) (*Storage, error) {
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = startData(db)
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func startData(db *sqlx.DB) error {
	var exist bool
	err := db.GetContext(context.Background(), &exist, `SELECT EXISTS(SELECT * FROM Users)`)

	if err != nil {
		return err
	}

	log.Printf("Data in users is %#v", exist)

	if exist {
		return nil
	}

	log.Printf("CREATING start data\n")

	pass, err := password.Hash("qwe")
	if err != nil {
		return err
	}

	_, err = db.ExecContext(context.Background(),
		`INSERT INTO Users VALUES ($1,$2,$3,$4,$5,$6,$7),($8,$9,$10,$11,$12,$13,$14),($15,$16,$17,$18,$19,$20,$21)`,
		uuid.New(), "Артём", "Куприянов", "Сергеевич", "nifle3@gmail.com", pass, entities.UserAdmin,
		uuid.New(), "Раниль", "Закиров", "Ильдусович", "homya@gmail.com", pass, entities.UserWorker,
		uuid.New(), "Антон", "Яковлев", "Дмитриевич", "mersya@gmail.com", pass, entities.UserWorker)

	return err
}
