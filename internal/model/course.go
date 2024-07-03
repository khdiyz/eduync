package model

import "time"

type Course struct {
	Id          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Photo       string     `json:"photo" db:"photo"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type CourseCreateRequest struct {
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
}

type CourseUpdateRequest struct {
	Id          int64  `json:"-" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
}

// Exam type models
type CourseExamType struct {
	Id        int64      `json:"id" db:"id"`
	CourseId  int64      `json:"-" db:"course_id"`
	Name      string     `json:"name" db:"name"`
	MaxBall   float32    `json:"max_ball" db:"max_ball"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type ExamTypeCreateRequest struct {
	CourseId int64   `json:"-" db:"course_id"`
	Name     string  `json:"name" db:"name" validate:"required"`
	MaxBall  float32 `json:"max_ball" db:"max_ball" max:"1000000"`
}

type ExamTypeUpdateRequest struct {
	Id       int64   `json:"-" db:"id"`
	CourseId int64   `json:"course_id" db:"course_id"`
	Name     string  `json:"name" db:"name"`
	MaxBall  float32 `json:"max_ball" db:"max_ball"`
}
