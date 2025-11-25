package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID               uuid.UUID   `db:"id" json:"id"`
	UserID           uuid.UUID   `db:"user_id" json:"user_id"`
	ProductName      string      `db:"product_name" json:"product_name"`
	Slug             string      `db:"slug" json:"slug"`
	SKU              *string     `db:"sku" json:"sku,omitempty"`
	Price            float64     `db:"price" json:"price"`
	DiscountPercent  int         `db:"discount_percent" json:"discount_percent"`
	ShortDescription *string     `db:"short_description" json:"short_description,omitempty"`
	Description      *string     `db:"description" json:"description,omitempty"`
	WeightGram       int         `db:"weight_gram" json:"weight_gram"`
	StockQuantity    int         `db:"stock_quantity" json:"stock_quantity"`
	IsPublished      bool        `db:"is_published" json:"is_published"`
	CreatedAt        time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt        *time.Time  `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt        *time.Time  `db:"deleted_at" json:"deleted_at,omitempty"`
	Categories       []Category  `json:"categories,omitempty"`
	Galleries        []ProductGallery `json:"galleries,omitempty"`
}

type ProductGallery struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	ProductID    uuid.UUID  `db:"product_id" json:"product_id"`
	ImagePath    string     `db:"image_path" json:"image_path"`
	DisplayOrder int        `db:"display_order" json:"display_order"`
	IsThumbnail  bool       `db:"is_thumbnail" json:"is_thumbnail"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	ProductName      string   `json:"product_name" validate:"required,min=2,max=255"`
	Slug             string   `json:"slug" validate:"required,max=255"`
	SKU              *string  `json:"sku,omitempty" validate:"omitempty,max=100"`
	Price            float64  `json:"price" validate:"required,min=0"`
	DiscountPercent  int      `json:"discount_percent" validate:"min=0,max=100"`
	ShortDescription *string  `json:"short_description,omitempty" validate:"omitempty,max=500"`
	Description      *string  `json:"description,omitempty"`
	WeightGram       int      `json:"weight_gram" validate:"required,min=0"`
	StockQuantity    int      `json:"stock_quantity" validate:"required,min=0"`
	IsPublished      *bool    `json:"is_published,omitempty"`
	CategoryIDs      []int    `json:"category_ids,omitempty"`
}

type UpdateProductRequest struct {
	ProductName      *string  `json:"product_name,omitempty" validate:"omitempty,min=2,max=255"`
	Slug             *string  `json:"slug,omitempty" validate:"omitempty,max=255"`
	SKU              *string  `json:"sku,omitempty" validate:"omitempty,max=100"`
	Price            *float64 `json:"price,omitempty" validate:"omitempty,min=0"`
	DiscountPercent  *int     `json:"discount_percent,omitempty" validate:"omitempty,min=0,max=100"`
	ShortDescription *string  `json:"short_description,omitempty" validate:"omitempty,max=500"`
	Description      *string  `json:"description,omitempty"`
	WeightGram       *int     `json:"weight_gram,omitempty" validate:"omitempty,min=0"`
	StockQuantity    *int     `json:"stock_quantity,omitempty" validate:"omitempty,min=0"`
	IsPublished      *bool    `json:"is_published,omitempty"`
	CategoryIDs      []int    `json:"category_ids,omitempty"`
}

type CreateProductGalleryRequest struct {
	ProductID    string `json:"product_id" validate:"required"`
	DisplayOrder int    `json:"display_order" validate:"min=0"`
	IsThumbnail  *bool  `json:"is_thumbnail,omitempty"`
}

type UpdateProductGalleryRequest struct {
	DisplayOrder *int  `json:"display_order,omitempty" validate:"omitempty,min=0"`
	IsThumbnail  *bool `json:"is_thumbnail,omitempty"`
}
