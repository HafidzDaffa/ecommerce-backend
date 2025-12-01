package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type ApplicationFeeRepository interface {
	Create(fee *domain.ApplicationFee) error
	GetByID(id uuid.UUID) (*domain.ApplicationFee, error)
	GetAll(isActive *bool, page, perPage int) ([]domain.ApplicationFee, int, error)
	GetActiveByType(feeType domain.FeeType) (*domain.ApplicationFee, error)
	Update(fee *domain.ApplicationFee) error
	Delete(id uuid.UUID) error
}
