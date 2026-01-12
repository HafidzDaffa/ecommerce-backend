-- +migrate Down
DROP INDEX IF EXISTS idx_categories_deleted_at;
DROP INDEX IF EXISTS idx_categories_is_active;
DROP INDEX IF EXISTS idx_categories_slug;
DROP TABLE IF EXISTS categories;
