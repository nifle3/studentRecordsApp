package service

import (
	"context"
	"io"
	"net/http"
	"studentRecordsApp/pkg/customError"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entities"
)

type (
	ApplicationDb interface {
		Get(ctx context.Context) ([]entities.Application, error)
		GetForUser(ctx context.Context, userId uuid.UUID) ([]entities.Application, error)
		GetById(ctx context.Context, id, userID uuid.UUID) (entities.Application, error)
		GetOne(ctx context.Context, id uuid.UUID) (entities.Application, error)
		Add(ctx context.Context, application entities.Application) error
		Update(ctx context.Context, application entities.Application) error
		UpdateStatus(ctx context.Context, applicationID uuid.UUID, status string) error
		Delete(ctx context.Context, id, userId uuid.UUID) error
		CheckLinkWithUserId(ctx context.Context, id uuid.UUID, link string) (bool, error)
		GetWithInfo(ctx context.Context) ([]entities.ApplicationWithInfo, error)
	}

	ApplicationFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Application struct {
	db ApplicationDb
	fs ApplicationFS
}

func NewApplication(db ApplicationDb, fs ApplicationFS) Application {
	return Application{
		db: db,
		fs: fs,
	}
}

func (a *Application) GetAll(ctx context.Context) ([]entities.Application, *customError.Http) {
	result, err := a.db.Get(ctx)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func (a *Application) GetAllWithInfo(ctx context.Context) ([]entities.ApplicationWithInfo, *customError.Http) {
	result, err := a.db.GetWithInfo(ctx)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func (a *Application) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]entities.Application, *customError.Http) {
	result, err := a.db.GetForUser(ctx, userId)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func (a *Application) GetByIdAndUserId(ctx context.Context, id, userId uuid.UUID) (entities.Application, *customError.Http) {
	application, err := a.db.GetById(ctx, id, userId)
	if err != nil {
		return entities.Application{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	_, err = a.fs.Get(ctx, application.Link)
	if err != nil {
		return entities.Application{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	return application, nil
}

func (a *Application) GetById(ctx context.Context, id uuid.UUID) (entities.Application, *customError.Http) {
	application, err := a.db.GetOne(ctx, id)
	if err != nil {
		return entities.Application{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	_, err = a.fs.Get(ctx, application.Link)
	if err != nil {
		return entities.Application{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	return application, nil
}

func (a *Application) Add(ctx context.Context, application entities.Application, size int64) *customError.Http {
	application.Status = entities.ApplicationCreated

	if !application.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	id := uuid.New()
	application.CreatedAt = time.Now()
	application.Link = id.String()
	application.Id = id

	if err := a.fs.Add(ctx, application.Link, size, application.File); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := a.db.Add(ctx, application); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (a *Application) Update(ctx context.Context, application entities.Application, size int64) *customError.Http {
	if !application.CheckStatus() {
		return customError.New(http.StatusBadRequest, "Invalid status")
	}

	if !application.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	if err := a.db.Update(ctx, application); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := a.fs.Update(ctx, application.File, size, application.Link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (a *Application) Delete(ctx context.Context, applicationID, userId uuid.UUID, link string) *customError.Http {
	if err := a.db.Delete(ctx, applicationID, userId); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := a.fs.Delete(ctx, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (a *Application) ChangeStatusToFinish(ctx context.Context, applicationID uuid.UUID) *customError.Http {
	if err := a.db.UpdateStatus(ctx, applicationID, entities.ApplicationClosed); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (a *Application) Download(ctx context.Context, link string) ([]byte, *customError.Http) {
	file, err := a.fs.Get(ctx, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return file, nil
}

func (a *Application) DownloadWithCheckId(ctx context.Context, link string, id uuid.UUID) ([]byte, *customError.Http) {
	result, err := a.db.CheckLinkWithUserId(ctx, id, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	if !result {
		return nil, customError.New(http.StatusUnauthorized, "Invalid link")
	}

	file, err := a.fs.Get(ctx, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return file, nil
}
