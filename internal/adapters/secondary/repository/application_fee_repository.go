package repository

import (
	"database/sql"
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type applicationFeeRepository struct {
	db *sqlx.DB
}

func NewApplicationFeeRepository(db *sqlx.DB) ports.ApplicationFeeRepository {
	return &applicationFeeRepository{
		db: db,
	}
}

func (r *applicationFeeRepository) Create(fee *domain.ApplicationFee) error {
	query := `
		INSERT INTO application_fees (fee_name, fee_type, fee_value, description, is_active, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		fee.FeeName,
		fee.FeeType,
		fee.FeeValue,
		fee.Description,
		fee.IsActive,
		fee.CreatedBy,
		time.Now(),
	).Scan(&fee.ID)

	return err
}

func (r *applicationFeeRepository) GetByID(id uuid.UUID) (*domain.ApplicationFee, error) {
	var fee domain.ApplicationFee
	query := `
		SELECT id, fee_name, fee_type, fee_value, description, is_active, created_by, 
		       created_at, updated_at, deleted_at
		FROM application_fees
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.db.Get(&fee, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("application fee not found")
		}
		return nil, err
	}

	return &fee, nil
}

func (r *applicationFeeRepository) GetAll(isActive *bool, page, perPage int) ([]domain.ApplicationFee, int, error) {
	var fees []domain.ApplicationFee
	var total int

	// Build base query
	query := `
		SELECT id, fee_name, fee_type, fee_value, description, is_active, created_by, 
		       created_at, updated_at, deleted_at
		FROM application_fees
		WHERE deleted_at IS NULL
	`

	countQuery := `
		SELECT COUNT(*)
		FROM application_fees
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	// Add is_active filter if provided
	if isActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", argPos)
		countQuery += fmt.Sprintf(" AND is_active = $%d", argPos)
		args = append(args, *isActive)
		argPos++
	}

	// Get total count
	countArgs := args
	err := r.db.Get(&total, countQuery, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	// Add ordering and pagination
	query += " ORDER BY created_at DESC"
	if perPage > 0 {
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
		args = append(args, perPage, (page-1)*perPage)
	}

	// Execute query
	err = r.db.Select(&fees, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return fees, total, nil
}

func (r *applicationFeeRepository) GetActiveByType(feeType domain.FeeType) (*domain.ApplicationFee, error) {
	var fee domain.ApplicationFee
	query := `
		SELECT id, fee_name, fee_type, fee_value, description, is_active, created_by, 
		       created_at, updated_at, deleted_at
		FROM application_fees
		WHERE fee_type = $1 AND is_active = true AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT 1
	`

	err := r.db.Get(&fee, query, feeType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no active fee found
		}
		return nil, err
	}

	return &fee, nil
}

func (r *applicationFeeRepository) Update(fee *domain.ApplicationFee) error {
	query := `
		UPDATE application_fees
		SET fee_name = $1, fee_type = $2, fee_value = $3, description = $4, 
		    is_active = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(
		query,
		fee.FeeName,
		fee.FeeType,
		fee.FeeValue,
		fee.Description,
		fee.IsActive,
		time.Now(),
		fee.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("application fee not found")
	}

	return nil
}

func (r *applicationFeeRepository) Delete(id uuid.UUID) error {
	query := `
		UPDATE application_fees
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("application fee not found")
	}

	return nil
}
