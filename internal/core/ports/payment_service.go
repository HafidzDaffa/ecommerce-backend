package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(userID uuid.UUID, req *domain.CreatePaymentRequest) (*domain.PaymentResponse, error)
	GetPaymentByID(userID uuid.UUID, paymentID uuid.UUID) (*domain.PaymentTransaction, error)
	GetPaymentByOrderID(userID uuid.UUID, orderID uuid.UUID) (*domain.PaymentTransaction, error)
	GetUserPayments(userID uuid.UUID, page, perPage int) (*domain.PaymentListResponse, error)
	GetAllPayments(page, perPage int) (*domain.PaymentListResponse, error)
	HandleXenditCallback(payload *domain.XenditCallbackPayload) error
	CheckPaymentStatus(userID uuid.UUID, paymentID uuid.UUID) (*domain.PaymentTransaction, error)
	CancelPayment(userID uuid.UUID, paymentID uuid.UUID) error
}
