package sqlEntities

import (
	"github.com/google/uuid"
	"time"
)

type Application struct {
	Id          uuid.UUID `db:'id'`
	StudentId   uuid.UUID `db:'student_id''`
	ContactInfo string    `db:'contact_info'`
	Text        string    `db:'application_text'`
	Status      string    `db:'application_status'`
	CreatedAt   time.Time `db:'created_at'`
}
