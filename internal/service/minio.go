package service

import (
	"edusync/internal/storage"
	"edusync/pkg/logger"
	"io"
)

type MinioService struct {
	storage storage.Storage
	logger  logger.Logger
}

func NewMinioService(storage storage.Storage, logger logger.Logger) *MinioService {
	return &MinioService{storage: storage, logger: logger}
}

func (m *MinioService) UploadImage(image io.Reader, imageSize int64, contextType string) (file storage.File, err error) {
	return m.storage.UploadStorage.UploadImage(image, imageSize, contextType)
}

func (m *MinioService) UploadDoc(doc io.Reader, docSize int64, contextType string) (file storage.File, err error) {
	return m.storage.UploadStorage.UploadDoc(doc, docSize, contextType)
}

func (m *MinioService) UploadExcel(doc io.Reader, docSize int64, contextType string) (file storage.File, err error) {
	return m.storage.UploadStorage.UploadExcel(doc, docSize, contextType)
}
