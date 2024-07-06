package model

import "time"

type Enrollment struct {
	Id        int64      `db:"id"`
	StudentId int64      `db:"student_id"`
	GroupId   int64      `db:"group_id"`
	JoinDate  time.Time  `db:"join_date"`
	LeftDate  *time.Time `db:"left_date"`
	Status    string     `db:"status"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type EnrollmentCreateRequest struct {
	StudentId int64     `db:"student_id"`
	GroupId   int64     `db:"group_id"`
	JoinDate  time.Time `db:"join_date"`
	Status    string    `db:"status"`
}

type EnrollmentUpdateRequest struct {
	Id        int64      `db:"id"`
	StudentId int64      `db:"student_id"`
	GroupId   int64      `db:"group_id"`
	JoinDate  time.Time  `db:"join_date"`
	LeftDate  *time.Time `db:"left_date"`
	Status    string     `db:"status"`
}
