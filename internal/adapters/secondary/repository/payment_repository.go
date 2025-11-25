package repository

import (
	"database/sql"
	"ecommerce-backend/internal/core/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment *domain.PaymentTransaction) error {
	query := `
		INSERT INTO payment_transactions (
			id, order_id, user_id, xendit_external_id, payment_method,
			payment_channel, payment_status, amount, paid_amount, admin_fee,
			total_amount, invoice_url, checkout_url, expired_at, description,
			payment_details, xendit_response, xendit_invoice_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		) RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		payment.ID,
		payment.OrderID,
		payment.UserID,
		payment.XenditExternalID,
		payment.PaymentMethod,
		payment.PaymentChannel,
		payment.PaymentStatus,
		payment.Amount,
		payment.PaidAmount,
		payment.AdminFee,
		payment.TotalAmount,
		payment.InvoiceURL,
		payment.CheckoutURL,
		payment.ExpiredAt,
		payment.Description,
		payment.PaymentDetails,
		payment.XenditResponse,
		payment.XenditInvoiceID,
	).Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)
}

func (r *PaymentRepository) GetByID(id uuid.UUID) (*domain.PaymentTransaction, error) {
	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, payment_details, xendit_response,
			created_at, updated_at, deleted_at
		FROM payment_transactions
		WHERE id = $1 AND deleted_at IS NULL
	`

	var payment domain.PaymentTransaction
	err := r.db.Get(&payment, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) GetByOrderID(orderID uuid.UUID) (*domain.PaymentTransaction, error) {
	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, payment_details, xendit_response,
			created_at, updated_at, deleted_at
		FROM payment_transactions
		WHERE order_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT 1
	`

	var payment domain.PaymentTransaction
	err := r.db.Get(&payment, query, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) GetByExternalID(externalID string) (*domain.PaymentTransaction, error) {
	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, payment_details, xendit_response,
			created_at, updated_at, deleted_at
		FROM payment_transactions
		WHERE xendit_external_id = $1 AND deleted_at IS NULL
	`

	var payment domain.PaymentTransaction
	err := r.db.Get(&payment, query, externalID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) GetByInvoiceID(invoiceID string) (*domain.PaymentTransaction, error) {
	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, payment_details, xendit_response,
			created_at, updated_at, deleted_at
		FROM payment_transactions
		WHERE xendit_invoice_id = $1 AND deleted_at IS NULL
	`

	var payment domain.PaymentTransaction
	err := r.db.Get(&payment, query, invoiceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) Update(payment *domain.PaymentTransaction) error {
	query := `
		UPDATE payment_transactions SET
			payment_method = $1,
			payment_channel = $2,
			payment_status = $3,
			paid_amount = $4,
			paid_at = $5,
			payment_details = $6,
			xendit_response = $7,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $8 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(
		query,
		payment.PaymentMethod,
		payment.PaymentChannel,
		payment.PaymentStatus,
		payment.PaidAmount,
		payment.PaidAt,
		payment.PaymentDetails,
		payment.XenditResponse,
		payment.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("payment not found")
	}

	return nil
}

func (r *PaymentRepository) UpdateStatus(id uuid.UUID, status domain.PaymentStatus, paidAt *string, paidAmount float64) error {
	query := `
		UPDATE payment_transactions SET
			payment_status = $1,
			paid_at = $2,
			paid_amount = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query, status, paidAt, paidAmount, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("payment not found")
	}

	return nil
}

func (r *PaymentRepository) GetUserPayments(userID uuid.UUID, page, perPage int) ([]domain.PaymentTransaction, int, error) {
	offset := (page - 1) * perPage

	countQuery := `
		SELECT COUNT(*) FROM payment_transactions
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	var total int
	err := r.db.Get(&total, countQuery, userID)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, created_at, updated_at
		FROM payment_transactions
		WHERE user_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	var payments []domain.PaymentTransaction
	err = r.db.Select(&payments, query, userID, perPage, offset)
	if err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}

func (r *PaymentRepository) GetAllPayments(page, perPage int) ([]domain.PaymentTransaction, int, error) {
	offset := (page - 1) * perPage

	countQuery := `SELECT COUNT(*) FROM payment_transactions WHERE deleted_at IS NULL`
	var total int
	err := r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			id, order_id, user_id, xendit_invoice_id, xendit_external_id,
			payment_method, payment_channel, payment_status, amount, paid_amount,
			admin_fee, total_amount, invoice_url, checkout_url, paid_at,
			expired_at, description, created_at, updated_at
		FROM payment_transactions
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var payments []domain.PaymentTransaction
	err = r.db.Select(&payments, query, perPage, offset)
	if err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}

func (r *PaymentRepository) Delete(id uuid.UUID) error {
	query := `
		UPDATE payment_transactions SET
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("payment not found")
	}

	return nil
}
