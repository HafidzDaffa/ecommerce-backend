-- Seed Categories
-- This will populate the categories table with default product categories
-- Images are sourced from Unsplash (free to use)

INSERT INTO categories (id, category_name, slug, icon, image_path, is_active, created_at) VALUES
(1, 'Electronics', 'electronics', '⚡', 'https://images.unsplash.com/photo-1498049794561-7780e7231661?w=800&q=80', true, NOW()),
(2, 'Fashion', 'fashion', '👔', 'https://images.unsplash.com/photo-1445205170230-053b83016050?w=800&q=80', true, NOW()),
(3, 'Home & Living', 'home-living', '🏠', 'https://images.unsplash.com/photo-1484101403633-562f891dc89a?w=800&q=80', true, NOW()),
(4, 'Beauty & Health', 'beauty-health', '💄', 'https://images.unsplash.com/photo-1596462502278-27bfdc403348?w=800&q=80', true, NOW()),
(5, 'Sports & Outdoor', 'sports-outdoor', '⚽', 'https://images.unsplash.com/photo-1461896836934-ffe607ba8211?w=800&q=80', true, NOW()),
(6, 'Books & Stationery', 'books-stationery', '📚', 'https://images.unsplash.com/photo-1495446815901-a7297e633e8d?w=800&q=80', true, NOW()),
(7, 'Toys & Games', 'toys-games', '🎮', 'https://images.unsplash.com/photo-1493711662062-fa541adb3fc8?w=800&q=80', true, NOW()),
(8, 'Food & Beverage', 'food-beverage', '🍔', 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=800&q=80', true, NOW()),
(9, 'Automotive', 'automotive', '🚗', 'https://images.unsplash.com/photo-1492144534655-ae79c964c9d7?w=800&q=80', true, NOW()),
(10, 'Baby & Kids', 'baby-kids', '👶', 'https://images.unsplash.com/photo-1515488042361-ee00e0ddd4e4?w=800&q=80', true, NOW()),
(11, 'Pet Supplies', 'pet-supplies', '🐾', 'https://images.unsplash.com/photo-1450778869180-41d0601e046e?w=800&q=80', true, NOW()),
(12, 'Office Supplies', 'office-supplies', '💼', 'https://images.unsplash.com/photo-1497366216548-37526070297c?w=800&q=80', true, NOW()),
(13, 'Garden & Outdoor', 'garden-outdoor', '🌱', 'https://images.unsplash.com/photo-1416879595882-3373a0480b5b?w=800&q=80', true, NOW()),
(14, 'Musical Instruments', 'musical-instruments', '🎸', 'https://images.unsplash.com/photo-1511379938547-c1f69419868d?w=800&q=80', true, NOW()),
(15, 'Jewelry & Accessories', 'jewelry-accessories', '💎', 'https://images.unsplash.com/photo-1515562141207-7a88fb7ce338?w=800&q=80', true, NOW()),
(16, 'Arts & Crafts', 'arts-crafts', '🎨', 'https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=800&q=80', true, NOW()),
(17, 'Furniture', 'furniture', '🛋️', 'https://images.unsplash.com/photo-1555041469-a586c61ea9bc?w=800&q=80', true, NOW()),
(18, 'Computer & Laptops', 'computer-laptops', '💻', 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=800&q=80', true, NOW()),
(19, 'Mobile Phones', 'mobile-phones', '📱', 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=800&q=80', true, NOW()),
(20, 'Cameras & Photography', 'cameras-photography', '📷', 'https://images.unsplash.com/photo-1502920917128-1aa500764cbd?w=800&q=80', true, NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for categories to ensure next auto-generated ID is correct
SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));
