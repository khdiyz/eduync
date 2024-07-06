package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type EnrollmentReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewEnrollmentReaderRepo(db *sqlx.DB, logger logger.Logger) *EnrollmentReaderRepo {
	return &EnrollmentReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *EnrollmentReaderRepo) GetEnrollmentByStudentIdAndGroupId(studentId, groupId int64) (model.Enrollment, error) {
	var enrollment model.Enrollment

	query := `
	SELECT
		id,
		student_id,
		group_id,
		join_date,
		left_date,
		status,
		created_at,
		updated_at
	FROM enrollments 
	WHERE
		student_id = $1
		AND group_id = $2
		AND deleted_at IS NULL;`

	err := r.db.Get(&enrollment, query, studentId, groupId)
	if err != nil {
		r.logger.Error(err)
		return model.Enrollment{}, err
	}

	return enrollment, nil
}
