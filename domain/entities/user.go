package entities

import "time"

type User struct {
	ID                   int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email                string     `json:"email" gorm:"uniqueIndex;not null;size:255"`
	PasswordHash         string     `json:"-" gorm:"not null;size:255"`
	FullName             string     `json:"full_name" gorm:"not null;size:255"`
	Phone                string     `json:"phone,omitempty" gorm:"size:20"`
	AvatarURL            string     `json:"avatar_url,omitempty" gorm:"size:500"`
	Gender               string     `json:"gender,omitempty" gorm:"size:50"`
	DateOfBirth          *time.Time `json:"date_of_birth,omitempty" gorm:"type:date"`
	IsActive             bool       `json:"is_active" gorm:"default:true"`
	IsVerified           bool       `json:"is_verified" gorm:"default:false"`
	VerificationToken    string     `json:"-" gorm:"size:255"`
	RoleID               int        `json:"role_id" gorm:"default:1"`
	ResetPasswordToken   string     `json:"-" gorm:"size:255"`
	ResetPasswordExpires *time.Time `json:"-" gorm:"type:timestamp"`
	LastLoginAt          *time.Time `json:"last_login_at,omitempty" gorm:"type:timestamp"`
	CreatedAt            time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	Role *Role `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null;size:50"`
	Description string    `json:"description,omitempty" gorm:"size:255"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Users []User `json:"users,omitempty" gorm:"foreignKey:RoleID"`
}
