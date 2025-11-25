package repository

import (
	"database/sql"
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (email, password_hash, full_name, phone_number, role_id, is_email_verified, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.PhoneNumber,
		user.RoleID,
		user.IsEmailVerified,
		user.IsActive,
		time.Now(),
	).Scan(&user.ID)
	
	return err
}

func (r *userRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, email, password_hash, full_name, phone_number, avatar_url, role_id, 
		       is_email_verified, is_active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, email, password_hash, full_name, phone_number, avatar_url, role_id, 
		       is_email_verified, is_active, last_login_at, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	
	err := r.db.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	query := `
		UPDATE users
		SET email = $1, full_name = $2, phone_number = $3, avatar_url = $4, 
		    is_email_verified = $5, is_active = $6, updated_at = $7
		WHERE id = $8
	`
	
	_, err := r.db.Exec(
		query,
		user.Email,
		user.FullName,
		user.PhoneNumber,
		user.AvatarURL,
		user.IsEmailVerified,
		user.IsActive,
		time.Now(),
		user.ID,
	)
	
	return err
}

func (r *userRepository) UpdateLastLogin(id uuid.UUID) error {
	query := `UPDATE users SET last_login_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}
