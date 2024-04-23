package casts

import (
	"context"
	"github.com/google/uuid"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func DocumentEntiteToSql(document entities.StudentsDocument, _ context.Context) (sqlEntities.StudentsDocument, error) {
	id, err := uuid.Parse(document.Id)
	if err != nil {
		return sqlEntities.StudentsDocument{}, err
	}

	studentId, err := uuid.Parse(document.StudentId)
	if err != nil {
		return sqlEntities.StudentsDocument{}, err
	}

	return sqlEntities.StudentsDocument{
		Id:        id,
		StudentId: studentId,
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
	}, nil
}

func DocumentSqlToEntite(document sqlEntities.StudentsDocument, _ context.Context) entities.StudentsDocument {
	return entities.StudentsDocument{
		Id:        document.Id.String(),
		StudentId: document.StudentId.String(),
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
	}
}
