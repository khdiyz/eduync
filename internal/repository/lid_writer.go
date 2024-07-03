package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type LidWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewLidWriterRepo(db *sqlx.DB, logger logger.Logger) *LidWriterRepo {
	return &LidWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *LidWriterRepo) Create(input model.LidCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO lids (
		full_name,
		phone_number,
		course_id,
		description
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	err := r.db.Get(&id, query, input.FullName, input.PhoneNumber, input.CourseId, input.Description)
	if err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *LidWriterRepo) Update(input model.LidUpdateRequest) error {
	query := `
	UPDATE lids
	SET
		full_name = :full_name,
		phone_number = :phone_number,
		course_id = :course_id,
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

func (r *LidWriterRepo) Delete(id int64) error {
	query := `
	UPDATE lids
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
