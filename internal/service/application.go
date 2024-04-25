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

type Application struct {
	db *ApplicationDb
}

func NewApplication(db *ApplicationDb) Application {
	return Application{
		db: db,
	}
}

func (a *Application) GetAll(ctx context.Context) ([]entities.Application, error) {
	return (*a.db).GetApplications(ctx)
}

func (a *Application) GetAllForUser(userId string, ctx context.Context) ([]entities.Application, error) {
	return (*a.db).GetApplicationForUser(userId, ctx)
}

func (a *Application) GetById(id string, ctx context.Context) (entities.Application, error) {
	return (*a.db).GetApplicationById(id, ctx)
}

func (a *Application) Add(application entities.Application, ctx context.Context) error {
	if !application.CheckStatus() {
		return fmt.Errorf("400")
	}

	if !application.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	application.CreatedAt = time.Now()

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

func (a *Application) Delete(id, userID string, ctx context.Context) error {
	return (*a.db).DeleteApplication(id, userID, ctx)
}

func (a *Application) ChangeStatusToFinish(application entities.Application, ctx context.Context) error {
	application.Status = "Закрыт"

	return (*a.db).UpdateApplication(application, ctx)
}
