package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	CreateOrderItem(item *domain.OrderItem) error
	GetByID(id uuid.UUID) (*domain.Order, error)
	GetByUser(userID uuid.UUID, page, perPage int) ([]domain.Order, int, error)
	GetAll(page, perPage int) ([]domain.Order, int, error)
	GetOrderItems(orderID uuid.UUID) ([]domain.OrderItem, error)
	GetOrderStatuses() ([]domain.OrderStatus, error)
	UpdateStatus(orderID uuid.UUID, statusID int, trackingNumber *string) error
	Update(order *domain.Order) error
	Delete(id uuid.UUID) error
}
