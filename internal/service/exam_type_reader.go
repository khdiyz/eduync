package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type ExamTypeReaderService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewExamTypeReaderService(repo repository.Repository, logger logger.Logger) *ExamTypeReaderService {
	return &ExamTypeReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ExamTypeReaderService) GetList(courseId int64, pagination *model.Pagination) ([]model.CourseExamType, error) {
	examTypes, err := s.repo.CourseRepo.ExamTypeReader.GetList(courseId, pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return examTypes, nil
}

func (s *ExamTypeReaderService) GetById(courseId int64, examTypeId int64) (model.CourseExamType, error) {
	examType, err := s.repo.CourseRepo.ExamTypeReader.GetById(courseId, examTypeId)
	if err != nil {
		return model.CourseExamType{}, serviceError(err, codes.Internal)
	}

	return examType, nil
}
