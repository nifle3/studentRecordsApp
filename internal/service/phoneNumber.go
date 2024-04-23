package service

import (
	"context"
	"studentRecordsApp/internal/service/entites"
)

type PhoneNumberDb interface {
	GetPhoneNumbers(userId string, ctx context.Context) ([]entities.PhoneNumber, error)
	AddPhoneNumber(number entities.PhoneNumber, ctx context.Context) error
	UpdatePhoneNumber(number entities.PhoneNumber, ctx context.Context) error
	DeletePhoneNumber(number entities.PhoneNumber, ctx context.Context) error
}

type PhoneNumber struct {
	db *PhoneNumberDb
}
