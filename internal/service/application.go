package service

import (
	"context"
	"io"
	"net/http"
	"studentRecordsApp/pkg/customError"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
)

type (
	ApplicationDb interface {
		GetApplications(ctx context.Context) ([]entities.Application, error)
		GetApplicationForUser(ctx context.Context, userId uuid.UUID) ([]entities.Application, error)
		GetApplicationById(ctx context.Context, id uuid.UUID) (entities.Application, error)
		AddApplication(ctx context.Context, application entities.Application) error
		UpdateApplication(ctx context.Context, application entities.Application) error
		UpdateStatusApplication(ctx context.Context, applicationID uuid.UUID, status string) error
		DeleteApplication(ctx context.Context, id, userId uuid.UUID) error
	}

	ApplicationFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Application struct {
	db *ApplicationDb
	fs *ApplicationFS
}

func NewApplication(db ApplicationDb, fs ApplicationFS) Application {
	return Application{
		db: &db,
		fs: &fs,
	}
}

func (a *Application) GetAll(ctx context.Context) ([]entities.Application, error) {
	return (*a.db).GetApplications(ctx)
}

func (a *Application) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]entities.Application, error) {
	return (*a.db).GetApplicationForUser(ctx, userId)
}

func (a *Application) GetById(ctx context.Context, id uuid.UUID) (entities.Application, error) {
	application, err := (*a.db).GetApplicationById(ctx, id)
	if err != nil {
		return entities.Application{}, err
	}

	_, err = (*a.fs).Get(ctx, application.Link)
	if err != nil {
		return entities.Application{}, err
	}

	return application, nil
}

func (a *Application) Add(ctx context.Context, application entities.Application, size int64) error {
	if !application.CheckStatus() {
		return customError.New(http.StatusBadRequest, "Has an invalid status")
	}

	if !application.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	id := uuid.New()
	application.CreatedAt = time.Now()
	application.Link = id.String()
	application.Id = id

	if err := (*a.fs).Add(ctx, application.Link, size, application.File); err != nil {
		return err
	}

	return (*a.db).AddApplication(ctx, application)
}

func (a *Application) Update(ctx context.Context, application entities.Application) error {
	if !application.CheckStatus() {
		return customError.New(http.StatusBadRequest, "Invalid status")
	}

	if !application.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	return (*a.db).UpdateApplication(ctx, application)
}

func (a *Application) Delete(ctx context.Context, applicationID, userId uuid.UUID, link string) error {
	if err := (*a.db).DeleteApplication(ctx, applicationID, userId); err != nil {
		return err
	}

	return (*a.fs).Delete(ctx, link)
}

func (a *Application) ChangeStatusToFinish(applicationID uuid.UUID, ctx context.Context) error {
	return (*a.db).UpdateStatusApplication(ctx, applicationID, entities.ApplicationClosed)
}
