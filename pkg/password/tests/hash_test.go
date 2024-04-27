package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentRecordsApp/pkg/password"
	"testing"
)

func TestHash(t *testing.T) {
	data := []struct {
		password      string
		expectedError interface{}
	}{
		{
			"qwe123",
			nil,
		},
		{
			"",
			struct{}{},
		},
		{
			"qweasdasdzxc123",
			nil,
		},
	}

	for idx, value := range data {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.password), func(t *testing.T) {
			result, err := password.Hash(value.password)
			result1, err := password.Hash(value.password)
			if value.expectedError == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}

			assert.Equal(t, result, result1)
		})
	}
}

func TestCheckHash(t *testing.T) {
	data := []struct {
		password      string
		hashPassword  string
		expectedError interface{}
	}{
		{
			"qwe123",
			"qwe123",
			nil,
		},
		{
			"qwe123",
			"qwe",
			struct{}{},
		},
		{
			"q",
			"q",
			nil,
		},
		{
			"qwe123",
			"123",
			struct{}{},
		},
	}

	var err error

	for idx, value := range data {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.password), func(t *testing.T) {
			value.hashPassword, err = password.Hash(value.hashPassword)
			if err != nil {
				t.Fatal("Hash func doesn't work")
			}

			err = password.CheckHash(value.password, []byte(value.hashPassword))
			if value.expectedError == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
