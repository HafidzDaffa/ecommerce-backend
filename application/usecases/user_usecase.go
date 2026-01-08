package usecases

import (
	"errors"

	"github.com/yourusername/ecommerce-go-vue/backend/common/utils"
	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
	"github.com/yourusername/ecommerce-go-vue/backend/domain/repositories"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/config"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (string, *entities.User, error)
	GetUserByID(id int64) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int64) error
	ListUsers(offset, limit int) ([]*entities.User, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
	cfg      *config.Config
}

func NewUserUseCase(userRepo repositories.UserRepository, cfg *config.Config) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (u *userUseCase) Register(user *entities.User) error {
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword
	return u.userRepo.Create(user)
}

func (u *userUseCase) Login(email, password string) (string, *entities.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return "", nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return "", nil, errors.New("user account is inactive")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.RoleID, u.cfg.JWTSecret, u.cfg.JWTExpiry)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (u *userUseCase) GetUserByID(id int64) (*entities.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) UpdateUser(user *entities.User) error {
	return u.userRepo.Update(user)
}

func (u *userUseCase) DeleteUser(id int64) error {
	return u.userRepo.Delete(id)
}

func (u *userUseCase) ListUsers(offset, limit int) ([]*entities.User, error) {
	return u.userRepo.List(offset, limit)
}
