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

func NewDocument(ctx context.Context, client *minio.Client) *Document {
	exists, err := client.BucketExists(ctx, documentBucket)
	if err != nil {
		return nil
	}

	if !exists {
		err = client.MakeBucket(ctx, documentBucket, minio.MakeBucketOptions{})
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
	info, err := object.Stat()
	if err != nil {
		return nil, err
	}

	result := make([]byte, 0, info.Size)
	for {
		_, err := object.Read(result)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *Document) Add(ctx context.Context, name string, size int64, file io.Reader) error {
	_, err := s.client.PutObject(ctx, documentBucket, name, file, size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
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
