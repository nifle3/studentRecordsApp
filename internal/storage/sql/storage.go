package sql

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"studentRecordsApp/internal/service"
	"studentRecordsApp/internal/service/entites"
)

var (
	_ service.StudentDb     = (*Storage)(nil)
	_ service.DocumentDb    = (*Storage)(nil)
	_ service.PhoneNumberDb = (*Storage)(nil)
	_ service.UserDb        = (*Storage)(nil)
	_ service.ApplicationDb = (*Storage)(nil)
)

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

	err = startData(db)
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func startData(db *sql.DB) error {
	var exist bool
	err := db.QueryRowContext(context.Background(), `SELECT EXISTS(SELECT * FROM Users)`).Scan(&exist)

	if err != nil {
		return err
	}

	log.Printf("Data in users is %#v", exist)

	if exist {
		return nil
	}

	log.Printf("CREATING start data\n")
	user1 := entities.User{Password: "qwe123"}
	user1.HashPassword()

	_, err = db.ExecContext(context.Background(),
		`INSERT INTO Users VALUES ($1,$2,$3,$4,$5,$6,$7,$8),($9,$10,$11,$12,$13,$14,$15,$16),($17,$18,$19,$20,$21,$22,$23,$24)`,
		uuid.New(), "Артём", "Куприянов", "Сергеевич", "nifle3@gmail.com", user1.Password, "123543123", entities.UserAdmin,
		uuid.New(), "Раниль", "Закиров", "Ильдусович", "homya@gmail.com", user1.Password, "123-45123", entities.UserWorker,
		uuid.New(), "Антон", "Яковлев", "Дмитриевич", "mersya@gmail.com", user1.Password, "123-123123", entities.UserWorker)

	return err
}
