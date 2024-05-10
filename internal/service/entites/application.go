package entities

import (
	"github.com/google/uuid"
	"io"
	"time"
)

type Application struct {
	Id          uuid.UUID
	StudentId   uuid.UUID
	ContactInfo string
	Name        string
	Text        string
	// Status must be "Создан" or "Закрыт"
	Status    string
	CreatedAt time.Time
	Link      string
	File      io.Reader
}

func (a *Application) CheckIsNotEmpty() bool {
	return a.StudentId != uuid.Nil &&
		a.ContactInfo != "" &&
		a.Text != "" &&
		a.Status != ""
}

func (a *Application) CheckStatus() bool {
	return a.Status == ApplicationClosed || a.Status == ApplicationCreated
}
