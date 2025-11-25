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

type ratingRepository struct {
	db *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) ports.RatingRepository {
	return &ratingRepository{
		db: db,
	}
}

func (r *ratingRepository) Create(rating *domain.ProductRating) error {
	query := `
		INSERT INTO product_ratings (product_id, user_id, order_id, rating, comment, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		rating.ProductID,
		rating.UserID,
		rating.OrderID,
		rating.Rating,
		rating.Comment,
		time.Now(),
	).Scan(&rating.ID)
	
	return err
}

func (r *ratingRepository) GetByID(id uuid.UUID) (*domain.ProductRating, error) {
	var rating domain.ProductRating
	query := `
		SELECT id, product_id, user_id, order_id, rating, comment, created_at, updated_at, deleted_at
		FROM product_ratings
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&rating, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("rating not found")
	}
	
	return &rating, err
}

func (r *ratingRepository) GetByProduct(productID uuid.UUID, page, perPage int) ([]domain.ProductRating, int, error) {
	offset := (page - 1) * perPage
	
	var ratings []domain.ProductRating
	query := `
		SELECT id, product_id, user_id, order_id, rating, comment, created_at, updated_at, deleted_at
		FROM product_ratings
		WHERE product_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	err := r.db.Select(&ratings, query, productID, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	
	var total int
	countQuery := `SELECT COUNT(*) FROM product_ratings WHERE product_id = $1 AND deleted_at IS NULL`
	err = r.db.Get(&total, countQuery, productID)
	if err != nil {
		return nil, 0, err
	}
	
	return ratings, total, nil
}

func (r *ratingRepository) GetByUser(userID uuid.UUID) ([]domain.ProductRating, error) {
	var ratings []domain.ProductRating
	query := `
		SELECT id, product_id, user_id, order_id, rating, comment, created_at, updated_at, deleted_at
		FROM product_ratings
		WHERE user_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.Select(&ratings, query, userID)
	if err != nil {
		return nil, err
	}
	
	return ratings, nil
}

func (r *ratingRepository) GetByProductUserOrder(productID, userID, orderID uuid.UUID) (*domain.ProductRating, error) {
	var rating domain.ProductRating
	query := `
		SELECT id, product_id, user_id, order_id, rating, comment, created_at, updated_at, deleted_at
		FROM product_ratings
		WHERE product_id = $1 AND user_id = $2 AND order_id = $3 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&rating, query, productID, userID, orderID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	
	return &rating, err
}

func (r *ratingRepository) GetStats(productID uuid.UUID) (*domain.ProductRatingStats, error) {
	var stats domain.ProductRatingStats
	query := `
		SELECT 
			$1 as product_id,
			COALESCE(AVG(rating), 0) as average_rating,
			COUNT(*) as total_ratings,
			COUNT(*) FILTER (WHERE rating = 1) as rating_1_count,
			COUNT(*) FILTER (WHERE rating = 2) as rating_2_count,
			COUNT(*) FILTER (WHERE rating = 3) as rating_3_count,
			COUNT(*) FILTER (WHERE rating = 4) as rating_4_count,
			COUNT(*) FILTER (WHERE rating = 5) as rating_5_count
		FROM product_ratings
		WHERE product_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&stats, query, productID)
	if err != nil {
		return nil, err
	}
	
	return &stats, nil
}

func (r *ratingRepository) Update(rating *domain.ProductRating) error {
	now := time.Now()
	rating.UpdatedAt = &now
	
	query := `
		UPDATE product_ratings
		SET rating = $1, comment = $2, updated_at = $3
		WHERE id = $4 AND deleted_at IS NULL
	`
	
	result, err := r.db.Exec(query, rating.Rating, rating.Comment, now, rating.ID)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("rating not found")
	}
	
	return nil
}

func (r *ratingRepository) Delete(id uuid.UUID) error {
	now := time.Now()
	query := `
		UPDATE product_ratings
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`
	
	result, err := r.db.Exec(query, now, id)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("rating not found")
	}
	
	return nil
}
