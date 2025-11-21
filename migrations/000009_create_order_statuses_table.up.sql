CREATE TABLE IF NOT EXISTS order_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(255),
    slug VARCHAR(255) UNIQUE,
    color VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_order_statuses_slug ON order_statuses(slug);
