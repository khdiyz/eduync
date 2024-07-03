package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type LidReaderRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewLidReaderRepo(db *sqlx.DB, logger logger.Logger) *LidReaderRepo {
	return &LidReaderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *LidReaderRepo) GetList(pagination *model.Pagination) ([]model.Lid, error) {
	var (
		lids []model.Lid
		err  error
	)

	countQuery := "SELECT count(id) FROM lids WHERE deleted_at IS NULL;"
	err = getListCount(r.db, &r.logger, pagination, countQuery, nil)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		full_name,
		COALESCE(phone_number, '') AS phone_number,
		course_id,
		COALESCE(description, '') AS description,
		created_at,
		updated_at
	FROM lids
	WHERE
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	if err = r.db.Select(&lids, query, pagination.Limit, pagination.Offset); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return lids, nil
}

func (r *LidReaderRepo) GetById(id int64) (model.Lid, error) {
	var lid model.Lid

	query := `
	SELECT
			id,
			full_name,
			COALESCE(phone_number, '') AS phone_number,
			course_id,
			COALESCE(description, '') AS description,
			created_at,
			updated_at
		FROM lids
		WHERE
			id = $1
			AND deleted_at IS NULL;`

	if err := r.db.Get(&lid, query, id); err != nil {
		r.logger.Error(err)
		return model.Lid{}, err
	}

	return lid, nil
}
