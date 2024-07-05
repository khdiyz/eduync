package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type GroupReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewGroupReaderRepo(db *sqlx.DB, logger logger.Logger) *GroupReaderRepo {
	return &GroupReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *GroupReaderRepo) GetList(pagination *model.Pagination) ([]model.Group, error) {
	var (
		groups []model.Group
		err    error
	)

	countQuery := "SELECT count(id) FROM groups WHERE deleted_at IS NULL;"
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
		course_id,
		(SELECT name FROM courses WHERE id = course_id) AS course_name,
		teacher_id,
		(SELECT full_name FROM users WHERE id = teacher_id) AS teacher_name,
		created_at,
		updated_at
	FROM groups
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&groups, query, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return groups, nil
}

func (r *GroupReaderRepo) GetById(id int64) (model.Group, error) {
	var group model.Group

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description,
		course_id,
		(SELECT name FROM courses WHERE id = course_id) AS course_name,
		teacher_id,
		(SELECT full_name FROM users WHERE id = teacher_id) AS teacher_name,
		created_at,
		updated_at
	FROM groups
	WHERE
		deleted_at IS NULL
		AND id = $1;`

	if err := r.db.Get(&group, query, id); err != nil {
		return model.Group{}, err
	}

	return group, nil
}
