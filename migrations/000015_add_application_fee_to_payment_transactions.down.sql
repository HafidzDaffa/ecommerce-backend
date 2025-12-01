-- Remove application_fee columns from payment_transactions
ALTER TABLE payment_transactions DROP COLUMN IF EXISTS application_fee_amount;
ALTER TABLE payment_transactions DROP COLUMN IF EXISTS application_fee_id;
