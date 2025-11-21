# Repository Adapters (Secondary/Output)

This directory contains implementations of repository interfaces for data persistence.

## Purpose
- Implement output ports (repository interfaces)
- Handle database operations using PostgreSQL
- Transform domain entities to/from database models

## Structure
```
repository/
├── postgres/
│   ├── user_repository.go
│   ├── product_repository.go
│   └── order_repository.go
└── models/        # Database models/DTOs
```

## Example
```go
type userRepository struct {
    db *sqlx.DB
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
    query := `INSERT INTO users (id, email, full_name) VALUES ($1, $2, $3)`
    _, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.FullName)
    return err
}
```
