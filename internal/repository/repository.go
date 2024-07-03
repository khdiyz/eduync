package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepo
	RoleRepo
	CourseRepo
	LidRepo
}

func NewRepository(db *sqlx.DB, logger logger.Logger) *Repository {
	return &Repository{
		UserRepo: UserRepo{
			UserReader: NewUserReaderRepo(db, logger),
			UserWriter: NewUserWriterRepo(db, logger),
		},
		RoleRepo: RoleRepo{
			RoleReader: NewRoleReaderRepo(db, logger),
		},
		CourseRepo: CourseRepo{
			CourseWriter:   NewCourseWriterRepo(db, logger),
			CourseReader:   NewCourseReaderRepo(db, logger),
			ExamTypeWriter: NewExamTypeWriterRepo(db, logger),
			ExamTypeReader: NewExamTypeReaderRepo(db, logger),
		},
		LidRepo: LidRepo{
			LidReader: NewLidReaderRepo(db, logger),
			LidWriter: NewLidWriterRepo(db, logger),
		},
	}
}

// User Repo
type UserRepo struct {
	UserReader
	UserWriter
}

type UserReader interface {
	GetByUsername(username string) (model.User, error)
	GetById(id int64) (model.User, error)
}

type UserWriter interface {
	Create(input model.UserCreateRequest) (int64, error)
}

// Role Repo
type RoleRepo struct {
	RoleReader
}

type RoleReader interface {
	GetList(pagination *model.Pagination) ([]model.Role, error)
	GetById(id int64) (model.Role, error)
}

// Course Repo
type CourseRepo struct {
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

// Lids Repo
type LidRepo struct {
	LidReader
	LidWriter
}

type LidReader interface {
	GetList(pagination *model.Pagination) ([]model.Lid, error)
	GetById(id int64) (model.Lid, error)
}

type LidWriter interface {
	Create(input model.LidCreateRequest) (int64, error)
	Update(input model.LidUpdateRequest) error
	Delete(id int64) error
}
