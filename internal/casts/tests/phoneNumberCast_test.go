package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func TestPhoneNumberSqlToEntitie(t *testing.T) {
	t.Parallel()
	uuId := uuid.New()

	testingData := []struct {
		Name     string
		Input    sqlEntities.PhoneNumber
		Expected entities.PhoneNumber
	}{
		{
			Name: "with +7",
			Input: sqlEntities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				CountryCode: "+7",
				CityCode:    "904",
				Code:        "7667071",
				Description: "lorem",
			},
			Expected: entities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				Phone:       "+79047667071",
				Description: "lorem",
			},
		},
		{
			Name: "with 8",
			Input: sqlEntities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				CountryCode: "8",
				CityCode:    "904",
				Code:        "7667071",
				Description: "lorem",
			},
			Expected: entities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				Phone:       "89047667071",
				Description: "lorem",
			},
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("test %d: %s", idx, value.Name), func(t *testing.T) {
			t.Parallel()

			result := casts.PhoneNumberSqlToEntitie(context.Background(), value.Input)
			assert.Equal(t, value.Expected, result)
		})
	}
}

func TestPhoneNumberEntitieToSql(t *testing.T) {
	t.Parallel()
	uuId := uuid.New()

	testingData := []struct {
		Name     string
		Input    entities.PhoneNumber
		Expected sqlEntities.PhoneNumber
	}{
		{
			Name: "with +7",
			Input: entities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				Phone:       "+79047667071",
				Description: "lorem",
			},
			Expected: sqlEntities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				CountryCode: "+7",
				CityCode:    "904",
				Code:        "7667071",
				Description: "lorem",
			},
		},
		{
			Name: "with 8",
			Input: entities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				Phone:       "89047667071",
				Description: "lorem",
			},
			Expected: sqlEntities.PhoneNumber{
				Id:          uuId,
				StudentId:   uuId,
				CountryCode: "8",
				CityCode:    "904",
				Code:        "7667071",
				Description: "lorem",
			},
		},
	}

	for idx, value := range testingData {
		t.Run(fmt.Sprintf("test %d: %s", idx, value.Name), func(t *testing.T) {
			t.Parallel()

			result := casts.PhoneNumberEntitieToSql(context.Background(), value.Input)
			assert.Equal(t, value.Expected, result)
		})
	}
}
