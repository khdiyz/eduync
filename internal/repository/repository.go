package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepo
	RoleRepo
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
	}
}

// USER REPO
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

// ROLE REPO
type RoleRepo struct {
	RoleReader
}

type RoleReader interface {
	GetList(pagination *model.Pagination) ([]model.Role, error)
	GetById(id int64) (model.Role, error)
}
