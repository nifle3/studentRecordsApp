package entities

import (
	"strings"
	"time"
)

type Document struct {
	Id        string
	StudentId string
	Name      string
	Type      string
	Link      string
	File      []byte
	CreatedAt time.Time
}

func (d *Document) CheckIsNotEmpty() bool {
	return strings.Compare(d.StudentId, "") != 0 &&
		strings.Compare(d.Name, "") != 0 &&
		strings.Compare(d.Type, "") != 0
}
