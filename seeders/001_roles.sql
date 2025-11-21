-- Seed Roles
-- This will populate the roles table with default user roles

INSERT INTO roles (id, slug, name, created_at) VALUES
(1, 'customer', 'Customer', NOW()),
(2, 'seller', 'Seller', NOW()),
(3, 'admin', 'Admin', NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for roles to ensure next auto-generated ID is correct
SELECT setval('roles_id_seq', (SELECT MAX(id) FROM roles));
