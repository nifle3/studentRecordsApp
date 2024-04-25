package entities

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        string
	FirstName string
	LastName  string
	Surname   string
	Email     string
	Password  string
	Medicine  string
	// Role must be "Сотрудник" or "Админ"
	Role string
}

func (u *User) CheckIsNotEmpty() bool {
	return u.FirstName != "" &&
		u.LastName != "" &&
		u.Surname != "" &&
		u.Email != "" &&
		u.Password != "" &&
		u.Medicine != ""
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
