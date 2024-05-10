package service

import (
	"context"
	"io"
	"net/http"
	"studentRecordsApp/pkg/customError"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/pkg/password"
)

type (
	StudentDB interface {
		GetStudents(ctx context.Context) ([]entities.Student, error)
		GetStudent(ctx context.Context, id uuid.UUID) (entities.Student, error)
		AddStudent(ctx context.Context, student entities.Student) (uuid.UUID, error)
		UpdateStudent(ctx context.Context, student entities.Student) error
		DeleteStudent(ctx context.Context, id uuid.UUID) error
	}

	StudentPhoneDB interface {
		AddStudentPhone(ctx context.Context, phone entities.PhoneNumber) error
	}

	StudentFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Student struct {
	db      *StudentDB
	phoneDB *StudentPhoneDB
	fs      *StudentFS
}

func NewStudent(db *StudentDB, fs *StudentFS, phoneDB *StudentPhoneDB) Student {
	return Student{
		db:      db,
		fs:      fs,
		phoneDB: phoneDB,
	}
}

func (s Student) Add(ctx context.Context, student entities.Student, size int64) error {
	if !s.checkCorrectStudent(student, ctx) {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	id := uuid.New()
	student.Id = id
	student.LinkPhoto = id.String()

	var err error
	student.Password, err = password.Hash(student.Password)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := (*s.fs).Add(ctx, student.LinkPhoto, size, student.Photo); err != nil {
		return err
	}

	_, err = (*s.db).AddStudent(ctx, student)
	if err != nil {
		return err
	}

	for _, value := range student.PhoneNumbers {
		value.StudentId = id
		value.Id = uuid.New()

		if err := value.CheckCorrectNumber(); err != nil {
			return customError.New(http.StatusBadRequest, "Has an invalid phone number")
		}

		if err := (*s.phoneDB).AddStudentPhone(ctx, value); err != nil {
			return err
		}
	}

	return nil
}

func (s Student) Update(ctx context.Context, student entities.Student, size int64) error {
	if !s.checkCorrectStudent(student, ctx) {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	if err := (*s.fs).Update(ctx, student.Photo, size, student.LinkPhoto); err != nil {
		return err
	}

	return (*s.db).UpdateStudent(ctx, student)
}

func (s Student) checkCorrectStudent(student entities.Student, _ context.Context) bool {
	return student.CheckIsNotEmpty() && student.CheckNumber() && student.CheckPassportSeria() &&
		student.CheckBirthdate() &&
		student.CheckPassword()
}

func (s Student) Delete(ctx context.Context, id uuid.UUID) error {
	return (*s.db).DeleteStudent(ctx, id)
}

func (s Student) Get(ctx context.Context, id uuid.UUID) (entities.Student, error) {
	student, err := (*s.db).GetStudent(ctx, id)
	if err != nil {
		return entities.Student{}, err
	}

	_, err = (*s.fs).Get(ctx, student.LinkPhoto)

	return student, err
}

func (s Student) GetAll(ctx context.Context) ([]entities.Student, error) {
	return (*s.db).GetStudents(ctx)
}
