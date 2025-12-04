package dtos

import "time"

// Response DTOs
type UserSummary struct {
	ID        int64      `json:"id"`
	FullName  string     `json:"fullName"`
	UserName  string     `json:"userName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	JobTitle  string     `json:"jobTitle"`
	Status    *string    `json:"status"`
	LastLogin *time.Time `json:"lastLogin"`
}

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Avatar    *string    `json:"avatar"`
	JobTitle  string     `json:"jobTitle"`
	Status    *string    `json:"status"`
	LastLogin *time.Time `json:"lastLogin"`
	CreatedAt time.Time  `json:"createdAt"`
}

// Request DTOs
type CreateUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"fullName" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	JobTitle string `json:"jobTitle" validate:"required"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email" validate:"omitempty,email"`
	FullName *string `json:"fullName"`
	Phone    *string `json:"phone"`
	Avatar   *string `json:"avatar"`
	JobTitle *string `json:"jobTitle"`
	Status   *string `json:"status"`
}
