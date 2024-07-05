package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type GroupWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewGroupWriterRepo(db *sqlx.DB, logger logger.Logger) *GroupWriterRepo {
	return &GroupWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *GroupWriterRepo) Create(input model.GroupCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO groups (
		name,
		course_id,
		teacher_id,
		description
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	if err := r.db.Get(&id, query,
		input.Name,
		input.CourseId,
		input.TeacherId,
		input.Description,
	); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *GroupWriterRepo) Update(input model.GroupUpdateRequest) error {
	query := `
	UPDATE groups
	SET
		name = :name,
		course_id = :course_id,
		teacher_id = :teacher_id,
		description = :description,
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

func (r *GroupWriterRepo) Delete(id int64) error {
	query := `
	UPDATE groups
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
