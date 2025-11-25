package services

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"ecommerce-backend/internal/infrastructure/auth"
	"errors"
	"strings"
)

type authService struct {
	userRepo   ports.UserRepository
	jwtService *auth.JWTService
}

func NewAuthService(userRepo ports.UserRepository, jwtService *auth.JWTService) ports.AuthService {
	return &authService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *authService) Register(req *domain.RegisterRequest) (*domain.User, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	
	if req.Password == "" || len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	if req.RoleID < 1 || req.RoleID > 3 {
		return nil, errors.New("invalid role_id, must be 1 (customer), 2 (seller), or 3 (admin)")
	}

	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &domain.User{
		Email:           req.Email,
		PasswordHash:    hashedPassword,
		FullName:        req.FullName,
		PhoneNumber:     req.PhoneNumber,
		RoleID:          req.RoleID,
		IsEmailVerified: false,
		IsActive:        true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (s *authService) Login(req *domain.LoginRequest) (*domain.LoginResponse, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	if !auth.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	if err := s.userRepo.UpdateLastLogin(user.ID); err != nil {
		return nil, errors.New("failed to update last login")
	}

	return &domain.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
