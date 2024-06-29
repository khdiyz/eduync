package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type UserWriterRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewUserWriterRepo(db *sqlx.DB, logger logger.Logger) *UserWriterRepo {
	return &UserWriterRepo{
		db:     db,
		logger: logger,
	}
}

func (r *UserWriterRepo) Create(input model.UserCreateReq) (int64, error) {
	var id int64

	query := `
	INSERT INTO users (
		full_name,
		phone_number,
		birth_date,
		role_id,
		username,
		password
	) VALUES (
		$1, $2, $3, $4, $5, $6 
	) RETURNING id;`

	if err := r.db.QueryRow(query,
		input.FullName,
		input.PhoneNumber,
		input.BirthDate,
		input.RoleId,
		input.Username,
		input.Password,
	).Scan(&id); err != nil {
		r.logger.Error(err)
		return 0, err
	}

	return id, nil
}
