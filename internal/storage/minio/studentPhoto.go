package minio

import (
	"context"
	"io"
	"studentRecordsApp/internal/service"

	"github.com/minio/minio-go/v7"
)

const studentPhotoBucket = "student_photo"

var _ service.StudentFS = (*StudentPhoto)(nil)
var _ service.StudentFS = &StudentPhoto{}

type StudentPhoto struct {
	client *minio.Client
}

func NewStudentPhoto(ctx context.Context, client *minio.Client) *StudentPhoto {
	exists, err := client.BucketExists(ctx, studentPhotoBucket)
	if err != nil {
		return nil
	}

	if !exists {
		err = client.MakeBucket(ctx, studentPhotoBucket, minio.MakeBucketOptions{})
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
