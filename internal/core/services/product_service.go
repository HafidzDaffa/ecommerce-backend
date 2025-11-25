package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

type productService struct {
	productRepo        ports.ProductRepository
	productGalleryRepo ports.ProductGalleryRepository
	storageService     ports.StorageService
}

func NewProductService(
	productRepo ports.ProductRepository,
	productGalleryRepo ports.ProductGalleryRepository,
	storageService ports.StorageService,
) ports.ProductService {
	return &productService{
		productRepo:        productRepo,
		productGalleryRepo: productGalleryRepo,
		storageService:     storageService,
	}
}

func (s *productService) CreateProduct(userID uuid.UUID, req *domain.CreateProductRequest) (*domain.Product, error) {
	product := &domain.Product{
		UserID:           userID,
		ProductName:      req.ProductName,
		Slug:             req.Slug,
		SKU:              req.SKU,
		Price:            req.Price,
		DiscountPercent:  req.DiscountPercent,
		ShortDescription: req.ShortDescription,
		Description:      req.Description,
		WeightGram:       req.WeightGram,
		StockQuantity:    req.StockQuantity,
		IsPublished:      true,
	}

	if req.IsPublished != nil {
		product.IsPublished = *req.IsPublished
	}

	err := s.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	if len(req.CategoryIDs) > 0 {
		err = s.productRepo.AddCategories(product.ID, req.CategoryIDs)
		if err != nil {
			return nil, err
		}
	}

	categories, _ := s.productRepo.GetCategories(product.ID)
	product.Categories = categories

	return product, nil
}

func (s *productService) GetProductByID(id uuid.UUID) (*domain.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	categories, _ := s.productRepo.GetCategories(id)
	product.Categories = categories

	galleries, _ := s.productGalleryRepo.GetByProductID(id)
	product.Galleries = galleries

	return product, nil
}

func (s *productService) GetProductBySlug(slug string) (*domain.Product, error) {
	product, err := s.productRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	categories, _ := s.productRepo.GetCategories(product.ID)
	product.Categories = categories

	galleries, _ := s.productGalleryRepo.GetByProductID(product.ID)
	product.Galleries = galleries

	return product, nil
}

func (s *productService) GetAllProducts(page, limit int, isPublished *bool) ([]domain.Product, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, total, err := s.productRepo.GetAll(page, limit, isPublished)
	if err != nil {
		return nil, 0, err
	}

	for i := range products {
		categories, _ := s.productRepo.GetCategories(products[i].ID)
		products[i].Categories = categories

		galleries, _ := s.productGalleryRepo.GetByProductID(products[i].ID)
		products[i].Galleries = galleries
	}

	return products, total, nil
}

func (s *productService) GetProductsByCategoryID(categoryID int, page, limit int) ([]domain.Product, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, total, err := s.productRepo.GetByCategoryID(categoryID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	for i := range products {
		categories, _ := s.productRepo.GetCategories(products[i].ID)
		products[i].Categories = categories

		galleries, _ := s.productGalleryRepo.GetByProductID(products[i].ID)
		products[i].Galleries = galleries
	}

	return products, total, nil
}

func (s *productService) UpdateProduct(id uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.ProductName != nil {
		product.ProductName = *req.ProductName
	}

	if req.Slug != nil {
		product.Slug = *req.Slug
	}

	if req.SKU != nil {
		product.SKU = req.SKU
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.DiscountPercent != nil {
		product.DiscountPercent = *req.DiscountPercent
	}

	if req.ShortDescription != nil {
		product.ShortDescription = req.ShortDescription
	}

	if req.Description != nil {
		product.Description = req.Description
	}

	if req.WeightGram != nil {
		product.WeightGram = *req.WeightGram
	}

	if req.StockQuantity != nil {
		product.StockQuantity = *req.StockQuantity
	}

	if req.IsPublished != nil {
		product.IsPublished = *req.IsPublished
	}

	err = s.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	if len(req.CategoryIDs) > 0 {
		currentCategories, _ := s.productRepo.GetCategories(id)
		var currentCategoryIDs []int
		for _, cat := range currentCategories {
			currentCategoryIDs = append(currentCategoryIDs, cat.ID)
		}

		s.productRepo.RemoveCategories(id, currentCategoryIDs)
		s.productRepo.AddCategories(id, req.CategoryIDs)
	}

	categories, _ := s.productRepo.GetCategories(id)
	product.Categories = categories

	galleries, _ := s.productGalleryRepo.GetByProductID(id)
	product.Galleries = galleries

	return product, nil
}

func (s *productService) DeleteProduct(id uuid.UUID) error {
	galleries, _ := s.productGalleryRepo.GetByProductID(id)
	for _, gallery := range galleries {
		s.storageService.DeleteFile(gallery.ImagePath)
		s.productGalleryRepo.Delete(gallery.ID)
	}

	return s.productRepo.Delete(id)
}

func (s *productService) AddProductGallery(req *domain.CreateProductGalleryRequest, imageFile *multipart.FileHeader) (*domain.ProductGallery, error) {
	if imageFile == nil {
		return nil, errors.New("image file is required")
	}

	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	_, err = s.productRepo.GetByID(productID)
	if err != nil {
		return nil, err
	}

	imageURL, err := s.storageService.UploadFile(imageFile, "products")
	if err != nil {
		return nil, err
	}

	gallery := &domain.ProductGallery{
		ProductID:    productID,
		ImagePath:    imageURL,
		DisplayOrder: req.DisplayOrder,
		IsThumbnail:  false,
	}

	if req.IsThumbnail != nil {
		gallery.IsThumbnail = *req.IsThumbnail
	}

	err = s.productGalleryRepo.Create(gallery)
	if err != nil {
		s.storageService.DeleteFile(imageURL)
		return nil, err
	}

	return gallery, nil
}

func (s *productService) GetProductGalleries(productID uuid.UUID) ([]domain.ProductGallery, error) {
	return s.productGalleryRepo.GetByProductID(productID)
}

func (s *productService) UpdateProductGallery(id uuid.UUID, req *domain.UpdateProductGalleryRequest) (*domain.ProductGallery, error) {
	gallery, err := s.productGalleryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.DisplayOrder != nil {
		gallery.DisplayOrder = *req.DisplayOrder
	}

	if req.IsThumbnail != nil {
		gallery.IsThumbnail = *req.IsThumbnail
	}

	err = s.productGalleryRepo.Update(gallery)
	if err != nil {
		return nil, err
	}

	return gallery, nil
}

func (s *productService) DeleteProductGallery(id uuid.UUID) error {
	gallery, err := s.productGalleryRepo.GetByID(id)
	if err != nil {
		return err
	}

	s.storageService.DeleteFile(gallery.ImagePath)

	return s.productGalleryRepo.Delete(id)
}
