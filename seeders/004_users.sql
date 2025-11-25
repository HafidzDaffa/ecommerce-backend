-- Seed Users (Admin, Seller, Customer)
-- This will populate the users table with default users for testing
-- All passwords are: password123

INSERT INTO users (id, email, password_hash, full_name, phone_number, role_id, is_email_verified, is_active, created_at) VALUES
-- Admin User (role_id: 3) - Password: password123
('a0000000-0000-0000-0000-000000000001', 'admin@ecommerce.com', '$2a$10$Z5UnNu2kAT3qgLCWbhyhgu7NCamR6X/9SKl.fd8/I9U/mOTjWQv76', 'Admin User', '081234567890', 3, true, true, NOW()),

-- Seller User (role_id: 2) - Password: password123
('b0000000-0000-0000-0000-000000000002', 'seller@ecommerce.com', '$2a$10$Z5UnNu2kAT3qgLCWbhyhgu7NCamR6X/9SKl.fd8/I9U/mOTjWQv76', 'Seller User', '081234567891', 2, true, true, NOW()),

-- Customer User (role_id: 1) - Password: password123
('c0000000-0000-0000-0000-000000000003', 'customer@ecommerce.com', '$2a$10$Z5UnNu2kAT3qgLCWbhyhgu7NCamR6X/9SKl.fd8/I9U/mOTjWQv76', 'Customer User', '081234567892', 1, true, true, NOW())
ON CONFLICT (id) DO NOTHING;

-- Note: The password hash above is a placeholder. 
-- In production, use properly hashed passwords.
-- You can generate a bcrypt hash using Go:
-- password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
