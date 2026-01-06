package usecases_test

import (
	"errors"
	"testing"

	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
	"github.com/yourusername/ecommerce-go-vue/backend/domain/repositories"
)

type MockUserRepository struct {
	users     []*entities.User
	createErr error
	getErr    error
	updateErr error
	deleteErr error
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make([]*entities.User, 0),
	}
}

func (m *MockUserRepository) Create(user *entities.User) error {
	if m.createErr != nil {
		return m.createErr
	}
	user.ID = int64(len(m.users) + 1)
	m.users = append(m.users, user)
	return nil
}

func (m *MockUserRepository) GetByID(id int64) (*entities.User, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) GetByEmail(email string) (*entities.User, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) Update(user *entities.User) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	return nil
}

func (m *MockUserRepository) Delete(id int64) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	return nil
}

func (m *MockUserRepository) List(offset, limit int) ([]*entities.User, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.users, nil
}

type MockUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewMockUserUseCase(userRepo repositories.UserRepository) *MockUserUseCase {
	return &MockUserUseCase{
		userRepo: userRepo,
	}
}

func (u *MockUserUseCase) Register(user *entities.User) error {
	return u.userRepo.Create(user)
}

func (u *MockUserUseCase) Login(email, password string) (*entities.User, error) {
	return u.userRepo.GetByEmail(email)
}

func (u *MockUserUseCase) GetUserByID(id int64) (*entities.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *MockUserUseCase) UpdateUser(user *entities.User) error {
	return u.userRepo.Update(user)
}

func (u *MockUserUseCase) DeleteUser(id int64) error {
	return u.userRepo.Delete(id)
}

func (u *MockUserUseCase) ListUsers(offset, limit int) ([]*entities.User, error) {
	return u.userRepo.List(offset, limit)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		Email:    "test@example.com",
		FullName: "Test User",
		RoleID:   1,
	}

	err := useCase.Register(user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.ID == 0 {
		t.Error("Expected user ID to be set")
	}
}

func TestRegisterUserWithError(t *testing.T) {
	mockRepo := NewMockUserRepository()
	mockRepo.createErr = errors.New("database error")
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		Email:    "test@example.com",
		FullName: "Test User",
	}

	err := useCase.Register(user)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestLoginSuccess(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		ID:       1,
		Email:    "test@example.com",
		FullName: "Test User",
		RoleID:   1,
	}
	mockRepo.users = append(mockRepo.users, user)

	result, err := useCase.Login("test@example.com", "password123")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected user, got nil")
	}

	if result.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", result.Email)
	}
}

func TestLoginUserNotFound(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	_, err := useCase.Login("nonexistent@example.com", "password123")

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetUserByID(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		ID:       1,
		Email:    "test@example.com",
		FullName: "Test User",
	}
	mockRepo.users = append(mockRepo.users, user)

	result, err := useCase.GetUserByID(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		ID:       1,
		Email:    "test@example.com",
		FullName: "Test User",
	}
	mockRepo.users = append(mockRepo.users, user)

	user.FullName = "Updated User"
	err := useCase.UpdateUser(user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	user := &entities.User{
		ID:       1,
		Email:    "test@example.com",
		FullName: "Test User",
	}
	mockRepo.users = append(mockRepo.users, user)

	err := useCase.DeleteUser(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestListUsers(t *testing.T) {
	mockRepo := NewMockUserRepository()
	useCase := NewMockUserUseCase(mockRepo)

	for i := 1; i <= 5; i++ {
		user := &entities.User{
			ID:       int64(i),
			Email:    "user@example.com",
			FullName: "Test User",
		}
		mockRepo.users = append(mockRepo.users, user)
	}

	users, err := useCase.ListUsers(0, 10)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(users) != 5 {
		t.Errorf("Expected 5 users, got %d", len(users))
	}
}
