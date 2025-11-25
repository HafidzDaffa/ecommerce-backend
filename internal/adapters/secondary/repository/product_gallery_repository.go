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

type productGalleryRepository struct {
	db *sqlx.DB
}

func NewProductGalleryRepository(db *sqlx.DB) ports.ProductGalleryRepository {
	return &productGalleryRepository{
		db: db,
	}
}

func (r *productGalleryRepository) Create(gallery *domain.ProductGallery) error {
	query := `
		INSERT INTO product_galleries (product_id, image_path, display_order, is_thumbnail, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		gallery.ProductID,
		gallery.ImagePath,
		gallery.DisplayOrder,
		gallery.IsThumbnail,
		time.Now(),
	).Scan(&gallery.ID)
	
	return err
}

func (r *productGalleryRepository) GetByID(id uuid.UUID) (*domain.ProductGallery, error) {
	var gallery domain.ProductGallery
	query := `
		SELECT id, product_id, image_path, display_order, is_thumbnail, created_at, updated_at
		FROM product_galleries
		WHERE id = $1
	`
	
	err := r.db.Get(&gallery, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product gallery not found")
		}
		return nil, err
	}
	
	return &gallery, nil
}

func (r *productGalleryRepository) GetByProductID(productID uuid.UUID) ([]domain.ProductGallery, error) {
	var galleries []domain.ProductGallery
	query := `
		SELECT id, product_id, image_path, display_order, is_thumbnail, created_at, updated_at
		FROM product_galleries
		WHERE product_id = $1
		ORDER BY display_order ASC, created_at ASC
	`
	
	err := r.db.Select(&galleries, query, productID)
	if err != nil {
		return nil, err
	}
	
	return galleries, nil
}

func (r *productGalleryRepository) Update(gallery *domain.ProductGallery) error {
	query := `
		UPDATE product_galleries
		SET display_order = $1, is_thumbnail = $2, updated_at = $3
		WHERE id = $4
	`
	
	_, err := r.db.Exec(
		query,
		gallery.DisplayOrder,
		gallery.IsThumbnail,
		time.Now(),
		gallery.ID,
	)
	
	return err
}

func (r *productGalleryRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM product_galleries WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
