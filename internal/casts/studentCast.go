package casts

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func StudentEntitieToSql(student entities.Student, _ context.Context) (sqlEntities.Student, error) {
	id, err := uuid.Parse(student.Id)
	if err != nil {
		return sqlEntities.Student{}, err
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
