package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/helper"
	"edusync/pkg/logger"

	"google.golang.org/grpc/codes"
)

type StudentWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewStudentWriterService(repo repository.Repository, logger logger.Logger) *StudentWriterService {
	return &StudentWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *StudentWriterService) Create(input model.StudentCreateRequest) (int64, error) {
	if input.BirthYear != "" {
		isValid, err := helper.IsValidBirthYear(input.BirthYear)
		if !isValid && err != nil {
			return 0, serviceError(err, codes.InvalidArgument)
		}
	}

	if input.ParentPhone != "" {
		isValid, err := helper.IsValidPhoneNumber(input.ParentPhone)
		if !isValid && err != nil {
			return 0, serviceError(err, codes.InvalidArgument)
		}
	}

	id, err := s.repo.StudentRepo.StudentWriter.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *StudentWriterService) Update(input model.StudentUpdateRequest) error {
	if input.BirthYear != "" {
		isValid, err := helper.IsValidBirthYear(input.BirthYear)
		if !isValid && err != nil {
			return serviceError(err, codes.InvalidArgument)
		}
	}

	if input.ParentPhone != "" {
		isValid, err := helper.IsValidPhoneNumber(input.ParentPhone)
		if !isValid && err != nil {
			return serviceError(err, codes.InvalidArgument)
		}
	}

	err := s.repo.StudentRepo.StudentWriter.Update(input)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *StudentWriterService) Delete(id int64) error {
	err := s.repo.StudentRepo.StudentWriter.Delete(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
