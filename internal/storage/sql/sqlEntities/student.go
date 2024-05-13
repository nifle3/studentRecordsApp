package sqlEntities

import (
	"github.com/google/uuid"
	"time"
)

type Student struct {
	Id              uuid.UUID `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	Surname         string    `db:"surname"`
	PassportSeria   int       `db:"passport_seria"`
	PassportNumber  int       `db:"passport_number"`
	BirthDate       time.Time `db:"birth_date"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	Country         string    `db:"country"`
	City            string    `db:"city"`
	Street          string    `db:"street"`
	HouseNumber     int       `db:"house"`
	ApartmentNumber int       `db:"apartment"`
	EnrollYear      time.Time `db:"enroll_year"`
	Specialization  string    `db:"specialization"`
	LinkPhoto       string    `db:"link_photo"`
	Group           int       `db:"_group"`
	Course          int       `db:"course"`
}
