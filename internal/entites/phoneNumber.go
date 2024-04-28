package entities

import (
	"regexp"
)

type PhoneNumber struct {
	Id          string
	StudentId   string
	Phone       string
	Description string
}

func (p *PhoneNumber) CheckCorrectNumber() (bool, error) {
	if len(p.Phone) > 12 {
		return false, nil
	}

	reg, err := regexp.Compile(`^(/+7\d{10})|(\d{11})$`)
	if err != nil {
		return false, err
	}

	return reg.MatchString(p.Phone), nil
}

func (p *PhoneNumber) CheckIsNotEmpty() bool {
	return p.StudentId != "" &&
		p.Phone != "" &&
		p.Description != ""
}
