-- Seed data for categories table
INSERT INTO categories (category_name, slug, icon, is_active, created_at) VALUES
('Electronics', 'electronics', '📱', TRUE, NOW()),
('Fashion', 'fashion', '👕', TRUE, NOW()),
('Home & Garden', 'home-garden', '🏠', TRUE, NOW()),
('Sports & Outdoors', 'sports-outdoors', '⚽', TRUE, NOW()),
('Books & Media', 'books-media', '📚', TRUE, NOW()),
('Toys & Games', 'toys-games', '🎮', TRUE, NOW()),
('Health & Beauty', 'health-beauty', '💄', TRUE, NOW()),
('Automotive', 'automotive', '🚗', TRUE, NOW()),
('Food & Beverages', 'food-beverages', '🍔', TRUE, NOW()),
('Office Supplies', 'office-supplies', '✏️', TRUE, NOW());

-- Display inserted categories
SELECT * FROM categories ORDER BY id;
