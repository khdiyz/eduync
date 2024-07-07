package service

import (
	"edusync/internal/model"
	"edusync/internal/repository"
	"edusync/internal/storage"
	"edusync/pkg/logger"
	"io"
	"time"
)

type Service struct {
	Authorization
	Minio

	UserService
	RoleService
	CourseService
	LidService
	GroupService
	StudentService
}

func NewService(repos repository.Repository, storage storage.Storage, logger logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.UserRepo, logger),
		Minio:         NewMinioService(storage, logger),

		UserService: UserService{
			UserReader: NewUserReaderService(repos.UserRepo, logger),
			UserWriter: NewUserWriterService(repos, logger),
		},
		RoleService: RoleService{
			RoleReader: NewRoleReaderService(repos.RoleRepo, logger),
		},
		CourseService: CourseService{
			CourseWriter: NewCourseWriterService(repos, logger),
			CourseReader: NewCourseReaderService(repos, logger),

			ExamTypeWriter: NewExamTypeWriterService(repos, logger),
			ExamTypeReader: NewExamTypeReaderService(repos, logger),
		},
		LidService: LidService{
			LidWriter: NewLidWriterService(repos, logger),
			LidReader: NewLidReaderService(repos, logger),
		},
		GroupService: GroupService{
			GroupReader: NewGroupReaderService(repos, logger),
			GroupWriter: NewGroupWriterService(repos, logger),
		},
		StudentService: StudentService{
			StudentReader: NewStudentReaderService(repos, logger),
			StudentWriter: NewStudentWriterService(repos, logger),
		},
	}
}

type Authorization interface {
	CreateToken(user model.User, tokenType string, expiresAt time.Time) (*model.Token, error)
	GenerateTokens(user model.User) (*model.Token, *model.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(input model.LoginRequest) (*model.Token, *model.Token, error)
}

type Minio interface {
	UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error)
	UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error)
	UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error)
}

// User Service
type UserService struct {
	UserReader
	UserWriter
}

type UserReader interface {
	GetByUsername(username string) (model.User, error)
	GetById(id int64) (model.User, error)
	GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error)
}

type UserWriter interface {
	Create(input model.UserCreateRequest) (int64, error)
}

// Role Service
type RoleService struct {
	RoleReader
}

type RoleReader interface {
	GetList(pagination *model.Pagination) ([]model.Role, error)
}

// Course Service
type CourseService struct {
	CourseWriter
	CourseReader

	ExamTypeWriter
	ExamTypeReader
}

type CourseReader interface {
	GetList(pagination *model.Pagination) ([]model.Course, error)
	GetById(id int64) (model.Course, error)
}

type CourseWriter interface {
	Create(input model.CourseCreateRequest) (int64, error)
	Update(input model.CourseUpdateRequest) error
	Delete(id int64) error
}

type ExamTypeWriter interface {
	Create(input model.ExamTypeCreateRequest) (int64, error)
	Update(input model.ExamTypeUpdateRequest) error
	Delete(courseId int64, id int64) error
}

type ExamTypeReader interface {
	GetList(courseId int64, pagination *model.Pagination) ([]model.CourseExamType, error)
	GetById(courseId int64, examTypeId int64) (model.CourseExamType, error)
}

// Lid Service
type LidService struct {
	LidWriter
	LidReader
}

type LidWriter interface {
	Create(input model.LidCreateRequest) (int64, error)
	Update(input model.LidUpdateRequest) error
	Delete(id int64) error
}

type LidReader interface {
	GetList(pagination *model.Pagination) ([]model.Lid, error)
	GetById(id int64) (model.Lid, error)
}

// Group Service
type GroupService struct {
	GroupReader
	GroupWriter
}

type GroupReader interface {
	GetList(pagination *model.Pagination) ([]model.Group, error)
	GetById(id int64) (model.Group, error)
	GetGroupStudents(request model.GetGroupStudentsRequest) ([]model.Student, error)
}

type GroupWriter interface {
	Create(input model.GroupCreateRequest) (int64, error)
	Update(input model.GroupUpdateRequest) error
	Delete(id int64) error
	// Student Actions
	JoinStudent(input model.JoinStudentRequest) error
	LeftStudent(input model.LeftStudentRequest) error
	FreezeStudent(input model.FreezeStudentRequest) error
	UnfreezeStudent(input model.UnfreezeStudentRequest) error
}

// Student Service
type StudentService struct {
	StudentReader
	StudentWriter
}

type StudentReader interface {
	GetList(pagination *model.Pagination) ([]model.Student, error)
	GetById(id int64) (model.Student, error)
	GetActions(studentId int64) ([]model.StudentAction, error)
}

type StudentWriter interface {
	Create(input model.StudentCreateRequest) (int64, error)
	Update(input model.StudentUpdateRequest) error
	Delete(id int64) error
}
