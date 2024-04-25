package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"studentRecordsApp/internal/service/entites"
)

func TestHashingPassword(t *testing.T) {
	testingData := []struct {
		name     string
		expected interface{}
		password string
	}{
		{
			name:     "empty password",
			expected: nil,
			password: "",
		},
		{
			name:     "too long password",
			expected: struct{}{},
			password: strings.Repeat("a", 10000),
		},
		{
			name:     "symbol only password",
			expected: nil,
			password: "asdasdasdASD",
		},
		{
			name:     "base password",
			expected: nil,
			password: "1_&%qweqwe1As",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testUser := entities.User{Password: value.password}
			result := testUser.HashPassword()

			if value.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
			}
		})
	}
}

func TestCheckHash(t *testing.T) {
	password1 := "1_&%qweqwe1As"
	hash1, _ := bcrypt.GenerateFromPassword([]byte(password1), 10)

	testingData := []struct {
		name     string
		expected interface{}
		password string
		hash     string
	}{
		{
			"valid password",
			nil,
			password1,
			string(hash1),
		},
		{
			"invalid password",
			struct{}{},
			"1_&%qweqwe1As",
			"$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testUser := entities.User{Password: value.hash}
			result := testUser.CheckHash(value.password)

			if value.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
			}
		})
	}
}
