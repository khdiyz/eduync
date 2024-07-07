package repository

import (
	"edusync/internal/constants"
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

func (r *GroupReaderRepo) GetGroupStudents(request model.GetGroupStudentsRequest) ([]model.Student, error) {
	var (
		students     []model.Student
		err          error
		statusActive = string(constants.EnrollmentStatusActive)
		statusFrozen = string(constants.EnrollmentStatusFrozen)
		statusLeft   = "LEFT"
	)

	countQuery := `
	SELECT 
		count(s.id) 
	FROM students s 
	JOIN 
		enrollments e ON s.id = e.student_id 
	WHERE 
		e.group_id = $1
		AND s.deleted_at IS NULL
		AND e.deleted_at IS NULL `

	query := `
	SELECT
		s.id,
		s.full_name,
		s.phone_number,
		COALESCE(s.parent_phone, '') AS parent_phone,
		COALESCE(s.address, '') AS address,
		COALESCE(s.birth_year, '') AS birth_year,
		s.created_at,
		s.updated_at
	FROM students s
	JOIN enrollments e ON s.id = e.student_id
	WHERE
		e.group_id = $1
		AND s.deleted_at IS NULL
		AND e.deleted_at IS NULL `

	switch request.StudentType {
	case statusActive:
		countQuery = countQuery + " AND e.status = '" + statusActive + "' "
		query = query + " AND e.status = '" + statusActive + "' "
	case statusFrozen:
		countQuery = countQuery + " AND e.status = '" + statusFrozen + "' "
		query = query + " AND e.status = '" + statusFrozen + "' "
	case statusLeft:
		countQuery = countQuery + " AND e.left_date IS NOT NULL AND e.status = '" + string(constants.EnrollmentStatusInactive) + "' "
		query = query + " AND e.left_date IS NOT NULL AND e.status = '" + string(constants.EnrollmentStatusInactive) + "' "
	}

	query += " LIMIT $2 OFFSET $3 "

	err = getListCount(r.db, &r.logger, request.Pagination, countQuery, []interface{}{request.GroupId})
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	err = r.db.Select(&students, query, request.GroupId, request.Pagination.Limit, request.Pagination.Offset)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return students, nil
}
