package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(userID uuid.UUID, req *domain.CreateOrderRequest) (*domain.Order, error)
	GetOrder(userID, orderID uuid.UUID) (*domain.Order, error)
	GetUserOrders(userID uuid.UUID, page, perPage int) (*domain.OrderListResponse, error)
	GetAllOrders(page, perPage int) (*domain.OrderListResponse, error)
	UpdateOrderStatus(orderID uuid.UUID, req *domain.UpdateOrderStatusRequest) error
	CancelOrder(userID, orderID uuid.UUID) error
	GetOrderStatuses() ([]domain.OrderStatus, error)
}
