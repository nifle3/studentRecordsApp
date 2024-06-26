package sqlEntities

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Surname   string    `db:"surname"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"user_role"`
}
