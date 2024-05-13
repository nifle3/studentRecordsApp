package minio

import (
	"context"
	"io"
	"studentRecordsApp/internal/service"

	"github.com/minio/minio-go/v7"
)

const studentPhotoBucket = "student-photo"

var _ service.StudentFS = (*StudentPhoto)(nil)
var _ service.StudentFS = &StudentPhoto{}

type StudentPhoto struct {
	client *minio.Client
}

func MustNewStudentPhoto(ctx context.Context, client *minio.Client) *StudentPhoto {
	result, err := client.BucketExists(ctx, studentPhotoBucket)

	if !result {
		err = client.MakeBucket(ctx, studentPhotoBucket, minio.MakeBucketOptions{})
		if err != nil {
			panic(err.Error())
		}
	}

	return &StudentPhoto{
		client: client,
	}
}

func (s *StudentPhoto) Get(ctx context.Context, link string) ([]byte, error) {
	object, err := s.client.GetObject(ctx, studentPhotoBucket, link, minio.GetObjectOptions{})
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

func (s *StudentPhoto) Add(ctx context.Context, name string, size int64, file io.Reader) error {
	_, err := s.client.PutObject(ctx, studentPhotoBucket, name, file, size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentPhoto) Delete(ctx context.Context, link string) error {
	return s.client.RemoveObject(ctx, studentPhotoBucket, link, minio.RemoveObjectOptions{})
}

func (s *StudentPhoto) Update(ctx context.Context, file io.Reader, size int64, link string) error {
	if err := s.Delete(ctx, link); err != nil {
		return err
	}

	return s.Add(ctx, link, size, file)
}
