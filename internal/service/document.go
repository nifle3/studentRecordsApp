package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"studentRecordsApp/pkg/customError"
	"sync"
	"time"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
)

type (
	DocumentDb interface {
		GetStudentsDocumentById(ctx context.Context, id uuid.UUID, userId uuid.UUID) (entities.Document, error)
		GetStudentsDocumentsForUser(ctx context.Context, userId uuid.UUID) ([]entities.Document, error)
		AddStudentsDocument(ctx context.Context, document entities.Document) error
		DeleteStudentsDocument(ctx context.Context, id uuid.UUID, userId uuid.UUID) error
		UpdateStudentsDocument(ctx context.Context, document entities.Document) error
	}

	DocumentFS interface {
		Get(ctx context.Context, link string) ([]byte, error)
		Add(ctx context.Context, name string, size int64, file io.Reader) error
		Delete(ctx context.Context, link string) error
		Update(ctx context.Context, file io.Reader, size int64, link string) error
	}
)

type Document struct {
	db *DocumentDb
	fs *DocumentFS
}

func NewDocument(db DocumentDb, fs DocumentFS) Document {
	return Document{
		db: &db,
		fs: &fs,
	}
}

func (d *Document) Add(ctx context.Context, document entities.Document, size int64) error {
	if !document.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	id := uuid.New()
	document.Id = id
	document.Link = id.String()
	document.CreatedAt = time.Now()

	if err := (*d.fs).Add(ctx, document.Name, size, document.File); err != nil {
		return err
	}

	return (*d.db).AddStudentsDocument(ctx, document)
}

func (d *Document) Update(ctx context.Context, document entities.Document, size int64) error {
	if document.CheckIsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has an empty field")
	}

	if err := (*d.fs).Update(ctx, document.File, size, document.Link); err != nil {
		return err
	}

	return (*d.db).UpdateStudentsDocument(ctx, document)
}

func (d *Document) Delete(ctx context.Context, id, userId uuid.UUID, link string) error {
	err := (*d.db).DeleteStudentsDocument(ctx, id, userId)
	if err != nil {
		return err
	}

	return (*d.fs).Delete(ctx, link)
}

func (d *Document) GetById(ctx context.Context, id, userId uuid.UUID) (entities.Document, error) {
	document, err := (*d.db).GetStudentsDocumentById(ctx, id, userId)
	if err != nil {
		return entities.Document{}, err
	}

	_, err = (*d.fs).Get(ctx, document.Link)
	if err != nil {
		return entities.Document{}, err
	}

	return document, nil
}

func (d *Document) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]entities.Document, error) {
	documents, err := (*d.db).GetStudentsDocumentsForUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(documents))

	errChan := make(chan error)
	defer close(errChan)

	quitChan := make(chan interface{})
	defer close(quitChan)

	for idx := range documents {
		go func() {
			defer wg.Done()

			_, err = (*d.fs).Get(ctx, documents[idx].Link)
			if err != nil {
				errChan <- fmt.Errorf("500")
			}
		}()
	}

	go func() {
		wg.Wait()
		quitChan <- struct{}{}
	}()

	for {
		select {
		case <-quitChan:
			return documents, nil
		case err := <-errChan:
			return nil, err
		}
	}
}
