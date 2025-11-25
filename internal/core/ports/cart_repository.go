package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type CartRepository interface {
	Create(cart *domain.Cart) error
	GetByID(id uuid.UUID) (*domain.Cart, error)
	GetByUserAndProduct(userID, productID uuid.UUID) (*domain.Cart, error)
	GetByUser(userID uuid.UUID) ([]domain.Cart, error)
	GetSelectedByUser(userID uuid.UUID) ([]domain.Cart, error)
	Update(cart *domain.Cart) error
	Delete(id uuid.UUID) error
	DeleteByIDs(userID uuid.UUID, ids []uuid.UUID) error
	ClearCart(userID uuid.UUID) error
}
