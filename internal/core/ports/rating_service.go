package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type RatingService interface {
	CreateRating(userID uuid.UUID, req *domain.CreateRatingRequest) (*domain.ProductRating, error)
	GetProductRatings(productID uuid.UUID, page, perPage int) ([]domain.ProductRating, int, error)
	GetUserRatings(userID uuid.UUID) ([]domain.ProductRating, error)
	GetRatingStats(productID uuid.UUID) (*domain.ProductRatingStats, error)
	UpdateRating(userID, ratingID uuid.UUID, req *domain.UpdateRatingRequest) (*domain.ProductRating, error)
	DeleteRating(userID, ratingID uuid.UUID) error
}
