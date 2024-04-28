package sql

import (
	"context"
	"studentRecordsApp/internal/entites"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s *Storage) GetApplications(ctx context.Context) ([]entities.entities, error) {
	results := make([]sqlEntities.Application, 0)
	err := s.db.SelectContext(ctx, &results, "SELECT * FROM Applications ORDER BY created_at")
	if err != nil {
		return nil, err
	}

	resultsEntitie := make([]entities.Application, 0, len(results))
	for _, value := range results {
		resultsEntitie = append(resultsEntitie, casts.ApplicationSqlToService(value, ctx))
	}

	return resultsEntitie, nil
}

func (s *Storage) GetApplicationForUser(userId string, ctx context.Context) ([]entities.Application, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM Applications WHERE id =$1 ORDER BY created_at", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]entities.Application, 0)

	for rows.Next() {
		var app sqlEntities.Application
		err := rows.Scan(&app)
		if err != nil {
			return nil, err
		}

		results = append(results, casts.ApplicationSqlToService(app, ctx))
	}

	return results, nil
}

func (s *Storage) GetApplicationById(id string, ctx context.Context) (entities.Application, error) {
	var app sqlEntities.Application

	err := s.db.GetContext(ctx, &app, "SELECT * FROM Applications WHERE id =$1;", id)
	if err != nil {
		return entities.Application{}, err
	}

	return casts.ApplicationSqlToService(app, ctx), nil
}

func (s *Storage) AddApplication(application entities.Application, ctx context.Context) error {
	sqlApplication, err := casts.ApplicationServiceToSqlWithOutId(application, ctx)
	if err != nil {
		return err
	}

	sqlApplication.Id = uuid.New()

	_, err = s.db.ExecContext(ctx, `
        INSERT INTO Applications (id, student_id, contact_info, application_text, application_status, created_at) 
        VALUES ($1, $2, $3, $4, $5, $6);`,
		sqlApplication.Id, sqlApplication.StudentId, sqlApplication.Text, sqlApplication.Status, sqlApplication.CreatedAt)

	return err
}

func (s *Storage) UpdateApplication(application entities.Application, ctx context.Context) error {
	sqlApplication, err := casts.ApplicationServiceToSql(application, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `
        UPDATE Applications SET student_id =$1, contact_info =$2, application_text =$3, application_status =$4, created_at =$5
        WHERE id =$6;`,
		sqlApplication.StudentId, sqlApplication.Text, sqlApplication.Status, sqlApplication.CreatedAt, sqlApplication.Id)

	return err
}

func (s *Storage) DeleteApplication(id, userId string, ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM Applications WHERE id =$1 AND student_id =$2;", id, userId)

	return err
}
