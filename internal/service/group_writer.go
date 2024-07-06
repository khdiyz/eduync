package service

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/pkg/logger"
	"errors"
	"time"

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

func (s *GroupWriterService) JoinStudent(input model.JoinStudentRequest) error {
	if input.JoinDate.IsZero() {
		return serviceError(errors.New("join date required"), codes.InvalidArgument)
	} else if input.JoinDate.After(time.Now()) {
		return serviceError(errors.New("join date must be before now"), codes.InvalidArgument)
	}

	enrollment, err := s.repo.EnrollmentReader.GetEnrollmentByStudentIdAndGroupId(input.StudentId, input.GroupId)
	if err == nil && enrollment.Status != string(constants.EnrollmentStatusInactive) {
		return serviceError(errors.New("student already in the group"), codes.InvalidArgument)
	}

	if err == nil && enrollment.LeftDate != nil && enrollment.Status == string(constants.EnrollmentStatusInactive) {
		if err = s.repo.EnrollmentWriter.UpdateEnrollment(model.EnrollmentUpdateRequest{
			Id:        enrollment.Id,
			StudentId: enrollment.StudentId,
			GroupId:   enrollment.GroupId,
			JoinDate:  enrollment.JoinDate,
			LeftDate:  nil,
			Status:    string(constants.EnrollmentStatusActive),
		}); err != nil {
			return serviceError(err, codes.Internal)
		}

		return nil
	}

	if _, err = s.repo.EnrollmentWriter.CreateEnrollment(model.EnrollmentCreateRequest{
		StudentId: input.StudentId,
		GroupId:   input.GroupId,
		JoinDate:  input.JoinDate,
		Status:    string(constants.EnrollmentStatusActive),
	}); err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *GroupWriterService) LeftStudent(input model.LeftStudentRequest) error {
	if input.LeftDate.IsZero() {
		return serviceError(errors.New("left date required"), codes.InvalidArgument)
	} else if input.LeftDate.After(time.Now()) {
		return serviceError(errors.New("left date must be before now"), codes.InvalidArgument)
	}

	enrollment, err := s.repo.EnrollmentReader.GetEnrollmentByStudentIdAndGroupId(input.StudentId, input.GroupId)
	if err != nil {
		return serviceError(err, codes.NotFound)
	}

	if enrollment.LeftDate != nil || enrollment.Status == string(constants.EnrollmentStatusInactive) {
		return serviceError(errors.New("student is not group"), codes.InvalidArgument)
	}

	if err = s.repo.EnrollmentWriter.UpdateEnrollment(model.EnrollmentUpdateRequest{
		Id:        enrollment.Id,
		StudentId: enrollment.StudentId,
		GroupId:   enrollment.GroupId,
		JoinDate:  enrollment.JoinDate,
		LeftDate:  &input.LeftDate,
		Status:    string(constants.EnrollmentStatusInactive),
	}); err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
