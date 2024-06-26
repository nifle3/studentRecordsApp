package service

import (
	"context"
	"io"
	"net/http"
	"studentRecordsApp/pkg/customError"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/pkg/password"
)

type (
	StudentDB interface {
		GetAll(ctx context.Context) ([]entities.Student, error)
		Get(ctx context.Context, id uuid.UUID) (entities.Student, error)
		GetLinkById(ctx context.Context, id uuid.UUID) (string, error)
		Add(ctx context.Context, student entities.Student) error
		Update(ctx context.Context, student entities.Student) error
		Delete(ctx context.Context, id uuid.UUID) error
	}

	StudentPhoneDB interface {
		Add(ctx context.Context, phone entities.PhoneNumber) error
		GetForUser(ctx context.Context, id uuid.UUID) ([]entities.PhoneNumber, error)
		Update(ctx context.Context, phone entities.PhoneNumber) error
	}

	StudentFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Student struct {
	db      StudentDB
	phoneDB StudentPhoneDB
	fs      StudentFS
}

func NewStudent(db StudentDB, fs StudentFS, phoneDB StudentPhoneDB) Student {
	return Student{
		db:      db,
		fs:      fs,
		phoneDB: phoneDB,
	}
}

func (s Student) Add(ctx context.Context, student entities.Student, size int64) *customError.Http {
	if !s.checkCorrectStudent(student, ctx) {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	id := uuid.New()
	student.Id = id
	student.LinkPhoto = id.String()
	student.EnrollYear = time.Now()

	var err error
	student.Password, err = password.Hash(student.Password)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := s.fs.Add(ctx, student.LinkPhoto, size, student.Photo); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	err = s.db.Add(ctx, student)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	for _, value := range student.PhoneNumbers {
		value.StudentId = id
		value.Id = uuid.New()

		if err := value.CheckCorrectNumber(); err != nil {
			return customError.New(http.StatusBadRequest, "Has an invalid phone number")
		}

		if err := s.phoneDB.Add(ctx, value); err != nil {
			return customError.New(http.StatusInternalServerError, err.Error())
		}
	}

	return nil
}

func (s Student) checkCorrectStudent(student entities.Student, _ context.Context) bool {
	return student.CheckIsNotEmpty() && student.CheckNumber() && student.CheckPassportSeria() &&
		student.CheckBirthdate() &&
		student.CheckPassword() && student.CheckEmail()
}

func (s Student) Update(ctx context.Context, student entities.Student) *customError.Http {
	if !s.checkCorrectStudentForUpdate(student, ctx) {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	if err := s.db.Update(ctx, student); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s Student) checkCorrectStudentForUpdate(student entities.Student, _ context.Context) bool {
	return student.CheckIsNotEmptyForUpdate() && student.CheckNumber() && student.CheckPassportSeria() &&
		student.CheckBirthdate() && student.CheckEmail()
}

func (s Student) UpdateImage(ctx context.Context, file io.Reader, size int64, link string) *customError.Http {
	if err := s.fs.Update(ctx, file, size, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s Student) Delete(ctx context.Context, id uuid.UUID, link string) *customError.Http {
	if err := s.db.Delete(ctx, id); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := s.fs.Delete(ctx, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s Student) Get(ctx context.Context, id uuid.UUID) (entities.Student, *customError.Http) {
	student, err := s.db.Get(ctx, id)
	if err != nil {
		return entities.Student{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	student.PhoneNumbers, err = s.phoneDB.GetForUser(ctx, id)
	if err != nil {
		return entities.Student{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	return student, nil
}

func (s Student) GetImage(ctx context.Context, link string) ([]byte, *customError.Http) {
	file, err := s.fs.Get(ctx, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return file, nil
}

func (s Student) GetImageById(ctx context.Context, id uuid.UUID) ([]byte, *customError.Http) {
	link, err := s.db.GetLinkById(ctx, id)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	file, err := s.fs.Get(ctx, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return file, nil
}

func (s Student) GetAll(ctx context.Context) ([]entities.Student, *customError.Http) {
	result, err := s.db.GetAll(ctx)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}
