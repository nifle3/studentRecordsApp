package sql

import (
	"context"
	"studentRecordsApp/internal/service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

var _ service.StudentDB = (*Student)(nil)

type Student struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) *Student {
	return &Student{
		db: db,
	}
}

func (s *Student) GetAll(ctx context.Context) ([]entities.Student, error) {
	result := make([]sqlEntities.Student, 0)
	err := s.db.SelectContext(ctx, &result, `SELECT * FROM Students;`)
	if err != nil {
		return nil, err
	}

	results := make([]entities.Student, 0)

	for _, value := range result {
		results = append(results, casts.StudentSqlToEntitie(ctx, value))
	}

	return results, nil
}

func (s *Student) Get(ctx context.Context, id uuid.UUID) (entities.Student, error) {
	var result sqlEntities.Student
	err := s.db.GetContext(ctx, &result, `SELECT * FROM Students WHERE id = $1 LIMIT 1;`, id)
	if err != nil {
		return entities.Student{}, err
	}

	return casts.StudentSqlToEntitie(ctx, result), nil
}

func (s *Student) GetLinkById(ctx context.Context, id uuid.UUID) (string, error) {
	var result string
	if err := s.db.GetContext(ctx, &result, `SELECT link_photo FROM Students WHERE id = $1 LIMIT 1;`, id); err != nil {
		return "", err
	}
	return result, nil
}

func (s *Student) Add(ctx context.Context, student entities.Student) error {
	sqlStudent := casts.StudentEntiteToSql(ctx, student)
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Students (id, first_name, last_name, surname, passport_seria, passport_number,
                    birth_date, email, password, country, city, street, house, apartment, enroll_year, 
                    specialization, link_photo, course, _group) 
                    VALUES(
                            $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19
                   );`, sqlStudent.Id, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname,
		sqlStudent.PassportSeria, sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email,
		sqlStudent.Password, sqlStudent.Country, sqlStudent.City, sqlStudent.Street,
		sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear, sqlStudent.Specialization,
		sqlStudent.LinkPhoto, sqlStudent.Course, sqlStudent.Group)

	return err
}

func (s *Student) Update(ctx context.Context, student entities.Student) error {
	sqlStudent := casts.StudentEntiteToSql(ctx, student)

	_, err := s.db.ExecContext(ctx,
		`UPDATE Students SET first_name =$1, last_name =$2, surname =$3, passport_seria =$4, passport_number =$5,
                    birth_date =$6, email =$7, country =$8, city =$9, street =$10, house =$11, apartment =$12, 
                    enroll_year =$13, specialization =$14
                WHERE id =$15;`, sqlStudent.FirstName, sqlStudent.LastName, sqlStudent.Surname, sqlStudent.PassportSeria,
		sqlStudent.PassportNumber, sqlStudent.BirthDate, sqlStudent.Email, sqlStudent.Country,
		sqlStudent.City, sqlStudent.Street, sqlStudent.HouseNumber, sqlStudent.ApartmentNumber, sqlStudent.EnrollYear,
		sqlStudent.Specialization, sqlStudent.Id)

	return err
}

func (s *Student) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM Students WHERE id =$1;`, id)

	return err
}

func (s *Student) Auth(ctx context.Context, email string) (uuid.UUID, string, error) {
	var student sqlEntities.Student

	err := s.db.GetContext(ctx, &student, `SELECT * FROM Students WHERE email =$1;`, email)

	return student.Id, student.Password, err
}
