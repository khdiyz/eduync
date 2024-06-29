package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepo
}

func NewRepository(db *sqlx.DB, logger logger.Logger) *Repository {
	return &Repository{
		UserRepo: UserRepo{
			UserReader: NewUserReaderRepo(db, logger),
			UserWriter: NewUserWriterRepo(db, logger),
		},
	}
}

// USER REPO
type UserRepo struct {
	UserReader
	UserWriter
}

type UserReader interface {
	GetByUsernameAndPassword(input model.UserLoginReq) (model.User, error)
}

type UserWriter interface {
	Create(input model.UserCreateReq) (int64, error)
}
