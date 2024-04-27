package casts

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/transport/server/jsonStruct"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func StudentEntitieToSql(student entities.Student, _ context.Context) (sqlEntities.Student, error) {
	id, err := uuid.Parse(student.Id)
	if err != nil {
		return sqlEntities.Student{}, nil
	}

	return sqlEntities.Student{
		Id:              id,
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Password:        student.Password,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
	}, nil
}

func StudentEntitieToSqlWithoutId(student entities.Student, _ context.Context) sqlEntities.Student {
	return sqlEntities.Student{
		Id:              uuid.UUID{},
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Password:        student.Password,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
	}
}

func StudentSqlToEntitie(student sqlEntities.Student, _ context.Context) entities.Student {
	return entities.Student{
		Id:              student.Id.String(),
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Password:        student.Password,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
	}
}

func StudentSqlToEntitieWithPhone(student sqlEntities.Student, phones []sqlEntities.PhoneNumber, _ context.Context) entities.Student {
	entitesPhone := make([]entities.PhoneNumber, 0, len(phones))

	for _, value := range phones {
		entitesPhone = append(entitesPhone, PhoneNumberSqlToService(value, context.Background()))
	}

	return entities.Student{
		Id:              student.Id.String(),
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Password:        student.Password,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
	}
}

func StudentEntitieToJsonShort(student entities.Student, _ context.Context) jsonStruct.StudentShort {
	return jsonStruct.StudentShort{
		Id:         student.Id,
		FirstName:  student.FirstName,
		LastName:   student.LastName,
		Surname:    student.Surname,
		EnrollYear: student.EnrollYear,
		Photo:      student.Photo,
	}
}

func StudentEntitieToJson(student entities.Student, _ context.Context) jsonStruct.StudentLong {
	return jsonStruct.StudentLong{
		Id:              student.Id,
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
		Photo:           student.Photo,
	}
}

func StudentJsonWithoutLinkToEntitie(student jsonStruct.StudentLongWithoutLink, _ context.Context) entities.Student {
	return entities.Student{
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Password:        student.Password,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
		Photo:           student.Photo,
	}
}

func StudentJsonToEntitie(student jsonStruct.StudentLong, _ context.Context) entities.Student {
	return entities.Student{
		Id:              student.Id,
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		Surname:         student.Surname,
		PassportSeria:   student.PassportSeria,
		PassportNumber:  student.PassportNumber,
		BirthDate:       student.BirthDate,
		Email:           student.Email,
		Country:         student.Country,
		City:            student.City,
		Street:          student.Street,
		HouseNumber:     student.HouseNumber,
		ApartmentNumber: student.ApartmentNumber,
		EnrollYear:      student.EnrollYear,
		Specialization:  student.Specialization,
		OrderNumber:     student.OrderNumber,
		LinkPhoto:       student.LinkPhoto,
		Photo:           student.Photo,
	}
}
