CREATE TABLE IF NOT EXISTS product_galleries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    image_path VARCHAR(500) NOT NULL,
    display_order INT DEFAULT 0,
    is_thumbnail BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    CONSTRAINT fk_product_galleries_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_product_galleries_product_id ON product_galleries(product_id);
CREATE INDEX IF NOT EXISTS idx_product_galleries_display_order ON product_galleries(display_order);
