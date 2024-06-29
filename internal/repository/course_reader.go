package repository

import (
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type CourseReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewCourseReaderRepo(db *sqlx.DB, logger logger.Logger) *CourseReaderRepo {
	return &CourseReaderRepo{
		db:     db,
		logger: logger,
	}
}
