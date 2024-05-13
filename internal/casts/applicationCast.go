package casts

import (
	"context"
	"studentRecordsApp/internal/transport/server/httpEntity"

	"studentRecordsApp/internal/service/entities"
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

func ApplicationEntitieToApplicationGet(_ context.Context, app entities.Application) (httpEntity.ApplicationGet, error) {
	id, err := UuidToString(app.Id)
	if err != nil {
		return httpEntity.ApplicationGet{}, err
	}

	return httpEntity.ApplicationGet{
		Id:          id,
		ContactInfo: app.ContactInfo,
		Name:        app.Name,
		Text:        app.Text,
		Status:      app.Status,
		CreatedAt:   app.CreatedAt,
		Link:        app.Link,
	}, nil
}

func ApplicationWithInfoToEntite(_ context.Context, app sqlEntities.ApplicationWithInfo) entities.ApplicationWithInfo {
	return entities.ApplicationWithInfo{
		Id:          app.Id,
		StudentId:   app.StudentId,
		ContactInfo: app.ContactInfo,
		Name:        app.Name,
		Text:        app.Text,
		Status:      app.Status,
		CreatedAt:   app.CreatedAt,
		Link:        app.Link,
		FIO:         app.Surname + " " + app.FirstName + " " + app.LastName,
		Course:      app.Course,
		Group:       app.Group,
	}
}

func ApplicationWithInfoEntitieToHttp(_ context.Context, app entities.ApplicationWithInfo) (httpEntity.ApplicationWithInfo, error) {
	id, err := UuidToString(app.Id)
	if err != nil {
		return httpEntity.ApplicationWithInfo{}, err
	}

	return httpEntity.ApplicationWithInfo{
		Id:          id,
		ContactInfo: app.ContactInfo,
		Name:        app.Name,
		Text:        app.Text,
		Status:      app.Status,
		CreatedAt:   app.CreatedAt,
		Link:        app.Link,
		FIO:         app.FIO,
		Course:      app.Course,
		Group:       app.Group,
	}, nil
}
