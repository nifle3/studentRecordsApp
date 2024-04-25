package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"studentRecordsApp/internal/service/entites"
)

func TestCheckStatus(t *testing.T) {
	testingData := []struct {
		name     string
		expected bool
		status   string
	}{
		{
			"valid status",
			true,
			"Создана",
		},
		{
			"valid status",
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
			testingApplication := entities.Application{Status: value.status}
			result := testingApplication.CheckStatus()
			assert.Equal(t, value.expected, result)
		})
	}
}
