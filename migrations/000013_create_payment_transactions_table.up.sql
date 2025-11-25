CREATE TABLE IF NOT EXISTS payment_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Xendit transaction details
    xendit_invoice_id VARCHAR(255) UNIQUE,
    xendit_external_id VARCHAR(255) UNIQUE NOT NULL,
    
    -- Payment information
    payment_method VARCHAR(50),
    payment_channel VARCHAR(100),
    payment_status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    
    -- Amount details
    amount DECIMAL(15,2) NOT NULL,
    paid_amount DECIMAL(15,2) DEFAULT 0,
    admin_fee DECIMAL(15,2) DEFAULT 0,
    total_amount DECIMAL(15,2) NOT NULL,
    
    -- Payment URLs
    invoice_url TEXT,
    checkout_url TEXT,
    
    -- Payment dates
    paid_at TIMESTAMP,
    expired_at TIMESTAMP,
    
    -- Additional information
    description TEXT,
    payment_details JSONB,
    xendit_response JSONB,
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    -- Indexes
    CONSTRAINT fk_payment_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_payment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Indexes for better query performance
CREATE INDEX idx_payment_order_id ON payment_transactions(order_id);
CREATE INDEX idx_payment_user_id ON payment_transactions(user_id);
CREATE INDEX idx_payment_status ON payment_transactions(payment_status);
CREATE INDEX idx_payment_xendit_invoice_id ON payment_transactions(xendit_invoice_id);
CREATE INDEX idx_payment_xendit_external_id ON payment_transactions(xendit_external_id);
CREATE INDEX idx_payment_created_at ON payment_transactions(created_at);

-- Payment status check constraint
ALTER TABLE payment_transactions 
ADD CONSTRAINT chk_payment_status 
CHECK (payment_status IN ('PENDING', 'PAID', 'EXPIRED', 'FAILED', 'CANCELLED'));

-- Comment on table
COMMENT ON TABLE payment_transactions IS 'Stores payment transaction information integrated with Xendit payment gateway';
COMMENT ON COLUMN payment_transactions.xendit_invoice_id IS 'Invoice ID from Xendit';
COMMENT ON COLUMN payment_transactions.xendit_external_id IS 'External ID for tracking in Xendit';
COMMENT ON COLUMN payment_transactions.payment_status IS 'Payment status: PENDING, PAID, EXPIRED, FAILED, CANCELLED';
COMMENT ON COLUMN payment_transactions.payment_details IS 'Additional payment details in JSON format';
COMMENT ON COLUMN payment_transactions.xendit_response IS 'Raw response from Xendit API in JSON format';
