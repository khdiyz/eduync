package model

import "time"

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
