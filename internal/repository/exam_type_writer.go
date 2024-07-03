package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type ExamTypeWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewExamTypeWriterRepo(db *sqlx.DB, logger logger.Logger) *ExamTypeWriterRepo {
	return &ExamTypeWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *ExamTypeWriterRepo) Create(input model.ExamTypeCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO course_exam_types(
		course_id,
		name,
		max_ball
	) VALUES (
		$1, $2, $3
	) RETURNING id;`

	err := r.db.Get(&id, query, input.CourseId, input.Name, input.MaxBall)
	if err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *ExamTypeWriterRepo) Update(input model.ExamTypeUpdateRequest) error {
	query := `
	UPDATE course_exam_types
	SET
		course_id = :course_id,
		name = :name,
		max_ball = :max_ball,
		updated_at = now()
	WHERE 
		id = :id
		AND deleted_at IS NULL;`

	row, err := r.db.NamedExec(query, input)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrorNoRowsAffected
	}

	return nil
}

func (r *ExamTypeWriterRepo) Delete(id int64) error {
	query := `
	UPDATE course_exam_types
	SET
		deleted_at = now()
	WHERE 
		id = $1
		AND deleted_at IS NULL;`

	row, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrorNoRowsAffected
	}

	return nil
}
