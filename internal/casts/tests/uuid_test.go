package tests

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"studentRecordsApp/internal/casts"
)

func TestUuidToString(t *testing.T) {
	t.Parallel()

	testingData := []uuid.UUID{
		uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			_, err := casts.UuidToString(value)
			assert.Nil(t, err)
		})
	}
}

func TestStringToUuid(t *testing.T) {
	t.Parallel()

	testingData := []uuid.UUID{
		uuid.New(), uuid.New(), uuid.New(), uuid.New(),
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			stringValue, _ := casts.UuidToString(value)
			result, err := casts.StringToUuid(stringValue)
			assert.Nil(t, err)
			assert.Equal(t, value, result)
		})
	}

}
