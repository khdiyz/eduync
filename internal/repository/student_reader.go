package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type StudentReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewStudentReaderRepo(db *sqlx.DB, logger logger.Logger) *StudentReaderRepo {
	return &StudentReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *StudentReaderRepo) GetList(pagination *model.Pagination) ([]model.Student, error) {
	var (
		students []model.Student
		err      error
	)

	countQuery := "SELECT count(id) FROM students WHERE deleted_at IS NULL;"
	err = getListCount(r.db, &r.logger, pagination, countQuery, nil)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		full_name,
		phone_number,
		COALESCE(parent_phone, '') AS parent_phone,
		COALESCE(address, '') AS address,
		COALESCE(birth_year, '') AS birth_year,
		created_at,
		updated_at
	FROM students
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&students, query, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return students, nil
}

func (r *StudentReaderRepo) GetById(id int64) (model.Student, error) {
	var student model.Student

	query := `
	SELECT
		id,
		full_name,
		phone_number,
		COALESCE(parent_phone, '') AS parent_phone,
		COALESCE(address, '') AS address,
		COALESCE(birth_year, '') AS birth_year,
		created_at,
		updated_at
	FROM students
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&student, query, id); err != nil {
		return model.Student{}, err
	}

	return student, nil
}
