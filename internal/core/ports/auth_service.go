package ports

import "ecommerce-backend/internal/core/domain"

type AuthService interface {
	Register(req *domain.RegisterRequest) (*domain.User, error)
	Login(req *domain.LoginRequest) (*domain.LoginResponse, error)
}
