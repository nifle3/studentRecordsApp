package minio

import "context"

func (s *Storage) GetPhotoStudentFile(link string, ctx context.Context) ([]byte, error) {
	return s.getFile(studentPhotoBucket, link, ctx)
}

func (s *Storage) AddPhotoStudentFile(name string, file []byte, ctx context.Context) (string, error) {
	return s.addFile(studentPhotoBucket, name, file, ctx)
}

func (s *Storage) DeletePhotoStudentFile(link string, ctx context.Context) error {
	return s.deleteFile(studentPhotoBucket, link, ctx)
}

func (s *Storage) UpdatePhotoStudentFile(file []byte, link string, ctx context.Context) (string, error) {
	return s.updateFile(studentPhotoBucket, file, link, ctx)
}
