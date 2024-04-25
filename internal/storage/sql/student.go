package sql

import (
	"context"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s *Storage) GetStudents(ctx context.Context) ([]entities.Student, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM Users`)
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
	err = s.db.QueryRowContext(ctx, `SELECT * FROM Students WHERE id = ?`, uuId).Scan(&result)
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
                            ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?
                   )`, sqlStudent.Id, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname,
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
		`UPDATE Students SET first_name =?, last_name =?, surname =?, passport_seria =?, passport_number =?,
                    birth_date =?, email =?, country =?, city =?, street =?, house =?, apartment =?, 
                    enroll_year =?, specialization =?, enroll_order_number =?
                WHERE id =?`, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname, sqlStudent.PassportSeria,
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

	_, err = s.db.ExecContext(ctx, `DELETE FROM Students WHERE id =?`, uuId)

	return err
}

func (s *Storage) GetStudentByEmail(email string, ctx context.Context) (entities.Student, error) {
	var student sqlEntities.Student

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Students WHERE email =?`, email).Scan(&student)

	return casts.StudentSqlToEntitie(student, ctx), err
}
