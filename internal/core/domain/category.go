package domain

import "time"

type Category struct {
	ID           int        `db:"id" json:"id"`
	CategoryName string     `db:"category_name" json:"category_name"`
	Slug         *string    `db:"slug" json:"slug,omitempty"`
	Icon         *string    `db:"icon" json:"icon,omitempty"`
	ImagePath    *string    `db:"image_path" json:"image_path,omitempty"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type CreateCategoryRequest struct {
	CategoryName string  `json:"category_name" validate:"required,min=2,max=255"`
	Slug         *string `json:"slug,omitempty" validate:"omitempty,max=255"`
	Icon         *string `json:"icon,omitempty" validate:"omitempty,max=255"`
	ImagePath    *string `json:"image_path,omitempty" validate:"omitempty,max=500"`
	IsActive     *bool   `json:"is_active,omitempty"`
}

type UpdateCategoryRequest struct {
	CategoryName *string `json:"category_name,omitempty" validate:"omitempty,min=2,max=255"`
	Slug         *string `json:"slug,omitempty" validate:"omitempty,max=255"`
	Icon         *string `json:"icon,omitempty" validate:"omitempty,max=255"`
	ImagePath    *string `json:"image_path,omitempty" validate:"omitempty,max=500"`
	IsActive     *bool   `json:"is_active,omitempty"`
}
