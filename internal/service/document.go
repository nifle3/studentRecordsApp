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
	DocumentDb interface {
		Get(ctx context.Context, id uuid.UUID, userId uuid.UUID) (entities.Document, error)
		GetForUser(ctx context.Context, userId uuid.UUID) ([]entities.Document, error)
		Add(ctx context.Context, document entities.Document) error
		Delete(ctx context.Context, id uuid.UUID) error
		DeleteWithUserId(ctx context.Context, id, userId uuid.UUID) error
		Update(ctx context.Context, document entities.Document) error
	}

	DocumentFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Document struct {
	db DocumentDb
	fs DocumentFS
}

func NewDocument(db DocumentDb, fs DocumentFS) Document {
	return Document{
		db: db,
		fs: fs,
	}
}

func (d *Document) Add(ctx context.Context, document entities.Document, size int64) *customError.Http {
	if !document.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	id := uuid.New()
	document.Id = id
	document.Link = id.String()
	document.CreatedAt = time.Now()

	if err := d.fs.Add(ctx, document.Link, size, document.File); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := d.db.Add(ctx, document); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (d *Document) Update(ctx context.Context, document entities.Document) *customError.Http {
	if document.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	if err := d.db.Update(ctx, document); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (d *Document) UpdateFile(ctx context.Context, link string, size int64, file io.Reader) *customError.Http {
	if err := d.fs.Update(ctx, file, size, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (d *Document) Delete(ctx context.Context, id uuid.UUID, link string) *customError.Http {
	if err := d.db.Delete(ctx, id); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := d.fs.Delete(ctx, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (d *Document) DeleteByUserdId(ctx context.Context, id, userId uuid.UUID, link string) *customError.Http {
	if err := d.db.DeleteWithUserId(ctx, id, userId); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	if err := d.fs.Delete(ctx, link); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (d *Document) GetById(ctx context.Context, id, userId uuid.UUID) (entities.Document, *customError.Http) {
	document, err := d.db.Get(ctx, id, userId)
	if err != nil {
		return entities.Document{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	return document, nil
}

func (d *Document) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]entities.Document, *customError.Http) {
	documents, err := d.db.GetForUser(ctx, userId)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return documents, nil
}

func (d *Document) DownloadDocument(ctx context.Context, link string) ([]byte, *customError.Http) {
	result, err := d.fs.Get(ctx, link)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}
