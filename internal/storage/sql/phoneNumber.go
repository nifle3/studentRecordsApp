package sql

import (
	"context"
	"studentRecordsApp/internal/entites"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s *Storage) GetPhoneNumbers(userId string, ctx context.Context) ([]entities.entities, error) {
	uuStudentId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	sqlResults := make([]sqlEntities.PhoneNumber, 0)
	err = s.db.SelectContext(ctx, &sqlResults, "SELECT * FROM PhoneNumbers WHERE student_id = $1;", uuStudentId)
	if err != nil {
		return nil, err
	}

	results := make([]entities.PhoneNumber, 0, len(sqlResults))
	for _, result := range sqlResults {
		results = append(results, casts.PhoneNumberSqlToService(result, ctx))
	}

	return results, nil
}

func (s *Storage) AddPhoneNumber(number entities.PhoneNumber, ctx context.Context) error {
	sqlPhone, err := casts.PhoneNumberServiceToSqlWithoutId(number, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `INSERT INTO PhoneNumbers (id, student_id, country_code, city_code, code, description) 
							VALUES ($1, $2, $3, $4, $5, $6);`,
		sqlPhone.Id, sqlPhone.StudentId, sqlPhone.CountryCode, sqlPhone.CityCode, sqlPhone.Code, sqlPhone.Description)

	return err
}

func (s *Storage) UpdatePhoneNumber(number entities.PhoneNumber, ctx context.Context) error {
	sqlPhone, err := casts.PhoneNumberServiceToSql(number, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `UPDATE PhoneNumbers SET country_code =$1, city_code =$2, code =$3, description =$4 WHERE id =$5;`,
		sqlPhone.CountryCode, sqlPhone.CityCode, sqlPhone.Code, sqlPhone.Description, sqlPhone.Id)

	return err
}

func (s *Storage) DeletePhoneNumber(id, studentId string, ctx context.Context) error {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	uuuStudentId, err := uuid.Parse(studentId)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `DELETE FROM PhoneNumbers WHERE id =$1 AND student_id =$2;`, uuId, uuuStudentId)
	return err
}
