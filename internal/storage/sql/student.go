package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

type Student struct {
	db *sqlx.DB
}

func (s *Student) GetStudents(ctx context.Context) ([]entities.Student, error) {
	result := make([]sqlEntities.Student, 0)
	err := s.db.SelectContext(ctx, &result, `SELECT * FROM Users;`)
	if err != nil {
		return nil, err
	}

	results := make([]entities.Student, 0)

	for _, value := range result {
		results = append(results, casts.StudentSqlToEntitie(ctx, value))
	}

	return results, nil
}

func (s *Student) GetStudent(ctx context.Context, id uuid.UUID) (entities.Student, error) {
	var result sqlEntities.Student
	err := s.db.GetContext(ctx, &result, `SELECT * FROM Students WHERE id = $1 LIMIT 1;`, id)
	if err != nil {
		return entities.Student{}, err
	}

	return casts.StudentSqlToEntitie(ctx, result), nil
}

func (s *Student) AddStudent(ctx context.Context, student entities.Student) error {
	sqlStudent := casts.StudentEntiteToSql(ctx, student)
	var id uuid.UUID
	err := s.db.GetContext(ctx, &id,
		`INSERT INTO Students (id, first_name, last_name, surname, passport_seria, passport_number,
                    birth_date, email, password, country, city, street, house, apartment, enroll_year, 
                    specialization, enroll_order_number) 
                    VALUES(
                            $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$13,$14,$15,$16
                   ) RETURNING id;`, sqlStudent.Id, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname,
		sqlStudent.PassportSeria, sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email,
		sqlStudent.Password, sqlStudent.Country, sqlStudent.City, sqlStudent.Street,
		sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear, sqlStudent.Specialization)

	return err
}

func (s *Student) UpdateStudent(ctx context.Context, student entities.Student) error {
	sqlStudent := casts.StudentEntiteToSql(ctx, student)

	_, err := s.db.ExecContext(ctx,
		`UPDATE Students SET first_name =$1, last_name =$2, surname =$3, passport_seria =$4, passport_number =$5,
                    birth_date =$6, email =$7, country =$8, city =$9, street =$10, house =$11, apartment =$12, 
                    enroll_year =$13, specialization =$14, enroll_order_number =$15
                WHERE id =$16;`, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname, sqlStudent.PassportSeria,
		sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email, sqlStudent.Password, sqlStudent.Country,
		sqlStudent.City, sqlStudent.Street, sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear,
		sqlStudent.Specialization, sqlStudent.Id)

	return err
}

func (s *Student) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM Students WHERE id =$1;`, id)

	return err
}

func (s *Student) GetStudentByEmail(ctx context.Context, email string) (entities.Student, error) {
	var student sqlEntities.Student

	err := s.db.GetContext(ctx, &student, `SELECT * FROM Students WHERE email =$1;`, email)

	return casts.StudentSqlToEntitie(ctx, student), err
}
