package jsonStruct

import "time"

type Document struct {
	Id        string    `json:"id"`
	StudentId string    `json:"student_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Link      string    `json:"link"`
	File      []byte    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentWithoutFile struct {
	Id        string    `json:"id"`
	StudentId string    `json:"student_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentForAdded struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	File      []byte `json:"file"`
}
