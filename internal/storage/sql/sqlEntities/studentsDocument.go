package sqlEntities

import (
	"github.com/google/uuid"
	"time"
)

type StudentsDocument struct {
	Id        uuid.UUID `db:"id"`
	StudentId uuid.UUID `db:"student_id"`
	Name      string    `db:"document_name"`
	Type      string    `db:"document_type"`
	Link      string    `db:"document_link_s3"`
	CreatedAt time.Time `db:"created_at"`
}
