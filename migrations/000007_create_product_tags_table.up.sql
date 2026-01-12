-- +migrate Up
CREATE TABLE product_tags (
    product_id BIGINT NOT NULL,
    tag_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, tag_id),
    CONSTRAINT fk_product_tags_products FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT fk_product_tags_tags FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
