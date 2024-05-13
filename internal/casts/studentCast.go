package casts

import (
	"context"
	"studentRecordsApp/internal/transport/server/httpEntity"

	"studentRecordsApp/internal/service/entities"
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

func StudentEntitieToHttpSelf(ctx context.Context, student entities.Student) httpEntity.StudentSelf {
	httpPhone := make([]httpEntity.PhoneNumberShort, 0, len(student.PhoneNumbers))
	for _, value := range student.PhoneNumbers {
		httpPhone = append(httpPhone, PhoneNumberEntiteToShort(ctx, value))
	}

	return httpEntity.StudentSelf{
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
		Group:           student.Group,
		Course:          student.Course,
		PhoneNumbers:    httpPhone,
	}
}

func StudentEntitieToStudentShort(_ context.Context, student entities.Student) (httpEntity.StudentShort, error) {
	id, err := UuidToString(student.Id)
	if err != nil {
		return httpEntity.StudentShort{}, err
	}

	return httpEntity.StudentShort{
		Id:             id,
		FirstName:      student.FirstName,
		LastName:       student.LastName,
		Surname:        student.Surname,
		EnrollYear:     student.EnrollYear,
		Specialization: student.Specialization,
		Group:          student.Group,
		Course:         student.Course,
	}, nil
}

func StudentEntitieToHttp(ctx context.Context, student entities.Student) (httpEntity.Student, error) {
	httpPhone := make([]httpEntity.PhoneNumberShort, 0, len(student.PhoneNumbers))
	for _, value := range student.PhoneNumbers {
		httpPhone = append(httpPhone, PhoneNumberEntiteToShort(ctx, value))
	}

	id, err := UuidToString(student.Id)
	if err != nil {
		return httpEntity.Student{}, err
	}

	return httpEntity.Student{
		Id:              id,
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
		Group:           student.Group,
		Course:          student.Course,
		PhoneNumbers:    httpPhone,
		Link:            student.LinkPhoto,
	}, nil
}
