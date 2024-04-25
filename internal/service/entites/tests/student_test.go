package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"studentRecordsApp/internal/service/entites"
)

func TestCheckBirthdate(t *testing.T) {
	testingData := []struct {
		name      string
		expected  bool
		birthdate time.Time
	}{{
		"valid birthdate",
		true,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
		{
			"birthdate is now",
			false,
			time.Now(),
		},
		{
			"student is 10 years old",
			false,
			time.Now().AddDate(-10, 1, 1),
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testingStudent := entities.Student{BirthDate: value.birthdate}

			result := testingStudent.CheckBirthdate()
			assert.Equal(t, value.expected, result)
		})
	}
}

func CheckPassportSeria(t *testing.T) {
	testingData := []struct {
		name     string
		expected bool
		seria    int
	}{
		{
			"valid seria",
			true,
			1234,
		},
		{
			"too short seria",
			false,
			999,
		},
		{
			"too long seria",
			false,
			10000,
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testingStudent := entities.Student{PassportSeria: value.seria}

			result := testingStudent.CheckPassportSeria()
			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckNumber(t *testing.T) {
	testingData := []struct {
		name     string
		expected bool
		input    int
	}{
		{
			"just numbers",
			true,
			789231,
		},
		{
			"too short number",
			false,
			99999,
		},
		{
			"too long number",
			false,
			1000000,
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testingStudent := entities.Student{PassportNumber: value.input}
			result := testingStudent.CheckNumber()

			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testingData := []struct {
		name     string
		expected bool
		password string
	}{
		{
			"valid password",
			true,
			"sqa1llsjcaz",
		},
		{
			"too short password",
			false,
			"1234q",
		},
		{
			"password with only spaces",
			false,
			"    ",
		},
		{
			"password with spaces",
			false,
			"123 123",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testingStudent := entities.Student{Password: value.password}
			result := testingStudent.CheckPassword()

			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckEmail(t *testing.T) {
	testingData := []struct {
		name     string
		expected bool
		email    string
	}{
		{
			"valid email",
			true,
			"nifle.3@mail.ru",
		},
		{
			"valid email 2",
			true,
			"nifle3@gmail.com",
		},
		{
			"email without @",
			false,
			"nifle3gmail.com",
		},
		{
			"email without .",
			false,
			"nifle3@mailru",
		},
		{
			"email with . before @",
			false,
			"nife3.@mailru",
		},
		{
			"email with spaces",
			false,
			"nifle 3@gmail.com",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			testingStudent := entities.Student{Email: value.email}
			result, err := testingStudent.CheckEmail()

			assert.Equal(t, value.expected, result)
			assert.Nil(t, err)
		})
	}
}
