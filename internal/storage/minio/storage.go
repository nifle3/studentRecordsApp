package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"studentRecordsApp/internal/service"
)

var (
	_ service.DocumentFS    = (*Storage)(nil)
	_ service.StudentFS     = (*Storage)(nil)
	_ service.ApplicationFS = (*Storage)(nil)
)

const (
	documentBucket     = "Document"
	studentPhotoBucket = "StudentPhoto"
	applicationBucket  = "Application"
)

type Storage struct {
	client *minio.Client
}

func New(endpoint, password, login string, isSecure bool, ctx context.Context) (*Storage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(login, password, ""),
		Secure: isSecure,
	})
	if err != nil {
		return nil, err
	}

	if client.IsOffline() {
		return nil, err
	}

	if err := checkBucketExist(client, applicationBucket, ctx); err != nil {
		return nil, err
	}

	if err := checkBucketExist(client, documentBucket, ctx); err != nil {
		return nil, err
	}

	if err := checkBucketExist(client, studentPhotoBucket, ctx); err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
	}, nil
}

func checkBucketExist(client *minio.Client, bucketName string, ctx context.Context) error {
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	return err

}

func (s *Storage) getFile(bucketName string, link string, ctx context.Context) ([]byte, error) {
	object, err := s.client.GetObject(ctx, bucketName, link, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	result := make([]byte, 0)

	_, err = object.Read(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Storage) addFile(bucketName string, name string, file []byte, ctx context.Context) (string, error) {
	reader := bytes.NewReader(file)
	result, err := s.client.PutObject(ctx, bucketName, name, reader, int64(len(file)),
		minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return result.Key, nil
}

func (s *Storage) deleteFile(bucketName string, link string, ctx context.Context) error {
	return s.client.RemoveObjectTagging(ctx, bucketName, link, minio.RemoveObjectTaggingOptions{})
}

func (s *Storage) updateFile(bucketName string, file []byte, link string, ctx context.Context) (string, error) {
	err := s.client.RemoveObjectTagging(ctx, bucketName, link, minio.RemoveObjectTaggingOptions{})
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(file)
	result, err := s.client.PutObject(ctx, bucketName, link, reader, int64(len(file)),
		minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return result.Key, nil
}
