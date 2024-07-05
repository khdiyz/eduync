package service

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"errors"

	"google.golang.org/grpc/codes"
)

var (
	errTeacherDoesNotExists = errors.New("mentor with this id does not exist")
	errUserIsNotMentor      = errors.New("user is not mentor")
)

type GroupWriterService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewGroupWriterService(repo repository.Repository, logger logger.Logger) *GroupWriterService {
	return &GroupWriterService{
		repo:   repo,
		logger: logger,
	}
}

func (s *GroupWriterService) Create(input model.GroupCreateRequest) (int64, error) {
	_, err := s.repo.CourseRepo.CourseReader.GetById(input.CourseId)
	if err != nil {
		return 0, serviceError(errCourseDoesNotExists, codes.InvalidArgument)
	}

	user, err := s.repo.UserRepo.UserReader.GetById(input.TeacherId)
	if err != nil {
		return 0, serviceError(errTeacherDoesNotExists, codes.InvalidArgument)
	}

	if user.RoleName != string(constants.RoleMentor) {
		return 0, serviceError(errUserIsNotMentor, codes.InvalidArgument)
	}

	id, err := s.repo.GroupRepo.GroupWriter.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *GroupWriterService) Update(input model.GroupUpdateRequest) error {
	_, err := s.repo.CourseRepo.CourseReader.GetById(input.CourseId)
	if err != nil {
		return serviceError(errCourseDoesNotExists, codes.InvalidArgument)
	}

	user, err := s.repo.UserRepo.UserReader.GetById(input.TeacherId)
	if err != nil {
		return serviceError(errTeacherDoesNotExists, codes.InvalidArgument)
	}

	if user.RoleName != string(constants.RoleMentor) {
		return serviceError(errUserIsNotMentor, codes.InvalidArgument)
	}

	err = s.repo.GroupRepo.GroupWriter.Update(input)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *GroupWriterService) Delete(id int64) error {
	err := s.repo.GroupRepo.GroupWriter.Delete(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
