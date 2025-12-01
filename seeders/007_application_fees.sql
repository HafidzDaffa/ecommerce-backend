-- Application Fees Seeder
-- Insert sample application fees for testing

-- Platform Fee (Percentage) - 2.5% platform fee
INSERT INTO application_fees (id, fee_name, fee_type, fee_value, description, is_active, created_at)
VALUES 
(gen_random_uuid(), 'Platform Fee (Percentage)', 'PERCENTAGE', 2.5, 'Platform service fee calculated as percentage of order amount', true, CURRENT_TIMESTAMP);

-- Platform Fee (Fixed) - Rp 5000 fixed fee
INSERT INTO application_fees (id, fee_name, fee_type, fee_value, description, is_active, created_at)
VALUES 
(gen_random_uuid(), 'Platform Fee (Fixed)', 'FIXED', 5000, 'Fixed platform service fee per transaction', false, CURRENT_TIMESTAMP);

-- Express Service Fee - 3.5% for express orders
INSERT INTO application_fees (id, fee_name, fee_type, fee_value, description, is_active, created_at)
VALUES 
(gen_random_uuid(), 'Express Service Fee', 'PERCENTAGE', 3.5, 'Additional fee for express delivery orders', false, CURRENT_TIMESTAMP);

-- Processing Fee - Rp 2500 fixed processing fee
INSERT INTO application_fees (id, fee_name, fee_type, fee_value, description, is_active, created_at)
VALUES 
(gen_random_uuid(), 'Processing Fee', 'FIXED', 2500, 'Transaction processing fee', false, CURRENT_TIMESTAMP);
