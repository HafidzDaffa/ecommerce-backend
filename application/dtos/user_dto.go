package dtos

type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	FullName    string `json:"full_name" validate:"required"`
	Phone       string `json:"phone"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	AvatarURL   string `json:"avatar_url"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type UserResponse struct {
	ID          int64   `json:"id"`
	Email       string  `json:"email"`
	FullName    string  `json:"full_name"`
	Phone       string  `json:"phone"`
	AvatarURL   string  `json:"avatar_url"`
	Gender      string  `json:"gender"`
	DateOfBirth *string `json:"date_of_birth"`
	IsActive    bool    `json:"is_active"`
	IsVerified  bool    `json:"is_verified"`
	RoleID      int     `json:"role_id"`
	LastLoginAt *string `json:"last_login_at"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
