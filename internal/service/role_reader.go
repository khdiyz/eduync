package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"edusync/pkg/response"

	"google.golang.org/grpc/codes"
)

type RoleReaderService struct {
	repo   repository.RoleRepo
	logger logger.Logger
}

func NewRoleReaderService(repo repository.RoleRepo, logger logger.Logger) *RoleReaderService {
	return &RoleReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *RoleReaderService) GetList(pagination *model.Pagination) ([]model.Role, error) {
	roles, err := s.repo.RoleReader.GetList(pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	return roles, nil
}
