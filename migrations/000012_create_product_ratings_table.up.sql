CREATE TABLE IF NOT EXISTS product_ratings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    user_id UUID NOT NULL,
    order_id UUID NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_product_ratings_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT fk_product_ratings_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_product_ratings_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT uq_product_ratings_product_user_order UNIQUE (product_id, user_id, order_id)
);

CREATE INDEX IF NOT EXISTS idx_product_ratings_product_id ON product_ratings(product_id);
CREATE INDEX IF NOT EXISTS idx_product_ratings_user_id ON product_ratings(user_id);
CREATE INDEX IF NOT EXISTS idx_product_ratings_order_id ON product_ratings(order_id);
CREATE INDEX IF NOT EXISTS idx_product_ratings_deleted_at ON product_ratings(deleted_at);
