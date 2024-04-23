package entities

import (
	"time"
)

type Student struct {
	Id              string
	FirstName       string
	LastName        string
	Surname         string
	PassportSeria   int
	PassportNumber  int
	BirthDate       time.Time
	Email           string
	Password        string
	Country         string
	City            string
	Street          string
	HouseNumber     int
	ApartmentNumber int
	EnrollYear      int
	Specialization  string
	OrderNumber     string
}
