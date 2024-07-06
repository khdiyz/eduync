package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type EnrollmentWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewEnrollmentWriterRepo(db *sqlx.DB, logger logger.Logger) *EnrollmentWriterRepo {
	return &EnrollmentWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *EnrollmentWriterRepo) CreateEnrollment(input model.EnrollmentCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO enrollments (
		student_id,
		group_id,
		join_date,
		status
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	if err := r.db.Get(&id, query,
		input.StudentId,
		input.GroupId,
		input.JoinDate,
		input.Status,
	); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *EnrollmentWriterRepo) UpdateEnrollment(input model.EnrollmentUpdateRequest) error {
	query := `
	UPDATE enrollments
	SET
		student_id = :student_id,
		group_id = :group_id,
		join_date = :join_date,
		left_date = :left_date,
		status = :status,
		updated_at = NOW()
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
