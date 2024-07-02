package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type UserReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewUserReaderRepo(db *sqlx.DB, logger logger.Logger) *UserReaderRepo {
	return &UserReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *UserReaderRepo) GetByUsername(username string) (model.User, error) {
	var user model.User

	query := `
	SELECT
			id,
			full_name,
			phone_number,
			birth_date,
			role_id,
			username,
			password,
			created_at,
			updated_at
		FROM users
		WHERE
			username = $1
			AND deleted_at IS NULL;`

	err := r.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.FullName,
		&user.PhoneNumber,
		&user.BirthDate,
		&user.RoleId,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		r.logger.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UserReaderRepo) GetById(id int64) (model.User, error) {
	var user model.User

	query := `
	SELECT
			id,
			full_name,
			phone_number,
			birth_date,
			role_id,
			username,
			password,
			created_at,
			updated_at
		FROM users
		WHERE
			id = $1
			AND deleted_at IS NULL;`

	if err := r.db.Get(&user, query, id); err != nil {
		r.logger.Error(err)
		return model.User{}, err
	}

	return user, nil
}
