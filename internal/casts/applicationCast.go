package casts

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func ApplicationServiceToSql(application entities.Application, _ context.Context) (sqlEntities.Application, error) {
	id, err := uuid.Parse(application.Id)
	if err != nil {
		return sqlEntities.Application{}, err
	}

	studentId, err := uuid.Parse(application.StudentId)
	if err != nil {
		return sqlEntities.Application{}, err
	}

	return sqlEntities.Application{
		Id:        id,
		StudentId: studentId,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
	}, nil
}

func ApplicationSqlToService(application sqlEntities.Application, _ context.Context) entities.Application {
	return entities.Application{
		Id:        application.Id.String(),
		StudentId: application.StudentId.String(),
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
	}
}
