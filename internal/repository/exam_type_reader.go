package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type ExamTypeReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewExamTypeReaderRepo(db *sqlx.DB, logger logger.Logger) *ExamTypeReaderRepo {
	return &ExamTypeReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *ExamTypeReaderRepo) GetList(courseId int64, pagination *model.Pagination) ([]model.CourseExamType, error) {
	var (
		examTypes []model.CourseExamType
		err       error
	)

	countQuery := "SELECT count(id) FROM course_exam_types WHERE course_id = $1 AND deleted_at IS NULL;"
	err = getListCount(r.db, &r.logger, pagination, countQuery, []interface{}{courseId})
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		course_id,
		name,
		max_ball,
		created_at,
		updated_at
	FROM course_exam_types
	WHERE
		course_id = $1
		AND	deleted_at IS NULL
	LIMIT $2 OFFSET $3;`

	if err = r.db.Select(&examTypes, query, courseId, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return examTypes, nil
}

func (r *ExamTypeReaderRepo) GetById(id int64) (model.CourseExamType, error) {
	var examType model.CourseExamType

	query := `
	SELECT
		id,
		course_id,
		name,
		max_ball,
		created_at,
		updated_at
	FROM course_exam_types
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&examType, query, id); err != nil {
		return model.CourseExamType{}, err
	}

	return examType, nil
}
