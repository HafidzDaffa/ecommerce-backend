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

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ports.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *domain.Product) error {
	query := `
		INSERT INTO products (user_id, product_name, slug, sku, price, discount_percent, 
		                      short_description, description, weight_gram, stock_quantity, 
		                      is_published, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		product.UserID,
		product.ProductName,
		product.Slug,
		product.SKU,
		product.Price,
		product.DiscountPercent,
		product.ShortDescription,
		product.Description,
		product.WeightGram,
		product.StockQuantity,
		product.IsPublished,
		time.Now(),
	).Scan(&product.ID)
	
	return err
}

func (r *productRepository) GetByID(id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	query := `
		SELECT id, user_id, product_name, slug, sku, price, discount_percent, 
		       short_description, description, weight_gram, stock_quantity, is_published,
		       created_at, updated_at, deleted_at
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	
	return &product, nil
}

func (r *productRepository) GetBySlug(slug string) (*domain.Product, error) {
	var product domain.Product
	query := `
		SELECT id, user_id, product_name, slug, sku, price, discount_percent, 
		       short_description, description, weight_gram, stock_quantity, is_published,
		       created_at, updated_at, deleted_at
		FROM products
		WHERE slug = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&product, query, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	
	return &product, nil
}

func (r *productRepository) GetAll(page, limit int, isPublished *bool) ([]domain.Product, int, error) {
	var products []domain.Product
	
	query := `
		SELECT id, user_id, product_name, slug, sku, price, discount_percent, 
		       short_description, description, weight_gram, stock_quantity, is_published,
		       created_at, updated_at, deleted_at
		FROM products
		WHERE deleted_at IS NULL
	`
	
	args := []interface{}{}
	argPos := 1
	
	if isPublished != nil {
		query += fmt.Sprintf(` AND is_published = $%d`, argPos)
		args = append(args, *isPublished)
		argPos++
	}
	
	query += ` ORDER BY created_at DESC`
	
	countQuery := `SELECT COUNT(*) FROM products WHERE deleted_at IS NULL`
	if isPublished != nil {
		countQuery += ` AND is_published = $1`
	}
	
	var total int
	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	
	offset := (page - 1) * limit
	query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, argPos, argPos+1)
	args = append(args, limit, offset)
	
	err = r.db.Select(&products, query, args...)
	if err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

func (r *productRepository) GetByCategoryID(categoryID int, page, limit int) ([]domain.Product, int, error) {
	var products []domain.Product
	
	query := `
		SELECT DISTINCT p.id, p.user_id, p.product_name, p.slug, p.sku, p.price, 
		       p.discount_percent, p.short_description, p.description, p.weight_gram, 
		       p.stock_quantity, p.is_published, p.created_at, p.updated_at, p.deleted_at
		FROM products p
		INNER JOIN product_categories pc ON p.id = pc.product_id
		WHERE pc.category_id = $1 AND p.deleted_at IS NULL AND p.is_published = TRUE
		ORDER BY p.created_at DESC
	`
	
	countQuery := `
		SELECT COUNT(DISTINCT p.id)
		FROM products p
		INNER JOIN product_categories pc ON p.id = pc.product_id
		WHERE pc.category_id = $1 AND p.deleted_at IS NULL AND p.is_published = TRUE
	`
	
	var total int
	err := r.db.Get(&total, countQuery, categoryID)
	if err != nil {
		return nil, 0, err
	}
	
	offset := (page - 1) * limit
	query += ` LIMIT $2 OFFSET $3`
	
	err = r.db.Select(&products, query, categoryID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

func (r *productRepository) Update(product *domain.Product) error {
	query := `
		UPDATE products
		SET product_name = $1, slug = $2, sku = $3, price = $4, discount_percent = $5,
		    short_description = $6, description = $7, weight_gram = $8, stock_quantity = $9,
		    is_published = $10, updated_at = $11
		WHERE id = $12 AND deleted_at IS NULL
	`
	
	_, err := r.db.Exec(
		query,
		product.ProductName,
		product.Slug,
		product.SKU,
		product.Price,
		product.DiscountPercent,
		product.ShortDescription,
		product.Description,
		product.WeightGram,
		product.StockQuantity,
		product.IsPublished,
		time.Now(),
		product.ID,
	)
	
	return err
}

func (r *productRepository) Delete(id uuid.UUID) error {
	query := `UPDATE products SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

func (r *productRepository) AddCategories(productID uuid.UUID, categoryIDs []int) error {
	if len(categoryIDs) == 0 {
		return nil
	}
	
	query := `
		INSERT INTO product_categories (product_id, category_id)
		VALUES ($1, $2)
		ON CONFLICT (product_id, category_id) DO NOTHING
	`
	
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	
	for _, categoryID := range categoryIDs {
		_, err := tx.Exec(query, productID, categoryID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	return tx.Commit()
}

func (r *productRepository) RemoveCategories(productID uuid.UUID, categoryIDs []int) error {
	if len(categoryIDs) == 0 {
		return nil
	}
	
	query := `DELETE FROM product_categories WHERE product_id = $1 AND category_id = ANY($2)`
	_, err := r.db.Exec(query, productID, categoryIDs)
	return err
}

func (r *productRepository) GetCategories(productID uuid.UUID) ([]domain.Category, error) {
	var categories []domain.Category
	query := `
		SELECT c.id, c.category_name, c.slug, c.icon, c.image_path, c.is_active, 
		       c.created_at, c.updated_at
		FROM categories c
		INNER JOIN product_categories pc ON c.id = pc.category_id
		WHERE pc.product_id = $1
		ORDER BY c.category_name ASC
	`
	
	err := r.db.Select(&categories, query, productID)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}
