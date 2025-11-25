package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductRating struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	ProductID uuid.UUID  `db:"product_id" json:"product_id"`
	UserID    uuid.UUID  `db:"user_id" json:"user_id"`
	OrderID   uuid.UUID  `db:"order_id" json:"order_id"`
	Rating    int        `db:"rating" json:"rating"`
	Comment   *string    `db:"comment" json:"comment,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	User      *User      `db:"-" json:"user,omitempty"`
	Product   *Product   `db:"-" json:"product,omitempty"`
}

type CreateRatingRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	OrderID   uuid.UUID `json:"order_id" validate:"required"`
	Rating    int       `json:"rating" validate:"required,min=1,max=5"`
	Comment   *string   `json:"comment"`
}

type UpdateRatingRequest struct {
	Rating  *int    `json:"rating" validate:"omitempty,min=1,max=5"`
	Comment *string `json:"comment"`
}

type ProductRatingStats struct {
	ProductID     uuid.UUID `json:"product_id"`
	AverageRating float64   `json:"average_rating"`
	TotalRatings  int       `json:"total_ratings"`
	Rating1Count  int       `json:"rating_1_count"`
	Rating2Count  int       `json:"rating_2_count"`
	Rating3Count  int       `json:"rating_3_count"`
	Rating4Count  int       `json:"rating_4_count"`
	Rating5Count  int       `json:"rating_5_count"`
}
