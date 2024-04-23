package entities

import (
	"time"
)

type Application struct {
	Id          string
	StudentId   string
	ContactInfo string
	Text        string
	Status      string
	CreatedAt   time.Time
}
