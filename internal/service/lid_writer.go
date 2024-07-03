package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type LidWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewLidWriterService(repo repository.Repository, logger logger.Logger) *LidWriterService {
	return &LidWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *LidWriterService) Create(input model.LidCreateRequest) (int64, error) {
	_, err := s.repo.CourseRepo.CourseReader.GetById(input.CourseId)
	if err != nil {
		return 0, serviceError(errCourseDoesNotExists, codes.InvalidArgument)
	}

	id, err := s.repo.LidRepo.LidWriter.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *LidWriterService) Update(input model.LidUpdateRequest) error {
	_, err := s.repo.CourseRepo.CourseReader.GetById(input.CourseId)
	if err != nil {
		return serviceError(errCourseDoesNotExists, codes.InvalidArgument)
	}

	err = s.repo.LidRepo.LidWriter.Update(input)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *LidWriterService) Delete(id int64) error {
	err := s.repo.LidRepo.LidWriter.Delete(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
