package casts

import (
	"context"
	"strings"
	"studentRecordsApp/internal/transport/server/httpEntity"
	"studentRecordsApp/pkg/stringMethod"

    "studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func PhoneNumberSqlToEntitie(_ context.Context, phone sqlEntities.PhoneNumber) entities.PhoneNumber {
	return entities.PhoneNumber{
		Id:          phone.Id,
		StudentId:   phone.StudentId,
		Phone:       strings.Join([]string{phone.CountryCode, phone.CityCode, phone.Code}, ""),
		Description: phone.Description,
	}
}

func PhoneNumberEntitieToSql(_ context.Context, phone entities.PhoneNumber) sqlEntities.PhoneNumber {
	phone.Phone = stringMethod.Reverse(phone.Phone)

	return sqlEntities.PhoneNumber{
		Id:          phone.Id,
		StudentId:   phone.StudentId,
		CountryCode: stringMethod.Reverse(phone.Phone[10:]),
		CityCode:    stringMethod.Reverse(phone.Phone[7:10]),
		Code:        stringMethod.Reverse(phone.Phone[0:7]),
		Description: phone.Description,
	}
}

func PhoneNumberEntiteToShort(_ context.Context, phone entities.PhoneNumber) httpEntity.PhoneNumberShort {
	return httpEntity.PhoneNumberShort{
		Phone:       phone.Phone,
		Description: phone.Description,
	}
}
