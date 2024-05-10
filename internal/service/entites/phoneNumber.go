package entities

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type PhoneNumber struct {
	Id          uuid.UUID
	StudentId   uuid.UUID
	Phone       string
	Description string
}

func (p *PhoneNumber) CheckCorrectNumber() error {
	if len(p.Phone) > 12 {
		return errors.New("phone have invalid length")
	}

	reg, err := regexp.Compile(`^(/+7\d{10})|(\d{11})$`)
	if err != nil {
		return err
	}

	if reg.MatchString(p.Phone) {
		return errors.New("invalid phone number")
	}

	return nil
}

func (p *PhoneNumber) CheckIsNotEmpty() bool {
	return p.StudentId != uuid.Nil &&
		p.Phone != "" && !strings.Contains(p.Phone, " ") &&
		p.Description != "" && !strings.Contains(p.Description, " ")
}
