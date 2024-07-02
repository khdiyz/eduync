package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/internal/storage"
	"edusync/pkg/logger"
	"io"
	"time"
)

type Service struct {
	Authorization
	Minio

	UserReader
	UserWriter

	RoleReader
}

func NewService(repos repository.Repository, storage storage.Storage, logger logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.UserRepo, logger),
		Minio:         NewMinioService(storage, logger),

		UserReader: NewUserReaderService(repos.UserRepo, logger),
		UserWriter: NewUserWriterService(repos, logger),

		RoleReader: NewRoleReaderService(repos.RoleRepo, logger),
	}
}

type Authorization interface {
	CreateToken(user model.User, tokenType string, expiresAt time.Time) (*model.Token, error)
	GenerateTokens(user model.User) (*model.Token, *model.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(input model.LoginRequest) (*model.Token, *model.Token, error)
}

type Minio interface {
	UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error)
	UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error)
	UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error)
}

// User Service
type UserReader interface {
	GetByUsername(username string) (model.User, error)
	GetById(id int64) (model.User, error)
}

type UserWriter interface {
	Create(input model.UserCreateRequest) (int64, error)
}

// Role Service
type RoleReader interface {
	GetList(pagination *model.Pagination) ([]model.Role, error)
}
