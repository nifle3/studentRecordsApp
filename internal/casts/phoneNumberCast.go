package casts

import (
	"context"
	"studentRecordsApp/internal/transport/server/jsonStruct"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
	"studentRecordsApp/pkg/stringMethod"
)

func PhoneNumberServiceToSql(number entities.PhoneNumber, _ context.Context) (sqlEntities.PhoneNumber, error) {
	id, err := uuid.Parse(number.Id)
	if err != nil {
		return sqlEntities.PhoneNumber{}, err
	}

	studentId, err := uuid.Parse(number.StudentId)
	if err != nil {
		return sqlEntities.PhoneNumber{}, err
	}

	reversePhone := stringMethod.Reverse(number.Phone)

	countryCode := reversePhone[:7]
	cityCode := reversePhone[7:10]
	code := reversePhone[10:]

	return sqlEntities.PhoneNumber{
		Id:          id,
		StudentId:   studentId,
		CountryCode: countryCode,
		CityCode:    cityCode,
		Code:        code,
		Description: number.Description,
	}, nil
}

func PhoneNumberServiceToSqlWithoutId(number entities.PhoneNumber, _ context.Context) (sqlEntities.PhoneNumber, error) {
	studentId, err := uuid.Parse(number.StudentId)
	if err != nil {
		return sqlEntities.PhoneNumber{}, err
	}

	reversePhone := stringMethod.Reverse(number.Phone)

	countryCode := reversePhone[:7]
	cityCode := reversePhone[7:10]
	code := reversePhone[10:]

	return sqlEntities.PhoneNumber{
		Id:          uuid.UUID{},
		StudentId:   studentId,
		CountryCode: countryCode,
		CityCode:    cityCode,
		Code:        code,
		Description: number.Description,
	}, nil
}

func PhoneNumberSqlToService(number sqlEntities.PhoneNumber, _ context.Context) entities.PhoneNumber {
	return entities.PhoneNumber{
		Id:          number.Id.String(),
		StudentId:   number.StudentId.String(),
		Phone:       number.CountryCode + number.CityCode + number.Code,
		Description: number.Description,
	}
}

func PhoneNumberEntitiesToJson(number entities.PhoneNumber, _ context.Context) jsonStruct.PhoneNumber {
	return jsonStruct.PhoneNumber{
		Id:          number.Id,
		StudentId:   number.StudentId,
		Phone:       number.Phone,
		Description: number.Description,
	}
}

func PhoneNumberJsonToEntities(number jsonStruct.PhoneNumberWithoutId, _ context.Context) entities.PhoneNumber {
	return entities.PhoneNumber{
		StudentId:   number.StudentId,
		Phone:       number.Phone,
		Description: number.Description,
	}
}

func PhoneNumberJsonLongToEntities(number jsonStruct.PhoneNumber, _ context.Context) entities.PhoneNumber {
	return entities.PhoneNumber{
		Id:          number.Id,
		StudentId:   number.StudentId,
		Phone:       number.Phone,
		Description: number.Description,
	}
}
