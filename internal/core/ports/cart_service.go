package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type CartService interface {
	AddToCart(userID uuid.UUID, req *domain.AddToCartRequest) (*domain.Cart, error)
	GetCart(userID uuid.UUID) (*domain.CartSummary, error)
	UpdateCartItem(userID, cartID uuid.UUID, req *domain.UpdateCartRequest) (*domain.Cart, error)
	RemoveFromCart(userID, cartID uuid.UUID) error
	RemoveMultiple(userID uuid.UUID, cartIDs []uuid.UUID) error
	ClearCart(userID uuid.UUID) error
}
