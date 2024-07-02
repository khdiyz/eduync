package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"edusync/pkg/response"

	"google.golang.org/grpc/codes"
)

type UserReaderService struct {
	repo   repository.UserRepo
	logger logger.Logger
}

func NewUserReaderService(repo repository.UserRepo, logger logger.Logger) *UserReaderService {
	return &UserReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserReaderService) GetByUsername(username string) (model.User, error) {
	user, err := s.repo.UserReader.GetByUsername(username)
	if err != nil {
		return model.User{}, response.ServiceError(err, codes.Internal)
	}

	return user, err
}

func (s *UserReaderService) GetById(id int64) (model.User, error) {
	user, err := s.repo.UserReader.GetById(id)
	if err != nil {
		return model.User{}, response.ServiceError(err, codes.Internal)
	}

	return user, err
}
