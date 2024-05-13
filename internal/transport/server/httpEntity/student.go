package httpEntity

import (
	"time"
)

type StudentSelf struct {
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Surname         string             `json:"surname"`
	PassportSeria   int                `json:"passport_seria"`
	PassportNumber  int                `json:"passport_number"`
	BirthDate       time.Time          `json:"birth_date"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Country         string             `json:"country"`
	City            string             `json:"city"`
	Street          string             `json:"street"`
	HouseNumber     int                `json:"house_number"`
	ApartmentNumber int                `json:"apartment_number"`
	EnrollYear      time.Time          `json:"enroll_year"`
	Specialization  string             `json:"specialization"`
	Course          int                `json:"course"`
	Group           int                `json:"group"`
	PhoneNumbers    []PhoneNumberShort `json:"phoneNumbers"`
}

type Student struct {
	Id              string             `json:"id"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Surname         string             `json:"surname"`
	PassportSeria   int                `json:"passport_seria"`
	PassportNumber  int                `json:"passport_number"`
	BirthDate       time.Time          `json:"birth_date"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Country         string             `json:"country"`
	City            string             `json:"city"`
	Street          string             `json:"street"`
	HouseNumber     int                `json:"house_number"`
	ApartmentNumber int                `json:"apartment_number"`
	EnrollYear      time.Time          `json:"enroll_year"`
	Specialization  string             `json:"specialization"`
	Course          int                `json:"course"`
	Group           int                `json:"group"`
	Link            string             `json:"link"`
	PhoneNumbers    []PhoneNumberShort `json:"phoneNumbers"`
}

type StudentShort struct {
	Id             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Surname        string    `json:"surname"`
	EnrollYear     time.Time `json:"enroll_year"`
	Specialization string    `json:"specialization"`
	Course         int       `json:"course"`
	Group          int       `json:"group"`
}

type StudentEnter struct {
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Surname         string             `json:"surname"`
	PassportSeria   int                `json:"passport_seria"`
	PassportNumber  int                `json:"passport_number"`
	BirthDate       time.Time          `json:"birth_date"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Country         string             `json:"country"`
	City            string             `json:"city"`
	Street          string             `json:"street"`
	HouseNumber     int                `json:"house_number"`
	ApartmentNumber int                `json:"apartment_number"`
	Specialization  string             `json:"specialization"`
	Course          int                `json:"course"`
	Group           int                `json:"group"`
	PhoneNumbers    []PhoneNumberShort `json:"phoneNumbers"`
}

type PhoneNumberShort struct {
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type PhoneNumber struct {
	Id          string `json:"id"`
	StudentId   string `json:"student_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
