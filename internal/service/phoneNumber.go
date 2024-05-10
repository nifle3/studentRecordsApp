package service

import (
	"context"
	"net/http"
	"studentRecordsApp/pkg/customError"

	"github.com/google/uuid"

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

func NewPhoneNumber(db PhoneNumberDb) PhoneNumber {
	return PhoneNumber{
		db: &db,
	}
}

func (p *PhoneNumber) GetAllForUser(userId string, ctx context.Context) ([]entities.PhoneNumber, error) {
	return (*p.db).GetPhoneNumbers(userId, ctx)
}

func (p *PhoneNumber) Add(number entities.PhoneNumber, ctx context.Context) error {
	if !number.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	if err := number.CheckCorrectNumber(); err != nil {
		return customError.New(http.StatusBadRequest, "Has an invalid number")
	}

	number.Id = uuid.New()

	return (*p.db).AddPhoneNumber(number, ctx)
}

func (p *PhoneNumber) Update(number entities.PhoneNumber, ctx context.Context) error {
	if !number.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	if err := number.CheckCorrectNumber(); err != nil {
		return customError.New(http.StatusBadRequest, "Has an invalid number")
	}

	return (*p.db).UpdatePhoneNumber(number, ctx)
}

func (p *PhoneNumber) Delete(id, userId string, ctx context.Context) error {
	return (*p.db).DeletePhoneNumber(id, userId, ctx)
}
