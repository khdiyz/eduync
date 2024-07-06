package model

import "time"

type Group struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	CourseId   int64  `json:"course_id" db:"course_id"`
	CourseName string `json:"course_name" db:"course_name"`

	TeacherId   int64  `json:"mentor_id" db:"teacher_id"`
	TeacherName string `json:"mentor_name" db:"teacher_name"`

	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type GroupCreateRequest struct {
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description"`
	CourseId    int64  `json:"course_id" db:"course_id"`
	TeacherId   int64  `json:"mentor_id" db:"teacher_id"`
}

type GroupUpdateRequest struct {
	Id          int64  `json:"-" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description"`
	CourseId    int64  `json:"course_id" db:"course_id"`
	TeacherId   int64  `json:"mentor_id" db:"teacher_id"`
}

type JoinStudentRequest struct {
	GroupId   int64     `json:"-"`
	StudentId int64     `json:"-"`
	JoinDate  time.Time `json:"join_date"`
}

type LeftStudentRequest struct {
	GroupId   int64     `json:"-"`
	StudentId int64     `json:"-"`
	LeftDate  time.Time `json:"left_date"`
}
