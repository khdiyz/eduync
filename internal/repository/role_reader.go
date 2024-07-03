package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type RoleReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewRoleReaderRepo(db *sqlx.DB, logger logger.Logger) *RoleReaderRepo {
	return &RoleReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *RoleReaderRepo) GetList(pagination *model.Pagination) ([]model.Role, error) {
	var (
		roles []model.Role
		err   error
	)

	countQuery := "SELECT count(id) FROM roles WHERE deleted_at IS NULL;"
	err = getListCount(r.db, &r.logger, pagination, countQuery, nil)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		created_at,
		updated_at
	FROM roles
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&roles, query, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return roles, nil
}

func (r *RoleReaderRepo) GetById(id int64) (model.Role, error) {
	var role model.Role

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		created_at,
		updated_at
	FROM roles
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&role, query, id); err != nil {
		return model.Role{}, err
	}

	return role, nil
}
