package repository

import (
	"edusync/internal/constants"
	"edusync/internal/model"
	"edusync/pkg/logger"
	"fmt"

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
	INSERT INTO courses (
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

func (r *CourseWriterRepo) Update(input model.CourseUpdateRequest) error {
	query := `
	UPDATE courses
	SET
		name = :name,
		description = :description,
		photo = :photo,
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

		fmt.Println("kirdi ")
		r.logger.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrorNoRowsAffected
	}

	return nil
}

func (r *CourseWriterRepo) Delete(id int64) error {
	query := `
	UPDATE courses
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
