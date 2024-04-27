package tests

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentRecordsApp/pkg/email"
	"testing"
)

func TestIsCorrect(t *testing.T) {
	testingData := []struct {
		name     string
		expected error
		email    string
	}{
		{
			"valid email",
			nil,
			"nifle.3@mail.ru",
		},
		{
			"valid email 2",
			nil,
			"nifle3@gmail.com",
		},
		{
			"email without @",
			errors.New("Invalid email"),
			"nifle3gmail.com",
		},
		{
			"email without .",
			errors.New("Invalid email"),
			"nifle3@mailru",
		},
		{
			"email with . before @",
			errors.New("Invalid email"),
			"nife3.@mailru",
		},
		{
			"email with spaces",
			errors.New("Invalid email"),
			"nifle 3@gmail.com",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			assert.Equal(t, email.IsCorrect(value.email), value.expected)
		})
	}
}
