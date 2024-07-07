package model

import (
	"edusync/internal/constants"
	"time"
)

type Student struct {
	Id          int64  `json:"id" db:"id"`
	FullName    string `json:"full_name" db:"full_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	ParentPhone string `json:"parent_phone" db:"parent_phone"`
	Address     string `json:"address" db:"address"`
	BirthYear   string `json:"birth_year" db:"birth_year"`

	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type StudentCreateRequest struct {
	FullName    string `json:"full_name" db:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required,uzbphone"`
	ParentPhone string `json:"parent_phone" db:"parent_phone"`
	Address     string `json:"address" db:"address"`
	BirthYear   string `json:"birth_year" db:"birth_year"`
}

type StudentUpdateRequest struct {
	Id          int64  `json:"-" db:"id"`
	FullName    string `json:"full_name" db:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required"`
	ParentPhone string `json:"parent_phone" db:"parent_phone"`
	Address     string `json:"address" db:"address"`
	BirthYear   string `json:"birth_year" db:"birth_year"`
}

// Student Actions
type StudentAction struct {
	Id         int64      `json:"id" db:"id"`
	StudentId  int64      `json:"-" db:"student_id"`
	GroupId    int64      `json:"group_id" db:"group_id"`
	GroupName  string     `json:"group_name" db:"group_name"`
	ActionName string     `json:"action_name" db:"action_name"`
	ActionDate time.Time  `json:"action_date" db:"action_date"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

type StudentActionCreateRequest struct {
	StudentId  int64
	GroupId    int64
	ActionName constants.StudentAction
	ActionDate time.Time
}
