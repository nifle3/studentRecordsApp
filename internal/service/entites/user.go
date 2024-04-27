package entities

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
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

func (u *User) CheckIsNotEmpty() bool {
	return u.FirstName != "" &&
		u.LastName != "" &&
		u.Surname != "" &&
		u.Email != "" &&
		u.Password != ""
}

func (u *User) CheckEmail() (bool, error) {
	regex, err := regexp.Compile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-z]+$`)
	if err != nil {
		return false, err
	}

	return regex.MatchString(u.Email), nil
}

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashPassword)
	return nil
}

func (u *User) CheckHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) CheckRole() bool {
	return u.Role == UserAdmin || u.Role == UserWorker
}
