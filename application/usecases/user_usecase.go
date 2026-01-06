package usecases

import (
	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
	"github.com/yourusername/ecommerce-go-vue/backend/domain/repositories"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (*entities.User, error)
	GetUserByID(id int64) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int64) error
	ListUsers(offset, limit int) ([]*entities.User, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Register(user *entities.User) error {
	return u.userRepo.Create(user)
}

func (u *userUseCase) Login(email, password string) (*entities.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
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
