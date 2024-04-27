package casts

import (
	"context"
	"studentRecordsApp/internal/transport/server/jsonStruct"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func DocumentEntiteToSql(document entities.Document, _ context.Context) (sqlEntities.StudentsDocument, error) {
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

func DocumentEntiteToSqlWithoutId(document entities.Document, _ context.Context) (sqlEntities.StudentsDocument, error) {
	studentId, err := uuid.Parse(document.StudentId)
	if err != nil {
		return sqlEntities.StudentsDocument{}, err
	}

	return sqlEntities.StudentsDocument{
		Id:        uuid.UUID{},
		StudentId: studentId,
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
	}, nil
}

func DocumentSqlToEntite(document sqlEntities.StudentsDocument, _ context.Context) entities.Document {
	return entities.Document{
		Id:        document.Id.String(),
		StudentId: document.StudentId.String(),
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
		File:      nil,
	}
}

func DocumentEntiteToJson(document entities.Document, _ context.Context) jsonStruct.Document {
	return jsonStruct.Document{
		Id:        document.Id,
		StudentId: document.StudentId,
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
		File:      document.File,
	}
}

func DocumentEntiteToJsonShort(document entities.Document, _ context.Context) jsonStruct.DocumentWithoutFile {
	return jsonStruct.DocumentWithoutFile{
		Id:        document.Id,
		StudentId: document.StudentId,
		Name:      document.Name,
		Type:      document.Type,
		Link:      document.Link,
		CreatedAt: document.CreatedAt,
	}
}

func DocumentForAddedToEntite(document jsonStruct.DocumentForAdded, _ context.Context) entities.Document {
	return entities.Document{
		StudentId: document.StudentId,
		Name:      document.Name,
		Type:      document.Type,
		File:      document.File,
	}
}
