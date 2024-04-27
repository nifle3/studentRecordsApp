package email

import (
	"errors"
	"regexp"
)

func IsCorrect(email string) (err error) {
	regex, err := regexp.Compile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-z]+$`)
	if err != nil {
		return err
	}

	if !regex.MatchString(email) {
		return errors.New("Invalid email")
	}

	return nil
}
