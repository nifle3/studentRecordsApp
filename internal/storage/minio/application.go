package minio

import "context"

func (s *Storage) GetApplicationFile(link string, ctx context.Context) ([]byte, error) {
	return s.getFile(documentBucket, link, ctx)
}

func (s *Storage) AddApplicationFile(name string, file []byte, ctx context.Context) (string, error) {
	return s.addFile(documentBucket, name, file, ctx)
}

func (s *Storage) DeleteApplicationFile(link string, ctx context.Context) error {
	return s.deleteFile(documentBucket, link, ctx)
}

func (s *Storage) UpdateApplicationFile(file []byte, link string, ctx context.Context) (string, error) {
	return s.updateFile(documentBucket, file, link, ctx)
}
