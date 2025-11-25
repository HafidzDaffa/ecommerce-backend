package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *domain.PaymentTransaction) error
	GetByID(id uuid.UUID) (*domain.PaymentTransaction, error)
	GetByOrderID(orderID uuid.UUID) (*domain.PaymentTransaction, error)
	GetByExternalID(externalID string) (*domain.PaymentTransaction, error)
	GetByInvoiceID(invoiceID string) (*domain.PaymentTransaction, error)
	Update(payment *domain.PaymentTransaction) error
	UpdateStatus(id uuid.UUID, status domain.PaymentStatus, paidAt *string, paidAmount float64) error
	GetUserPayments(userID uuid.UUID, page, perPage int) ([]domain.PaymentTransaction, int, error)
	GetAllPayments(page, perPage int) ([]domain.PaymentTransaction, int, error)
	Delete(id uuid.UUID) error
}
