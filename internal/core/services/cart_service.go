package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"

	"github.com/google/uuid"
)

type cartService struct {
	cartRepo    ports.CartRepository
	productRepo ports.ProductRepository
}

func NewCartService(cartRepo ports.CartRepository, productRepo ports.ProductRepository) ports.CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *cartService) AddToCart(userID uuid.UUID, req *domain.AddToCartRequest) (*domain.Cart, error) {
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if product.StockQuantity < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	existingCart, err := s.cartRepo.GetByUserAndProduct(userID, req.ProductID)
	if err != nil {
		return nil, err
	}

	if existingCart != nil {
		existingCart.Quantity += req.Quantity
		if product.StockQuantity < existingCart.Quantity {
			return nil, errors.New("insufficient stock")
		}
		existingCart.Note = req.Note
		existingCart.IsSelected = true
		
		if err := s.cartRepo.Update(existingCart); err != nil {
			return nil, err
		}
		return existingCart, nil
	}

	cart := &domain.Cart{
		UserID:     userID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
		Note:       req.Note,
		IsSelected: true,
	}

	if err := s.cartRepo.Create(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *cartService) GetCart(userID uuid.UUID) (*domain.CartSummary, error) {
	carts, err := s.cartRepo.GetByUser(userID)
	if err != nil {
		return nil, err
	}

	var cartResponses []domain.CartResponse
	var totalAmount, selectedAmount float64
	var totalItems, selectedItems int

	for _, cart := range carts {
		product, err := s.productRepo.GetByID(cart.ProductID)
		if err != nil {
			continue
		}

		subtotal := product.Price * float64(cart.Quantity)
		
		cartResponse := domain.CartResponse{
			ID:         cart.ID,
			ProductID:  cart.ProductID,
			Quantity:   cart.Quantity,
			Note:       cart.Note,
			IsSelected: cart.IsSelected,
			Product:    product,
			Subtotal:   subtotal,
		}

		cartResponses = append(cartResponses, cartResponse)
		totalAmount += subtotal
		totalItems++

		if cart.IsSelected {
			selectedAmount += subtotal
			selectedItems++
		}
	}

	return &domain.CartSummary{
		Items:          cartResponses,
		TotalItems:     totalItems,
		SelectedItems:  selectedItems,
		TotalAmount:    totalAmount,
		SelectedAmount: selectedAmount,
	}, nil
}

func (s *cartService) UpdateCartItem(userID, cartID uuid.UUID, req *domain.UpdateCartRequest) (*domain.Cart, error) {
	cart, err := s.cartRepo.GetByID(cartID)
	if err != nil {
		return nil, err
	}

	if cart.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if req.Quantity != nil {
		product, err := s.productRepo.GetByID(cart.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if product.StockQuantity < *req.Quantity {
			return nil, errors.New("insufficient stock")
		}

		cart.Quantity = *req.Quantity
	}

	if req.Note != nil {
		cart.Note = req.Note
	}

	if req.IsSelected != nil {
		cart.IsSelected = *req.IsSelected
	}

	if err := s.cartRepo.Update(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *cartService) RemoveFromCart(userID, cartID uuid.UUID) error {
	cart, err := s.cartRepo.GetByID(cartID)
	if err != nil {
		return err
	}

	if cart.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.cartRepo.Delete(cartID)
}

func (s *cartService) RemoveMultiple(userID uuid.UUID, cartIDs []uuid.UUID) error {
	return s.cartRepo.DeleteByIDs(userID, cartIDs)
}

func (s *cartService) ClearCart(userID uuid.UUID) error {
	return s.cartRepo.ClearCart(userID)
}
