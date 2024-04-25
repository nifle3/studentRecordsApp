package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"studentRecordsApp/internal/service/entites"
)

func TestCheckCorrectPhone(t *testing.T) {
	testData := []struct {
		name     string
		expected bool
		input    entities.PhoneNumber
	}{
		{
			"just numbers",
			false,
			entities.PhoneNumber{
				Phone: "123456789",
			},
		},
		{
			"empty string",
			false,
			entities.PhoneNumber{
				Phone: "",
			},
		},
		{
			"just symbols",
			false,
			entities.PhoneNumber{
				Phone: "aaaaaaaaaa",
			},
		},
		{
			"very long numbers",
			false,
			entities.PhoneNumber{
				Phone: "123123123123123123123",
			},
		},
		{
			"number and symbols",
			false,
			entities.PhoneNumber{
				Phone: "+8qweqweqweqwe",
			},
		},
		{
			"short number",
			false,
			entities.PhoneNumber{
				Phone: "8903",
			},
		},
		{
			"correct number with +",
			true,
			entities.PhoneNumber{
				Phone: "+79047667071",
			},
		},
		{
			"correct number without +",
			true,
			entities.PhoneNumber{
				Phone: "89047667071",
			},
		},
		{
			"correct number",
			true,
			entities.PhoneNumber{
				Phone: "29047667071",
			},
		},
	}

	for idx, value := range testData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			result, err := value.input.CheckCorrectNumber()

			assert.Equal(t, value.expected, result)
			assert.Nil(t, err)
		})
	}

}
