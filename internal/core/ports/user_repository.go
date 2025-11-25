package ports

import (
	"ecommerce-backend/internal/core/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uuid.UUID) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	UpdateLastLogin(id uuid.UUID) error
}
