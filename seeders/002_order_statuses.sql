-- Seed Order Statuses
-- This will populate the order_statuses table with default statuses

INSERT INTO order_statuses (id, status_name, slug, color, created_at) VALUES
(1, 'Pending', 'pending', '#FFA500', NOW()),
(2, 'Processing', 'processing', '#1E90FF', NOW()),
(3, 'Shipped', 'shipped', '#9370DB', NOW()),
(4, 'Delivered', 'delivered', '#32CD32', NOW()),
(5, 'Cancelled', 'cancelled', '#DC143C', NOW()),
(6, 'Refunded', 'refunded', '#808080', NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for order_statuses to ensure next auto-generated ID is correct
SELECT setval('order_statuses_id_seq', (SELECT MAX(id) FROM order_statuses));
