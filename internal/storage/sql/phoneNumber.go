package sql

import (
	"context"
	"studentRecordsApp/internal/service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

var _ service.StudentPhoneDB = (*Phone)(nil)

type Phone struct {
	db *sqlx.DB
}

func NewPhone(db *sqlx.DB) *Phone {
	return &Phone{
		db: db,
	}
}

func (p *Phone) GetForUser(ctx context.Context, userId uuid.UUID) ([]entities.PhoneNumber, error) {
	var sqlResults []sqlEntities.PhoneNumber
	err := p.db.SelectContext(ctx, &sqlResults, "SELECT * FROM PhoneNumbers WHERE student_id = $1", userId)
	if err != nil {
		return nil, err
	}

	results := make([]entities.PhoneNumber, 0, len(sqlResults))
	for idx := range sqlResults {
		results = append(results, casts.PhoneNumberSqlToEntitie(ctx, sqlResults[idx]))
	}

	return results, nil
}

func (p *Phone) Add(ctx context.Context, number entities.PhoneNumber) error {
	sqlPhone := casts.PhoneNumberEntitieToSql(ctx, number)

	_, err := p.db.ExecContext(ctx, `INSERT INTO PhoneNumbers (id, student_id, country_code, city_code, code, description) 
							VALUES ($1, $2, $3, $4, $5, $6);`,
		sqlPhone.Id, sqlPhone.StudentId, sqlPhone.CountryCode, sqlPhone.CityCode, sqlPhone.Code, sqlPhone.Description)

	return err
}

func (p *Phone) Update(ctx context.Context, number entities.PhoneNumber) error {
	sqlPhone := casts.PhoneNumberEntitieToSql(ctx, number)

	_, err := p.db.ExecContext(ctx, `UPDATE PhoneNumbers SET country_code =$1, city_code =$2, code =$3, description =$4 WHERE id =$5`,
		sqlPhone.CountryCode, sqlPhone.CityCode, sqlPhone.Code, sqlPhone.Description, sqlPhone.Id)

	return err
}

func (p *Phone) Delete(ctx context.Context, id uuid.UUID) error {

	_, err := p.db.ExecContext(ctx, `DELETE FROM PhoneNumbers WHERE id =$1`, id)
	return err
}
