package model

import "time"

type User struct {
	Id          int64  `json:"id" db:"id"`
	FullName    string `json:"full_name" db:"full_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	BirthDate   string `json:"birth_date" db:"birth_date"`

	RoleId   int64  `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`

	Username  string     `json:"username" db:"username"`
	Password  string     `json:"-" db:"password"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UserCreateRequest struct {
	FullName    string    `json:"full_name" db:"full_name" validate:"required"`
	PhoneNumber string    `json:"phone_number" db:"phone_number" validate:"uzbphone,required"`
	BirthDate   time.Time `json:"birth_date" db:"birth_date"`
	RoleId      int64     `json:"role_id" db:"role_id"`
	Username    string    `json:"username" db:"username" validate:"required"`
	Password    string    `json:"password" db:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" default:"superadmin"`
	Password string `json:"password" default:"P@$$w0rd2o24"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserFilter struct {
	RoleId int64
}
