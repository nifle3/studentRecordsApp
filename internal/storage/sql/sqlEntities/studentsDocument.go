package sqlEntities

import (
	"github.com/google/uuid"
	"time"
)

type StudentsDocument struct {
	Id        uuid.UUID `db:"id"`
	StudentId uuid.UUID `db:"student_id"`
	Name      string    `db:"_name"`
	Type      string    `db:"_type"`
	Link      string    `db:"link"`
	CreatedAt time.Time `db:"created_at"`
}
