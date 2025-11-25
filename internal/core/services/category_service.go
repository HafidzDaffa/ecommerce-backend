package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"mime/multipart"
	"strings"
)

type categoryService struct {
	categoryRepo   ports.CategoryRepository
	storageService ports.StorageService
}

func NewCategoryService(categoryRepo ports.CategoryRepository, storageService ports.StorageService) ports.CategoryService {
	return &categoryService{
		categoryRepo:   categoryRepo,
		storageService: storageService,
	}
}

func (s *categoryService) CreateCategory(req *domain.CreateCategoryRequest, imageFile *multipart.FileHeader) (*domain.Category, error) {
	category := &domain.Category{
		CategoryName: req.CategoryName,
		Slug:         req.Slug,
		Icon:         req.Icon,
		ImagePath:    req.ImagePath,
		IsActive:     true,
	}

	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if category.Slug == nil || *category.Slug == "" {
		slug := generateSlug(req.CategoryName)
		category.Slug = &slug
	}

	if imageFile != nil {
		imageURL, err := s.storageService.UploadFile(imageFile, "categories")
		if err != nil {
			return nil, err
		}
		category.ImagePath = &imageURL
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategoryByID(id int) (*domain.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *categoryService) GetCategoryBySlug(slug string) (*domain.Category, error) {
	return s.categoryRepo.GetBySlug(slug)
}

func (s *categoryService) GetAllCategories(isActive *bool) ([]domain.Category, error) {
	return s.categoryRepo.GetAll(isActive)
}

func (s *categoryService) UpdateCategory(id int, req *domain.UpdateCategoryRequest, imageFile *multipart.FileHeader) (*domain.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.CategoryName != nil {
		category.CategoryName = *req.CategoryName
	}

	if req.Slug != nil {
		category.Slug = req.Slug
	}

	if req.Icon != nil {
		category.Icon = req.Icon
	}

	if req.ImagePath != nil {
		category.ImagePath = req.ImagePath
	}

	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if imageFile != nil {
		if category.ImagePath != nil && *category.ImagePath != "" {
			s.storageService.DeleteFile(*category.ImagePath)
		}

		imageURL, err := s.storageService.UploadFile(imageFile, "categories")
		if err != nil {
			return nil, err
		}
		category.ImagePath = &imageURL
	}

	err = s.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(id int) error {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return err
	}

	if category.ImagePath != nil && *category.ImagePath != "" {
		s.storageService.DeleteFile(*category.ImagePath)
	}

	return s.categoryRepo.Delete(id)
}

func generateSlug(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")
	text = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, text)
	return text
}
