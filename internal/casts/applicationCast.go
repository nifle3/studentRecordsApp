package casts

import (
	"context"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/internal/transport/server/jsonStruct"

	"github.com/google/uuid"

	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func ApplicationServiceToSql(application entities.entities, _ context.Context) (sqlEntities.Application, error) {
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
		Name:      application.Name,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		Link:      application.Link,
	}, nil
}

func ApplicationServiceToSqlWithOutId(application entities.Application, _ context.Context) (sqlEntities.Application, error) {
	studentId, err := uuid.Parse(application.StudentId)
	if err != nil {
		return sqlEntities.Application{}, err
	}

	return sqlEntities.Application{
		Id:        uuid.UUID{},
		StudentId: studentId,
		Name:      application.Name,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		Link:      application.Link,
	}, nil
}

func ApplicationSqlToService(application sqlEntities.Application, _ context.Context) entities.Application {
	return entities.Application{
		Id:        application.Id.String(),
		StudentId: application.StudentId.String(),
		Name:      application.Name,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		Link:      application.Link,
	}
}

func ApplicationJsonAddedToEntites(added jsonStruct.ApplicationAdded, _ context.Context) entities.Application {
	return entities.Application{
		StudentId:   added.StudentId,
		ContactInfo: added.ContactInfo,
		Name:        added.Name,
		Text:        added.Text,
		File:        added.File,
	}
}

func ApplicationEntitesToJson(application entities.Application, _ context.Context) jsonStruct.Application {
	return jsonStruct.Application{
		Id:        application.Id,
		StudentId: application.StudentId,
		Name:      application.Name,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		Link:      application.Link,
		File:      application.File,
	}
}

func ApplicationJsonToEntites(application jsonStruct.Application, _ context.Context) entities.Application {
	return entities.Application{
		Id:        application.Id,
		StudentId: application.StudentId,
		Name:      application.Name,
		Text:      application.Text,
		Status:    application.Status,
		CreatedAt: application.CreatedAt,
		Link:      application.Link,
		File:      application.File,
	}
}
