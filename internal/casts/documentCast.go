package casts

import (
	"context"

	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
	"studentRecordsApp/internal/transport/server/httpEntity"
)

func DocumentSqlToEntite(_ context.Context, doc sqlEntities.StudentsDocument) entities.Document {
	return entities.Document{
		Id:        doc.Id,
		StudentId: doc.StudentId,
		Name:      doc.Name,
		Type:      doc.Type,
		Link:      doc.Link,
		CreatedAt: doc.CreatedAt,
		File:      nil,
	}
}

func DocumentEntiteToSql(_ context.Context, doc entities.Document) sqlEntities.StudentsDocument {
	return sqlEntities.StudentsDocument{
		Id:        doc.Id,
		StudentId: doc.StudentId,
		Name:      doc.Name,
		Type:      doc.Type,
		Link:      doc.Link,
		CreatedAt: doc.CreatedAt,
	}
}

func DocumentEntitieToDocumentSelf(_ context.Context, doc entities.Document) (httpEntity.DocumentSelf, error) {
	id, err := UuidToString(doc.Id)
	if err != nil {
		return httpEntity.DocumentSelf{}, err
	}

	return httpEntity.DocumentSelf{
		Id:        id,
		Name:      doc.Name,
		Type:      doc.Type,
		Link:      doc.Link,
		CreatedAt: doc.CreatedAt,
	}, nil
}
