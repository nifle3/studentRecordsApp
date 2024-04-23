package entities

import "time"

type StudentsDocument struct {
	Id        string
	StudentId string
	Name      string
	Type      string
	Link      string
	CreatedAt time.Time
}
