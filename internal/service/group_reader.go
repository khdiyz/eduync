package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"errors"

	"google.golang.org/grpc/codes"
)

type GroupReaderService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewGroupReaderService(repo repository.Repository, logger logger.Logger) *GroupReaderService {
	return &GroupReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *GroupReaderService) GetList(pagination *model.Pagination) ([]model.Group, error) {
	groups, err := s.repo.GroupRepo.GroupReader.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return groups, nil
}

func (s *GroupReaderService) GetById(id int64) (model.Group, error) {
	group, err := s.repo.GroupRepo.GroupReader.GetById(id)
	if err != nil {
		return model.Group{}, serviceError(err, codes.Internal)
	}

	return group, nil
}

func (s *GroupReaderService) GetGroupStudents(request model.GetGroupStudentsRequest) ([]model.Student, error) {
	if request.StudentType != "ACTIVE" && request.StudentType != "FROZEN" && request.StudentType != "LEFT" {
		return nil, serviceError(errors.New("invalid student type"), codes.InvalidArgument)
	}

	students, err := s.repo.GroupReader.GetGroupStudents(request)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return students, nil
}
