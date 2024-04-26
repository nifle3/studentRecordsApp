package sql

import (
	"context"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s *Storage) GetStudents(ctx context.Context) ([]entities.Student, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM Users;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]entities.Student, 0)

	for rows.Next() {
		var result sqlEntities.Student
		err := rows.Scan(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, casts.StudentSqlToEntitie(result, ctx))
	}

	return results, nil
}

func (s *Storage) GetStudent(id string, ctx context.Context) (entities.Student, error) {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return entities.Student{}, err
	}

	var result sqlEntities.Student
	err = s.db.QueryRowContext(ctx, `SELECT * FROM Students WHERE id = $1;`, uuId).Scan(&result)
	if err != nil {
		return entities.Student{}, err
	}

	return casts.StudentSqlToEntitie(result, ctx), nil
}

func (s *Storage) AddStudent(student entities.Student, ctx context.Context) error {
	sqlStudent := casts.StudentEntitieToSqlWithoutId(student, ctx)
	sqlStudent.Id = uuid.New()

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Students (id, first_name, last_name, surname, passport_seria, passport_number,
                    birth_date, email, password, country, city, street, house, apartment, enroll_year, 
                    specialization, enroll_order_number) 
                    VALUES(
                            $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$13,$14,$15,$16
                   );`, sqlStudent.Id, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname,
		sqlStudent.PassportSeria, sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email,
		sqlStudent.Password, sqlStudent.Country, sqlStudent.City, sqlStudent.Street,
		sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear, sqlStudent.Specialization,
		sqlStudent.OrderNumber)

	return err
}

func (s *Storage) UpdateStudent(student entities.Student, ctx context.Context) error {
	sqlStudent, err := casts.StudentEntitieToSql(student, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		`UPDATE Students SET first_name =$1, last_name =$2, surname =$3, passport_seria =$4, passport_number =$5,
                    birth_date =$6, email =$7, country =$8, city =$9, street =$10, house =$11, apartment =$12, 
                    enroll_year =$13, specialization =$14, enroll_order_number =$15
                WHERE id =$16;`, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname, sqlStudent.PassportSeria,
		sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email, sqlStudent.Password, sqlStudent.Country,
		sqlStudent.City, sqlStudent.Street, sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear,
		sqlStudent.Specialization, sqlStudent.OrderNumber, sqlStudent.Id)

	return err
}

func (s *Storage) DeleteStudent(id string, ctx context.Context) error {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `DELETE FROM Students WHERE id =$1;`, uuId)

	return err
}

func (s *Storage) GetStudentByEmail(email string, ctx context.Context) (entities.Student, error) {
	var student sqlEntities.Student

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Students WHERE email =$1;`, email).Scan(&student)

	return casts.StudentSqlToEntitie(student, ctx), err
}
