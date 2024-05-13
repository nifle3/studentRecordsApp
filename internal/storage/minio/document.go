package minio

import (
	"context"
	"io"
	"studentRecordsApp/internal/service"

	"github.com/minio/minio-go/v7"
)

const documentBucket = "document"

var _ service.DocumentFS = (*Document)(nil)
var _ service.DocumentFS = &Document{}

type Document struct {
	client *minio.Client
}

func MustNewDocument(ctx context.Context, client *minio.Client) *Document {
	result, err := client.BucketExists(ctx, documentBucket)
	if err != nil {
		panic(err.Error())
	}

	if !result {
		err = client.MakeBucket(ctx, documentBucket, minio.MakeBucketOptions{})
		if err != nil {
			panic(err.Error())
		}
	}

	return &Document{
		client: client,
	}
}

func (s *Document) Get(ctx context.Context, link string) ([]byte, error) {
	object, err := s.client.GetObject(ctx, documentBucket, link, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer object.Close()

	result, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Document) Add(ctx context.Context, name string, size int64, file io.Reader) error {
	_, err := s.client.PutObject(ctx, documentBucket, name, file, size, minio.PutObjectOptions{
		ContentType: "application/pdf",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Document) Delete(ctx context.Context, link string) error {
	return s.client.RemoveObject(ctx, documentBucket, link, minio.RemoveObjectOptions{})
}

func (s *Document) Update(ctx context.Context, file io.Reader, size int64, link string) error {
	if err := s.Delete(ctx, link); err != nil {
		return err
	}

	return s.Add(ctx, link, size, file)
}
