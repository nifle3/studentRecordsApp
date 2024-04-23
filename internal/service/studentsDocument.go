package service

import (
	"context"
	entities "studentRecordsApp/internal/service/entites"
)

type StudentsDocumentDb interface {
	GetStudentsDocument(id string, ctx context.Context) (entities.StudentsDocument, error)
	GetStudentsDocuments(ctx context.Context) ([]entities.StudentsDocument, error)
	AddStudentsDocument(document entities.StudentsDocument, ctx context.Context) error
	DeleteStudentsDocument(id string, ctx context.Context) error
	UpdateStudentsDocument(document entities.StudentsDocument, ctx context.Context) error
}

type StudentsDocument struct {
	db *StudentsDocumentDb
}
