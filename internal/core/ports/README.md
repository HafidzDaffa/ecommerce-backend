# Ports Layer

This directory contains interfaces that define the boundaries of the application.

## Types of Ports

### Input Ports (Services)
Interfaces that define use cases/business operations
```go
type UserService interface {
    CreateUser(ctx context.Context, user *domain.User) error
    GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
```

### Output Ports (Repositories)
Interfaces for data persistence operations
```go
type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
```
