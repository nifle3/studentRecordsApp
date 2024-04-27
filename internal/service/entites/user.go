package entities

import (
	"strings"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
	Surname   string
	Email     string
	Password  string
	// Role must be "Сотрудник" or "Админ"
	Role string
}

func (u *User) IsNotEmpty() bool {
	return u.FirstName != "" && !strings.Contains(u.FirstName, " ") &&
		u.LastName != "" && !strings.Contains(u.LastName, " ") &&
		u.Surname != "" && !strings.Contains(u.Surname, " ") &&
		u.Email != "" && !strings.Contains(u.Email, " ") &&
		u.Password != "" && !strings.Contains(u.Password, " ") &&
		u.Role != "" && !strings.Contains(u.Role, " ")
}

func (u *User) IsRoleCorrect() bool {
	return u.Role == UserAdmin || u.Role == UserWorker
}
