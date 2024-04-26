package minio

import "context"

func (s *Storage) GetDocumentFile(link string, ctx context.Context) ([]byte, error) {
	return s.getFile(documentBucket, link, ctx)
}

func (s *Storage) AddDocumentFile(name string, file []byte, ctx context.Context) (string, error) {
	return s.addFile(documentBucket, name, file, ctx)
}

func (s *Storage) DeleteDocumentFile(link string, ctx context.Context) error {
	return s.deleteFile(documentBucket, link, ctx)
}

func (s *Storage) UpdateDocumentFile(file []byte, link string, ctx context.Context) (string, error) {
	return s.updateFile(documentBucket, file, link, ctx)
}
