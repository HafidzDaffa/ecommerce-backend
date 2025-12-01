-- Add application_fee_id column to payment_transactions
ALTER TABLE payment_transactions 
ADD COLUMN application_fee_id UUID REFERENCES application_fees(id) ON DELETE SET NULL;

-- Add application_fee_amount column to payment_transactions
ALTER TABLE payment_transactions 
ADD COLUMN application_fee_amount DECIMAL(15,2) DEFAULT 0;

-- Create index for application_fee_id
CREATE INDEX idx_payment_application_fee_id ON payment_transactions(application_fee_id);

-- Comments
COMMENT ON COLUMN payment_transactions.application_fee_id IS 'Reference to the application fee used for this transaction';
COMMENT ON COLUMN payment_transactions.application_fee_amount IS 'Calculated application fee amount for this transaction';
