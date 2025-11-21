# Services Layer

This directory contains the business logic implementation (use cases).

## Purpose
- Implement input ports (service interfaces)
- Orchestrate business operations
- Use output ports (repositories) to persist data
- Contains application-specific business rules

## Example
```go
type userService struct {
    userRepo ports.UserRepository
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
    // Business logic here
    return s.userRepo.Create(ctx, user)
}
```
