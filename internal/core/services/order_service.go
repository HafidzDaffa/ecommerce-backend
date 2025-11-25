package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"math"

	"github.com/google/uuid"
)

type orderService struct {
	orderRepo   ports.OrderRepository
	cartRepo    ports.CartRepository
	productRepo ports.ProductRepository
}

func NewOrderService(orderRepo ports.OrderRepository, cartRepo ports.CartRepository, productRepo ports.ProductRepository) ports.OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) CreateOrder(userID uuid.UUID, req *domain.CreateOrderRequest) (*domain.Order, error) {
	if len(req.CartIDs) == 0 {
		return nil, errors.New("cart items are required")
	}

	var carts []domain.Cart
	for _, cartID := range req.CartIDs {
		cart, err := s.cartRepo.GetByID(cartID)
		if err != nil {
			return nil, errors.New("cart item not found")
		}

		if cart.UserID != userID {
			return nil, errors.New("unauthorized cart access")
		}

		carts = append(carts, *cart)
	}

	var productTotal float64
	var orderItems []domain.OrderItem

	for _, cart := range carts {
		product, err := s.productRepo.GetByID(cart.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if product.StockQuantity < cart.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.ProductName)
		}

		itemTotal := product.Price * float64(cart.Quantity)
		productTotal += itemTotal

		orderItem := domain.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     product.Price,
			Note:      cart.Note,
		}
		orderItems = append(orderItems, orderItem)
	}

	order := &domain.Order{
		UserID:          userID,
		TotalAmount:     productTotal + req.ShippingCost,
		ShippingCost:    req.ShippingCost,
		ProductTotal:    productTotal,
		AddressLine:     req.AddressLine,
		PostalCode:      req.PostalCode,
		ProvinceID:      req.ProvinceID,
		CityID:          req.CityID,
		SubdistrictID:   req.SubdistrictID,
		ProvinceName:    req.ProvinceName,
		CityName:        req.CityName,
		SubdistrictName: req.SubdistrictName,
		ShippingCourier: req.ShippingCourier,
		ShippingService: req.ShippingService,
		OrderStatusID:   1,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, errors.New("failed to create order")
	}

	for _, item := range orderItems {
		item.OrderID = order.ID
		if err := s.orderRepo.CreateOrderItem(&item); err != nil {
			return nil, errors.New("failed to create order item")
		}
	}

	if err := s.cartRepo.DeleteByIDs(userID, req.CartIDs); err != nil {
		return nil, errors.New("failed to clear cart items")
	}

	order.Items = orderItems
	return order, nil
}

func (s *orderService) GetOrder(userID, orderID uuid.UUID) (*domain.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	items, err := s.orderRepo.GetOrderItems(orderID)
	if err != nil {
		return nil, err
	}

	for i := range items {
		product, err := s.productRepo.GetByID(items[i].ProductID)
		if err == nil {
			items[i].Product = product
		}
	}

	order.Items = items

	statuses, err := s.orderRepo.GetOrderStatuses()
	if err == nil {
		for _, status := range statuses {
			if status.ID == order.OrderStatusID {
				order.OrderStatus = &status
				break
			}
		}
	}

	return order, nil
}

func (s *orderService) GetUserOrders(userID uuid.UUID, page, perPage int) (*domain.OrderListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	orders, total, err := s.orderRepo.GetByUser(userID, page, perPage)
	if err != nil {
		return nil, err
	}

	statuses, _ := s.orderRepo.GetOrderStatuses()

	for i := range orders {
		items, err := s.orderRepo.GetOrderItems(orders[i].ID)
		if err == nil {
			orders[i].Items = items
		}

		for _, status := range statuses {
			if status.ID == orders[i].OrderStatusID {
				orders[i].OrderStatus = &status
				break
			}
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &domain.OrderListResponse{
		Orders:     orders,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *orderService) GetAllOrders(page, perPage int) (*domain.OrderListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	orders, total, err := s.orderRepo.GetAll(page, perPage)
	if err != nil {
		return nil, err
	}

	statuses, _ := s.orderRepo.GetOrderStatuses()

	for i := range orders {
		items, err := s.orderRepo.GetOrderItems(orders[i].ID)
		if err == nil {
			orders[i].Items = items
		}

		for _, status := range statuses {
			if status.ID == orders[i].OrderStatusID {
				orders[i].OrderStatus = &status
				break
			}
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &domain.OrderListResponse{
		Orders:     orders,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *orderService) UpdateOrderStatus(orderID uuid.UUID, req *domain.UpdateOrderStatusRequest) error {
	_, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	return s.orderRepo.UpdateStatus(orderID, req.OrderStatusID, req.TrackingNumber)
}

func (s *orderService) CancelOrder(userID, orderID uuid.UUID) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("unauthorized")
	}

	if order.OrderStatusID > 2 {
		return errors.New("order cannot be cancelled")
	}

	return s.orderRepo.Delete(orderID)
}

func (s *orderService) GetOrderStatuses() ([]domain.OrderStatus, error) {
	return s.orderRepo.GetOrderStatuses()
}
