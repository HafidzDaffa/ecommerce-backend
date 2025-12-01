package domain

import (
	"time"

	"github.com/google/uuid"
)

// FeeType represents the type of application fee
type FeeType string

const (
	FeeTypePercentage FeeType = "PERCENTAGE"
	FeeTypeFixed      FeeType = "FIXED"
)

// ApplicationFee represents an application fee configuration
type ApplicationFee struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	FeeName     string     `db:"fee_name" json:"fee_name"`
	FeeType     FeeType    `db:"fee_type" json:"fee_type"`
	FeeValue    float64    `db:"fee_value" json:"fee_value"`
	Description *string    `db:"description" json:"description,omitempty"`
	IsActive    bool       `db:"is_active" json:"is_active"`
	CreatedBy   *uuid.UUID `db:"created_by" json:"created_by,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// CreateApplicationFeeRequest represents the request to create an application fee
type CreateApplicationFeeRequest struct {
	FeeName     string  `json:"fee_name" validate:"required,min=3,max=255"`
	FeeType     FeeType `json:"fee_type" validate:"required,oneof=PERCENTAGE FIXED"`
	FeeValue    float64 `json:"fee_value" validate:"required,gt=0"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// UpdateApplicationFeeRequest represents the request to update an application fee
type UpdateApplicationFeeRequest struct {
	FeeName     *string  `json:"fee_name,omitempty" validate:"omitempty,min=3,max=255"`
	FeeType     *FeeType `json:"fee_type,omitempty" validate:"omitempty,oneof=PERCENTAGE FIXED"`
	FeeValue    *float64 `json:"fee_value,omitempty" validate:"omitempty,gt=0"`
	Description *string  `json:"description,omitempty"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

// ApplicationFeeResponse represents the response for application fee
type ApplicationFeeResponse struct {
	ID          uuid.UUID  `json:"id"`
	FeeName     string     `json:"fee_name"`
	FeeType     FeeType    `json:"fee_type"`
	FeeValue    float64    `json:"fee_value"`
	Description *string    `json:"description,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// ApplicationFeeListResponse represents the response for listing application fees
type ApplicationFeeListResponse struct {
	Fees       []ApplicationFee `json:"fees"`
	Total      int              `json:"total"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	TotalPages int              `json:"total_pages"`
}

// CalculateFeeAmount calculates the fee amount based on the fee type and value
func (f *ApplicationFee) CalculateFeeAmount(baseAmount float64) float64 {
	if f.FeeType == FeeTypePercentage {
		return baseAmount * (f.FeeValue / 100)
	}
	return f.FeeValue
}
