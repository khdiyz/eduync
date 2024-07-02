package model

import "time"

type Course struct {
	Id          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	Photo       *string    `json:"photo" db:"photo"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type CourseCreateRequest struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
}
