package entities

import (
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id        uuid.UUID
	StudentId uuid.UUID
	Name      string
	Type      string
	Link      string
	File      io.Reader
	CreatedAt time.Time
}

func (d *Document) CheckIsNotEmpty() bool {
	return d.StudentId != uuid.Nil &&
		d.Name != "" && !strings.Contains(d.Name, " ") &&
		d.Type != "" && !strings.Contains(d.Type, " ")
}
