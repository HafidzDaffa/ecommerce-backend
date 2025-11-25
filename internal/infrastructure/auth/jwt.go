package auth

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/infrastructure/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	RoleID int       `json:"role_id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	config *config.Config
}

func NewJWTService(cfg *config.Config) *JWTService {
	return &JWTService{
		config: cfg,
	}
}

func (s *JWTService) GenerateToken(user *domain.User) (string, error) {
	if s.config.JWT.Secret == "" {
		return "", errors.New("JWT secret is not configured")
	}

	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JWT.Expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWT.Secret))
}

func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	if s.config.JWT.Secret == "" {
		return nil, errors.New("JWT secret is not configured")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
