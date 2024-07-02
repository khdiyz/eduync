package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/helper"
	"edusync/pkg/logger"
	"edusync/pkg/response"
	"errors"

	"google.golang.org/grpc/codes"
)

type UserWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewUserWriterService(repo repository.Repository, logger logger.Logger) *UserWriterService {
	return &UserWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserWriterService) Create(input model.UserCreateRequest) (int64, error) {
	_, err := s.repo.RoleReader.GetById(input.RoleId)
	if err != nil {
		return 0, response.ServiceError(errors.New("role with this id does not exist"), codes.InvalidArgument)
	}

	hash, err := helper.GenerateHash(input.Password)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}
	input.Password = hash

	id, err := s.repo.UserWriter.Create(input)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}
