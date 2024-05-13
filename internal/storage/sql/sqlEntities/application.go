package sqlEntities

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	Id          uuid.UUID `db:"id"`
	StudentId   uuid.UUID `db:"student_id"`
	ContactInfo string    `db:"contact_info"`
	Name        string    `db:"_name"`
	Text        string    `db:"_text"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	Link        string    `db:"link"`
}

type ApplicationWithInfo struct {
	Id          uuid.UUID `db:"id"`
	StudentId   uuid.UUID `db:"student_id"`
	ContactInfo string    `db:"contact_info"`
	Name        string    `db:"_name"`
	Text        string    `db:"_text"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	Link        string    `db:"link"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Surname     string    `db:"surname"`
	Course      int       `db:"course"`
	Group       int       `db:"_group"`
}
