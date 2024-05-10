package casts

import (
	"context"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func ApplicationServiceToSQL(_ context.Context, app entities.Application) sqlEntities.Application {
	return sqlEntities.Application{
		Id:          app.Id,
		StudentId:   app.StudentId,
		ContactInfo: app.ContactInfo,
		Name:        app.Name,
		Text:        app.Text,
		Status:      app.Status,
		CreatedAt:   app.CreatedAt,
		Link:        app.Link,
	}
}

func ApplicationSqlToService(_ context.Context, app sqlEntities.Application) entities.Application {
	return entities.Application{
		Id:          app.Id,
		StudentId:   app.StudentId,
		ContactInfo: app.ContactInfo,
		Name:        app.Name,
		Text:        app.Text,
		Status:      app.Status,
		CreatedAt:   app.CreatedAt,
		Link:        app.Link,
		File:        nil,
	}
}
