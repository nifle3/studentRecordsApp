package tests

import (
	"fmt"
	"studentRecordsApp/internal/entites"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckBirthdate(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			testingStudent := entities.entities{BirthDate: value.birthdate}

			result := testingStudent.CheckBirthdate()
			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckPassportSeria(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			testingStudent := entities.Student{PassportSeria: value.seria}

			result := testingStudent.CheckPassportSeria()
			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckNumber(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			testingStudent := entities.Student{PassportNumber: value.input}
			result := testingStudent.CheckNumber()

			assert.Equal(t, value.expected, result)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			testingStudent := entities.Student{Password: value.password}
			result := testingStudent.CheckPassword()

			assert.Equal(t, value.expected, result)
		})
	}
}
