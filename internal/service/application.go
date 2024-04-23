package service

import (
	"context"
	"studentRecordsApp/internal/service/entites"
)

type ApplicationDb interface {
	GetApplications(ctx context.Context) ([]entities.Application, error)
	GetApplicationById(id string, ctx context.Context) (entities.Application, error)
	AddApplication(application entities.Application, ctx context.Context) error
	UpdateApplication(application entities.Application, ctx context.Context) error
	DeleteApplication(id string, ctx context.Context) error
}

type Application struct {
	db *ApplicationDb
}
