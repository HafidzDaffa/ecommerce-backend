CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    shipping_cost DECIMAL(15, 2) NOT NULL,
    product_total DECIMAL(15, 2) NOT NULL,
    address_line TEXT NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    province_id INT NOT NULL,
    city_id INT NOT NULL,
    subdistrict_id INT,
    province_name VARCHAR(255) NOT NULL,
    city_name VARCHAR(255) NOT NULL,
    subdistrict_name VARCHAR(255),
    shipping_courier VARCHAR(100),
    shipping_service VARCHAR(100),
    tracking_number VARCHAR(255),
    order_status_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    approved_at TIMESTAMP,
    delivered_at TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_orders_order_status FOREIGN KEY (order_status_id) REFERENCES order_statuses(id) ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_order_status_id ON orders(order_status_id);
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);
CREATE INDEX IF NOT EXISTS idx_orders_deleted_at ON orders(deleted_at);
