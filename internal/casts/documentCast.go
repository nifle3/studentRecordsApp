package casts

import (
	"context"
	entities "studentRecordsApp/internal/service/entites"

	"studentRecordsApp/internal/storage/sql/sqlEntities"
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
