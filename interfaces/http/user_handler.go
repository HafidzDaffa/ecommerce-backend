package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/ecommerce-go-vue/backend/application/dtos"
	"github.com/yourusername/ecommerce-go-vue/backend/application/usecases"
	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dtos.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user := &entities.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		FullName:     req.FullName,
		Phone:        req.Phone,
		Gender:       req.Gender,
		RoleID:       1,
		IsActive:     true,
		IsVerified:   false,
	}

	if req.Gender == "" {
		user.Gender = "other"
	}

	if err := h.userUseCase.Register(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"data": fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
		},
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, user, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"data": dtos.AuthResponse{
			Token: token,
			User: fiber.Map{
				"id":          user.ID,
				"email":       user.Email,
				"full_name":   user.FullName,
				"role_id":     user.RoleID,
				"is_active":   user.IsActive,
				"is_verified": user.IsVerified,
			},
		},
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req dtos.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}

	if err := h.userUseCase.UpdateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	if err := h.userUseCase.DeleteUser(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"data": fiber.Map{
			"id": id,
		},
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	users, err := h.userUseCase.ListUsers(offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"users": users,
			"page":  page,
			"limit": limit,
		},
	})
}
