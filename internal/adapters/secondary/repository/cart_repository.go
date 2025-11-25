package repository

import (
	"database/sql"
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) ports.CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) Create(cart *domain.Cart) error {
	query := `
		INSERT INTO carts (user_id, product_id, quantity, note, is_selected, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		cart.UserID,
		cart.ProductID,
		cart.Quantity,
		cart.Note,
		cart.IsSelected,
		time.Now(),
	).Scan(&cart.ID)
	
	return err
}

func (r *cartRepository) GetByID(id uuid.UUID) (*domain.Cart, error) {
	var cart domain.Cart
	query := `
		SELECT id, user_id, product_id, quantity, note, is_selected, created_at, updated_at
		FROM carts
		WHERE id = $1
	`
	
	err := r.db.Get(&cart, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("cart not found")
	}
	
	return &cart, err
}

func (r *cartRepository) GetByUserAndProduct(userID, productID uuid.UUID) (*domain.Cart, error) {
	var cart domain.Cart
	query := `
		SELECT id, user_id, product_id, quantity, note, is_selected, created_at, updated_at
		FROM carts
		WHERE user_id = $1 AND product_id = $2
	`
	
	err := r.db.Get(&cart, query, userID, productID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	
	return &cart, err
}

func (r *cartRepository) GetByUser(userID uuid.UUID) ([]domain.Cart, error) {
	var carts []domain.Cart
	query := `
		SELECT id, user_id, product_id, quantity, note, is_selected, created_at, updated_at
		FROM carts
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	
	err := r.db.Select(&carts, query, userID)
	if err != nil {
		return nil, err
	}
	
	return carts, nil
}

func (r *cartRepository) GetSelectedByUser(userID uuid.UUID) ([]domain.Cart, error) {
	var carts []domain.Cart
	query := `
		SELECT id, user_id, product_id, quantity, note, is_selected, created_at, updated_at
		FROM carts
		WHERE user_id = $1 AND is_selected = true
		ORDER BY created_at DESC
	`
	
	err := r.db.Select(&carts, query, userID)
	if err != nil {
		return nil, err
	}
	
	return carts, nil
}

func (r *cartRepository) Update(cart *domain.Cart) error {
	query := `
		UPDATE carts
		SET quantity = $1, note = $2, is_selected = $3, updated_at = $4
		WHERE id = $5
	`
	
	now := time.Now()
	cart.UpdatedAt = &now
	
	result, err := r.db.Exec(query, cart.Quantity, cart.Note, cart.IsSelected, now, cart.ID)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("cart not found")
	}
	
	return nil
}

func (r *cartRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM carts WHERE id = $1`
	
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("cart not found")
	}
	
	return nil
}

func (r *cartRepository) DeleteByIDs(userID uuid.UUID, ids []uuid.UUID) error {
	query := `DELETE FROM carts WHERE user_id = $1 AND id = ANY($2)`
	
	_, err := r.db.Exec(query, userID, ids)
	return err
}

func (r *cartRepository) ClearCart(userID uuid.UUID) error {
	query := `DELETE FROM carts WHERE user_id = $1`
	
	_, err := r.db.Exec(query, userID)
	return err
}
