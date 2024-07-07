package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type StudentWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewStudentWriterRepo(db *sqlx.DB, logger logger.Logger) *StudentWriterRepo {
	return &StudentWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *StudentWriterRepo) Create(input model.StudentCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO students (
		full_name,
		phone_number,
		parent_phone,
		address,
		birth_year
	) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	if err := r.db.Get(&id, query,
		input.FullName,
		input.PhoneNumber,
		input.ParentPhone,
		input.Address,
		input.BirthYear,
	); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *StudentWriterRepo) Update(input model.StudentUpdateRequest) error {
	query := `
	UPDATE students
	SET
		full_name = :full_name,
		phone_number = :phone_number,
		parent_phone = :parent_phone,
		address = :address,
		birth_year = :birth_year,
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

func (r *StudentWriterRepo) Delete(id int64) error {
	query := `
	UPDATE students
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

func (r *StudentWriterRepo) CreateAction(input model.StudentActionCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO student_actions (
		student_id,
		group_id,
		action_name,
		action_date
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	if err := r.db.Get(&id, query,
		input.StudentId,
		input.GroupId,
		input.ActionName,
		input.ActionDate,
	); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}
