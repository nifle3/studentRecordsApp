package minio

import (
	"context"
	"io"
	"studentRecordsApp/internal/service"

	"github.com/minio/minio-go/v7"
)

const applicationBucket = "application"

var _ service.ApplicationFS = (*Application)(nil)
var _ service.ApplicationFS = &Application{}

type Application struct {
	client *minio.Client
}

func MustNewApplication(ctx context.Context, client *minio.Client) *Application {
	result, err := client.BucketExists(ctx, applicationBucket)
	if err != nil {
		panic(err.Error())
	}

	if !result {
		err = client.MakeBucket(ctx, applicationBucket, minio.MakeBucketOptions{})
		if err != nil {
			panic(err.Error())
		}
	}

	return &Application{
		client: client,
	}
}

func (s *Application) Get(ctx context.Context, link string) ([]byte, error) {
	object, err := s.client.GetObject(ctx, applicationBucket, link, minio.GetObjectOptions{})
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

func (s *Application) Add(ctx context.Context, name string, size int64, file io.Reader) error {
	_, err := s.client.PutObject(ctx, applicationBucket, name, file, size, minio.PutObjectOptions{
		ContentType: "application/pdf",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Application) Delete(ctx context.Context, link string) error {
	return s.client.RemoveObject(ctx, applicationBucket, link, minio.RemoveObjectOptions{})
}

func (s *Application) Update(ctx context.Context, file io.Reader, size int64, link string) error {
	if err := s.Delete(ctx, link); err != nil {
		return err
	}

	return s.Add(ctx, link, size, file)
}
