package entities

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
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
	Phones          []PhoneNumber
}

func (s *Student) CheckIsNotEmpty() bool {
	return s.FirstName != "" &&
		s.LastName != "" &&
		s.Surname != "" &&
		s.PassportSeria != 0 &&
		s.PassportNumber != 0 &&
		s.BirthDate != time.Time{} &&
		s.Email != "" &&
		s.Password != "" &&
		s.Country != "" &&
		s.City != "" &&
		s.Street != "" &&
		s.HouseNumber != 0 &&
		s.ApartmentNumber != 0 &&
		s.EnrollYear != 0 &&
		s.Specialization != "" &&
		s.OrderNumber != ""
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

func (s *Student) CheckEmail() (bool, error) {
	regex, err := regexp.Compile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-z]+$`)
	if err != nil {
		return false, err
	}

	return regex.MatchString(s.Email), nil
}

func (s *Student) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), 10)
	if err != nil {
		return err
	}

	s.Password = string(hashPassword)
	return nil
}

func (s *Student) CheckHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
}
