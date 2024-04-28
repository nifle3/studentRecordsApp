package tests

import (
	"fmt"
	"studentRecordsApp/internal/entites"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStatus(t *testing.T) {
	t.Parallel()

	testingData := []struct {
		name     string
		expected bool
		status   string
	}{
		{
			"valid status 1",
			true,
			"Создана",
		},
		{
			"valid status 2",
			true,
			"Закрыта",
		},
		{
			"invalid value",
			false,
			"rejected",
		},
		{
			"invalid status",
			false,
			"invalid",
		},
		{
			"empty status",
			false,
			"",
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
			t.Parallel()
			testingApplication := entities.entities{Status: value.status}
			result := testingApplication.CheckStatus()
			assert.Equal(t, value.expected, result)
		})
	}
}
