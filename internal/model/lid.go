package model

import "time"

type Lid struct {
	Id          int64      `json:"id" db:"id"`
	FullName    string     `json:"full_name" db:"full_name"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	CourseId    int64      `json:"course_id" db:"course_id"`
	Description string     `json:"description" db:"description"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type LidCreateRequest struct {
	FullName    string `json:"full_name" db:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"uzbphone"`
	CourseId    int64  `json:"course_id" db:"course_id"`
	Description string `json:"description" db:"description"`
}

type LidUpdateRequest struct {
	Id          int64  `json:"-" db:"id"`
	FullName    string `json:"full_name" db:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"uzbphone"`
	CourseId    int64  `json:"course_id" db:"course_id"`
	Description string `json:"description" db:"description"`
}
