package sqlEntities

import (
	"github.com/google/uuid"
	"time"
)

type Application struct {
	Id          uuid.UUID `db:"i"`
	StudentId   uuid.UUID `db:"student_id"`
	ContactInfo string    `db:"contact_inf"`
	Name        string    `db:"application_name"`
	Text        string    `db:"application_tex"`
	Status      string    `db:"application_status"`
	CreatedAt   time.Time `db:"created_a"`
	Link        string    `db:"link_to_application"`
}
