CREATE TABLE IF NOT EXISTS application_fees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fee_name VARCHAR(255) NOT NULL,
    fee_type VARCHAR(50) NOT NULL,
    fee_value DECIMAL(15,2) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    -- Constraints
    CONSTRAINT chk_fee_type CHECK (fee_type IN ('PERCENTAGE', 'FIXED')),
    CONSTRAINT chk_fee_value_positive CHECK (fee_value >= 0)
);

-- Indexes for better query performance
CREATE INDEX idx_application_fees_fee_type ON application_fees(fee_type);
CREATE INDEX idx_application_fees_is_active ON application_fees(is_active);
CREATE INDEX idx_application_fees_created_by ON application_fees(created_by);
CREATE INDEX idx_application_fees_created_at ON application_fees(created_at);

-- Comments
COMMENT ON TABLE application_fees IS 'Stores application fee configurations set by admin';
COMMENT ON COLUMN application_fees.fee_type IS 'Type of fee: PERCENTAGE (e.g., 2.5%) or FIXED (e.g., Rp 5000)';
COMMENT ON COLUMN application_fees.fee_value IS 'Fee value: percentage (0-100) or fixed amount';
COMMENT ON COLUMN application_fees.is_active IS 'Whether this fee configuration is currently active';
