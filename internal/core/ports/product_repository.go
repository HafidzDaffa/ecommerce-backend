package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id uuid.UUID) (*domain.Product, error)
	GetBySlug(slug string) (*domain.Product, error)
	GetAll(page, limit int, isPublished *bool) ([]domain.Product, int, error)
	GetByCategoryID(categoryID int, page, limit int) ([]domain.Product, int, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
	AddCategories(productID uuid.UUID, categoryIDs []int) error
	RemoveCategories(productID uuid.UUID, categoryIDs []int) error
	GetCategories(productID uuid.UUID) ([]domain.Category, error)
}

type ProductGalleryRepository interface {
	Create(gallery *domain.ProductGallery) error
	GetByID(id uuid.UUID) (*domain.ProductGallery, error)
	GetByProductID(productID uuid.UUID) ([]domain.ProductGallery, error)
	Update(gallery *domain.ProductGallery) error
	Delete(id uuid.UUID) error
}
