package repository

import (
	"edusync/internal/model"
	"edusync/pkg/logger"
	"strconv"
	"strings"

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
			(SELECT name FROM roles WHERE id = role_id) AS role_name,
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
		&user.RoleName,
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
			(SELECT name FROM roles WHERE id = role_id) AS role_name,
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

func (r *UserReaderRepo) GetList(pagination *model.Pagination, filters map[string]interface{}) ([]model.User, error) {
	var (
		users []model.User
		err   error
	)

	countQuery := "SELECT count(id) FROM users WHERE deleted_at IS NULL"

	query := `
	SELECT
		id,
		full_name,
		phone_number,
		birth_date,
		role_id,
		(SELECT name FROM roles WHERE id = role_id) AS role_name,
		username,
		password,
		created_at,
		updated_at
	FROM users
	WHERE
		deleted_at IS NULL`

	var filterClauses []string
	var args []interface{}
	var counter int

	if roleId, ok := filters["role-id"]; ok {
		counter++
		filterClauses = append(filterClauses, "role_id = $"+strconv.Itoa(counter))
		args = append(args, roleId)
	}

	if len(filterClauses) > 0 {
		countQuery += " AND " + strings.Join(filterClauses, " AND ")
		query += " AND " + strings.Join(filterClauses, " AND ")
	}

	query += " LIMIT $" + strconv.Itoa(counter+1) + " OFFSET $" + strconv.Itoa(counter+2)

	args = append(args, pagination.Limit, pagination.Offset)

	err = getListCount(r.db, &r.logger, pagination, countQuery, args[:len(args)-2])
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	if err = r.db.Select(&users, query, args...); err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return users, nil
}
