-- +migrate Down
DROP INDEX IF EXISTS idx_products_deleted_at;
DROP INDEX IF EXISTS idx_products_is_active;
DROP INDEX IF EXISTS idx_products_status;
DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_sku;
DROP INDEX IF EXISTS idx_products_slug;
DROP INDEX IF EXISTS idx_products_user_id;
DROP TABLE IF EXISTS products;
DROP TYPE IF EXISTS product_status_enum;
