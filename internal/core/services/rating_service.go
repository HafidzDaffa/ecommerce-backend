package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"

	"github.com/google/uuid"
)

type ratingService struct {
	ratingRepo  ports.RatingRepository
	orderRepo   ports.OrderRepository
	productRepo ports.ProductRepository
	userRepo    ports.UserRepository
}

func NewRatingService(ratingRepo ports.RatingRepository, orderRepo ports.OrderRepository, productRepo ports.ProductRepository, userRepo ports.UserRepository) ports.RatingService {
	return &ratingService{
		ratingRepo:  ratingRepo,
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (s *ratingService) CreateRating(userID uuid.UUID, req *domain.CreateRatingRequest) (*domain.ProductRating, error) {
	order, err := s.orderRepo.GetByID(req.OrderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if order.OrderStatusID < 4 {
		return nil, errors.New("can only rate delivered orders")
	}

	_, err = s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	items, err := s.orderRepo.GetOrderItems(req.OrderID)
	if err != nil {
		return nil, err
	}

	found := false
	for _, item := range items {
		if item.ProductID == req.ProductID {
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("product not found in order")
	}

	existing, err := s.ratingRepo.GetByProductUserOrder(req.ProductID, userID, req.OrderID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("rating already exists for this product")
	}

	if req.Rating < 1 || req.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	rating := &domain.ProductRating{
		ProductID: req.ProductID,
		UserID:    userID,
		OrderID:   req.OrderID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}

	if err := s.ratingRepo.Create(rating); err != nil {
		return nil, errors.New("failed to create rating")
	}

	return rating, nil
}

func (s *ratingService) GetProductRatings(productID uuid.UUID, page, perPage int) ([]domain.ProductRating, int, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	ratings, total, err := s.ratingRepo.GetByProduct(productID, page, perPage)
	if err != nil {
		return nil, 0, err
	}

	for i := range ratings {
		user, err := s.userRepo.GetByID(ratings[i].UserID)
		if err == nil {
			user.PasswordHash = ""
			ratings[i].User = user
		}
	}

	return ratings, total, nil
}

func (s *ratingService) GetUserRatings(userID uuid.UUID) ([]domain.ProductRating, error) {
	ratings, err := s.ratingRepo.GetByUser(userID)
	if err != nil {
		return nil, err
	}

	for i := range ratings {
		product, err := s.productRepo.GetByID(ratings[i].ProductID)
		if err == nil {
			ratings[i].Product = product
		}
	}

	return ratings, nil
}

func (s *ratingService) GetRatingStats(productID uuid.UUID) (*domain.ProductRatingStats, error) {
	return s.ratingRepo.GetStats(productID)
}

func (s *ratingService) UpdateRating(userID, ratingID uuid.UUID, req *domain.UpdateRatingRequest) (*domain.ProductRating, error) {
	rating, err := s.ratingRepo.GetByID(ratingID)
	if err != nil {
		return nil, err
	}

	if rating.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if req.Rating != nil {
		if *req.Rating < 1 || *req.Rating > 5 {
			return nil, errors.New("rating must be between 1 and 5")
		}
		rating.Rating = *req.Rating
	}

	if req.Comment != nil {
		rating.Comment = req.Comment
	}

	if err := s.ratingRepo.Update(rating); err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *ratingService) DeleteRating(userID, ratingID uuid.UUID) error {
	rating, err := s.ratingRepo.GetByID(ratingID)
	if err != nil {
		return err
	}

	if rating.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.ratingRepo.Delete(ratingID)
}
