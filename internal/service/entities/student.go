package entities

import (
	"io"
	"strings"
	"studentRecordsApp/pkg/email"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	Id              uuid.UUID
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
	EnrollYear      time.Time
	Specialization  string
	Course          int
	Group           int
	LinkPhoto       string
	Photo           io.Reader
	PhoneNumbers    []PhoneNumber
}

func (s *Student) CheckIsNotEmpty() bool {
	return s.FirstName != "" && !strings.Contains(s.FirstName, " ") &&
		s.LastName != "" && !strings.Contains(s.LastName, " ") &&
		s.Surname != "" && !strings.Contains(s.Surname, " ") &&
		s.PassportSeria != 0 &&
		s.PassportNumber != 0 &&
		s.Email != "" && !strings.Contains(s.Email, " ") &&
		s.Password != "" && !strings.Contains(s.Password, " ") &&
		s.Country != "" && !strings.Contains(s.Country, " ") &&
		s.City != "" && !strings.Contains(s.City, " ") &&
		s.Street != "" &&
		s.HouseNumber != 0 &&
		s.ApartmentNumber != 0 &&
		s.Specialization != "" && !strings.Contains(s.Specialization, " ")

}

func (s *Student) CheckBirthdate() bool {
	return s.BirthDate.Before(time.Now().AddDate(-16, 0, 0))
}

func (s *Student) CheckPassportSeria() bool {
	return s.PassportSeria > 999 && s.PassportSeria < 10000
}

func (s *Student) CheckNumber() bool {
	return s.PassportNumber > 99999 && s.PassportNumber < 1000000
}

func (s *Student) CheckPassword() bool {
	return len(s.Password) >= 6 && !strings.Contains(s.Password, " ")
}

func (s *Student) CheckEmail() bool {
	return email.IsCorrect(s.Email) == nil
}
