package repository

import (
	"database/sql"
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) ports.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(category *domain.Category) error {
	query := `
		INSERT INTO categories (category_name, slug, icon, image_path, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		category.CategoryName,
		category.Slug,
		category.Icon,
		category.ImagePath,
		category.IsActive,
		time.Now(),
	).Scan(&category.ID)
	
	return err
}

func (r *categoryRepository) GetByID(id int) (*domain.Category, error) {
	var category domain.Category
	query := `
		SELECT id, category_name, slug, icon, image_path, is_active, created_at, updated_at
		FROM categories
		WHERE id = $1
	`
	
	err := r.db.Get(&category, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	
	return &category, nil
}

func (r *categoryRepository) GetBySlug(slug string) (*domain.Category, error) {
	var category domain.Category
	query := `
		SELECT id, category_name, slug, icon, image_path, is_active, created_at, updated_at
		FROM categories
		WHERE slug = $1
	`
	
	err := r.db.Get(&category, query, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	
	return &category, nil
}

func (r *categoryRepository) GetAll(isActive *bool) ([]domain.Category, error) {
	var categories []domain.Category
	query := `
		SELECT id, category_name, slug, icon, image_path, is_active, created_at, updated_at
		FROM categories
	`
	
	args := []interface{}{}
	if isActive != nil {
		query += ` WHERE is_active = $1`
		args = append(args, *isActive)
	}
	
	query += ` ORDER BY category_name ASC`
	
	err := r.db.Select(&categories, query, args...)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}

func (r *categoryRepository) Update(category *domain.Category) error {
	query := `
		UPDATE categories
		SET category_name = $1, slug = $2, icon = $3, image_path = $4, 
		    is_active = $5, updated_at = $6
		WHERE id = $7
	`
	
	_, err := r.db.Exec(
		query,
		category.CategoryName,
		category.Slug,
		category.Icon,
		category.ImagePath,
		category.IsActive,
		time.Now(),
		category.ID,
	)
	
	return err
}

func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
