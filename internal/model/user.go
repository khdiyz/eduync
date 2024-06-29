package model

import "time"

type User struct {
	Id          int64      `json:"id" db:"id"`
	FullName    string     `json:"full_name" db:"full_name"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	BirthDate   string     `json:"birth_date" db:"birth_date"`
	RoleId      int64      `json:"role_id" db:"role_id"`
	Username    string     `json:"username" db:"username"`
	Password    string     `json:"-" db:"password"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type UserCreateReq struct {
	FullName    string `json:"full_name" db:"full_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	BirthDate   string `json:"birth_date" db:"birth_date"`
	RoleId      int64  `json:"role_id" db:"role_id"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"-" db:"password"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
