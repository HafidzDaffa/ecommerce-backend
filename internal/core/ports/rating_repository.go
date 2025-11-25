package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type RatingRepository interface {
	Create(rating *domain.ProductRating) error
	GetByID(id uuid.UUID) (*domain.ProductRating, error)
	GetByProduct(productID uuid.UUID, page, perPage int) ([]domain.ProductRating, int, error)
	GetByUser(userID uuid.UUID) ([]domain.ProductRating, error)
	GetByProductUserOrder(productID, userID, orderID uuid.UUID) (*domain.ProductRating, error)
	GetStats(productID uuid.UUID) (*domain.ProductRatingStats, error)
	Update(rating *domain.ProductRating) error
	Delete(id uuid.UUID) error
}
