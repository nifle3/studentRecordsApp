package jsonStruct

import "time"

type StudentShort struct {
	Id             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Surname        string `json:"surname"`
	EnrollYear     int    `json:"enroll_year"`
	Specialization string `json:"specialization"`
	Photo          []byte `json:"photo"`
}

type StudentLong struct {
	Id              string    `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Surname         string    `json:"surname"`
	PassportSeria   int       `json:"passport_seria"`
	PassportNumber  int       `json:"passport_number"`
	BirthDate       time.Time `json:"birth_date"`
	Email           string    `json:"email"`
	Country         string    `json:"country"`
	City            string    `json:"city"`
	Street          string    `json:"street"`
	HouseNumber     int       `json:"houseNumber"`
	ApartmentNumber int       `json:"apartmentNumber"`
	EnrollYear      int       `json:"enrollYear"`
	Specialization  string    `json:"specialization"`
	OrderNumber     string    `json:"order_number"`
	LinkPhoto       string    `json:"link_photo "`
	Photo           []byte    `json:"photo"`
}

type StudentLongWithoutLink struct {
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Surname         string    `json:"surname"`
	Password        string    `json:"password"`
	PassportSeria   int       `json:"passport_seria"`
	PassportNumber  int       `json:"passport_number"`
	BirthDate       time.Time `json:"birth_date"`
	Email           string    `json:"email"`
	Country         string    `json:"country"`
	City            string    `json:"city"`
	Street          string    `json:"street"`
	HouseNumber     int       `json:"house_number"`
	ApartmentNumber int       `json:"apartment_number"`
	EnrollYear      int       `json:"enroll_year"`
	Specialization  string    `json:"specialization"`
	OrderNumber     string    `json:"order_number"`
	Photo           []byte    `json:"photo"`
}
