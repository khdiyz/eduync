package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"edusync/pkg/response"

	"google.golang.org/grpc/codes"
)

type CourseWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewCourseWriterService(repo repository.Repository, logger logger.Logger) *CourseWriterService {
	return &CourseWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *CourseWriterService) Create(input model.CourseCreateRequest) (int64, error) {
	id, err := s.repo.CourseRepo.CourseWriter.Create(input)
	if err != nil {
		return 0, response.ServiceError(err, codes.Internal)
	}

	return id, nil
}

func (s *CourseWriterService) Update(input model.CourseUpdateRequest) error {
	err := s.repo.CourseRepo.CourseWriter.Update(input)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}

func (s *CourseWriterService) Delete(id int64) error {
	err := s.repo.CourseRepo.CourseWriter.Delete(id)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}
