package ports

import "ecommerce-backend/internal/core/domain"

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetByID(id int) (*domain.Category, error)
	GetBySlug(slug string) (*domain.Category, error)
	GetAll(isActive *bool) ([]domain.Category, error)
	Update(category *domain.Category) error
	Delete(id int) error
}
