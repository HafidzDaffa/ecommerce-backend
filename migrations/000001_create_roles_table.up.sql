CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_roles_slug ON roles(slug);
