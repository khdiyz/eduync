package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"errors"

	"google.golang.org/grpc/codes"
)

type ExamTypeWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewExamTypeWriterService(repo repository.Repository, logger logger.Logger) *ExamTypeWriterService {
	return &ExamTypeWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ExamTypeWriterService) Create(input model.ExamTypeCreateRequest) (int64, error) {
	_, err := s.repo.CourseRepo.CourseReader.GetById(input.CourseId)
	if err != nil {
		return 0, serviceError(errors.New("course with this id does not exist"), codes.InvalidArgument)
	}

	id, err := s.repo.CourseRepo.ExamTypeWriter.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *ExamTypeWriterService) Update(input model.ExamTypeUpdateRequest) error {
	err := s.repo.CourseRepo.ExamTypeWriter.Update(input)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *ExamTypeWriterService) Delete(courseId int64, id int64) error {
	err := s.repo.CourseRepo.ExamTypeWriter.Delete(courseId, id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
