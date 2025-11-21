CREATE TABLE IF NOT EXISTS user_addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    recipient_name VARCHAR(255) NOT NULL,
    recipient_phone VARCHAR(50) NOT NULL,
    label VARCHAR(100) DEFAULT 'Home',
    address_line TEXT NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    province_id INT NOT NULL,
    city_id INT NOT NULL,
    subdistrict_id INT,
    province_name VARCHAR(255) NOT NULL,
    city_name VARCHAR(255) NOT NULL,
    subdistrict_name VARCHAR(255),
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user_addresses_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_addresses_user_id ON user_addresses(user_id);
CREATE INDEX IF NOT EXISTS idx_user_addresses_deleted_at ON user_addresses(deleted_at);
