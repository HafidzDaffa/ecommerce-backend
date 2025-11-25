package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	Email           string     `db:"email" json:"email"`
	PasswordHash    string     `db:"password_hash" json:"-"`
	FullName        *string    `db:"full_name" json:"full_name,omitempty"`
	PhoneNumber     *string    `db:"phone_number" json:"phone_number,omitempty"`
	AvatarURL       *string    `db:"avatar_url" json:"avatar_url,omitempty"`
	RoleID          int        `db:"role_id" json:"role_id"`
	IsEmailVerified bool       `db:"is_email_verified" json:"is_email_verified"`
	IsActive        bool       `db:"is_active" json:"is_active"`
	LastLoginAt     *time.Time `db:"last_login_at" json:"last_login_at,omitempty"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type RegisterRequest struct {
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=6"`
	FullName    *string `json:"full_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	RoleID      int     `json:"role_id" validate:"required,min=1,max=3"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
