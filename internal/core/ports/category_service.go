package ports

import (
	"ecommerce-backend/internal/core/domain"
	"mime/multipart"
)

type CategoryService interface {
	CreateCategory(req *domain.CreateCategoryRequest, imageFile *multipart.FileHeader) (*domain.Category, error)
	GetCategoryByID(id int) (*domain.Category, error)
	GetCategoryBySlug(slug string) (*domain.Category, error)
	GetAllCategories(isActive *bool) ([]domain.Category, error)
	UpdateCategory(id int, req *domain.UpdateCategoryRequest, imageFile *multipart.FileHeader) (*domain.Category, error)
	DeleteCategory(id int) error
}
