package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type StudentReaderService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewStudentReaderService(repo repository.Repository, logger logger.Logger) *StudentReaderService {
	return &StudentReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *StudentReaderService) GetList(pagination *model.Pagination) ([]model.Student, error) {
	students, err := s.repo.StudentRepo.StudentReader.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return students, nil
}

func (s *StudentReaderService) GetById(id int64) (model.Student, error) {
	student, err := s.repo.StudentRepo.StudentReader.GetById(id)
	if err != nil {
		return model.Student{}, serviceError(err, codes.Internal)
	}

	return student, nil
}
