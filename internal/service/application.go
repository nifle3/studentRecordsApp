package service

import (
	"context"
	"fmt"
	"time"

	"studentRecordsApp/internal/service/entites"
)

type ApplicationDb interface {
	GetApplications(ctx context.Context) ([]entities.Application, error)
	GetApplicationForUser(userId string, ctx context.Context) ([]entities.Application, error)
	GetApplicationById(id string, ctx context.Context) (entities.Application, error)
	AddApplication(application entities.Application, ctx context.Context) error
	UpdateApplication(application entities.Application, ctx context.Context) error
	DeleteApplication(id, userId string, ctx context.Context) error
}

type ApplicationFS interface {
	GetApplicationFile(link string, ctx context.Context) ([]byte, error)
	AddApplicationFile(name string, file []byte, ctx context.Context) (string, error)
	DeleteApplicationFile(link string, ctx context.Context) error
	UpdateApplicationFile(file []byte, link string, ctx context.Context) (string, error)
}

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

func (a *Application) GetAllForUser(userId string, ctx context.Context) ([]entities.Application, error) {
	return (*a.db).GetApplicationForUser(userId, ctx)
}

func (a *Application) GetById(id string, ctx context.Context) (entities.Application, error) {
	application, err := (*a.db).GetApplicationById(id, ctx)
	if err != nil {
		return entities.Application{}, fmt.Errorf("500")
	}

	application.File, err = (*a.fs).GetApplicationFile(application.Link, ctx)
	if err != nil {
		return entities.Application{}, fmt.Errorf("500")
	}

	return application, nil
}

func (a *Application) Add(application entities.Application, ctx context.Context) error {
	if !application.CheckStatus() {
		return fmt.Errorf("400")
	}

	if !application.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	application.CreatedAt = time.Now()

	var err error
	application.Link, err = (*a.fs).AddApplicationFile(application.Name, application.File, ctx)
	if err != nil {
		return err
	}

	return (*a.db).AddApplication(application, ctx)
}

func (a *Application) Update(application entities.Application, ctx context.Context) error {
	if !application.CheckStatus() {
		return fmt.Errorf("400")
	}

	if !application.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	application.CreatedAt = time.Now()

	return (*a.db).UpdateApplication(application, ctx)
}

func (a *Application) Delete(application entities.Application, userId string, ctx context.Context) error {
	err := (*a.fs).DeleteApplicationFile(application.Link, ctx)
	if err != nil {
		return err
	}

	return (*a.db).DeleteApplication(application.Id, userId, ctx)
}

func (a *Application) ChangeStatusToFinish(application entities.Application, ctx context.Context) error {
	application.Status = "Закрыт"

	return (*a.db).UpdateApplication(application, ctx)
}
