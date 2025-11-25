-- Seed data for products table
-- Note: Replace the user_id with an actual UUID from your users table
-- You can get a user ID by running: SELECT id FROM users LIMIT 1;

DO $$
DECLARE
    user_uuid UUID;
    product_uuid UUID;
BEGIN
    -- Get the first user from the users table
    SELECT id INTO user_uuid FROM users WHERE role_id = 2 LIMIT 1;
    
    IF user_uuid IS NULL THEN
        RAISE EXCEPTION 'No user found in users table. Please create a user first.';
    END IF;

    -- Insert Electronics Products
    INSERT INTO products (user_id, product_name, slug, sku, price, discount_percent, short_description, description, weight_gram, stock_quantity, is_published, created_at)
    VALUES 
    (user_uuid, 'Smartphone XYZ Pro', 'smartphone-xyz-pro', 'ELEC-001', 7999000, 10, 'Latest flagship smartphone', 'High-end smartphone with advanced features, 6.7" display, 128GB storage', 200, 50, TRUE, NOW()),
    (user_uuid, 'Wireless Headphones', 'wireless-headphones', 'ELEC-002', 1299000, 15, 'Premium noise-cancelling headphones', 'Bluetooth 5.0, 30-hour battery life, active noise cancellation', 250, 100, TRUE, NOW()),
    (user_uuid, 'Smart Watch Series 5', 'smart-watch-series-5', 'ELEC-003', 3499000, 5, 'Fitness and health tracking watch', 'Heart rate monitor, GPS, water resistant, 48-hour battery', 50, 75, TRUE, NOW());

    -- Get the last inserted product ID for adding to categories
    SELECT id INTO product_uuid FROM products WHERE slug = 'smartphone-xyz-pro';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 1);

    SELECT id INTO product_uuid FROM products WHERE slug = 'wireless-headphones';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 1);

    SELECT id INTO product_uuid FROM products WHERE slug = 'smart-watch-series-5';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 1);

    -- Insert Fashion Products
    INSERT INTO products (user_id, product_name, slug, sku, price, discount_percent, short_description, description, weight_gram, stock_quantity, is_published, created_at)
    VALUES 
    (user_uuid, 'Classic Cotton T-Shirt', 'classic-cotton-tshirt', 'FASH-001', 149000, 0, 'Comfortable everyday t-shirt', '100% cotton, available in multiple colors, machine washable', 150, 200, TRUE, NOW()),
    (user_uuid, 'Denim Jeans', 'denim-jeans', 'FASH-002', 499000, 20, 'Classic fit denim jeans', 'Premium denim fabric, regular fit, multiple sizes available', 400, 150, TRUE, NOW());

    SELECT id INTO product_uuid FROM products WHERE slug = 'classic-cotton-tshirt';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 2);

    SELECT id INTO product_uuid FROM products WHERE slug = 'denim-jeans';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 2);

    -- Insert Home & Garden Products
    INSERT INTO products (user_id, product_name, slug, sku, price, discount_percent, short_description, description, weight_gram, stock_quantity, is_published, created_at)
    VALUES 
    (user_uuid, 'LED Desk Lamp', 'led-desk-lamp', 'HOME-001', 299000, 10, 'Adjustable brightness desk lamp', 'Energy-efficient LED, 3 brightness levels, flexible neck', 500, 80, TRUE, NOW()),
    (user_uuid, 'Indoor Plant Pot Set', 'indoor-plant-pot-set', 'HOME-002', 179000, 0, 'Ceramic plant pots (set of 3)', 'Modern design, drainage holes, includes saucers', 1200, 60, TRUE, NOW());

    SELECT id INTO product_uuid FROM products WHERE slug = 'led-desk-lamp';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 3);

    SELECT id INTO product_uuid FROM products WHERE slug = 'indoor-plant-pot-set';
    INSERT INTO product_categories (product_id, category_id) VALUES (product_uuid, 3);

    RAISE NOTICE 'Sample products seeded successfully!';
END $$;

-- Display inserted products
SELECT p.id, p.product_name, p.slug, p.price, p.stock_quantity, 
       STRING_AGG(c.category_name, ', ') as categories
FROM products p
LEFT JOIN product_categories pc ON p.id = pc.product_id
LEFT JOIN categories c ON pc.category_id = c.id
GROUP BY p.id, p.product_name, p.slug, p.price, p.stock_quantity
ORDER BY p.created_at DESC;
