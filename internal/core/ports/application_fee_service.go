package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type ApplicationFeeService interface {
	CreateApplicationFee(req *domain.CreateApplicationFeeRequest, createdBy uuid.UUID) (*domain.ApplicationFee, error)
	GetApplicationFeeByID(id uuid.UUID) (*domain.ApplicationFee, error)
	GetAllApplicationFees(isActive *bool, page, perPage int) (*domain.ApplicationFeeListResponse, error)
	GetActiveByType(feeType domain.FeeType) (*domain.ApplicationFee, error)
	UpdateApplicationFee(id uuid.UUID, req *domain.UpdateApplicationFeeRequest) (*domain.ApplicationFee, error)
	DeleteApplicationFee(id uuid.UUID) error
	CalculateFee(feeID uuid.UUID, baseAmount float64) (float64, error)
}
