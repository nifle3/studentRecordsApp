package service

import (
	"context"
	"fmt"
	"studentRecordsApp/internal/service/entites"
)

type StudentDb interface {
	GetStudents(ctx context.Context) ([]entities.Student, error)
	GetStudent(id string, ctx context.Context) (entities.Student, error)
	AddStudent(student entities.Student, ctx context.Context) error
	UpdateStudent(student entities.Student, ctx context.Context) error
	DeleteStudent(id string, ctx context.Context) error
	GetStudentByEmail(email string, ctx context.Context) (entities.Student, error)
}

type Student struct {
	db *StudentDb
}

func NewStudent(db *StudentDb) Student {
	return Student{
		db: db,
	}
}

func (s Student) Add(student entities.Student, ctx context.Context) error {
	if !s.checkCorrectStudent(student, ctx) {
		return fmt.Errorf("400")
	}

	err := student.HashPassword()
	if err != nil {
		return err
	}

	return (*s.db).AddStudent(student, ctx)
}

func (s Student) Update(student entities.Student, ctx context.Context) error {
	if !s.checkCorrectStudent(student, ctx) {
		return fmt.Errorf("400")
	}

	return (*s.db).UpdateStudent(student, ctx)
}

func (s Student) checkCorrectStudent(student entities.Student, _ context.Context) bool {
	emailResult, err := student.CheckEmail()
	if err != nil {
		return false
	}

	return student.CheckIsNotEmpty() && student.CheckNumber() && student.CheckPassportSeria() &&
		student.CheckBirthdate() && emailResult &&
		student.CheckPassword()
}

func (s Student) Delete(id string, ctx context.Context) error {
	return (*s.db).DeleteStudent(id, ctx)
}

func (s Student) Get(id string, ctx context.Context) (entities.Student, error) {
	return (*s.db).GetStudent(id, ctx)
}

func (s Student) GetAll(ctx context.Context) ([]entities.Student, error) {
	return (*s.db).GetStudents(ctx)
}

func (s Student) Login(email, password string, ctx context.Context) (entities.Student, bool, error) {
	result, err := (*s.db).GetStudentByEmail(email, ctx)
	if err != nil {
		return entities.Student{}, false, err
	}

	loginError := result.CheckHash(password)
	if loginError != nil {
		return entities.Student{}, false, loginError
	}

	return result, true, nil
}
