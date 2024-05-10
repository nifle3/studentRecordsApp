package casts

import (
	"context"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func StudentEntiteToSql(_ context.Context, student entities.Student) sqlEntities.Student {
	return sqlEntities.Student{
		Id:              student.Id,
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
		LinkPhoto:       student.LinkPhoto,
		Group:           student.Group,
		Course:          student.Course,
	}
}

func StudentSqlToEntitie(_ context.Context, student sqlEntities.Student) entities.Student {
	return entities.Student{
		Id:              student.Id,
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
		LinkPhoto:       student.LinkPhoto,
		Group:           student.Group,
		Course:          student.Course,
		Photo:           nil,
	}
}
