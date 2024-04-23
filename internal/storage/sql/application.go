package sql

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s Storage) AddApplication(application sqlEntities.Application, ctx context.Context) error {
	application.Id = uuid.New()
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Applications (id, student_id, contact_info, application_text, application_status, 
                          created_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`,
		application.Id, application.StudentId, application.Text, application.Status, application.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) UpdateApplication(application sqlEntities.Application, ctx context.Context) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE Applications SET student_id = $1, contact_info = $2, application_text = $3,
					application_status = $4, created_at = $5 
				WHERE id = $6`,
		application.StudentId, application.Text, application.Status, application.CreatedAt, application.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) GetApplicationById(id string, ctx context.Context) (sqlEntities.Application, error) {
	return sqlEntities.Application{}, nil
}

func (s Storage) GetApplications(ctx context.Context) ([]sqlEntities.Application, error) {
	return nil, nil
}
