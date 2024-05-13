package tests

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"studentRecordsApp/internal/service/entities"
)

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()

	dataSet := []struct {
		name     string
		expected bool
		input    entities.User
	}{
		{
			"correct input",
			true,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"empty first name",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"empty last name",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"empty surname",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"empty email",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"empty password",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "",
				Role:      "Админ",
			},
		},
		{
			"empty role",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "",
			},
		},
		{
			"full empty",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "",
				LastName:  "",
				Surname:   "",
				Email:     "",
				Password:  "",
				Role:      "",
			},
		},
		{
			"space first name",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: " ",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"space last name",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "  ",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"space surname",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   " ",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"space email",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "  ",
				Password:  "qwe123",
				Role:      "Админ",
			},
		},
		{
			"space password",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "  ",
				Role:      "Админ",
			},
		},
		{
			"space role",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: "John",
				LastName:  "Doe",
				Surname:   "Smith",
				Email:     "kuper115rus.2005@gmail.com",
				Password:  "qwe123",
				Role:      " ",
			},
		},
		{
			"full space",
			false,
			entities.User{
				Id:        uuid.Nil,
				FirstName: " ",
				LastName:  " ",
				Surname:   " ",
				Email:     " ",
				Password:  " ",
				Role:      " ",
			},
		},
	}

	for idx, value := range dataSet {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			t.Parallel()

			result := value.input.IsNotEmpty()
			assert.Equal(t, value.expected, result)
		})
	}
}
