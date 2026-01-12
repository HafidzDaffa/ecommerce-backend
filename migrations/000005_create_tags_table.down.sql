-- +migrate Down
DROP INDEX IF EXISTS idx_tags_name;
DROP INDEX IF EXISTS idx_tags_slug;
DROP TABLE IF EXISTS tags;
