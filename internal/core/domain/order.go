package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus struct {
	ID         int       `db:"id" json:"id"`
	StatusName string    `db:"status_name" json:"status_name"`
	Slug       string    `db:"slug" json:"slug"`
	Color      string    `db:"color" json:"color"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type Order struct {
	ID                uuid.UUID    `db:"id" json:"id"`
	UserID            uuid.UUID    `db:"user_id" json:"user_id"`
	TotalAmount       float64      `db:"total_amount" json:"total_amount"`
	ShippingCost      float64      `db:"shipping_cost" json:"shipping_cost"`
	ProductTotal      float64      `db:"product_total" json:"product_total"`
	AddressLine       string       `db:"address_line" json:"address_line"`
	PostalCode        string       `db:"postal_code" json:"postal_code"`
	ProvinceID        int          `db:"province_id" json:"province_id"`
	CityID            int          `db:"city_id" json:"city_id"`
	SubdistrictID     *int         `db:"subdistrict_id" json:"subdistrict_id,omitempty"`
	ProvinceName      string       `db:"province_name" json:"province_name"`
	CityName          string       `db:"city_name" json:"city_name"`
	SubdistrictName   *string      `db:"subdistrict_name" json:"subdistrict_name,omitempty"`
	ShippingCourier   *string      `db:"shipping_courier" json:"shipping_courier,omitempty"`
	ShippingService   *string      `db:"shipping_service" json:"shipping_service,omitempty"`
	TrackingNumber    *string      `db:"tracking_number" json:"tracking_number,omitempty"`
	OrderStatusID     int          `db:"order_status_id" json:"order_status_id"`
	CreatedAt         time.Time    `db:"created_at" json:"created_at"`
	ApprovedAt        *time.Time   `db:"approved_at" json:"approved_at,omitempty"`
	DeliveredAt       *time.Time   `db:"delivered_at" json:"delivered_at,omitempty"`
	DeletedAt         *time.Time   `db:"deleted_at" json:"deleted_at,omitempty"`
	UpdatedAt         *time.Time   `db:"updated_at" json:"updated_at,omitempty"`
	OrderStatus       *OrderStatus `db:"-" json:"order_status,omitempty"`
	Items             []OrderItem  `db:"-" json:"items,omitempty"`
}

type OrderItem struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	OrderID   uuid.UUID  `db:"order_id" json:"order_id"`
	ProductID uuid.UUID  `db:"product_id" json:"product_id"`
	Quantity  int        `db:"quantity" json:"quantity"`
	Price     float64    `db:"price" json:"price"`
	Note      *string    `db:"note" json:"note,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	Product   *Product   `db:"-" json:"product,omitempty"`
}

type CreateOrderRequest struct {
	CartIDs           []uuid.UUID `json:"cart_ids" validate:"required,min=1"`
	AddressLine       string      `json:"address_line" validate:"required"`
	PostalCode        string      `json:"postal_code" validate:"required"`
	ProvinceID        int         `json:"province_id" validate:"required"`
	CityID            int         `json:"city_id" validate:"required"`
	SubdistrictID     *int        `json:"subdistrict_id"`
	ProvinceName      string      `json:"province_name" validate:"required"`
	CityName          string      `json:"city_name" validate:"required"`
	SubdistrictName   *string     `json:"subdistrict_name"`
	ShippingCost      float64     `json:"shipping_cost" validate:"required,min=0"`
	ShippingCourier   *string     `json:"shipping_courier"`
	ShippingService   *string     `json:"shipping_service"`
}

type UpdateOrderStatusRequest struct {
	OrderStatusID  int     `json:"order_status_id" validate:"required"`
	TrackingNumber *string `json:"tracking_number"`
}

type OrderListResponse struct {
	Orders      []Order `json:"orders"`
	Total       int     `json:"total"`
	Page        int     `json:"page"`
	PerPage     int     `json:"per_page"`
	TotalPages  int     `json:"total_pages"`
}
