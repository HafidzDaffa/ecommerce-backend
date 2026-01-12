-- +migrate Down
DROP INDEX IF EXISTS idx_product_categories_category_id;
DROP TABLE IF EXISTS product_categories;
