package database

import (
	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *entities.User) error {
	return DB.Create(user).Error
}

func (r *UserRepository) GetByID(id int64) (*entities.User, error) {
	var user entities.User
	err := DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *entities.User) error {
	return DB.Save(user).Error
}

func (r *UserRepository) Delete(id int64) error {
	return DB.Delete(&entities.User{}, id).Error
}

func (r *UserRepository) List(offset, limit int) ([]*entities.User, error) {
	var users []*entities.User
	err := DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

type RoleRepository struct {
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (r *RoleRepository) Create(role *entities.Role) error {
	return DB.Create(role).Error
}

func (r *RoleRepository) GetByID(id int) (*entities.Role, error) {
	var role entities.Role
	err := DB.First(&role, id).Error
	return &role, err
}

func (r *RoleRepository) GetByName(name string) (*entities.Role, error) {
	var role entities.Role
	err := DB.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *RoleRepository) Update(role *entities.Role) error {
	return DB.Save(role).Error
}

func (r *RoleRepository) Delete(id int) error {
	return DB.Delete(&entities.Role{}, id).Error
}

func (r *RoleRepository) List(offset, limit int) ([]*entities.Role, error) {
	var roles []*entities.Role
	err := DB.Offset(offset).Limit(limit).Find(&roles).Error
	return roles, err
}
