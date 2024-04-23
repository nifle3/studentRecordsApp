package service

import (
	"context"
	entities "studentRecordsApp/internal/service/entites"
)

type StudentDb interface {
	GetStudents(ctx context.Context) ([]entities.Student, error)
	GetStudent(id int, ctx context.Context) (entities.Student, error)
	AddStudent(student entities.Student, phones []entities.PhoneNumber, ctx context.Context) error
	UpdateStudent(student entities.Student, ctx context.Context) error
	DeleteStudent(id int, ctx context.Context) error
}

type Student struct {
	db *StudentDb
}
