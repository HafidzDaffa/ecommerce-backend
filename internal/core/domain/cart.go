package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	UserID     uuid.UUID  `db:"user_id" json:"user_id"`
	ProductID  uuid.UUID  `db:"product_id" json:"product_id"`
	Quantity   int        `db:"quantity" json:"quantity"`
	Note       *string    `db:"note" json:"note"`
	IsSelected bool       `db:"is_selected" json:"is_selected"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Product    *Product   `db:"-" json:"product,omitempty"`
}

type CartResponse struct {
	ID         uuid.UUID `json:"id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Note       *string   `json:"note"`
	IsSelected bool      `json:"is_selected"`
	Product    *Product  `json:"product,omitempty"`
	Subtotal   float64   `json:"subtotal"`
}

type AddToCartRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
	Note      *string   `json:"note"`
}

type UpdateCartRequest struct {
	Quantity   *int    `json:"quantity" validate:"omitempty,min=1"`
	Note       *string `json:"note"`
	IsSelected *bool   `json:"is_selected"`
}

type CartSummary struct {
	Items            []CartResponse `json:"items"`
	TotalItems       int            `json:"total_items"`
	SelectedItems    int            `json:"selected_items"`
	TotalAmount      float64        `json:"total_amount"`
	SelectedAmount   float64        `json:"selected_amount"`
}
