package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"studentRecordsApp/internal/service/entites"
)

type DocumentDb interface {
	GetStudentsDocumentById(id string, userId string, ctx context.Context) (entities.Document, error)
	GetStudentsDocumentsForUser(userId string, ctx context.Context) ([]entities.Document, error)
	AddStudentsDocument(document entities.Document, ctx context.Context) error
	DeleteStudentsDocument(id string, userId string, ctx context.Context) error
	UpdateStudentsDocument(document entities.Document, ctx context.Context) error
}

type DocumentFileStorage interface {
	GetDocumentFile(link string, ctx context.Context) ([]byte, error)
	AddDocumentFile(file []byte, ctx context.Context) (string, error)
	DeleteDocumentFile(link string, ctx context.Context) error
	UpdateDocumentFile(file []byte, link string, ctx context.Context) error
}

type Document struct {
	db          *DocumentDb
	fileStorage *DocumentFileStorage
}

func NewsDocument(db *DocumentDb) Document {
	return Document{
		db: db,
	}
}

func (d *Document) Add(document entities.Document, ctx context.Context) error {
	if !document.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	var err error
	document.Link, err = (*d.fileStorage).AddDocumentFile(document.File, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	document.CreatedAt = time.Now()

	return (*d.db).AddStudentsDocument(document, ctx)
}

func (d *Document) Update(document entities.Document, userId string, ctx context.Context) error {
	if document.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	document.CreatedAt = time.Now()

	err := (*d.db).UpdateStudentsDocument(document, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	err = (*d.fileStorage).UpdateDocumentFile(document.File, document.Link, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	return nil
}

func (d *Document) Delete(id, link, userId string, ctx context.Context) error {
	err := (*d.db).DeleteStudentsDocument(id, userId, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	err = (*d.fileStorage).DeleteDocumentFile(link, ctx)
	if err != nil {
		return fmt.Errorf("500")
	}

	return nil
}

func (d *Document) GetById(id, userId string, ctx context.Context) (entities.Document, error) {
	document, err := (*d.db).GetStudentsDocumentById(id, userId, ctx)
	if err != nil {
		return entities.Document{}, fmt.Errorf("500")
	}

	document.File, err = (*d.fileStorage).GetDocumentFile(document.Link, ctx)
	if err != nil {
		return entities.Document{}, fmt.Errorf("500")
	}

	return document, nil
}

func (d *Document) GetAllForUser(userId string, ctx context.Context) ([]entities.Document, error) {
	documents, err := (*d.db).GetStudentsDocumentsForUser(userId, ctx)
	if err != nil {
		return nil, fmt.Errorf("500")
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

			documents[idx].File, err = (*d.fileStorage).GetDocumentFile(documents[idx].Link, ctx)
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
