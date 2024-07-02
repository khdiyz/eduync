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
			CourseWriter: NewCourseWriterRepo(db, logger),
			CourseReader: NewCourseReaderRepo(db, logger),
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
}

type CourseReader interface {
}

type CourseWriter interface {
	Create(input model.CourseCreateRequest) (int64, error)
}
