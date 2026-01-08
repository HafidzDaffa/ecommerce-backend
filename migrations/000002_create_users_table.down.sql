-- +migrate Down
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_is_verified;
DROP INDEX IF EXISTS idx_users_is_active;
DROP INDEX IF EXISTS idx_users_phone;
DROP INDEX IF EXISTS idx_users_email;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS gender_enum;
