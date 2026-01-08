-- +migrate Up
CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    avatar_url VARCHAR(500),
    gender gender_enum,
    date_of_birth DATE,
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    role_id INTEGER NOT NULL DEFAULT 1,
    reset_password_token VARCHAR(255),
    reset_password_expires TIMESTAMP,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_is_verified ON users(is_verified);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
