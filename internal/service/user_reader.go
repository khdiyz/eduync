package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"

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
		return model.User{}, serviceError(err, codes.Internal)
	}

	return user, err
}

func (s *UserReaderService) GetById(id int64) (model.User, error) {
	user, err := s.repo.UserReader.GetById(id)
	if err != nil {
		return model.User{}, serviceError(err, codes.Internal)
	}

	return user, err
}

func (s *UserReaderService) GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error) {
	users, err := s.repo.UserReader.GetList(pagination, filters)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return users, nil
}
