package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user account with email, password, and role
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Registration Request"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterRequest
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password, returns JWT token and sets session cookie (24 hours)
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "Login Request"
// @Success 200 {object} domain.LoginResponse "Login successful"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Debug logging
	log.Printf("🔐 Login attempt - Email: '%s' | Email length: %d | Password length: %d", req.Email, len(req.Email), len(req.Password))

	loginResp, err := h.authService.Login(&req)
	if err != nil {
		log.Printf("❌ Login failed - Email: '%s' | Error: %s", req.Email, err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	log.Printf("✅ Login successful - Email: '%s' | User ID: %s", req.Email, loginResp.User.ID)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    loginResp.Token,
		HTTPOnly: true,
		Secure:   false,
		MaxAge:   86400,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   loginResp.Token,
		"user":    loginResp.User,
	})
}

// Logout godoc
// @Summary Logout user
// @Description Clear the session cookie
// @Tags Authentication
// @Produce json
// @Success 200 {object} map[string]interface{} "Logout successful"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}

// Me godoc
// @Summary Get current user
// @Description Get current authenticated user information
// @Tags Authentication
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Current user info"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /auth/me [get]
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userID,
		"email":   c.Locals("email"),
		"role_id": c.Locals("role_id"),
	})
}
