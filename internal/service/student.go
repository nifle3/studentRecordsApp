package service

import (
	"context"
	"fmt"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/pkg/password"
)

type StudentDb interface {
	GetStudents(ctx context.Context) ([]entities.entities, error)
	GetStudent(id string, ctx context.Context) (entities.Student, error)
	AddStudent(student entities.Student, ctx context.Context) error
	UpdateStudent(student entities.Student, ctx context.Context) error
	DeleteStudent(id string, ctx context.Context) error
	GetStudentByEmail(email string, ctx context.Context) (entities.Student, error)
}

type StudentFS interface {
	GetPhotoStudentFile(link string, ctx context.Context) ([]byte, error)
	AddPhotoStudentFile(name string, file []byte, ctx context.Context) (string, error)
	DeletePhotoStudentFile(link string, ctx context.Context) error
	UpdatePhotoStudentFile(file []byte, link string, ctx context.Context) (string, error)
}

type Student struct {
	db *StudentDb
	fs *StudentFS
}

func NewStudent(db StudentDb, fs StudentFS) Student {
	return Student{
		db: &db,
		fs: &fs,
	}
}

func (s Student) Add(student entities.Student, ctx context.Context) error {
	if !s.checkCorrectStudent(student, ctx) {
		return fmt.Errorf("400")
	}

	var err error
	student.Password, err = password.Hash(student.Password)
	if err != nil {
		return fmt.Errorf("500")
	}

	student.LinkPhoto, err = (*s.fs).AddPhotoStudentFile(student.GetFullName(), student.Photo, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	err = (*s.db).AddStudent(student, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	return nil
}

func (s Student) Update(student entities.Student, ctx context.Context) error {
	if !s.checkCorrectStudent(student, ctx) {
		return fmt.Errorf("400")
	}

	var err error
	student.LinkPhoto, err = (*s.fs).UpdatePhotoStudentFile(student.Photo, student.GetFullName(), ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	err = (*s.db).UpdateStudent(student, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	return nil
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
	student, err := (*s.db).GetStudent(id, ctx)
	if err != nil {
		return entities.Student{}, err
	}

	student.Photo, err = (*s.fs).GetPhotoStudentFile(student.LinkPhoto, ctx)

	return student, err
}

func (s Student) GetAll(ctx context.Context) ([]entities.Student, error) {
	return (*s.db).GetStudents(ctx)
}

func (s Student) Login(email, pass string, ctx context.Context) (entities.Student, error) {
	result, err := (*s.db).GetStudentByEmail(email, ctx)
	if err != nil {
		return entities.Student{}, err
	}

	return result, password.CheckHash(pass, []byte(result.Password))
}
