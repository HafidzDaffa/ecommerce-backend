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

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) ports.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(order *domain.Order) error {
	query := `
		INSERT INTO orders (
			user_id, total_amount, shipping_cost, product_total, address_line, postal_code,
			province_id, city_id, subdistrict_id, province_name, city_name, subdistrict_name,
			shipping_courier, shipping_service, order_status_id, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		order.UserID,
		order.TotalAmount,
		order.ShippingCost,
		order.ProductTotal,
		order.AddressLine,
		order.PostalCode,
		order.ProvinceID,
		order.CityID,
		order.SubdistrictID,
		order.ProvinceName,
		order.CityName,
		order.SubdistrictName,
		order.ShippingCourier,
		order.ShippingService,
		order.OrderStatusID,
		time.Now(),
	).Scan(&order.ID)
	
	return err
}

func (r *orderRepository) CreateOrderItem(item *domain.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity, price, note, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
		item.Note,
		time.Now(),
	).Scan(&item.ID)
	
	return err
}

func (r *orderRepository) GetByID(id uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	query := `
		SELECT id, user_id, total_amount, shipping_cost, product_total, address_line, postal_code,
		       province_id, city_id, subdistrict_id, province_name, city_name, subdistrict_name,
		       shipping_courier, shipping_service, tracking_number, order_status_id,
		       created_at, approved_at, delivered_at, deleted_at, updated_at
		FROM orders
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&order, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("order not found")
	}
	
	return &order, err
}

func (r *orderRepository) GetByUser(userID uuid.UUID, page, perPage int) ([]domain.Order, int, error) {
	offset := (page - 1) * perPage
	
	var orders []domain.Order
	query := `
		SELECT id, user_id, total_amount, shipping_cost, product_total, address_line, postal_code,
		       province_id, city_id, subdistrict_id, province_name, city_name, subdistrict_name,
		       shipping_courier, shipping_service, tracking_number, order_status_id,
		       created_at, approved_at, delivered_at, deleted_at, updated_at
		FROM orders
		WHERE user_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	err := r.db.Select(&orders, query, userID, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	
	var total int
	countQuery := `SELECT COUNT(*) FROM orders WHERE user_id = $1 AND deleted_at IS NULL`
	err = r.db.Get(&total, countQuery, userID)
	if err != nil {
		return nil, 0, err
	}
	
	return orders, total, nil
}

func (r *orderRepository) GetAll(page, perPage int) ([]domain.Order, int, error) {
	offset := (page - 1) * perPage
	
	var orders []domain.Order
	query := `
		SELECT id, user_id, total_amount, shipping_cost, product_total, address_line, postal_code,
		       province_id, city_id, subdistrict_id, province_name, city_name, subdistrict_name,
		       shipping_courier, shipping_service, tracking_number, order_status_id,
		       created_at, approved_at, delivered_at, deleted_at, updated_at
		FROM orders
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.Select(&orders, query, perPage, offset)
	if err != nil {
		return nil, 0, err
	}
	
	var total int
	countQuery := `SELECT COUNT(*) FROM orders WHERE deleted_at IS NULL`
	err = r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}
	
	return orders, total, nil
}

func (r *orderRepository) GetOrderItems(orderID uuid.UUID) ([]domain.OrderItem, error) {
	var items []domain.OrderItem
	query := `
		SELECT id, order_id, product_id, quantity, price, note, created_at
		FROM order_items
		WHERE order_id = $1
	`
	
	err := r.db.Select(&items, query, orderID)
	if err != nil {
		return nil, err
	}
	
	return items, nil
}

func (r *orderRepository) GetOrderStatuses() ([]domain.OrderStatus, error) {
	var statuses []domain.OrderStatus
	query := `
		SELECT id, status_name, slug, color, created_at
		FROM order_statuses
		ORDER BY id
	`
	
	err := r.db.Select(&statuses, query)
	if err != nil {
		return nil, err
	}
	
	return statuses, nil
}

func (r *orderRepository) UpdateStatus(orderID uuid.UUID, statusID int, trackingNumber *string) error {
	now := time.Now()
	
	query := `
		UPDATE orders
		SET order_status_id = $1, tracking_number = $2, updated_at = $3
		WHERE id = $4 AND deleted_at IS NULL
	`
	
	result, err := r.db.Exec(query, statusID, trackingNumber, now, orderID)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("order not found")
	}
	
	return nil
}

func (r *orderRepository) Update(order *domain.Order) error {
	now := time.Now()
	order.UpdatedAt = &now
	
	query := `
		UPDATE orders
		SET total_amount = $1, shipping_cost = $2, product_total = $3, address_line = $4,
		    postal_code = $5, tracking_number = $6, order_status_id = $7, updated_at = $8
		WHERE id = $9 AND deleted_at IS NULL
	`
	
	result, err := r.db.Exec(
		query,
		order.TotalAmount,
		order.ShippingCost,
		order.ProductTotal,
		order.AddressLine,
		order.PostalCode,
		order.TrackingNumber,
		order.OrderStatusID,
		now,
		order.ID,
	)
	if err != nil {
		return err
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("order not found")
	}
	
	return nil
}

func (r *orderRepository) Delete(id uuid.UUID) error {
	now := time.Now()
	query := `
		UPDATE orders
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
		return errors.New("order not found")
	}
	
	return nil
}
