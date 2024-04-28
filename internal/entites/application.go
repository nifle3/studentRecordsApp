package entities

import (
	"time"
)

type Application struct {
	Id          string
	StudentId   string
	ContactInfo string
	Name        string
	Text        string
	// Status must be "Создан" or "Закрыт"
	Status    string
	CreatedAt time.Time
	Link      string
	File      []byte
}

func (a *Application) CheckIsNotEmpty() bool {
	return a.StudentId != "" &&
		a.ContactInfo != "" &&
		a.Text != "" &&
		a.Status != ""
}

func (a *Application) CheckStatus() bool {
	return a.Status == ApplicationClosed || a.Status == ApplicationCreated
}
