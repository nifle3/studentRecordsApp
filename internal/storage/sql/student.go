package sql

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s Storage) AddStudent(student sqlEntities.Student, phone []sqlEntities.PhoneNumber,
	studentsReport sqlEntities.StudentsReport, ctx context.Context) error {

	student.Id = uuid.New()
	studentsReport.Id = student.Id

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO Students (id, first_name, last_name, surname, passport_seria, passport_number, birth_date, email,
                              password, country, city, street, house, apartment)
        VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		student.Id, student.FirstName, student.LastName, student.Surname, student.PassportSeria, student.PassportNumber,
		student.BirthDate, student.Email, student.Password, student.Country, student.City, student.Street,
		student.HouseNumber, student.ApartmentNumber)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, item := range phone {
		item.Id = uuid.New()
		item.StudentId = student.Id

		_, err = tx.ExecContext(ctx, `
        	INSERT INTO PhoneNumbers (id, student_id, country_code, city_code, code, description)
        	VALUES(?, ?, ?, ?, ?, ?)`,
			item.Id, item.StudentId, item.CountryCode, item.CityCode, item.Code, item.Description)

		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO StudentsReports (id, enroll_year, specialization, enroll_order_number)
        VALUES(?,?,?,?)`,
		studentsReport.Id, studentsReport.EnrollYear, studentsReport.Specialization, studentsReport.EnrollYear)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (s Storage) UpdateStudent(student sqlEntities.Student, ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `
		UPDATE Students SET first_name=?, last_name=?, surname=?, passport_seria=?, passport_number=?,
		birth_date=?, email=?, password=?, country=?, city=?, street=?, house=?, apartment=? 
		WHERE id=?`,
		student.FirstName, student.LastName, student.Surname, student.PassportSeria, student.PassportNumber,
		student.BirthDate, student.Email, student.Password, student.Country, student.City, student.Street,
		student.HouseNumber, student.ApartmentNumber, student.Id)

	return err
}

func (s Storage) GetAllStudents(ctx context.Context) ([]sqlEntities.Student, error) {
	return nil, nil
}

func (s Storage) GetStudentById(id uuid.UUID, ctx context.Context) (sqlEntities.Student, error) {
	row, err := s.db.QueryContext(ctx, `
		SELECT * FROM Students WHERE id=?`, id)

	if err != nil {
		return sqlEntities.Student{}, err
	}

	var result sqlEntities.Student
	err = row.Scan(&result)
	if err != nil {
		return sqlEntities.Student{}, err
	}

	return result, nil
}
