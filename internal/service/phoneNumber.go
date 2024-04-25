package service

import (
	"context"
	"fmt"

	"studentRecordsApp/internal/service/entites"
)

type PhoneNumberDb interface {
	GetPhoneNumbers(userId string, ctx context.Context) ([]entities.PhoneNumber, error)
	AddPhoneNumber(number entities.PhoneNumber, ctx context.Context) error
	UpdatePhoneNumber(number entities.PhoneNumber, ctx context.Context) error
	DeletePhoneNumber(id, studentId string, ctx context.Context) error
}

type PhoneNumber struct {
	db *PhoneNumberDb
}

func NewPhoneNumber(db *PhoneNumberDb) PhoneNumber {
	return PhoneNumber{
		db: db,
	}
}

func (p *PhoneNumber) GetAllForUser(userId string, ctx context.Context) ([]entities.PhoneNumber, error) {
	return (*p.db).GetPhoneNumbers(userId, ctx)
}

func (p *PhoneNumber) Add(number entities.PhoneNumber, ctx context.Context) error {
	if !number.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	result, err := number.CheckCorrectNumber()
	if err != nil {
		return err
	}

	if !result {
		return fmt.Errorf("400")
	}

	return (*p.db).AddPhoneNumber(number, ctx)
}

func (p *PhoneNumber) Update(number entities.PhoneNumber, ctx context.Context) error {
	if !number.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	result, err := number.CheckCorrectNumber()
	if err != nil {
		return err
	}

	if !result {
		return fmt.Errorf("400")
	}

	return (*p.db).UpdatePhoneNumber(number, ctx)
}

func (p *PhoneNumber) Delete(number entities.PhoneNumber, ctx context.Context) error {
	return (*p.db).DeletePhoneNumber(number.Id, number.StudentId, ctx)
}
