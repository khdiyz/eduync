package service

import (
	"edusync/internal/config"
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/internal/storage"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type CourseReaderService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewCourseReaderService(repo repository.Repository, logger logger.Logger) *CourseReaderService {
	return &CourseReaderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *CourseReaderService) GetList(pagination *model.Pagination) ([]model.Course, error) {
	courses, err := s.repo.CourseRepo.CourseReader.GetList(pagination)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	cfg := config.GetConfig()

	for i := range courses {
		if courses[i].Photo != "" {
			courses[i].Photo = storage.GenerateLink(*cfg, courses[i].Photo)
		}
	}

	return courses, nil
}

func (s *CourseReaderService) GetById(id int64) (model.Course, error) {
	course, err := s.repo.CourseRepo.CourseReader.GetById(id)
	if err != nil {
		return model.Course{}, serviceError(err, codes.Internal)
	}

	if course.Photo != "" {
		course.Photo = storage.GenerateLink(*config.GetConfig(), course.Photo)
	}

	return course, nil
}
