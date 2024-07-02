package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type CourseWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewCourseWriterRepo(db *sqlx.DB, logger logger.Logger) *CourseWriterRepo {
	return &CourseWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *CourseWriterRepo) Create(input model.CourseCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO course (
		name,
		description,
		photo
	) VALUES ($1, $2, $3) RETURNING id;`

	if err := r.db.Get(&id, query,
		input.Name,
		input.Description,
		input.Photo,
	); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}
