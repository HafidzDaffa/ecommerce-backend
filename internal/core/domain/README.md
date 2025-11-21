# Domain Layer

This directory contains the core domain entities and business rules.

## Purpose
- Define domain entities (User, Product, Order, etc.)
- Contains business logic that is independent of any framework or infrastructure
- Pure Go structs with business rules

## Example
```go
type User struct {
    ID        uuid.UUID
    Email     string
    FullName  string
    Role      Role
    CreatedAt time.Time
}
```
