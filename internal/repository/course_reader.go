package repository

import (
	"edusync/internal/model"
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

func (r *CourseReaderRepo) GetList(pagination *model.Pagination) ([]model.Course, error) {
	var (
		courses []model.Course
		err     error
	)

	countQuery := "SELECT count(id) FROM courses WHERE deleted_at IS NULL;"
	err = getListCount(r.db, &r.logger, pagination, countQuery, nil)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		COALESCE(photo, '') AS photo,
		created_at,
		updated_at
	FROM courses
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&courses, query, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return courses, nil
}

func (r *CourseReaderRepo) GetById(id int64) (model.Course, error) {
	var course model.Course

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		COALESCE(photo, '') AS photo,
		created_at,
		updated_at
	FROM courses
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&course, query, id); err != nil {
		return model.Course{}, err
	}

	return course, nil
}
