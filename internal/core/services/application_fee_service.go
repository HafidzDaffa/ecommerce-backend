package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"math"

	"github.com/google/uuid"
)

type applicationFeeService struct {
	feeRepo ports.ApplicationFeeRepository
}

func NewApplicationFeeService(feeRepo ports.ApplicationFeeRepository) ports.ApplicationFeeService {
	return &applicationFeeService{
		feeRepo: feeRepo,
	}
}

func (s *applicationFeeService) CreateApplicationFee(req *domain.CreateApplicationFeeRequest, createdBy uuid.UUID) (*domain.ApplicationFee, error) {
	// Validate fee value based on fee type
	if req.FeeType == domain.FeeTypePercentage && req.FeeValue > 100 {
		return nil, errors.New("percentage fee value cannot exceed 100")
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	fee := &domain.ApplicationFee{
		FeeName:     req.FeeName,
		FeeType:     req.FeeType,
		FeeValue:    req.FeeValue,
		Description: req.Description,
		IsActive:    isActive,
		CreatedBy:   &createdBy,
	}

	err := s.feeRepo.Create(fee)
	if err != nil {
		return nil, err
	}

	return fee, nil
}

func (s *applicationFeeService) GetApplicationFeeByID(id uuid.UUID) (*domain.ApplicationFee, error) {
	return s.feeRepo.GetByID(id)
}

func (s *applicationFeeService) GetAllApplicationFees(isActive *bool, page, perPage int) (*domain.ApplicationFeeListResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}

	fees, total, err := s.feeRepo.GetAll(isActive, page, perPage)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &domain.ApplicationFeeListResponse{
		Fees:       fees,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}

func (s *applicationFeeService) GetActiveByType(feeType domain.FeeType) (*domain.ApplicationFee, error) {
	return s.feeRepo.GetActiveByType(feeType)
}

func (s *applicationFeeService) UpdateApplicationFee(id uuid.UUID, req *domain.UpdateApplicationFeeRequest) (*domain.ApplicationFee, error) {
	fee, err := s.feeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.FeeName != nil {
		fee.FeeName = *req.FeeName
	}

	if req.FeeType != nil {
		fee.FeeType = *req.FeeType
	}

	if req.FeeValue != nil {
		// Validate fee value based on fee type
		if fee.FeeType == domain.FeeTypePercentage && *req.FeeValue > 100 {
			return nil, errors.New("percentage fee value cannot exceed 100")
		}
		fee.FeeValue = *req.FeeValue
	}

	if req.Description != nil {
		fee.Description = req.Description
	}

	if req.IsActive != nil {
		fee.IsActive = *req.IsActive
	}

	err = s.feeRepo.Update(fee)
	if err != nil {
		return nil, err
	}

	return fee, nil
}

func (s *applicationFeeService) DeleteApplicationFee(id uuid.UUID) error {
	_, err := s.feeRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.feeRepo.Delete(id)
}

func (s *applicationFeeService) CalculateFee(feeID uuid.UUID, baseAmount float64) (float64, error) {
	fee, err := s.feeRepo.GetByID(feeID)
	if err != nil {
		return 0, err
	}

	if !fee.IsActive {
		return 0, errors.New("application fee is not active")
	}

	return fee.CalculateFeeAmount(baseAmount), nil
}
