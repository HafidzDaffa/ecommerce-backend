package repositories

import (
	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id int64) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id int64) error
	List(offset, limit int) ([]*entities.User, error)
}

type RoleRepository interface {
	Create(role *entities.Role) error
	GetByID(id int) (*entities.Role, error)
	GetByName(name string) (*entities.Role, error)
	Update(role *entities.Role) error
	Delete(id int) error
	List(offset, limit int) ([]*entities.Role, error)
}
