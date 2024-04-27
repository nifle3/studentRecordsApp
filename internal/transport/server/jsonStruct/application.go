package jsonStruct

import "time"

type ApplicationAdded struct {
	StudentId   string `json:"student_id"`
	ContactInfo string `json:"contact_info"`
	Name        string `json:"name"`
	Text        string `json:"text"`
	File        []byte `json:"file"`
}

type Application struct {
	Id          string
	StudentId   string
	ContactInfo string
	Name        string
	Text        string
	Status      string
	CreatedAt   time.Time
	Link        string
	File        []byte
}
