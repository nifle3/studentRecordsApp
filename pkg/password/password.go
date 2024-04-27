package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hashPassword), err
}

func CheckHash(password string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
