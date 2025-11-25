package ports

import (
	"ecommerce-backend/internal/core/domain"
	"mime/multipart"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(userID uuid.UUID, req *domain.CreateProductRequest) (*domain.Product, error)
	GetProductByID(id uuid.UUID) (*domain.Product, error)
	GetProductBySlug(slug string) (*domain.Product, error)
	GetAllProducts(page, limit int, isPublished *bool) ([]domain.Product, int, error)
	GetProductsByCategoryID(categoryID int, page, limit int) ([]domain.Product, int, error)
	UpdateProduct(id uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(id uuid.UUID) error
	
	AddProductGallery(req *domain.CreateProductGalleryRequest, imageFile *multipart.FileHeader) (*domain.ProductGallery, error)
	GetProductGalleries(productID uuid.UUID) ([]domain.ProductGallery, error)
	UpdateProductGallery(id uuid.UUID, req *domain.UpdateProductGalleryRequest) (*domain.ProductGallery, error)
	DeleteProductGallery(id uuid.UUID) error
}
